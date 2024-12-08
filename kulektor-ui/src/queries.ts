import { graphql } from "./gql";

export const LIST_BLUEPRINT = graphql(`
  query BlueprintsList($input: BlueprintsListInput!) {
    blueprintsList(input: $input) {
      items {
        id
        title
        description
        updatedAt
        createdAt
      }
      meta {
        total
        prevPage {
          limit
          offset
        }
        nextPage {
          limit
          offset
        }
      }
    }
  }
`);

export const DELETE_MY_COLLECTION = graphql(`
  mutation DeleteMyCollection($input: KSUID!) {
    deleteMyCollection(collectionId: $input) {
      success
    }
  }
`);

export const MY_NEW_COLLECTION = graphql(`
  mutation CreateMyCollection($input: CollectionInput!) {
    createMyCollection(input: $input) {
      success
      data {
        id
        title
        description
        type
        createdAt
        updatedAt
      }
    }
  }
`);

export const UPDATE_MY_COLLECTION = graphql(`
  mutation UpdateMyCollection($input: UpdateCollectionInput!) {
    updateMyCollection(input: $input) {
      success
      data {
        id
        title
        description
        type
        createdAt
        updatedAt
      }
    }
  }
`);

export const MY_COLLECTIONS_DETAIL = graphql(`
  query MyCollection($input: KSUID!) {
    myCollectionDetail(collectionID: $input) {
      id
      title
      description
      type
      createdAt
    }
  }
`);

export const MY_COLLECTIONS = graphql(`
  query MyCollections($input: CollectionsListInput!) {
    myCollectionsList(input: $input) {
      items {
        id
        title
        description
      }
      meta {
        total
        nextPage {
          limit
          offset
        }
        prevPage {
          limit
          offset
        }
      }
    }
  }
`);
