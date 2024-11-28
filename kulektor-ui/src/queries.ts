import { graphql } from "./gql";

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
