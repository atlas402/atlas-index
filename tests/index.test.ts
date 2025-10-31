import { describe, it, expect } from 'vitest';
import { AtlasIndex } from '../core/index';

describe('AtlasIndex', () => {
  it('should create instance', () => {
    const index = new AtlasIndex({
      facilitatorUrl: 'https://facilitator.payai.network',
    });
    expect(index).toBeDefined();
  });
});


