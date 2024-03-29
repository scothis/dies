//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021-2022 the original author or authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by diegen. DO NOT EDIT.

package v1

import (
	testing "dies.dev/testing"
	testingx "testing"
)

func TestWebhookClientConfigDie_MissingMethods(t *testingx.T) {
	die := WebhookClientConfigBlank
	ignore := []string{}
	diff := testing.DieFieldDiff(die).Delete(ignore...)
	if diff.Len() != 0 {
		t.Errorf("found missing fields for WebhookClientConfigDie: %s", diff.List())
	}
}

func TestServiceReferenceDie_MissingMethods(t *testingx.T) {
	die := ServiceReferenceBlank
	ignore := []string{}
	diff := testing.DieFieldDiff(die).Delete(ignore...)
	if diff.Len() != 0 {
		t.Errorf("found missing fields for ServiceReferenceDie: %s", diff.List())
	}
}

func TestRuleWithOperationsDie_MissingMethods(t *testingx.T) {
	die := RuleWithOperationsBlank
	ignore := []string{}
	diff := testing.DieFieldDiff(die).Delete(ignore...)
	if diff.Len() != 0 {
		t.Errorf("found missing fields for RuleWithOperationsDie: %s", diff.List())
	}
}

func TestRuleDie_MissingMethods(t *testingx.T) {
	die := RuleBlank
	ignore := []string{}
	diff := testing.DieFieldDiff(die).Delete(ignore...)
	if diff.Len() != 0 {
		t.Errorf("found missing fields for RuleDie: %s", diff.List())
	}
}

func TestMatchConditionDie_MissingMethods(t *testingx.T) {
	die := MatchConditionBlank
	ignore := []string{}
	diff := testing.DieFieldDiff(die).Delete(ignore...)
	if diff.Len() != 0 {
		t.Errorf("found missing fields for MatchConditionDie: %s", diff.List())
	}
}

func TestMutatingWebhookConfigurationDie_MissingMethods(t *testingx.T) {
	die := MutatingWebhookConfigurationBlank
	ignore := []string{"TypeMeta", "ObjectMeta"}
	diff := testing.DieFieldDiff(die).Delete(ignore...)
	if diff.Len() != 0 {
		t.Errorf("found missing fields for MutatingWebhookConfigurationDie: %s", diff.List())
	}
}

func TestMutatingWebhookDie_MissingMethods(t *testingx.T) {
	die := MutatingWebhookBlank
	ignore := []string{}
	diff := testing.DieFieldDiff(die).Delete(ignore...)
	if diff.Len() != 0 {
		t.Errorf("found missing fields for MutatingWebhookDie: %s", diff.List())
	}
}

func TestValidatingWebhookConfigurationDie_MissingMethods(t *testingx.T) {
	die := ValidatingWebhookConfigurationBlank
	ignore := []string{"TypeMeta", "ObjectMeta"}
	diff := testing.DieFieldDiff(die).Delete(ignore...)
	if diff.Len() != 0 {
		t.Errorf("found missing fields for ValidatingWebhookConfigurationDie: %s", diff.List())
	}
}

func TestValidatingWebhookDie_MissingMethods(t *testingx.T) {
	die := ValidatingWebhookBlank
	ignore := []string{}
	diff := testing.DieFieldDiff(die).Delete(ignore...)
	if diff.Len() != 0 {
		t.Errorf("found missing fields for ValidatingWebhookDie: %s", diff.List())
	}
}
