import { graphql } from "./gql";

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
        variant
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
      variant
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
