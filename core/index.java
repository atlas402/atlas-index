package com.atlas402.index.core;

import java.util.List;
import java.util.concurrent.CompletableFuture;

public class AtlasIndex {
    private final String facilitatorUrl;
    
    public AtlasIndex(String facilitatorUrl) {
        this.facilitatorUrl = facilitatorUrl;
    }
    
    public CompletableFuture<List<ServiceDiscovery>> discover(DiscoveryOptions options) {
        return CompletableFuture.supplyAsync(() -> {
            return List.of();
        });
    }
    
    public ServiceDiscovery getService(String serviceId) {
        return null;
    }
}

