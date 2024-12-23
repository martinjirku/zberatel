/* eslint-disable */
import * as types from './graphql';
import { TypedDocumentNode as DocumentNode } from '@graphql-typed-document-node/core';

/**
 * Map of all GraphQL operations in the project.
 *
 * This map has several performance disadvantages:
 * 1. It is not tree-shakeable, so it will include all operations in the project.
 * 2. It is not minifiable, so the string of a GraphQL query will be multiple times inside the bundle.
 * 3. It does not support dead code elimination, so it will add unused operations.
 *
 * Therefore it is highly recommended to use the babel or swc plugin for production.
 * Learn more about it here: https://the-guild.dev/graphql/codegen/plugins/presets/preset-client#reducing-bundle-size
 */
const documents = {
    "\n  query BlueprintsList($input: BlueprintsListInput!) {\n    blueprintsList(input: $input) {\n      items {\n        id\n        title\n        description\n        updatedAt\n        createdAt\n      }\n      meta {\n        total\n        prevPage {\n          limit\n          offset\n        }\n        nextPage {\n          limit\n          offset\n        }\n      }\n    }\n  }\n": types.BlueprintsListDocument,
    "\n  mutation DeleteMyCollection($input: KSUID!) {\n    deleteMyCollection(collectionId: $input) {\n      success\n    }\n  }\n": types.DeleteMyCollectionDocument,
    "\n  mutation CreateMyCollection($input: CollectionInput!) {\n    createMyCollection(input: $input) {\n      success\n      data {\n        id\n        title\n        description\n        type\n        createdAt\n        updatedAt\n      }\n    }\n  }\n": types.CreateMyCollectionDocument,
    "\n  mutation UpdateMyCollection($input: UpdateCollectionInput!) {\n    updateMyCollection(input: $input) {\n      success\n      data {\n        id\n        title\n        description\n        type\n        createdAt\n        updatedAt\n      }\n    }\n  }\n": types.UpdateMyCollectionDocument,
    "\n  query MyCollection($input: KSUID!) {\n    myCollectionDetail(collectionID: $input) {\n      id\n      title\n      description\n      type\n      createdAt\n    }\n  }\n": types.MyCollectionDocument,
    "\n  query MyCollections($input: CollectionsListInput!) {\n    myCollectionsList(input: $input) {\n      items {\n        id\n        title\n        description\n      }\n      meta {\n        total\n        nextPage {\n          limit\n          offset\n        }\n        prevPage {\n          limit\n          offset\n        }\n      }\n    }\n  }\n": types.MyCollectionsDocument,
};

/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 *
 *
 * @example
 * ```ts
 * const query = graphql(`query GetUser($id: ID!) { user(id: $id) { name } }`);
 * ```
 *
 * The query argument is unknown!
 * Please regenerate the types.
 */
export function graphql(source: string): unknown;

/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  query BlueprintsList($input: BlueprintsListInput!) {\n    blueprintsList(input: $input) {\n      items {\n        id\n        title\n        description\n        updatedAt\n        createdAt\n      }\n      meta {\n        total\n        prevPage {\n          limit\n          offset\n        }\n        nextPage {\n          limit\n          offset\n        }\n      }\n    }\n  }\n"): (typeof documents)["\n  query BlueprintsList($input: BlueprintsListInput!) {\n    blueprintsList(input: $input) {\n      items {\n        id\n        title\n        description\n        updatedAt\n        createdAt\n      }\n      meta {\n        total\n        prevPage {\n          limit\n          offset\n        }\n        nextPage {\n          limit\n          offset\n        }\n      }\n    }\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  mutation DeleteMyCollection($input: KSUID!) {\n    deleteMyCollection(collectionId: $input) {\n      success\n    }\n  }\n"): (typeof documents)["\n  mutation DeleteMyCollection($input: KSUID!) {\n    deleteMyCollection(collectionId: $input) {\n      success\n    }\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  mutation CreateMyCollection($input: CollectionInput!) {\n    createMyCollection(input: $input) {\n      success\n      data {\n        id\n        title\n        description\n        type\n        createdAt\n        updatedAt\n      }\n    }\n  }\n"): (typeof documents)["\n  mutation CreateMyCollection($input: CollectionInput!) {\n    createMyCollection(input: $input) {\n      success\n      data {\n        id\n        title\n        description\n        type\n        createdAt\n        updatedAt\n      }\n    }\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  mutation UpdateMyCollection($input: UpdateCollectionInput!) {\n    updateMyCollection(input: $input) {\n      success\n      data {\n        id\n        title\n        description\n        type\n        createdAt\n        updatedAt\n      }\n    }\n  }\n"): (typeof documents)["\n  mutation UpdateMyCollection($input: UpdateCollectionInput!) {\n    updateMyCollection(input: $input) {\n      success\n      data {\n        id\n        title\n        description\n        type\n        createdAt\n        updatedAt\n      }\n    }\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  query MyCollection($input: KSUID!) {\n    myCollectionDetail(collectionID: $input) {\n      id\n      title\n      description\n      type\n      createdAt\n    }\n  }\n"): (typeof documents)["\n  query MyCollection($input: KSUID!) {\n    myCollectionDetail(collectionID: $input) {\n      id\n      title\n      description\n      type\n      createdAt\n    }\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  query MyCollections($input: CollectionsListInput!) {\n    myCollectionsList(input: $input) {\n      items {\n        id\n        title\n        description\n      }\n      meta {\n        total\n        nextPage {\n          limit\n          offset\n        }\n        prevPage {\n          limit\n          offset\n        }\n      }\n    }\n  }\n"): (typeof documents)["\n  query MyCollections($input: CollectionsListInput!) {\n    myCollectionsList(input: $input) {\n      items {\n        id\n        title\n        description\n      }\n      meta {\n        total\n        nextPage {\n          limit\n          offset\n        }\n        prevPage {\n          limit\n          offset\n        }\n      }\n    }\n  }\n"];

export function graphql(source: string) {
  return (documents as any)[source] ?? {};
}

export type DocumentType<TDocumentNode extends DocumentNode<any, any>> = TDocumentNode extends DocumentNode<  infer TType,  any>  ? TType  : never;