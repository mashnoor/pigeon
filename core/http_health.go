package core

import (
	"github.com/mashnoor/pigeon/structures"

	"sync"
)

/***

{
  "query": {
    "bool": {
      "must": [],
      "filter": [
        {
          "multi_match": {
            "type": "best_fields",
            "query": "Payment completed successfully",
            "lenient": true
          }
        }
      ],
      "should": [],
      "must_not": []
    }
  }
}
*/

func checkHTTPHealth(service *structures.Service, wg *sync.WaitGroup) {

	//wg.Done()

}
