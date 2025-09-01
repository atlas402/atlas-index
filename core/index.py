from typing import Optional, Dict, Any, List
from dataclasses import dataclass
import aiohttp

@dataclass
class PaymentAccept:
    asset: str
    pay_to: str
    network: str
    max_amount_required: str
    scheme: str
    mime_type: str

@dataclass
class ServiceDiscovery:
    id: str
    name: str
    description: str
    endpoint: str
    category: str
    network: str
    accepts: List[PaymentAccept]
    metadata: Optional[Dict[str, Any]] = None

class AtlasIndex:
    def __init__(self, facilitator_url: str):
        self.facilitator_url = facilitator_url
        self.services: Dict[str, ServiceDiscovery] = {}
        
    async def discover(
        self,
        category: Optional[str] = None,
        network: Optional[str] = None,
        scheme: Optional[str] = None,
        limit: Optional[int] = None,
        offset: Optional[int] = None
    ) -> List[ServiceDiscovery]:
        url = f"{self.facilitator_url}/discovery/resources"
        params = {}
        
        if category:
            params['category'] = category
        if network:
            params['network'] = network
        if scheme:
            params['scheme'] = scheme
        if limit:
            params['limit'] = limit
        if offset:
            params['offset'] = offset
            
        async with aiohttp.ClientSession() as session:
            async with session.get(url, params=params) as response:
                if response.status != 200:
                    raise Exception(f"Discovery failed: {response.status}")
                
                data = await response.json()
                resources = data.get('resources', [])
                
                services = []
                for resource in resources:
                    service = ServiceDiscovery(
                        id=resource['id'],
                        name=resource['name'],
                        description=resource.get('description', ''),
                        endpoint=resource['endpoint'],
                        category=resource.get('category', ''),
                        network=resource.get('network', 'base'),
                        accepts=[
                            PaymentAccept(**accept) for accept in resource.get('accepts', [])
                        ],
                        metadata=resource.get('metadata'),
                    )
                    services.append(service)
                    self.services[service.id] = service
                    
                return services
    
    def get_service(self, service_id: str) -> Optional[ServiceDiscovery]:
        return self.services.get(service_id)


