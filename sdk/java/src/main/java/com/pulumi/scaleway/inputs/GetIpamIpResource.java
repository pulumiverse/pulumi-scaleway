// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetIpamIpResource extends com.pulumi.resources.InvokeArgs {

    public static final GetIpamIpResource Empty = new GetIpamIpResource();

    /**
     * The ID of the resource that the IP is bound to.
     * 
     */
    @Import(name="id")
    private @Nullable String id;

    /**
     * @return The ID of the resource that the IP is bound to.
     * 
     */
    public Optional<String> id() {
        return Optional.ofNullable(this.id);
    }

    /**
     * The type of the resource to get the IP from. [Documentation](https://pkg.go.dev/github.com/scaleway/scaleway-sdk-go@master/api/ipam/v1alpha1#pkg-constants) with type list.
     * 
     */
    @Import(name="type")
    private @Nullable String type;

    /**
     * @return The type of the resource to get the IP from. [Documentation](https://pkg.go.dev/github.com/scaleway/scaleway-sdk-go@master/api/ipam/v1alpha1#pkg-constants) with type list.
     * 
     */
    public Optional<String> type() {
        return Optional.ofNullable(this.type);
    }

    private GetIpamIpResource() {}

    private GetIpamIpResource(GetIpamIpResource $) {
        this.id = $.id;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetIpamIpResource defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetIpamIpResource $;

        public Builder() {
            $ = new GetIpamIpResource();
        }

        public Builder(GetIpamIpResource defaults) {
            $ = new GetIpamIpResource(Objects.requireNonNull(defaults));
        }

        /**
         * @param id The ID of the resource that the IP is bound to.
         * 
         * @return builder
         * 
         */
        public Builder id(@Nullable String id) {
            $.id = id;
            return this;
        }

        /**
         * @param type The type of the resource to get the IP from. [Documentation](https://pkg.go.dev/github.com/scaleway/scaleway-sdk-go@master/api/ipam/v1alpha1#pkg-constants) with type list.
         * 
         * @return builder
         * 
         */
        public Builder type(@Nullable String type) {
            $.type = type;
            return this;
        }

        public GetIpamIpResource build() {
            return $;
        }
    }

}