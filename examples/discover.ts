import { AtlasIndex } from './core/index';

async function main() {
  const index = new AtlasIndex({
    facilitatorUrl: 'https://facilitator.payai.network',
  });

  const services = await index.discover({
    category: 'AI',
    network: 'base',
  });

  console.log(`Found ${services.length} services`);
}

main().catch(console.error);


