// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.scaleway.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class IamPolicyRuleArgs extends com.pulumi.resources.ResourceArgs {

    public static final IamPolicyRuleArgs Empty = new IamPolicyRuleArgs();

    /**
     * ID of organization scoped to the rule.
     * 
     */
    @Import(name="organizationId")
    private @Nullable Output<String> organizationId;

    /**
     * @return ID of organization scoped to the rule.
     * 
     */
    public Optional<Output<String>> organizationId() {
        return Optional.ofNullable(this.organizationId);
    }

    /**
     * Names of permission sets bound to the rule.
     * 
     */
    @Import(name="permissionSetNames", required=true)
    private Output<List<String>> permissionSetNames;

    /**
     * @return Names of permission sets bound to the rule.
     * 
     */
    public Output<List<String>> permissionSetNames() {
        return this.permissionSetNames;
    }

    /**
     * List of project IDs scoped to the rule.
     * 
     */
    @Import(name="projectIds")
    private @Nullable Output<List<String>> projectIds;

    /**
     * @return List of project IDs scoped to the rule.
     * 
     */
    public Optional<Output<List<String>>> projectIds() {
        return Optional.ofNullable(this.projectIds);
    }

    private IamPolicyRuleArgs() {}

    private IamPolicyRuleArgs(IamPolicyRuleArgs $) {
        this.organizationId = $.organizationId;
        this.permissionSetNames = $.permissionSetNames;
        this.projectIds = $.projectIds;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(IamPolicyRuleArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private IamPolicyRuleArgs $;

        public Builder() {
            $ = new IamPolicyRuleArgs();
        }

        public Builder(IamPolicyRuleArgs defaults) {
            $ = new IamPolicyRuleArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param organizationId ID of organization scoped to the rule.
         * 
         * @return builder
         * 
         */
        public Builder organizationId(@Nullable Output<String> organizationId) {
            $.organizationId = organizationId;
            return this;
        }

        /**
         * @param organizationId ID of organization scoped to the rule.
         * 
         * @return builder
         * 
         */
        public Builder organizationId(String organizationId) {
            return organizationId(Output.of(organizationId));
        }

        /**
         * @param permissionSetNames Names of permission sets bound to the rule.
         * 
         * @return builder
         * 
         */
        public Builder permissionSetNames(Output<List<String>> permissionSetNames) {
            $.permissionSetNames = permissionSetNames;
            return this;
        }

        /**
         * @param permissionSetNames Names of permission sets bound to the rule.
         * 
         * @return builder
         * 
         */
        public Builder permissionSetNames(List<String> permissionSetNames) {
            return permissionSetNames(Output.of(permissionSetNames));
        }

        /**
         * @param permissionSetNames Names of permission sets bound to the rule.
         * 
         * @return builder
         * 
         */
        public Builder permissionSetNames(String... permissionSetNames) {
            return permissionSetNames(List.of(permissionSetNames));
        }

        /**
         * @param projectIds List of project IDs scoped to the rule.
         * 
         * @return builder
         * 
         */
        public Builder projectIds(@Nullable Output<List<String>> projectIds) {
            $.projectIds = projectIds;
            return this;
        }

        /**
         * @param projectIds List of project IDs scoped to the rule.
         * 
         * @return builder
         * 
         */
        public Builder projectIds(List<String> projectIds) {
            return projectIds(Output.of(projectIds));
        }

        /**
         * @param projectIds List of project IDs scoped to the rule.
         * 
         * @return builder
         * 
         */
        public Builder projectIds(String... projectIds) {
            return projectIds(List.of(projectIds));
        }

        public IamPolicyRuleArgs build() {
            $.permissionSetNames = Objects.requireNonNull($.permissionSetNames, "expected parameter 'permissionSetNames' to be non-null");
            return $;
        }
    }

}