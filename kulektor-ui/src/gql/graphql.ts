/* eslint-disable */
import { TypedDocumentNode as DocumentNode } from '@graphql-typed-document-node/core';
export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
export type MakeEmpty<T extends { [key: string]: unknown }, K extends keyof T> = { [_ in K]?: never };
export type Incremental<T> = T | { [P in keyof T]?: P extends ' $fragmentName' | '__typename' ? T[P] : never };
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: { input: string; output: string; }
  String: { input: string; output: string; }
  Boolean: { input: boolean; output: boolean; }
  Int: { input: number; output: number; }
  Float: { input: number; output: number; }
  Date: { input: any; output: any; }
  KSUID: { input: any; output: any; }
};

export type Collection = {
  __typename?: 'Collection';
  createdAt: Scalars['Date']['output'];
  description?: Maybe<Scalars['String']['output']>;
  id: Scalars['KSUID']['output'];
  title: Scalars['String']['output'];
  type?: Maybe<Scalars['String']['output']>;
  updatedAt: Scalars['Date']['output'];
  variant: CollectionVariant;
};

export enum CollectionField {
  Description = 'description',
  Title = 'title',
  Type = 'type',
  Variant = 'variant'
}

export type CollectionInput = {
  description?: InputMaybe<Scalars['String']['input']>;
  title: Scalars['String']['input'];
  type?: InputMaybe<Scalars['String']['input']>;
  variant?: InputMaybe<CollectionVariant>;
};

export enum CollectionVariant {
  Blueprint = 'BLUEPRINT',
  Normal = 'NORMAL'
}

export type CollectionsListInput = {
  paging: PagingInput;
};

export type CollectionsListResp = {
  __typename?: 'CollectionsListResp';
  items: Array<Collection>;
  meta: Meta;
};

export type CreateCollectionResp = {
  __typename?: 'CreateCollectionResp';
  data?: Maybe<Collection>;
  success: Scalars['Boolean']['output'];
};

export type DeleteMyCollectionResp = {
  __typename?: 'DeleteMyCollectionResp';
  success: Scalars['Boolean']['output'];
};

export type Meta = {
  __typename?: 'Meta';
  currentPage: Paging;
  nextPage?: Maybe<Paging>;
  prevPage?: Maybe<Paging>;
  total: Scalars['Int']['output'];
};

export type Mutation = {
  __typename?: 'Mutation';
  createMyCollection: CreateCollectionResp;
  deleteMyCollection: DeleteMyCollectionResp;
  updateMyCollection: UpdateCollectionResp;
};


export type MutationCreateMyCollectionArgs = {
  input: CollectionInput;
};


export type MutationDeleteMyCollectionArgs = {
  collectionId: Scalars['KSUID']['input'];
};


export type MutationUpdateMyCollectionArgs = {
  input: UpdateCollectionInput;
};

export type Paging = {
  __typename?: 'Paging';
  limit?: Maybe<Scalars['Int']['output']>;
  offset?: Maybe<Scalars['Int']['output']>;
};

export type PagingInput = {
  limit?: InputMaybe<Scalars['Int']['input']>;
  offset?: InputMaybe<Scalars['Int']['input']>;
};

export type Query = {
  __typename?: 'Query';
  collectionsList: CollectionsListResp;
  myCollectionDetail: Collection;
  myCollectionsList: CollectionsListResp;
};


export type QueryCollectionsListArgs = {
  input: CollectionsListInput;
};


export type QueryMyCollectionDetailArgs = {
  collectionID: Scalars['KSUID']['input'];
};


export type QueryMyCollectionsListArgs = {
  input: CollectionsListInput;
};

export enum Role {
  Admin = 'ADMIN',
  Collector = 'COLLECTOR',
  Public = 'PUBLIC'
}

export type UpdateCollectionInput = {
  collection: CollectionInput;
  fieldsToUpdate: Array<CollectionField>;
  id: Scalars['KSUID']['input'];
};

