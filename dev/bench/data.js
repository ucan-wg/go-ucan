window.BENCHMARK_DATA = {
  "lastUpdate": 1728491991974,
  "repoUrl": "https://github.com/ucan-wg/go-ucan",
  "entries": {
    "Go Benchmark": [
      {
        "commit": {
          "author": {
            "email": "batolettre@gmail.com",
            "name": "Michael Muré",
            "username": "MichaelMure"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "50ea43e3fa317d1e3b202c5e7ca176b35e19b924",
          "message": "Merge pull request #33 from ucan-wg/continuous-bench\n\ngha: add a continuous benchmark action",
          "timestamp": "2024-10-01T17:40:50+02:00",
          "tree_id": "1e61b975dacff32071db7329167226eda6aeaa00",
          "url": "https://github.com/ucan-wg/go-ucan/commit/50ea43e3fa317d1e3b202c5e7ca176b35e19b924"
        },
        "date": 1727797732051,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkGlob",
            "value": 42.37,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "27301147 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - ns/op",
            "value": 42.37,
            "unit": "ns/op",
            "extra": "27301147 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "27301147 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "27301147 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad",
            "value": 15176,
            "unit": "ns/op\t   14569 B/op\t     162 allocs/op",
            "extra": "76120 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - ns/op",
            "value": 15176,
            "unit": "ns/op",
            "extra": "76120 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - B/op",
            "value": 14569,
            "unit": "B/op",
            "extra": "76120 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - allocs/op",
            "value": 162,
            "unit": "allocs/op",
            "extra": "76120 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson",
            "value": 111582,
            "unit": "ns/op\t   15358 B/op\t     374 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - ns/op",
            "value": 111582,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - B/op",
            "value": 15358,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - allocs/op",
            "value": 374,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal",
            "value": 54064,
            "unit": "ns/op\t    9616 B/op\t     210 allocs/op",
            "extra": "22204 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - ns/op",
            "value": 54064,
            "unit": "ns/op",
            "extra": "22204 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - B/op",
            "value": 9616,
            "unit": "B/op",
            "extra": "22204 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - allocs/op",
            "value": 210,
            "unit": "allocs/op",
            "extra": "22204 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal",
            "value": 99378,
            "unit": "ns/op\t   14413 B/op\t     305 allocs/op",
            "extra": "12066 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - ns/op",
            "value": 99378,
            "unit": "ns/op",
            "extra": "12066 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - B/op",
            "value": 14413,
            "unit": "B/op",
            "extra": "12066 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - allocs/op",
            "value": 305,
            "unit": "allocs/op",
            "extra": "12066 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson",
            "value": 62659,
            "unit": "ns/op\t   15128 B/op\t     250 allocs/op",
            "extra": "19202 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - ns/op",
            "value": 62659,
            "unit": "ns/op",
            "extra": "19202 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - B/op",
            "value": 15128,
            "unit": "B/op",
            "extra": "19202 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - allocs/op",
            "value": 250,
            "unit": "allocs/op",
            "extra": "19202 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader",
            "value": 111723,
            "unit": "ns/op\t   15309 B/op\t     373 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - ns/op",
            "value": 111723,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - B/op",
            "value": 15309,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - allocs/op",
            "value": 373,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter",
            "value": 55260,
            "unit": "ns/op\t    9136 B/op\t     241 allocs/op",
            "extra": "21765 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - ns/op",
            "value": 55260,
            "unit": "ns/op",
            "extra": "21765 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - B/op",
            "value": 9136,
            "unit": "B/op",
            "extra": "21765 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - allocs/op",
            "value": 241,
            "unit": "allocs/op",
            "extra": "21765 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader",
            "value": 100876,
            "unit": "ns/op\t   14413 B/op\t     305 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - ns/op",
            "value": 100876,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - B/op",
            "value": 14413,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - allocs/op",
            "value": 305,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader",
            "value": 58579,
            "unit": "ns/op\t   13096 B/op\t     244 allocs/op",
            "extra": "19824 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - ns/op",
            "value": 58579,
            "unit": "ns/op",
            "extra": "19824 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - B/op",
            "value": 13096,
            "unit": "B/op",
            "extra": "19824 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - allocs/op",
            "value": 244,
            "unit": "allocs/op",
            "extra": "19824 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "batolettre@gmail.com",
            "name": "Michael Muré",
            "username": "MichaelMure"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "6b8fbcee0adc4a289da32668616a3f49bad526a1",
          "message": "Merge pull request #34 from ucan-wg/invocation-stub\n\ntoken: add invocation partial stub",
          "timestamp": "2024-10-02T11:02:32+02:00",
          "tree_id": "f11c633892fe52609f43e85a407013ded5a0ce7d",
          "url": "https://github.com/ucan-wg/go-ucan/commit/6b8fbcee0adc4a289da32668616a3f49bad526a1"
        },
        "date": 1727859793240,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkGlob",
            "value": 41.8,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "28909456 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - ns/op",
            "value": 41.8,
            "unit": "ns/op",
            "extra": "28909456 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "28909456 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "28909456 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad",
            "value": 15132,
            "unit": "ns/op\t   14569 B/op\t     162 allocs/op",
            "extra": "78162 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - ns/op",
            "value": 15132,
            "unit": "ns/op",
            "extra": "78162 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - B/op",
            "value": 14569,
            "unit": "B/op",
            "extra": "78162 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - allocs/op",
            "value": 162,
            "unit": "allocs/op",
            "extra": "78162 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson",
            "value": 109556,
            "unit": "ns/op\t   15358 B/op\t     374 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - ns/op",
            "value": 109556,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - B/op",
            "value": 15358,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - allocs/op",
            "value": 374,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal",
            "value": 53087,
            "unit": "ns/op\t    9616 B/op\t     210 allocs/op",
            "extra": "22490 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - ns/op",
            "value": 53087,
            "unit": "ns/op",
            "extra": "22490 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - B/op",
            "value": 9616,
            "unit": "B/op",
            "extra": "22490 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - allocs/op",
            "value": 210,
            "unit": "allocs/op",
            "extra": "22490 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal",
            "value": 97627,
            "unit": "ns/op\t   14414 B/op\t     305 allocs/op",
            "extra": "12273 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - ns/op",
            "value": 97627,
            "unit": "ns/op",
            "extra": "12273 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - B/op",
            "value": 14414,
            "unit": "B/op",
            "extra": "12273 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - allocs/op",
            "value": 305,
            "unit": "allocs/op",
            "extra": "12273 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson",
            "value": 58314,
            "unit": "ns/op\t   15128 B/op\t     250 allocs/op",
            "extra": "20526 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - ns/op",
            "value": 58314,
            "unit": "ns/op",
            "extra": "20526 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - B/op",
            "value": 15128,
            "unit": "B/op",
            "extra": "20526 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - allocs/op",
            "value": 250,
            "unit": "allocs/op",
            "extra": "20526 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader",
            "value": 108815,
            "unit": "ns/op\t   15310 B/op\t     373 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - ns/op",
            "value": 108815,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - B/op",
            "value": 15310,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - allocs/op",
            "value": 373,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter",
            "value": 55681,
            "unit": "ns/op\t    9136 B/op\t     241 allocs/op",
            "extra": "21964 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - ns/op",
            "value": 55681,
            "unit": "ns/op",
            "extra": "21964 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - B/op",
            "value": 9136,
            "unit": "B/op",
            "extra": "21964 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - allocs/op",
            "value": 241,
            "unit": "allocs/op",
            "extra": "21964 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader",
            "value": 97985,
            "unit": "ns/op\t   14413 B/op\t     305 allocs/op",
            "extra": "12271 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - ns/op",
            "value": 97985,
            "unit": "ns/op",
            "extra": "12271 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - B/op",
            "value": 14413,
            "unit": "B/op",
            "extra": "12271 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - allocs/op",
            "value": 305,
            "unit": "allocs/op",
            "extra": "12271 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader",
            "value": 57119,
            "unit": "ns/op\t   13096 B/op\t     244 allocs/op",
            "extra": "20989 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - ns/op",
            "value": 57119,
            "unit": "ns/op",
            "extra": "20989 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - B/op",
            "value": 13096,
            "unit": "B/op",
            "extra": "20989 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - allocs/op",
            "value": 244,
            "unit": "allocs/op",
            "extra": "20989 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "batolettre@gmail.com",
            "name": "Michael Muré",
            "username": "MichaelMure"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "8615f6c72bc7d2d1fbcf9e084721332ce557d56e",
          "message": "Merge pull request #35 from ucan-wg/cid-proposal\n\ntoken: don't store the CID in the token, symmetric API for sealed",
          "timestamp": "2024-10-02T13:42:06+02:00",
          "tree_id": "f9db21048fde99c92acf6eafeda506da75819967",
          "url": "https://github.com/ucan-wg/go-ucan/commit/8615f6c72bc7d2d1fbcf9e084721332ce557d56e"
        },
        "date": 1727869361557,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkGlob",
            "value": 42.22,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "28424530 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - ns/op",
            "value": 42.22,
            "unit": "ns/op",
            "extra": "28424530 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "28424530 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "28424530 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad",
            "value": 15063,
            "unit": "ns/op\t   14569 B/op\t     162 allocs/op",
            "extra": "78519 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - ns/op",
            "value": 15063,
            "unit": "ns/op",
            "extra": "78519 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - B/op",
            "value": 14569,
            "unit": "B/op",
            "extra": "78519 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - allocs/op",
            "value": 162,
            "unit": "allocs/op",
            "extra": "78519 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson",
            "value": 108585,
            "unit": "ns/op\t   15342 B/op\t     374 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - ns/op",
            "value": 108585,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - B/op",
            "value": 15342,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - allocs/op",
            "value": 374,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal",
            "value": 53039,
            "unit": "ns/op\t    9616 B/op\t     210 allocs/op",
            "extra": "22627 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - ns/op",
            "value": 53039,
            "unit": "ns/op",
            "extra": "22627 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - B/op",
            "value": 9616,
            "unit": "B/op",
            "extra": "22627 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - allocs/op",
            "value": 210,
            "unit": "allocs/op",
            "extra": "22627 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal",
            "value": 97280,
            "unit": "ns/op\t   14397 B/op\t     305 allocs/op",
            "extra": "12331 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - ns/op",
            "value": 97280,
            "unit": "ns/op",
            "extra": "12331 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - B/op",
            "value": 14397,
            "unit": "B/op",
            "extra": "12331 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - allocs/op",
            "value": 305,
            "unit": "allocs/op",
            "extra": "12331 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson",
            "value": 59753,
            "unit": "ns/op\t   15128 B/op\t     250 allocs/op",
            "extra": "20534 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - ns/op",
            "value": 59753,
            "unit": "ns/op",
            "extra": "20534 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - B/op",
            "value": 15128,
            "unit": "B/op",
            "extra": "20534 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - allocs/op",
            "value": 250,
            "unit": "allocs/op",
            "extra": "20534 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader",
            "value": 108276,
            "unit": "ns/op\t   15293 B/op\t     373 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - ns/op",
            "value": 108276,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - B/op",
            "value": 15293,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - allocs/op",
            "value": 373,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter",
            "value": 54751,
            "unit": "ns/op\t    9136 B/op\t     241 allocs/op",
            "extra": "21912 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - ns/op",
            "value": 54751,
            "unit": "ns/op",
            "extra": "21912 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - B/op",
            "value": 9136,
            "unit": "B/op",
            "extra": "21912 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - allocs/op",
            "value": 241,
            "unit": "allocs/op",
            "extra": "21912 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader",
            "value": 97988,
            "unit": "ns/op\t   14397 B/op\t     305 allocs/op",
            "extra": "12265 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - ns/op",
            "value": 97988,
            "unit": "ns/op",
            "extra": "12265 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - B/op",
            "value": 14397,
            "unit": "B/op",
            "extra": "12265 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - allocs/op",
            "value": 305,
            "unit": "allocs/op",
            "extra": "12265 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader",
            "value": 58597,
            "unit": "ns/op\t   13096 B/op\t     244 allocs/op",
            "extra": "20824 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - ns/op",
            "value": 58597,
            "unit": "ns/op",
            "extra": "20824 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - B/op",
            "value": 13096,
            "unit": "B/op",
            "extra": "20824 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - allocs/op",
            "value": 244,
            "unit": "allocs/op",
            "extra": "20824 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "batolettre@gmail.com",
            "name": "Michael Muré",
            "username": "MichaelMure"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "d60fb71156bcd1c3347323d03627ba9e509ed0c7",
          "message": "Merge pull request #22 from ucan-wg/container\n\nadd a token container with serialization as CARv1 file",
          "timestamp": "2024-10-08T11:40:00+02:00",
          "tree_id": "dcaedada721ca4a53fc4ec289909644281438720",
          "url": "https://github.com/ucan-wg/go-ucan/commit/d60fb71156bcd1c3347323d03627ba9e509ed0c7"
        },
        "date": 1728380464442,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkContainerSerialisation/car_write",
            "value": 5952,
            "unit": "ns/op\t   17928 B/op\t      59 allocs/op",
            "extra": "208119 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - ns/op",
            "value": 5952,
            "unit": "ns/op",
            "extra": "208119 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - B/op",
            "value": 17928,
            "unit": "B/op",
            "extra": "208119 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - allocs/op",
            "value": 59,
            "unit": "allocs/op",
            "extra": "208119 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read",
            "value": 977339,
            "unit": "ns/op\t  149321 B/op\t    2631 allocs/op",
            "extra": "1194 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - ns/op",
            "value": 977339,
            "unit": "ns/op",
            "extra": "1194 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - B/op",
            "value": 149321,
            "unit": "B/op",
            "extra": "1194 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - allocs/op",
            "value": 2631,
            "unit": "allocs/op",
            "extra": "1194 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write",
            "value": 12011,
            "unit": "ns/op\t   24232 B/op\t      61 allocs/op",
            "extra": "104894 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - ns/op",
            "value": 12011,
            "unit": "ns/op",
            "extra": "104894 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - B/op",
            "value": 24232,
            "unit": "B/op",
            "extra": "104894 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - allocs/op",
            "value": 61,
            "unit": "allocs/op",
            "extra": "104894 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read",
            "value": 1004959,
            "unit": "ns/op\t  151381 B/op\t    2633 allocs/op",
            "extra": "1196 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - ns/op",
            "value": 1004959,
            "unit": "ns/op",
            "extra": "1196 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - B/op",
            "value": 151381,
            "unit": "B/op",
            "extra": "1196 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - allocs/op",
            "value": 2633,
            "unit": "allocs/op",
            "extra": "1196 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write",
            "value": 4847,
            "unit": "ns/op\t   16368 B/op\t      25 allocs/op",
            "extra": "261568 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - ns/op",
            "value": 4847,
            "unit": "ns/op",
            "extra": "261568 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - B/op",
            "value": 16368,
            "unit": "B/op",
            "extra": "261568 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - allocs/op",
            "value": 25,
            "unit": "allocs/op",
            "extra": "261568 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read",
            "value": 947913,
            "unit": "ns/op\t  138775 B/op\t    2534 allocs/op",
            "extra": "1220 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - ns/op",
            "value": 947913,
            "unit": "ns/op",
            "extra": "1220 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - B/op",
            "value": 138775,
            "unit": "B/op",
            "extra": "1220 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - allocs/op",
            "value": 2534,
            "unit": "allocs/op",
            "extra": "1220 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write",
            "value": 8938,
            "unit": "ns/op\t   20720 B/op\t      26 allocs/op",
            "extra": "134644 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - ns/op",
            "value": 8938,
            "unit": "ns/op",
            "extra": "134644 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - B/op",
            "value": 20720,
            "unit": "B/op",
            "extra": "134644 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "134644 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read",
            "value": 966182,
            "unit": "ns/op\t  140846 B/op\t    2536 allocs/op",
            "extra": "1230 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - ns/op",
            "value": 966182,
            "unit": "ns/op",
            "extra": "1230 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - B/op",
            "value": 140846,
            "unit": "B/op",
            "extra": "1230 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - allocs/op",
            "value": 2536,
            "unit": "allocs/op",
            "extra": "1230 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob",
            "value": 42.09,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "26533195 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - ns/op",
            "value": 42.09,
            "unit": "ns/op",
            "extra": "26533195 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "26533195 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "26533195 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad",
            "value": 15020,
            "unit": "ns/op\t   14568 B/op\t     162 allocs/op",
            "extra": "79266 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - ns/op",
            "value": 15020,
            "unit": "ns/op",
            "extra": "79266 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - B/op",
            "value": 14568,
            "unit": "B/op",
            "extra": "79266 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - allocs/op",
            "value": 162,
            "unit": "allocs/op",
            "extra": "79266 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson",
            "value": 111332,
            "unit": "ns/op\t   15342 B/op\t     374 allocs/op",
            "extra": "9328 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - ns/op",
            "value": 111332,
            "unit": "ns/op",
            "extra": "9328 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - B/op",
            "value": 15342,
            "unit": "B/op",
            "extra": "9328 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - allocs/op",
            "value": 374,
            "unit": "allocs/op",
            "extra": "9328 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal",
            "value": 53708,
            "unit": "ns/op\t    9616 B/op\t     210 allocs/op",
            "extra": "22641 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - ns/op",
            "value": 53708,
            "unit": "ns/op",
            "extra": "22641 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - B/op",
            "value": 9616,
            "unit": "B/op",
            "extra": "22641 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - allocs/op",
            "value": 210,
            "unit": "allocs/op",
            "extra": "22641 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal",
            "value": 97648,
            "unit": "ns/op\t   14397 B/op\t     305 allocs/op",
            "extra": "12223 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - ns/op",
            "value": 97648,
            "unit": "ns/op",
            "extra": "12223 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - B/op",
            "value": 14397,
            "unit": "B/op",
            "extra": "12223 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - allocs/op",
            "value": 305,
            "unit": "allocs/op",
            "extra": "12223 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson",
            "value": 59582,
            "unit": "ns/op\t   15128 B/op\t     250 allocs/op",
            "extra": "20095 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - ns/op",
            "value": 59582,
            "unit": "ns/op",
            "extra": "20095 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - B/op",
            "value": 15128,
            "unit": "B/op",
            "extra": "20095 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - allocs/op",
            "value": 250,
            "unit": "allocs/op",
            "extra": "20095 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader",
            "value": 109148,
            "unit": "ns/op\t   15294 B/op\t     373 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - ns/op",
            "value": 109148,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - B/op",
            "value": 15294,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - allocs/op",
            "value": 373,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter",
            "value": 54425,
            "unit": "ns/op\t    9136 B/op\t     241 allocs/op",
            "extra": "21979 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - ns/op",
            "value": 54425,
            "unit": "ns/op",
            "extra": "21979 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - B/op",
            "value": 9136,
            "unit": "B/op",
            "extra": "21979 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - allocs/op",
            "value": 241,
            "unit": "allocs/op",
            "extra": "21979 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader",
            "value": 98852,
            "unit": "ns/op\t   14397 B/op\t     305 allocs/op",
            "extra": "12224 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - ns/op",
            "value": 98852,
            "unit": "ns/op",
            "extra": "12224 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - B/op",
            "value": 14397,
            "unit": "B/op",
            "extra": "12224 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - allocs/op",
            "value": 305,
            "unit": "allocs/op",
            "extra": "12224 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader",
            "value": 59612,
            "unit": "ns/op\t   13096 B/op\t     244 allocs/op",
            "extra": "20325 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - ns/op",
            "value": 59612,
            "unit": "ns/op",
            "extra": "20325 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - B/op",
            "value": 13096,
            "unit": "B/op",
            "extra": "20325 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - allocs/op",
            "value": 244,
            "unit": "allocs/op",
            "extra": "20325 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "batolettre@gmail.com",
            "name": "Michael Muré",
            "username": "MichaelMure"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "2a51d61b46fd3195fd3d8dc48d829f1bb0a09df9",
          "message": "Merge pull request #39 from ucan-wg/command-as-string\n\ncommand: make the type a string, for easier equality test",
          "timestamp": "2024-10-09T18:30:45+02:00",
          "tree_id": "073d1b9347bdb661e7a47a926e727f167f7af4ba",
          "url": "https://github.com/ucan-wg/go-ucan/commit/2a51d61b46fd3195fd3d8dc48d829f1bb0a09df9"
        },
        "date": 1728491495173,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkContainerSerialisation/car_write",
            "value": 5803,
            "unit": "ns/op\t   17928 B/op\t      59 allocs/op",
            "extra": "208149 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - ns/op",
            "value": 5803,
            "unit": "ns/op",
            "extra": "208149 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - B/op",
            "value": 17928,
            "unit": "B/op",
            "extra": "208149 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - allocs/op",
            "value": 59,
            "unit": "allocs/op",
            "extra": "208149 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read",
            "value": 989770,
            "unit": "ns/op\t  148843 B/op\t    2621 allocs/op",
            "extra": "1185 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - ns/op",
            "value": 989770,
            "unit": "ns/op",
            "extra": "1185 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - B/op",
            "value": 148843,
            "unit": "B/op",
            "extra": "1185 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - allocs/op",
            "value": 2621,
            "unit": "allocs/op",
            "extra": "1185 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write",
            "value": 11415,
            "unit": "ns/op\t   24232 B/op\t      61 allocs/op",
            "extra": "100453 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - ns/op",
            "value": 11415,
            "unit": "ns/op",
            "extra": "100453 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - B/op",
            "value": 24232,
            "unit": "B/op",
            "extra": "100453 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - allocs/op",
            "value": 61,
            "unit": "allocs/op",
            "extra": "100453 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read",
            "value": 1001742,
            "unit": "ns/op\t  150909 B/op\t    2623 allocs/op",
            "extra": "1191 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - ns/op",
            "value": 1001742,
            "unit": "ns/op",
            "extra": "1191 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - B/op",
            "value": 150909,
            "unit": "B/op",
            "extra": "1191 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - allocs/op",
            "value": 2623,
            "unit": "allocs/op",
            "extra": "1191 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write",
            "value": 4421,
            "unit": "ns/op\t   16368 B/op\t      25 allocs/op",
            "extra": "263764 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - ns/op",
            "value": 4421,
            "unit": "ns/op",
            "extra": "263764 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - B/op",
            "value": 16368,
            "unit": "B/op",
            "extra": "263764 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - allocs/op",
            "value": 25,
            "unit": "allocs/op",
            "extra": "263764 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read",
            "value": 970388,
            "unit": "ns/op\t  138303 B/op\t    2524 allocs/op",
            "extra": "1218 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - ns/op",
            "value": 970388,
            "unit": "ns/op",
            "extra": "1218 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - B/op",
            "value": 138303,
            "unit": "B/op",
            "extra": "1218 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - allocs/op",
            "value": 2524,
            "unit": "allocs/op",
            "extra": "1218 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write",
            "value": 8848,
            "unit": "ns/op\t   20720 B/op\t      26 allocs/op",
            "extra": "133581 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - ns/op",
            "value": 8848,
            "unit": "ns/op",
            "extra": "133581 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - B/op",
            "value": 20720,
            "unit": "B/op",
            "extra": "133581 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "133581 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read",
            "value": 985494,
            "unit": "ns/op\t  140367 B/op\t    2526 allocs/op",
            "extra": "1208 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - ns/op",
            "value": 985494,
            "unit": "ns/op",
            "extra": "1208 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - B/op",
            "value": 140367,
            "unit": "B/op",
            "extra": "1208 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - allocs/op",
            "value": 2526,
            "unit": "allocs/op",
            "extra": "1208 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob",
            "value": 42.15,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "29015739 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - ns/op",
            "value": 42.15,
            "unit": "ns/op",
            "extra": "29015739 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "29015739 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "29015739 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad",
            "value": 15715,
            "unit": "ns/op\t   14569 B/op\t     162 allocs/op",
            "extra": "77061 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - ns/op",
            "value": 15715,
            "unit": "ns/op",
            "extra": "77061 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - B/op",
            "value": 14569,
            "unit": "B/op",
            "extra": "77061 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - allocs/op",
            "value": 162,
            "unit": "allocs/op",
            "extra": "77061 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson",
            "value": 111724,
            "unit": "ns/op\t   15294 B/op\t     373 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - ns/op",
            "value": 111724,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - B/op",
            "value": 15294,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - allocs/op",
            "value": 373,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal",
            "value": 54235,
            "unit": "ns/op\t    9600 B/op\t     208 allocs/op",
            "extra": "22125 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - ns/op",
            "value": 54235,
            "unit": "ns/op",
            "extra": "22125 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - B/op",
            "value": 9600,
            "unit": "B/op",
            "extra": "22125 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - allocs/op",
            "value": 208,
            "unit": "allocs/op",
            "extra": "22125 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal",
            "value": 100569,
            "unit": "ns/op\t   14350 B/op\t     304 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - ns/op",
            "value": 100569,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - B/op",
            "value": 14350,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - allocs/op",
            "value": 304,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson",
            "value": 60128,
            "unit": "ns/op\t   15112 B/op\t     248 allocs/op",
            "extra": "20041 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - ns/op",
            "value": 60128,
            "unit": "ns/op",
            "extra": "20041 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - B/op",
            "value": 15112,
            "unit": "B/op",
            "extra": "20041 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - allocs/op",
            "value": 248,
            "unit": "allocs/op",
            "extra": "20041 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader",
            "value": 111112,
            "unit": "ns/op\t   15246 B/op\t     372 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - ns/op",
            "value": 111112,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - B/op",
            "value": 15246,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - allocs/op",
            "value": 372,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter",
            "value": 58007,
            "unit": "ns/op\t    9120 B/op\t     239 allocs/op",
            "extra": "21585 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - ns/op",
            "value": 58007,
            "unit": "ns/op",
            "extra": "21585 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - B/op",
            "value": 9120,
            "unit": "B/op",
            "extra": "21585 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - allocs/op",
            "value": 239,
            "unit": "allocs/op",
            "extra": "21585 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader",
            "value": 104613,
            "unit": "ns/op\t   14350 B/op\t     304 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - ns/op",
            "value": 104613,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - B/op",
            "value": 14350,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - allocs/op",
            "value": 304,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader",
            "value": 59921,
            "unit": "ns/op\t   13080 B/op\t     242 allocs/op",
            "extra": "20181 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - ns/op",
            "value": 59921,
            "unit": "ns/op",
            "extra": "20181 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - B/op",
            "value": 13080,
            "unit": "B/op",
            "extra": "20181 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - allocs/op",
            "value": 242,
            "unit": "allocs/op",
            "extra": "20181 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "michael.mure@consensys.net",
            "name": "Michael Muré",
            "username": "MichaelMure"
          },
          "committer": {
            "email": "michael.mure@consensys.net",
            "name": "Michael Muré",
            "username": "MichaelMure"
          },
          "distinct": true,
          "id": "100a510097e600ee542e1ca2c3162ef136cb9bdc",
          "message": "pkg/container: harden the CAR file round-trip with fuzzing",
          "timestamp": "2024-10-09T18:38:35+02:00",
          "tree_id": "de75bb5656a96302c7670c3f0b3661dbc1cf4c9f",
          "url": "https://github.com/ucan-wg/go-ucan/commit/100a510097e600ee542e1ca2c3162ef136cb9bdc"
        },
        "date": 1728491991117,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkContainerSerialisation/car_write",
            "value": 5923,
            "unit": "ns/op\t   17896 B/op\t      58 allocs/op",
            "extra": "206194 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - ns/op",
            "value": 5923,
            "unit": "ns/op",
            "extra": "206194 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - B/op",
            "value": 17896,
            "unit": "B/op",
            "extra": "206194 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "206194 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read",
            "value": 949584,
            "unit": "ns/op\t  148841 B/op\t    2621 allocs/op",
            "extra": "1231 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - ns/op",
            "value": 949584,
            "unit": "ns/op",
            "extra": "1231 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - B/op",
            "value": 148841,
            "unit": "B/op",
            "extra": "1231 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - allocs/op",
            "value": 2621,
            "unit": "allocs/op",
            "extra": "1231 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write",
            "value": 10962,
            "unit": "ns/op\t   24200 B/op\t      60 allocs/op",
            "extra": "109051 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - ns/op",
            "value": 10962,
            "unit": "ns/op",
            "extra": "109051 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - B/op",
            "value": 24200,
            "unit": "B/op",
            "extra": "109051 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - allocs/op",
            "value": 60,
            "unit": "allocs/op",
            "extra": "109051 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read",
            "value": 960665,
            "unit": "ns/op\t  150909 B/op\t    2623 allocs/op",
            "extra": "1226 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - ns/op",
            "value": 960665,
            "unit": "ns/op",
            "extra": "1226 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - B/op",
            "value": 150909,
            "unit": "B/op",
            "extra": "1226 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - allocs/op",
            "value": 2623,
            "unit": "allocs/op",
            "extra": "1226 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write",
            "value": 4376,
            "unit": "ns/op\t   16368 B/op\t      25 allocs/op",
            "extra": "273243 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - ns/op",
            "value": 4376,
            "unit": "ns/op",
            "extra": "273243 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - B/op",
            "value": 16368,
            "unit": "B/op",
            "extra": "273243 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - allocs/op",
            "value": 25,
            "unit": "allocs/op",
            "extra": "273243 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read",
            "value": 932288,
            "unit": "ns/op\t  138301 B/op\t    2524 allocs/op",
            "extra": "1267 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - ns/op",
            "value": 932288,
            "unit": "ns/op",
            "extra": "1267 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - B/op",
            "value": 138301,
            "unit": "B/op",
            "extra": "1267 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - allocs/op",
            "value": 2524,
            "unit": "allocs/op",
            "extra": "1267 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write",
            "value": 8582,
            "unit": "ns/op\t   20720 B/op\t      26 allocs/op",
            "extra": "139054 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - ns/op",
            "value": 8582,
            "unit": "ns/op",
            "extra": "139054 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - B/op",
            "value": 20720,
            "unit": "B/op",
            "extra": "139054 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "139054 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read",
            "value": 947744,
            "unit": "ns/op\t  140364 B/op\t    2526 allocs/op",
            "extra": "1250 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - ns/op",
            "value": 947744,
            "unit": "ns/op",
            "extra": "1250 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - B/op",
            "value": 140364,
            "unit": "B/op",
            "extra": "1250 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - allocs/op",
            "value": 2526,
            "unit": "allocs/op",
            "extra": "1250 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob",
            "value": 41.83,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "28611658 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - ns/op",
            "value": 41.83,
            "unit": "ns/op",
            "extra": "28611658 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "28611658 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "28611658 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad",
            "value": 14850,
            "unit": "ns/op\t   14569 B/op\t     162 allocs/op",
            "extra": "80373 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - ns/op",
            "value": 14850,
            "unit": "ns/op",
            "extra": "80373 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - B/op",
            "value": 14569,
            "unit": "B/op",
            "extra": "80373 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - allocs/op",
            "value": 162,
            "unit": "allocs/op",
            "extra": "80373 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson",
            "value": 109006,
            "unit": "ns/op\t   15294 B/op\t     373 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - ns/op",
            "value": 109006,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - B/op",
            "value": 15294,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - allocs/op",
            "value": 373,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal",
            "value": 52049,
            "unit": "ns/op\t    9600 B/op\t     208 allocs/op",
            "extra": "22779 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - ns/op",
            "value": 52049,
            "unit": "ns/op",
            "extra": "22779 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - B/op",
            "value": 9600,
            "unit": "B/op",
            "extra": "22779 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - allocs/op",
            "value": 208,
            "unit": "allocs/op",
            "extra": "22779 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal",
            "value": 96023,
            "unit": "ns/op\t   14349 B/op\t     304 allocs/op",
            "extra": "12540 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - ns/op",
            "value": 96023,
            "unit": "ns/op",
            "extra": "12540 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - B/op",
            "value": 14349,
            "unit": "B/op",
            "extra": "12540 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - allocs/op",
            "value": 304,
            "unit": "allocs/op",
            "extra": "12540 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson",
            "value": 57831,
            "unit": "ns/op\t   15112 B/op\t     248 allocs/op",
            "extra": "20840 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - ns/op",
            "value": 57831,
            "unit": "ns/op",
            "extra": "20840 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - B/op",
            "value": 15112,
            "unit": "B/op",
            "extra": "20840 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - allocs/op",
            "value": 248,
            "unit": "allocs/op",
            "extra": "20840 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader",
            "value": 107731,
            "unit": "ns/op\t   15245 B/op\t     372 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - ns/op",
            "value": 107731,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - B/op",
            "value": 15245,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - allocs/op",
            "value": 372,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter",
            "value": 56559,
            "unit": "ns/op\t    9120 B/op\t     239 allocs/op",
            "extra": "20695 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - ns/op",
            "value": 56559,
            "unit": "ns/op",
            "extra": "20695 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - B/op",
            "value": 9120,
            "unit": "B/op",
            "extra": "20695 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - allocs/op",
            "value": 239,
            "unit": "allocs/op",
            "extra": "20695 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader",
            "value": 96374,
            "unit": "ns/op\t   14349 B/op\t     304 allocs/op",
            "extra": "12436 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - ns/op",
            "value": 96374,
            "unit": "ns/op",
            "extra": "12436 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - B/op",
            "value": 14349,
            "unit": "B/op",
            "extra": "12436 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - allocs/op",
            "value": 304,
            "unit": "allocs/op",
            "extra": "12436 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader",
            "value": 56745,
            "unit": "ns/op\t   13080 B/op\t     242 allocs/op",
            "extra": "21183 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - ns/op",
            "value": 56745,
            "unit": "ns/op",
            "extra": "21183 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - B/op",
            "value": 13080,
            "unit": "B/op",
            "extra": "21183 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - allocs/op",
            "value": 242,
            "unit": "allocs/op",
            "extra": "21183 times\n4 procs"
          }
        ]
      }
    ]
  }
}