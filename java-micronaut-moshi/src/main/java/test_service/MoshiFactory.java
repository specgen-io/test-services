package test_service;

import com.squareup.moshi.Moshi;
import io.micronaut.context.annotation.*;
import jakarta.inject.Singleton;
import test_service.json.Json;

@Factory
public class MoshiFactory {
    @Singleton
    public Moshi getMoshi() {
        var moshiBuilder = new Moshi.Builder();
        Json.setupMoshiAdapters(moshiBuilder);
        return moshiBuilder.build();
    }
}