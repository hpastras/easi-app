/* tslint:disable */
/* eslint-disable */
// @generated
// This file was automatically generated and should not be edited.

// ====================================================
// GraphQL query operation: GetAccessibilityRequest
// ====================================================

export interface GetAccessibilityRequest_accessibilityRequest {
  __typename: "AccessibilityRequest";
  id: UUID;
  name: string;
  submittedAt: Time;
}

export interface GetAccessibilityRequest {
  accessibilityRequest: GetAccessibilityRequest_accessibilityRequest | null;
}

export interface GetAccessibilityRequestVariables {
  id: UUID;
}
