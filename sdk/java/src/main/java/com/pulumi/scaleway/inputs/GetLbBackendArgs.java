// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetLbBackendArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetLbBackendArgs Empty = new GetLbBackendArgs();

    /**
     * The backend id.
     * - Only one of `name` and `backend_id` should be specified.
     * 
     */
    @Import(name="backendId")
    private @Nullable Output<String> backendId;

    /**
     * @return The backend id.
     * - Only one of `name` and `backend_id` should be specified.
     * 
     */
    public Optional<Output<String>> backendId() {
        return Optional.ofNullable(this.backendId);
    }

    /**
     * The load-balancer ID this backend is attached to.
     * 
     */
    @Import(name="lbId")
    private @Nullable Output<String> lbId;

    /**
     * @return The load-balancer ID this backend is attached to.
     * 
     */
    public Optional<Output<String>> lbId() {
        return Optional.ofNullable(this.lbId);
    }

    /**
     * The name of the backend.
     * - When using the `name` you should specify the `lb-id`
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the backend.
     * - When using the `name` you should specify the `lb-id`
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    private GetLbBackendArgs() {}

    private GetLbBackendArgs(GetLbBackendArgs $) {
        this.backendId = $.backendId;
        this.lbId = $.lbId;
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetLbBackendArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetLbBackendArgs $;

        public Builder() {
            $ = new GetLbBackendArgs();
        }

        public Builder(GetLbBackendArgs defaults) {
            $ = new GetLbBackendArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param backendId The backend id.
         * - Only one of `name` and `backend_id` should be specified.
         * 
         * @return builder
         * 
         */
        public Builder backendId(@Nullable Output<String> backendId) {
            $.backendId = backendId;
            return this;
        }

        /**
         * @param backendId The backend id.
         * - Only one of `name` and `backend_id` should be specified.
         * 
         * @return builder
         * 
         */
        public Builder backendId(String backendId) {
            return backendId(Output.of(backendId));
        }

        /**
         * @param lbId The load-balancer ID this backend is attached to.
         * 
         * @return builder
         * 
         */
        public Builder lbId(@Nullable Output<String> lbId) {
            $.lbId = lbId;
            return this;
        }

        /**
         * @param lbId The load-balancer ID this backend is attached to.
         * 
         * @return builder
         * 
         */
        public Builder lbId(String lbId) {
            return lbId(Output.of(lbId));
        }

        /**
         * @param name The name of the backend.
         * - When using the `name` you should specify the `lb-id`
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the backend.
         * - When using the `name` you should specify the `lb-id`
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        public GetLbBackendArgs build() {
            return $;
        }
    }

}