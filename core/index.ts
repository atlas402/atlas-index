export interface ServiceDiscovery {
  id: string;
  name: string;
  description: string;
  endpoint: string;
  category: string;
  network: string;
  accepts: PaymentAccept[];
  metadata?: Record<string, any>;
}

export interface PaymentAccept {
  asset: string;
  payTo: string;
  network: string;
  maxAmountRequired: string;
  scheme: string;
  mimeType: string;
}

export class AtlasIndex {
  private facilitatorUrl: string;
  private serviceCache: Map<string, ServiceDiscovery> = new Map();

  constructor(config: { facilitatorUrl: string }) {
    this.facilitatorUrl = config.facilitatorUrl;
  }

  async discover(options: {
    category?: string;
    network?: 'base' | 'solana-mainnet';
    scheme?: string;
    limit?: number;
    offset?: number;
  } = {}): Promise<ServiceDiscovery[]> {
    const url = new URL(`${this.facilitatorUrl}/discovery/resources`);
    
    if (options.category) url.searchParams.set('category', options.category);
    if (options.network) url.searchParams.set('network', options.network);
    if (options.scheme) url.searchParams.set('scheme', options.scheme);
    if (options.limit) url.searchParams.set('limit', options.limit.toString());
    if (options.offset) url.searchParams.set('offset', options.offset.toString());

    const response = await fetch(url.toString());
    const data = await response.json();
    const services = data.resources || [];

    services.forEach((service: ServiceDiscovery) => {
      this.serviceCache.set(service.id, service);
    });

    return services;
  }

  getService(serviceId: string): ServiceDiscovery | undefined {
    return this.serviceCache.get(serviceId);
  }
}

export default AtlasIndex;

