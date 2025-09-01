import pytest
from atlas_index.core.index import AtlasIndex

@pytest.mark.asyncio
async def test_discover_services():
    index = AtlasIndex(facilitator_url='https://facilitator.payai.network')
    services = await index.discover(category='AI', network='base')
    assert isinstance(services, list)

