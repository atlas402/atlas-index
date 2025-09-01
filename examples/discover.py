import asyncio
from atlas_index.core.index import AtlasIndex

async def main():
    index = AtlasIndex(facilitator_url='https://facilitator.payai.network')
    
    services = await index.discover(category='AI', network='base')
    
    print(f"Found {len(services)} services")

if __name__ == '__main__':
    asyncio.run(main())