export type UpdateCollectionResp = {
  __typename?: 'UpdateCollectionResp';
  data?: Maybe<Collection>;
  success: Scalars['Boolean']['output'];
};

export type User = {
  __typename?: 'User';
  email: Scalars['String']['output'];
  role: Role;
  uid: Scalars['String']['output'];
};

export type DeleteMyCollectionMutationVariables = Exact<{
  input: Scalars['KSUID']['input'];
}>;


export type DeleteMyCollectionMutation = { __typename?: 'Mutation', deleteMyCollection: { __typename?: 'DeleteMyCollectionResp', success: boolean } };

export type CreateMyCollectionMutationVariables = Exact<{
  input: CollectionInput;
}>;


export type CreateMyCollectionMutation = { __typename?: 'Mutation', createMyCollection: { __typename?: 'CreateCollectionResp', success: boolean, data?: { __typename?: 'Collection', id: any, title: string, description?: string | null, type?: string | null, variant: CollectionVariant, createdAt: any, updatedAt: any } | null } };

export type MyCollectionQueryVariables = Exact<{
  input: Scalars['KSUID']['input'];
}>;


export type MyCollectionQuery = { __typename?: 'Query', myCollectionDetail: { __typename?: 'Collection', id: any, title: string, description?: string | null, type?: string | null, variant: CollectionVariant, createdAt: any } };

export type MyCollectionsQueryVariables = Exact<{
  input: CollectionsListInput;
}>;


export type MyCollectionsQuery = { __typename?: 'Query', myCollectionsList: { __typename?: 'CollectionsListResp', items: Array<{ __typename?: 'Collection', id: any, title: string, description?: string | null }>, meta: { __typename?: 'Meta', total: number, nextPage?: { __typename?: 'Paging', limit?: number | null, offset?: number | null } | null, prevPage?: { __typename?: 'Paging', limit?: number | null, offset?: number | null } | null } } };


export const DeleteMyCollectionDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"DeleteMyCollection"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"KSUID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"deleteMyCollection"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"collectionId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"success"}}]}}]}}]} as unknown as DocumentNode<DeleteMyCollectionMutation, DeleteMyCollectionMutationVariables>;
export const CreateMyCollectionDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"CreateMyCollection"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"CollectionInput"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"createMyCollection"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"success"}},{"kind":"Field","name":{"kind":"Name","value":"data"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"title"}},{"kind":"Field","name":{"kind":"Name","value":"description"}},{"kind":"Field","name":{"kind":"Name","value":"type"}},{"kind":"Field","name":{"kind":"Name","value":"variant"}},{"kind":"Field","name":{"kind":"Name","value":"createdAt"}},{"kind":"Field","name":{"kind":"Name","value":"updatedAt"}}]}}]}}]}}]} as unknown as DocumentNode<CreateMyCollectionMutation, CreateMyCollectionMutationVariables>;
export const MyCollectionDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"MyCollection"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"KSUID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"myCollectionDetail"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"collectionID"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"title"}},{"kind":"Field","name":{"kind":"Name","value":"description"}},{"kind":"Field","name":{"kind":"Name","value":"type"}},{"kind":"Field","name":{"kind":"Name","value":"variant"}},{"kind":"Field","name":{"kind":"Name","value":"createdAt"}}]}}]}}]} as unknown as DocumentNode<MyCollectionQuery, MyCollectionQueryVariables>;
export const MyCollectionsDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"MyCollections"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"CollectionsListInput"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"myCollectionsList"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"title"}},{"kind":"Field","name":{"kind":"Name","value":"description"}}]}},{"kind":"Field","name":{"kind":"Name","value":"meta"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"total"}},{"kind":"Field","name":{"kind":"Name","value":"nextPage"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"limit"}},{"kind":"Field","name":{"kind":"Name","value":"offset"}}]}},{"kind":"Field","name":{"kind":"Name","value":"prevPage"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"limit"}},{"kind":"Field","name":{"kind":"Name","value":"offset"}}]}}]}}]}}]}}]} as unknown as DocumentNode<MyCollectionsQuery, MyCollectionsQueryVariables>;