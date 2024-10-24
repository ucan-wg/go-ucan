window.BENCHMARK_DATA = {
  "lastUpdate": 1729807875909,
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
          "id": "9051e5250bd50bcceb7dc13a898e2b518018dc84",
          "message": "Merge pull request #40 from ucan-wg/dlg-example\n\ndelegation: make the examples more examply, less testy",
          "timestamp": "2024-10-14T12:52:52+02:00",
          "tree_id": "8220a4f978b9aa6ca748991ec77c2f6092ff47e0",
          "url": "https://github.com/ucan-wg/go-ucan/commit/9051e5250bd50bcceb7dc13a898e2b518018dc84"
        },
        "date": 1728903219800,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkContainerSerialisation/car_write",
            "value": 5582,
            "unit": "ns/op\t   17896 B/op\t      58 allocs/op",
            "extra": "192968 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - ns/op",
            "value": 5582,
            "unit": "ns/op",
            "extra": "192968 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - B/op",
            "value": 17896,
            "unit": "B/op",
            "extra": "192968 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "192968 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read",
            "value": 952866,
            "unit": "ns/op\t  148843 B/op\t    2621 allocs/op",
            "extra": "1238 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - ns/op",
            "value": 952866,
            "unit": "ns/op",
            "extra": "1238 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - B/op",
            "value": 148843,
            "unit": "B/op",
            "extra": "1238 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - allocs/op",
            "value": 2621,
            "unit": "allocs/op",
            "extra": "1238 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write",
            "value": 11007,
            "unit": "ns/op\t   24200 B/op\t      60 allocs/op",
            "extra": "108049 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - ns/op",
            "value": 11007,
            "unit": "ns/op",
            "extra": "108049 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - B/op",
            "value": 24200,
            "unit": "B/op",
            "extra": "108049 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - allocs/op",
            "value": 60,
            "unit": "allocs/op",
            "extra": "108049 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read",
            "value": 993321,
            "unit": "ns/op\t  150909 B/op\t    2623 allocs/op",
            "extra": "1203 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - ns/op",
            "value": 993321,
            "unit": "ns/op",
            "extra": "1203 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - B/op",
            "value": 150909,
            "unit": "B/op",
            "extra": "1203 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - allocs/op",
            "value": 2623,
            "unit": "allocs/op",
            "extra": "1203 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write",
            "value": 4383,
            "unit": "ns/op\t   16368 B/op\t      25 allocs/op",
            "extra": "265404 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - ns/op",
            "value": 4383,
            "unit": "ns/op",
            "extra": "265404 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - B/op",
            "value": 16368,
            "unit": "B/op",
            "extra": "265404 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - allocs/op",
            "value": 25,
            "unit": "allocs/op",
            "extra": "265404 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read",
            "value": 937988,
            "unit": "ns/op\t  138296 B/op\t    2524 allocs/op",
            "extra": "1260 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - ns/op",
            "value": 937988,
            "unit": "ns/op",
            "extra": "1260 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - B/op",
            "value": 138296,
            "unit": "B/op",
            "extra": "1260 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - allocs/op",
            "value": 2524,
            "unit": "allocs/op",
            "extra": "1260 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write",
            "value": 8621,
            "unit": "ns/op\t   20720 B/op\t      26 allocs/op",
            "extra": "137844 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - ns/op",
            "value": 8621,
            "unit": "ns/op",
            "extra": "137844 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - B/op",
            "value": 20720,
            "unit": "B/op",
            "extra": "137844 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "137844 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read",
            "value": 949706,
            "unit": "ns/op\t  140363 B/op\t    2526 allocs/op",
            "extra": "1242 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - ns/op",
            "value": 949706,
            "unit": "ns/op",
            "extra": "1242 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - B/op",
            "value": 140363,
            "unit": "B/op",
            "extra": "1242 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - allocs/op",
            "value": 2526,
            "unit": "allocs/op",
            "extra": "1242 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob",
            "value": 41.28,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "28536588 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - ns/op",
            "value": 41.28,
            "unit": "ns/op",
            "extra": "28536588 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "28536588 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "28536588 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad",
            "value": 14777,
            "unit": "ns/op\t   14569 B/op\t     162 allocs/op",
            "extra": "80704 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - ns/op",
            "value": 14777,
            "unit": "ns/op",
            "extra": "80704 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - B/op",
            "value": 14569,
            "unit": "B/op",
            "extra": "80704 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - allocs/op",
            "value": 162,
            "unit": "allocs/op",
            "extra": "80704 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson",
            "value": 108411,
            "unit": "ns/op\t   15294 B/op\t     373 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - ns/op",
            "value": 108411,
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
            "value": 54108,
            "unit": "ns/op\t    9600 B/op\t     208 allocs/op",
            "extra": "22686 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - ns/op",
            "value": 54108,
            "unit": "ns/op",
            "extra": "22686 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - B/op",
            "value": 9600,
            "unit": "B/op",
            "extra": "22686 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - allocs/op",
            "value": 208,
            "unit": "allocs/op",
            "extra": "22686 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal",
            "value": 96093,
            "unit": "ns/op\t   14349 B/op\t     304 allocs/op",
            "extra": "12433 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - ns/op",
            "value": 96093,
            "unit": "ns/op",
            "extra": "12433 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - B/op",
            "value": 14349,
            "unit": "B/op",
            "extra": "12433 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - allocs/op",
            "value": 304,
            "unit": "allocs/op",
            "extra": "12433 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson",
            "value": 57455,
            "unit": "ns/op\t   15112 B/op\t     248 allocs/op",
            "extra": "20838 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - ns/op",
            "value": 57455,
            "unit": "ns/op",
            "extra": "20838 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - B/op",
            "value": 15112,
            "unit": "B/op",
            "extra": "20838 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - allocs/op",
            "value": 248,
            "unit": "allocs/op",
            "extra": "20838 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader",
            "value": 107977,
            "unit": "ns/op\t   15245 B/op\t     372 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - ns/op",
            "value": 107977,
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
            "value": 53910,
            "unit": "ns/op\t    9120 B/op\t     239 allocs/op",
            "extra": "22077 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - ns/op",
            "value": 53910,
            "unit": "ns/op",
            "extra": "22077 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - B/op",
            "value": 9120,
            "unit": "B/op",
            "extra": "22077 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - allocs/op",
            "value": 239,
            "unit": "allocs/op",
            "extra": "22077 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader",
            "value": 97186,
            "unit": "ns/op\t   14349 B/op\t     304 allocs/op",
            "extra": "12362 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - ns/op",
            "value": 97186,
            "unit": "ns/op",
            "extra": "12362 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - B/op",
            "value": 14349,
            "unit": "B/op",
            "extra": "12362 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - allocs/op",
            "value": 304,
            "unit": "allocs/op",
            "extra": "12362 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader",
            "value": 59129,
            "unit": "ns/op\t   13080 B/op\t     242 allocs/op",
            "extra": "21042 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - ns/op",
            "value": 59129,
            "unit": "ns/op",
            "extra": "21042 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - B/op",
            "value": 13080,
            "unit": "B/op",
            "extra": "21042 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - allocs/op",
            "value": 242,
            "unit": "allocs/op",
            "extra": "21042 times\n4 procs"
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
          "id": "59da2d1a2c777ef19179a762331c9da5c8e1cb6b",
          "message": "Merge pull request #43 from ucan-wg/dlg-options\n\ndelegation: tune Nbf & Exp options",
          "timestamp": "2024-10-15T10:38:50+02:00",
          "tree_id": "ca4073a652fd30947906ea23822f7b6faee97f7b",
          "url": "https://github.com/ucan-wg/go-ucan/commit/59da2d1a2c777ef19179a762331c9da5c8e1cb6b"
        },
        "date": 1728981584660,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkContainerSerialisation/car_write",
            "value": 5613,
            "unit": "ns/op\t   17896 B/op\t      58 allocs/op",
            "extra": "195318 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - ns/op",
            "value": 5613,
            "unit": "ns/op",
            "extra": "195318 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - B/op",
            "value": 17896,
            "unit": "B/op",
            "extra": "195318 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "195318 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read",
            "value": 947017,
            "unit": "ns/op\t  148837 B/op\t    2621 allocs/op",
            "extra": "1224 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - ns/op",
            "value": 947017,
            "unit": "ns/op",
            "extra": "1224 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - B/op",
            "value": 148837,
            "unit": "B/op",
            "extra": "1224 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - allocs/op",
            "value": 2621,
            "unit": "allocs/op",
            "extra": "1224 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write",
            "value": 11078,
            "unit": "ns/op\t   24200 B/op\t      60 allocs/op",
            "extra": "106699 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - ns/op",
            "value": 11078,
            "unit": "ns/op",
            "extra": "106699 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - B/op",
            "value": 24200,
            "unit": "B/op",
            "extra": "106699 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - allocs/op",
            "value": 60,
            "unit": "allocs/op",
            "extra": "106699 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read",
            "value": 958854,
            "unit": "ns/op\t  150903 B/op\t    2623 allocs/op",
            "extra": "1219 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - ns/op",
            "value": 958854,
            "unit": "ns/op",
            "extra": "1219 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - B/op",
            "value": 150903,
            "unit": "B/op",
            "extra": "1219 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - allocs/op",
            "value": 2623,
            "unit": "allocs/op",
            "extra": "1219 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write",
            "value": 4440,
            "unit": "ns/op\t   16368 B/op\t      25 allocs/op",
            "extra": "255966 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - ns/op",
            "value": 4440,
            "unit": "ns/op",
            "extra": "255966 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - B/op",
            "value": 16368,
            "unit": "B/op",
            "extra": "255966 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - allocs/op",
            "value": 25,
            "unit": "allocs/op",
            "extra": "255966 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read",
            "value": 930774,
            "unit": "ns/op\t  138297 B/op\t    2524 allocs/op",
            "extra": "1262 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - ns/op",
            "value": 930774,
            "unit": "ns/op",
            "extra": "1262 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - B/op",
            "value": 138297,
            "unit": "B/op",
            "extra": "1262 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - allocs/op",
            "value": 2524,
            "unit": "allocs/op",
            "extra": "1262 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write",
            "value": 9257,
            "unit": "ns/op\t   20720 B/op\t      26 allocs/op",
            "extra": "135852 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - ns/op",
            "value": 9257,
            "unit": "ns/op",
            "extra": "135852 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - B/op",
            "value": 20720,
            "unit": "B/op",
            "extra": "135852 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "135852 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read",
            "value": 947392,
            "unit": "ns/op\t  140364 B/op\t    2526 allocs/op",
            "extra": "1252 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - ns/op",
            "value": 947392,
            "unit": "ns/op",
            "extra": "1252 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - B/op",
            "value": 140364,
            "unit": "B/op",
            "extra": "1252 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - allocs/op",
            "value": 2526,
            "unit": "allocs/op",
            "extra": "1252 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob",
            "value": 41.27,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "28767634 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - ns/op",
            "value": 41.27,
            "unit": "ns/op",
            "extra": "28767634 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "28767634 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "28767634 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad",
            "value": 14789,
            "unit": "ns/op\t   14568 B/op\t     162 allocs/op",
            "extra": "80943 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - ns/op",
            "value": 14789,
            "unit": "ns/op",
            "extra": "80943 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - B/op",
            "value": 14568,
            "unit": "B/op",
            "extra": "80943 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - allocs/op",
            "value": 162,
            "unit": "allocs/op",
            "extra": "80943 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson",
            "value": 108417,
            "unit": "ns/op\t   15293 B/op\t     373 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - ns/op",
            "value": 108417,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - B/op",
            "value": 15293,
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
            "value": 52198,
            "unit": "ns/op\t    9600 B/op\t     208 allocs/op",
            "extra": "22881 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - ns/op",
            "value": 52198,
            "unit": "ns/op",
            "extra": "22881 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - B/op",
            "value": 9600,
            "unit": "B/op",
            "extra": "22881 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - allocs/op",
            "value": 208,
            "unit": "allocs/op",
            "extra": "22881 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal",
            "value": 95978,
            "unit": "ns/op\t   14349 B/op\t     304 allocs/op",
            "extra": "12505 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - ns/op",
            "value": 95978,
            "unit": "ns/op",
            "extra": "12505 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - B/op",
            "value": 14349,
            "unit": "B/op",
            "extra": "12505 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - allocs/op",
            "value": 304,
            "unit": "allocs/op",
            "extra": "12505 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson",
            "value": 62242,
            "unit": "ns/op\t   15112 B/op\t     248 allocs/op",
            "extra": "21021 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - ns/op",
            "value": 62242,
            "unit": "ns/op",
            "extra": "21021 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - B/op",
            "value": 15112,
            "unit": "B/op",
            "extra": "21021 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - allocs/op",
            "value": 248,
            "unit": "allocs/op",
            "extra": "21021 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader",
            "value": 108363,
            "unit": "ns/op\t   15245 B/op\t     372 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - ns/op",
            "value": 108363,
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
            "value": 53682,
            "unit": "ns/op\t    9120 B/op\t     239 allocs/op",
            "extra": "22261 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - ns/op",
            "value": 53682,
            "unit": "ns/op",
            "extra": "22261 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - B/op",
            "value": 9120,
            "unit": "B/op",
            "extra": "22261 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - allocs/op",
            "value": 239,
            "unit": "allocs/op",
            "extra": "22261 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader",
            "value": 96766,
            "unit": "ns/op\t   14349 B/op\t     304 allocs/op",
            "extra": "12355 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - ns/op",
            "value": 96766,
            "unit": "ns/op",
            "extra": "12355 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - B/op",
            "value": 14349,
            "unit": "B/op",
            "extra": "12355 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - allocs/op",
            "value": 304,
            "unit": "allocs/op",
            "extra": "12355 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader",
            "value": 57477,
            "unit": "ns/op\t   13080 B/op\t     242 allocs/op",
            "extra": "21172 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - ns/op",
            "value": 57477,
            "unit": "ns/op",
            "extra": "21172 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - B/op",
            "value": 13080,
            "unit": "B/op",
            "extra": "21172 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - allocs/op",
            "value": 242,
            "unit": "allocs/op",
            "extra": "21172 times\n4 procs"
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
          "id": "aa4ad2fc10b7d8ee85d3ad2b02c8c222d363fa9f",
          "message": "Merge pull request #42 from ucan-wg/fluent-policy\n\npolicy: fluent construction",
          "timestamp": "2024-10-15T13:06:56+02:00",
          "tree_id": "211b9a7091724ac5f9b821b963feca59e552adaf",
          "url": "https://github.com/ucan-wg/go-ucan/commit/aa4ad2fc10b7d8ee85d3ad2b02c8c222d363fa9f"
        },
        "date": 1728990465642,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkContainerSerialisation/car_write",
            "value": 5888,
            "unit": "ns/op\t   17896 B/op\t      58 allocs/op",
            "extra": "216106 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - ns/op",
            "value": 5888,
            "unit": "ns/op",
            "extra": "216106 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - B/op",
            "value": 17896,
            "unit": "B/op",
            "extra": "216106 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "216106 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read",
            "value": 951101,
            "unit": "ns/op\t  148838 B/op\t    2621 allocs/op",
            "extra": "1228 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - ns/op",
            "value": 951101,
            "unit": "ns/op",
            "extra": "1228 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - B/op",
            "value": 148838,
            "unit": "B/op",
            "extra": "1228 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - allocs/op",
            "value": 2621,
            "unit": "allocs/op",
            "extra": "1228 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write",
            "value": 10978,
            "unit": "ns/op\t   24200 B/op\t      60 allocs/op",
            "extra": "108514 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - ns/op",
            "value": 10978,
            "unit": "ns/op",
            "extra": "108514 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - B/op",
            "value": 24200,
            "unit": "B/op",
            "extra": "108514 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - allocs/op",
            "value": 60,
            "unit": "allocs/op",
            "extra": "108514 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read",
            "value": 956728,
            "unit": "ns/op\t  150901 B/op\t    2623 allocs/op",
            "extra": "1233 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - ns/op",
            "value": 956728,
            "unit": "ns/op",
            "extra": "1233 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - B/op",
            "value": 150901,
            "unit": "B/op",
            "extra": "1233 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - allocs/op",
            "value": 2623,
            "unit": "allocs/op",
            "extra": "1233 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write",
            "value": 4366,
            "unit": "ns/op\t   16368 B/op\t      25 allocs/op",
            "extra": "266731 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - ns/op",
            "value": 4366,
            "unit": "ns/op",
            "extra": "266731 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - B/op",
            "value": 16368,
            "unit": "B/op",
            "extra": "266731 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - allocs/op",
            "value": 25,
            "unit": "allocs/op",
            "extra": "266731 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read",
            "value": 936983,
            "unit": "ns/op\t  138300 B/op\t    2524 allocs/op",
            "extra": "1250 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - ns/op",
            "value": 936983,
            "unit": "ns/op",
            "extra": "1250 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - B/op",
            "value": 138300,
            "unit": "B/op",
            "extra": "1250 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - allocs/op",
            "value": 2524,
            "unit": "allocs/op",
            "extra": "1250 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write",
            "value": 8634,
            "unit": "ns/op\t   20720 B/op\t      26 allocs/op",
            "extra": "136808 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - ns/op",
            "value": 8634,
            "unit": "ns/op",
            "extra": "136808 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - B/op",
            "value": 20720,
            "unit": "B/op",
            "extra": "136808 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "136808 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read",
            "value": 946874,
            "unit": "ns/op\t  140366 B/op\t    2526 allocs/op",
            "extra": "1240 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - ns/op",
            "value": 946874,
            "unit": "ns/op",
            "extra": "1240 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - B/op",
            "value": 140366,
            "unit": "B/op",
            "extra": "1240 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - allocs/op",
            "value": 2526,
            "unit": "allocs/op",
            "extra": "1240 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob",
            "value": 41.91,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "28631013 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - ns/op",
            "value": 41.91,
            "unit": "ns/op",
            "extra": "28631013 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "28631013 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "28631013 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad",
            "value": 14741,
            "unit": "ns/op\t   14569 B/op\t     162 allocs/op",
            "extra": "81099 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - ns/op",
            "value": 14741,
            "unit": "ns/op",
            "extra": "81099 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - B/op",
            "value": 14569,
            "unit": "B/op",
            "extra": "81099 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - allocs/op",
            "value": 162,
            "unit": "allocs/op",
            "extra": "81099 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson",
            "value": 107071,
            "unit": "ns/op\t   15102 B/op\t     369 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - ns/op",
            "value": 107071,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - B/op",
            "value": 15102,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - allocs/op",
            "value": 369,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal",
            "value": 51822,
            "unit": "ns/op\t    9600 B/op\t     208 allocs/op",
            "extra": "23054 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - ns/op",
            "value": 51822,
            "unit": "ns/op",
            "extra": "23054 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - B/op",
            "value": 9600,
            "unit": "B/op",
            "extra": "23054 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - allocs/op",
            "value": 208,
            "unit": "allocs/op",
            "extra": "23054 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal",
            "value": 95255,
            "unit": "ns/op\t   14157 B/op\t     300 allocs/op",
            "extra": "12613 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - ns/op",
            "value": 95255,
            "unit": "ns/op",
            "extra": "12613 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - B/op",
            "value": 14157,
            "unit": "B/op",
            "extra": "12613 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - allocs/op",
            "value": 300,
            "unit": "allocs/op",
            "extra": "12613 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson",
            "value": 57052,
            "unit": "ns/op\t   15112 B/op\t     248 allocs/op",
            "extra": "20968 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - ns/op",
            "value": 57052,
            "unit": "ns/op",
            "extra": "20968 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - B/op",
            "value": 15112,
            "unit": "B/op",
            "extra": "20968 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - allocs/op",
            "value": 248,
            "unit": "allocs/op",
            "extra": "20968 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader",
            "value": 106667,
            "unit": "ns/op\t   15053 B/op\t     368 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - ns/op",
            "value": 106667,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - B/op",
            "value": 15053,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - allocs/op",
            "value": 368,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter",
            "value": 53429,
            "unit": "ns/op\t    9120 B/op\t     239 allocs/op",
            "extra": "21702 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - ns/op",
            "value": 53429,
            "unit": "ns/op",
            "extra": "21702 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - B/op",
            "value": 9120,
            "unit": "B/op",
            "extra": "21702 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - allocs/op",
            "value": 239,
            "unit": "allocs/op",
            "extra": "21702 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader",
            "value": 100585,
            "unit": "ns/op\t   14157 B/op\t     300 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - ns/op",
            "value": 100585,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - B/op",
            "value": 14157,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - allocs/op",
            "value": 300,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader",
            "value": 56287,
            "unit": "ns/op\t   13080 B/op\t     242 allocs/op",
            "extra": "21322 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - ns/op",
            "value": 56287,
            "unit": "ns/op",
            "extra": "21322 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - B/op",
            "value": 13080,
            "unit": "B/op",
            "extra": "21322 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - allocs/op",
            "value": 242,
            "unit": "allocs/op",
            "extra": "21322 times\n4 procs"
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
          "id": "030db7ec0d512e30e4383827883c79b43f2f7695",
          "message": "invocation: fix comment",
          "timestamp": "2024-10-15T15:41:14+02:00",
          "tree_id": "149146e49e7e235dfd93b1116108b9a36b86e534",
          "url": "https://github.com/ucan-wg/go-ucan/commit/030db7ec0d512e30e4383827883c79b43f2f7695"
        },
        "date": 1728999736985,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkContainerSerialisation/car_write",
            "value": 5679,
            "unit": "ns/op\t   17896 B/op\t      58 allocs/op",
            "extra": "212125 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - ns/op",
            "value": 5679,
            "unit": "ns/op",
            "extra": "212125 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - B/op",
            "value": 17896,
            "unit": "B/op",
            "extra": "212125 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "212125 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read",
            "value": 971518,
            "unit": "ns/op\t  148839 B/op\t    2621 allocs/op",
            "extra": "1204 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - ns/op",
            "value": 971518,
            "unit": "ns/op",
            "extra": "1204 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - B/op",
            "value": 148839,
            "unit": "B/op",
            "extra": "1204 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - allocs/op",
            "value": 2621,
            "unit": "allocs/op",
            "extra": "1204 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write",
            "value": 11107,
            "unit": "ns/op\t   24200 B/op\t      60 allocs/op",
            "extra": "107204 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - ns/op",
            "value": 11107,
            "unit": "ns/op",
            "extra": "107204 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - B/op",
            "value": 24200,
            "unit": "B/op",
            "extra": "107204 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - allocs/op",
            "value": 60,
            "unit": "allocs/op",
            "extra": "107204 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read",
            "value": 989070,
            "unit": "ns/op\t  150907 B/op\t    2623 allocs/op",
            "extra": "1197 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - ns/op",
            "value": 989070,
            "unit": "ns/op",
            "extra": "1197 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - B/op",
            "value": 150907,
            "unit": "B/op",
            "extra": "1197 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - allocs/op",
            "value": 2623,
            "unit": "allocs/op",
            "extra": "1197 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write",
            "value": 4413,
            "unit": "ns/op\t   16368 B/op\t      25 allocs/op",
            "extra": "265068 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - ns/op",
            "value": 4413,
            "unit": "ns/op",
            "extra": "265068 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - B/op",
            "value": 16368,
            "unit": "B/op",
            "extra": "265068 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - allocs/op",
            "value": 25,
            "unit": "allocs/op",
            "extra": "265068 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read",
            "value": 944400,
            "unit": "ns/op\t  138296 B/op\t    2524 allocs/op",
            "extra": "1272 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - ns/op",
            "value": 944400,
            "unit": "ns/op",
            "extra": "1272 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - B/op",
            "value": 138296,
            "unit": "B/op",
            "extra": "1272 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - allocs/op",
            "value": 2524,
            "unit": "allocs/op",
            "extra": "1272 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write",
            "value": 9234,
            "unit": "ns/op\t   20720 B/op\t      26 allocs/op",
            "extra": "136378 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - ns/op",
            "value": 9234,
            "unit": "ns/op",
            "extra": "136378 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - B/op",
            "value": 20720,
            "unit": "B/op",
            "extra": "136378 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "136378 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read",
            "value": 957328,
            "unit": "ns/op\t  140365 B/op\t    2526 allocs/op",
            "extra": "1237 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - ns/op",
            "value": 957328,
            "unit": "ns/op",
            "extra": "1237 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - B/op",
            "value": 140365,
            "unit": "B/op",
            "extra": "1237 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - allocs/op",
            "value": 2526,
            "unit": "allocs/op",
            "extra": "1237 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob",
            "value": 41.55,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "28815086 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - ns/op",
            "value": 41.55,
            "unit": "ns/op",
            "extra": "28815086 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "28815086 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "28815086 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad",
            "value": 14756,
            "unit": "ns/op\t   14569 B/op\t     162 allocs/op",
            "extra": "80802 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - ns/op",
            "value": 14756,
            "unit": "ns/op",
            "extra": "80802 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - B/op",
            "value": 14569,
            "unit": "B/op",
            "extra": "80802 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - allocs/op",
            "value": 162,
            "unit": "allocs/op",
            "extra": "80802 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson",
            "value": 109203,
            "unit": "ns/op\t   15102 B/op\t     369 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - ns/op",
            "value": 109203,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - B/op",
            "value": 15102,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - allocs/op",
            "value": 369,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal",
            "value": 54092,
            "unit": "ns/op\t    9600 B/op\t     208 allocs/op",
            "extra": "22437 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - ns/op",
            "value": 54092,
            "unit": "ns/op",
            "extra": "22437 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - B/op",
            "value": 9600,
            "unit": "B/op",
            "extra": "22437 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - allocs/op",
            "value": 208,
            "unit": "allocs/op",
            "extra": "22437 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal",
            "value": 97363,
            "unit": "ns/op\t   14158 B/op\t     300 allocs/op",
            "extra": "12391 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - ns/op",
            "value": 97363,
            "unit": "ns/op",
            "extra": "12391 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - B/op",
            "value": 14158,
            "unit": "B/op",
            "extra": "12391 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - allocs/op",
            "value": 300,
            "unit": "allocs/op",
            "extra": "12391 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson",
            "value": 61286,
            "unit": "ns/op\t   15112 B/op\t     248 allocs/op",
            "extra": "20444 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - ns/op",
            "value": 61286,
            "unit": "ns/op",
            "extra": "20444 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - B/op",
            "value": 15112,
            "unit": "B/op",
            "extra": "20444 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - allocs/op",
            "value": 248,
            "unit": "allocs/op",
            "extra": "20444 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader",
            "value": 109122,
            "unit": "ns/op\t   15054 B/op\t     368 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - ns/op",
            "value": 109122,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - B/op",
            "value": 15054,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - allocs/op",
            "value": 368,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter",
            "value": 55359,
            "unit": "ns/op\t    9120 B/op\t     239 allocs/op",
            "extra": "21771 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - ns/op",
            "value": 55359,
            "unit": "ns/op",
            "extra": "21771 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - B/op",
            "value": 9120,
            "unit": "B/op",
            "extra": "21771 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - allocs/op",
            "value": 239,
            "unit": "allocs/op",
            "extra": "21771 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader",
            "value": 98148,
            "unit": "ns/op\t   14157 B/op\t     300 allocs/op",
            "extra": "12139 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - ns/op",
            "value": 98148,
            "unit": "ns/op",
            "extra": "12139 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - B/op",
            "value": 14157,
            "unit": "B/op",
            "extra": "12139 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - allocs/op",
            "value": 300,
            "unit": "allocs/op",
            "extra": "12139 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader",
            "value": 59429,
            "unit": "ns/op\t   13080 B/op\t     242 allocs/op",
            "extra": "20269 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - ns/op",
            "value": 59429,
            "unit": "ns/op",
            "extra": "20269 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - B/op",
            "value": 13080,
            "unit": "B/op",
            "extra": "20269 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - allocs/op",
            "value": 242,
            "unit": "allocs/op",
            "extra": "20269 times\n4 procs"
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
          "id": "c372f0793cb8dcf9ebc2b3fa3f2aa6dad5b27cb2",
          "message": "policy: match is now a method of Policy",
          "timestamp": "2024-10-15T16:37:24+02:00",
          "tree_id": "f72781e6d15a9a54d8ed1762acfc5db0cd60db85",
          "url": "https://github.com/ucan-wg/go-ucan/commit/c372f0793cb8dcf9ebc2b3fa3f2aa6dad5b27cb2"
        },
        "date": 1729003128006,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkContainerSerialisation/car_write",
            "value": 5688,
            "unit": "ns/op\t   17896 B/op\t      58 allocs/op",
            "extra": "212377 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - ns/op",
            "value": 5688,
            "unit": "ns/op",
            "extra": "212377 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - B/op",
            "value": 17896,
            "unit": "B/op",
            "extra": "212377 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "212377 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read",
            "value": 955089,
            "unit": "ns/op\t  148839 B/op\t    2621 allocs/op",
            "extra": "1230 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - ns/op",
            "value": 955089,
            "unit": "ns/op",
            "extra": "1230 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - B/op",
            "value": 148839,
            "unit": "B/op",
            "extra": "1230 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - allocs/op",
            "value": 2621,
            "unit": "allocs/op",
            "extra": "1230 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write",
            "value": 10993,
            "unit": "ns/op\t   24200 B/op\t      60 allocs/op",
            "extra": "107526 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - ns/op",
            "value": 10993,
            "unit": "ns/op",
            "extra": "107526 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - B/op",
            "value": 24200,
            "unit": "B/op",
            "extra": "107526 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - allocs/op",
            "value": 60,
            "unit": "allocs/op",
            "extra": "107526 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read",
            "value": 982875,
            "unit": "ns/op\t  150906 B/op\t    2623 allocs/op",
            "extra": "1192 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - ns/op",
            "value": 982875,
            "unit": "ns/op",
            "extra": "1192 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - B/op",
            "value": 150906,
            "unit": "B/op",
            "extra": "1192 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - allocs/op",
            "value": 2623,
            "unit": "allocs/op",
            "extra": "1192 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write",
            "value": 4498,
            "unit": "ns/op\t   16368 B/op\t      25 allocs/op",
            "extra": "226275 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - ns/op",
            "value": 4498,
            "unit": "ns/op",
            "extra": "226275 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - B/op",
            "value": 16368,
            "unit": "B/op",
            "extra": "226275 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - allocs/op",
            "value": 25,
            "unit": "allocs/op",
            "extra": "226275 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read",
            "value": 937746,
            "unit": "ns/op\t  138297 B/op\t    2524 allocs/op",
            "extra": "1252 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - ns/op",
            "value": 937746,
            "unit": "ns/op",
            "extra": "1252 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - B/op",
            "value": 138297,
            "unit": "B/op",
            "extra": "1252 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - allocs/op",
            "value": 2524,
            "unit": "allocs/op",
            "extra": "1252 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write",
            "value": 8675,
            "unit": "ns/op\t   20720 B/op\t      26 allocs/op",
            "extra": "137601 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - ns/op",
            "value": 8675,
            "unit": "ns/op",
            "extra": "137601 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - B/op",
            "value": 20720,
            "unit": "B/op",
            "extra": "137601 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "137601 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read",
            "value": 949442,
            "unit": "ns/op\t  140361 B/op\t    2526 allocs/op",
            "extra": "1236 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - ns/op",
            "value": 949442,
            "unit": "ns/op",
            "extra": "1236 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - B/op",
            "value": 140361,
            "unit": "B/op",
            "extra": "1236 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - allocs/op",
            "value": 2526,
            "unit": "allocs/op",
            "extra": "1236 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad",
            "value": 14722,
            "unit": "ns/op\t   14569 B/op\t     162 allocs/op",
            "extra": "81415 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - ns/op",
            "value": 14722,
            "unit": "ns/op",
            "extra": "81415 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - B/op",
            "value": 14569,
            "unit": "B/op",
            "extra": "81415 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - allocs/op",
            "value": 162,
            "unit": "allocs/op",
            "extra": "81415 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson",
            "value": 107593,
            "unit": "ns/op\t   15102 B/op\t     369 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - ns/op",
            "value": 107593,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - B/op",
            "value": 15102,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - allocs/op",
            "value": 369,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal",
            "value": 52069,
            "unit": "ns/op\t    9600 B/op\t     208 allocs/op",
            "extra": "22951 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - ns/op",
            "value": 52069,
            "unit": "ns/op",
            "extra": "22951 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - B/op",
            "value": 9600,
            "unit": "B/op",
            "extra": "22951 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - allocs/op",
            "value": 208,
            "unit": "allocs/op",
            "extra": "22951 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal",
            "value": 98689,
            "unit": "ns/op\t   14158 B/op\t     300 allocs/op",
            "extra": "12169 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - ns/op",
            "value": 98689,
            "unit": "ns/op",
            "extra": "12169 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - B/op",
            "value": 14158,
            "unit": "B/op",
            "extra": "12169 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - allocs/op",
            "value": 300,
            "unit": "allocs/op",
            "extra": "12169 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson",
            "value": 57357,
            "unit": "ns/op\t   15112 B/op\t     248 allocs/op",
            "extra": "20938 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - ns/op",
            "value": 57357,
            "unit": "ns/op",
            "extra": "20938 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - B/op",
            "value": 15112,
            "unit": "B/op",
            "extra": "20938 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - allocs/op",
            "value": 248,
            "unit": "allocs/op",
            "extra": "20938 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader",
            "value": 109376,
            "unit": "ns/op\t   15053 B/op\t     368 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - ns/op",
            "value": 109376,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - B/op",
            "value": 15053,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - allocs/op",
            "value": 368,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter",
            "value": 53617,
            "unit": "ns/op\t    9120 B/op\t     239 allocs/op",
            "extra": "22281 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - ns/op",
            "value": 53617,
            "unit": "ns/op",
            "extra": "22281 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - B/op",
            "value": 9120,
            "unit": "B/op",
            "extra": "22281 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - allocs/op",
            "value": 239,
            "unit": "allocs/op",
            "extra": "22281 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader",
            "value": 96071,
            "unit": "ns/op\t   14157 B/op\t     300 allocs/op",
            "extra": "12462 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - ns/op",
            "value": 96071,
            "unit": "ns/op",
            "extra": "12462 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - B/op",
            "value": 14157,
            "unit": "B/op",
            "extra": "12462 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - allocs/op",
            "value": 300,
            "unit": "allocs/op",
            "extra": "12462 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader",
            "value": 59489,
            "unit": "ns/op\t   13080 B/op\t     242 allocs/op",
            "extra": "21157 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - ns/op",
            "value": 59489,
            "unit": "ns/op",
            "extra": "21157 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - B/op",
            "value": 13080,
            "unit": "B/op",
            "extra": "21157 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - allocs/op",
            "value": 242,
            "unit": "allocs/op",
            "extra": "21157 times\n4 procs"
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
          "id": "2ad3aeb6da7f545626e353b25a9d79dea44a4b24",
          "message": "policy: match is now a method of Policy",
          "timestamp": "2024-10-15T16:53:06+02:00",
          "tree_id": "6c78a5b7b1a6756a695778e9d1cb6c967672c5cf",
          "url": "https://github.com/ucan-wg/go-ucan/commit/2ad3aeb6da7f545626e353b25a9d79dea44a4b24"
        },
        "date": 1729004044599,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkContainerSerialisation/car_write",
            "value": 5731,
            "unit": "ns/op\t   17896 B/op\t      58 allocs/op",
            "extra": "209125 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - ns/op",
            "value": 5731,
            "unit": "ns/op",
            "extra": "209125 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - B/op",
            "value": 17896,
            "unit": "B/op",
            "extra": "209125 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "209125 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read",
            "value": 975900,
            "unit": "ns/op\t  148842 B/op\t    2621 allocs/op",
            "extra": "1206 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - ns/op",
            "value": 975900,
            "unit": "ns/op",
            "extra": "1206 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - B/op",
            "value": 148842,
            "unit": "B/op",
            "extra": "1206 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - allocs/op",
            "value": 2621,
            "unit": "allocs/op",
            "extra": "1206 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write",
            "value": 11080,
            "unit": "ns/op\t   24200 B/op\t      60 allocs/op",
            "extra": "107016 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - ns/op",
            "value": 11080,
            "unit": "ns/op",
            "extra": "107016 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - B/op",
            "value": 24200,
            "unit": "B/op",
            "extra": "107016 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - allocs/op",
            "value": 60,
            "unit": "allocs/op",
            "extra": "107016 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read",
            "value": 976309,
            "unit": "ns/op\t  150903 B/op\t    2623 allocs/op",
            "extra": "1222 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - ns/op",
            "value": 976309,
            "unit": "ns/op",
            "extra": "1222 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - B/op",
            "value": 150903,
            "unit": "B/op",
            "extra": "1222 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - allocs/op",
            "value": 2623,
            "unit": "allocs/op",
            "extra": "1222 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write",
            "value": 4454,
            "unit": "ns/op\t   16368 B/op\t      25 allocs/op",
            "extra": "263181 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - ns/op",
            "value": 4454,
            "unit": "ns/op",
            "extra": "263181 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - B/op",
            "value": 16368,
            "unit": "B/op",
            "extra": "263181 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - allocs/op",
            "value": 25,
            "unit": "allocs/op",
            "extra": "263181 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read",
            "value": 944615,
            "unit": "ns/op\t  138300 B/op\t    2524 allocs/op",
            "extra": "1255 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - ns/op",
            "value": 944615,
            "unit": "ns/op",
            "extra": "1255 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - B/op",
            "value": 138300,
            "unit": "B/op",
            "extra": "1255 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - allocs/op",
            "value": 2524,
            "unit": "allocs/op",
            "extra": "1255 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write",
            "value": 8726,
            "unit": "ns/op\t   20720 B/op\t      26 allocs/op",
            "extra": "137562 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - ns/op",
            "value": 8726,
            "unit": "ns/op",
            "extra": "137562 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - B/op",
            "value": 20720,
            "unit": "B/op",
            "extra": "137562 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "137562 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read",
            "value": 979660,
            "unit": "ns/op\t  140369 B/op\t    2526 allocs/op",
            "extra": "1250 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - ns/op",
            "value": 979660,
            "unit": "ns/op",
            "extra": "1250 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - B/op",
            "value": 140369,
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
            "value": 42.07,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "28863916 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - ns/op",
            "value": 42.07,
            "unit": "ns/op",
            "extra": "28863916 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "28863916 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "28863916 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad",
            "value": 15088,
            "unit": "ns/op\t   14569 B/op\t     162 allocs/op",
            "extra": "80158 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - ns/op",
            "value": 15088,
            "unit": "ns/op",
            "extra": "80158 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - B/op",
            "value": 14569,
            "unit": "B/op",
            "extra": "80158 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - allocs/op",
            "value": 162,
            "unit": "allocs/op",
            "extra": "80158 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson",
            "value": 109992,
            "unit": "ns/op\t   15102 B/op\t     369 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - ns/op",
            "value": 109992,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - B/op",
            "value": 15102,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - allocs/op",
            "value": 369,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal",
            "value": 53913,
            "unit": "ns/op\t    9600 B/op\t     208 allocs/op",
            "extra": "21808 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - ns/op",
            "value": 53913,
            "unit": "ns/op",
            "extra": "21808 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - B/op",
            "value": 9600,
            "unit": "B/op",
            "extra": "21808 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - allocs/op",
            "value": 208,
            "unit": "allocs/op",
            "extra": "21808 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal",
            "value": 98474,
            "unit": "ns/op\t   14157 B/op\t     300 allocs/op",
            "extra": "12150 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - ns/op",
            "value": 98474,
            "unit": "ns/op",
            "extra": "12150 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - B/op",
            "value": 14157,
            "unit": "B/op",
            "extra": "12150 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - allocs/op",
            "value": 300,
            "unit": "allocs/op",
            "extra": "12150 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson",
            "value": 61232,
            "unit": "ns/op\t   15112 B/op\t     248 allocs/op",
            "extra": "20166 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - ns/op",
            "value": 61232,
            "unit": "ns/op",
            "extra": "20166 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - B/op",
            "value": 15112,
            "unit": "B/op",
            "extra": "20166 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - allocs/op",
            "value": 248,
            "unit": "allocs/op",
            "extra": "20166 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader",
            "value": 109719,
            "unit": "ns/op\t   15053 B/op\t     368 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - ns/op",
            "value": 109719,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - B/op",
            "value": 15053,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - allocs/op",
            "value": 368,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter",
            "value": 54266,
            "unit": "ns/op\t    9120 B/op\t     239 allocs/op",
            "extra": "20744 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - ns/op",
            "value": 54266,
            "unit": "ns/op",
            "extra": "20744 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - B/op",
            "value": 9120,
            "unit": "B/op",
            "extra": "20744 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - allocs/op",
            "value": 239,
            "unit": "allocs/op",
            "extra": "20744 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader",
            "value": 97451,
            "unit": "ns/op\t   14157 B/op\t     300 allocs/op",
            "extra": "12289 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - ns/op",
            "value": 97451,
            "unit": "ns/op",
            "extra": "12289 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - B/op",
            "value": 14157,
            "unit": "B/op",
            "extra": "12289 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - allocs/op",
            "value": 300,
            "unit": "allocs/op",
            "extra": "12289 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader",
            "value": 57107,
            "unit": "ns/op\t   13080 B/op\t     242 allocs/op",
            "extra": "20964 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - ns/op",
            "value": 57107,
            "unit": "ns/op",
            "extra": "20964 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - B/op",
            "value": 13080,
            "unit": "B/op",
            "extra": "20964 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - allocs/op",
            "value": 242,
            "unit": "allocs/op",
            "extra": "20964 times\n4 procs"
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
          "id": "081d382028a720206c4e3303e66c13bb52bf3121",
          "message": "selector: Select is now a method",
          "timestamp": "2024-10-15T17:26:49+02:00",
          "tree_id": "11f8c90bc4951df8944c09e35088ce910c088889",
          "url": "https://github.com/ucan-wg/go-ucan/commit/081d382028a720206c4e3303e66c13bb52bf3121"
        },
        "date": 1729006072713,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkContainerSerialisation/car_write",
            "value": 5705,
            "unit": "ns/op\t   17896 B/op\t      58 allocs/op",
            "extra": "204608 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - ns/op",
            "value": 5705,
            "unit": "ns/op",
            "extra": "204608 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - B/op",
            "value": 17896,
            "unit": "B/op",
            "extra": "204608 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "204608 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read",
            "value": 970735,
            "unit": "ns/op\t  148844 B/op\t    2621 allocs/op",
            "extra": "1202 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - ns/op",
            "value": 970735,
            "unit": "ns/op",
            "extra": "1202 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - B/op",
            "value": 148844,
            "unit": "B/op",
            "extra": "1202 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - allocs/op",
            "value": 2621,
            "unit": "allocs/op",
            "extra": "1202 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write",
            "value": 11792,
            "unit": "ns/op\t   24200 B/op\t      60 allocs/op",
            "extra": "106664 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - ns/op",
            "value": 11792,
            "unit": "ns/op",
            "extra": "106664 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - B/op",
            "value": 24200,
            "unit": "B/op",
            "extra": "106664 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - allocs/op",
            "value": 60,
            "unit": "allocs/op",
            "extra": "106664 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read",
            "value": 991614,
            "unit": "ns/op\t  150906 B/op\t    2623 allocs/op",
            "extra": "1195 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - ns/op",
            "value": 991614,
            "unit": "ns/op",
            "extra": "1195 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - B/op",
            "value": 150906,
            "unit": "B/op",
            "extra": "1195 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - allocs/op",
            "value": 2623,
            "unit": "allocs/op",
            "extra": "1195 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write",
            "value": 4490,
            "unit": "ns/op\t   16368 B/op\t      25 allocs/op",
            "extra": "265090 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - ns/op",
            "value": 4490,
            "unit": "ns/op",
            "extra": "265090 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - B/op",
            "value": 16368,
            "unit": "B/op",
            "extra": "265090 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - allocs/op",
            "value": 25,
            "unit": "allocs/op",
            "extra": "265090 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read",
            "value": 956458,
            "unit": "ns/op\t  138301 B/op\t    2524 allocs/op",
            "extra": "1225 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - ns/op",
            "value": 956458,
            "unit": "ns/op",
            "extra": "1225 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - B/op",
            "value": 138301,
            "unit": "B/op",
            "extra": "1225 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - allocs/op",
            "value": 2524,
            "unit": "allocs/op",
            "extra": "1225 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write",
            "value": 8816,
            "unit": "ns/op\t   20720 B/op\t      26 allocs/op",
            "extra": "124540 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - ns/op",
            "value": 8816,
            "unit": "ns/op",
            "extra": "124540 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - B/op",
            "value": 20720,
            "unit": "B/op",
            "extra": "124540 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "124540 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read",
            "value": 971607,
            "unit": "ns/op\t  140365 B/op\t    2526 allocs/op",
            "extra": "1219 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - ns/op",
            "value": 971607,
            "unit": "ns/op",
            "extra": "1219 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - B/op",
            "value": 140365,
            "unit": "B/op",
            "extra": "1219 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - allocs/op",
            "value": 2526,
            "unit": "allocs/op",
            "extra": "1219 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob",
            "value": 41.97,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "28580660 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - ns/op",
            "value": 41.97,
            "unit": "ns/op",
            "extra": "28580660 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "28580660 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "28580660 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad",
            "value": 15179,
            "unit": "ns/op\t   14569 B/op\t     162 allocs/op",
            "extra": "79766 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - ns/op",
            "value": 15179,
            "unit": "ns/op",
            "extra": "79766 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - B/op",
            "value": 14569,
            "unit": "B/op",
            "extra": "79766 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - allocs/op",
            "value": 162,
            "unit": "allocs/op",
            "extra": "79766 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson",
            "value": 114668,
            "unit": "ns/op\t   15102 B/op\t     369 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - ns/op",
            "value": 114668,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - B/op",
            "value": 15102,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - allocs/op",
            "value": 369,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal",
            "value": 53816,
            "unit": "ns/op\t    9600 B/op\t     208 allocs/op",
            "extra": "22155 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - ns/op",
            "value": 53816,
            "unit": "ns/op",
            "extra": "22155 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - B/op",
            "value": 9600,
            "unit": "B/op",
            "extra": "22155 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - allocs/op",
            "value": 208,
            "unit": "allocs/op",
            "extra": "22155 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal",
            "value": 97870,
            "unit": "ns/op\t   14158 B/op\t     300 allocs/op",
            "extra": "12304 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - ns/op",
            "value": 97870,
            "unit": "ns/op",
            "extra": "12304 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - B/op",
            "value": 14158,
            "unit": "B/op",
            "extra": "12304 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - allocs/op",
            "value": 300,
            "unit": "allocs/op",
            "extra": "12304 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson",
            "value": 59186,
            "unit": "ns/op\t   15112 B/op\t     248 allocs/op",
            "extra": "20228 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - ns/op",
            "value": 59186,
            "unit": "ns/op",
            "extra": "20228 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - B/op",
            "value": 15112,
            "unit": "B/op",
            "extra": "20228 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - allocs/op",
            "value": 248,
            "unit": "allocs/op",
            "extra": "20228 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader",
            "value": 109594,
            "unit": "ns/op\t   15054 B/op\t     368 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - ns/op",
            "value": 109594,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - B/op",
            "value": 15054,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - allocs/op",
            "value": 368,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter",
            "value": 56489,
            "unit": "ns/op\t    9120 B/op\t     239 allocs/op",
            "extra": "21669 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - ns/op",
            "value": 56489,
            "unit": "ns/op",
            "extra": "21669 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - B/op",
            "value": 9120,
            "unit": "B/op",
            "extra": "21669 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - allocs/op",
            "value": 239,
            "unit": "allocs/op",
            "extra": "21669 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader",
            "value": 98847,
            "unit": "ns/op\t   14157 B/op\t     300 allocs/op",
            "extra": "12214 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - ns/op",
            "value": 98847,
            "unit": "ns/op",
            "extra": "12214 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - B/op",
            "value": 14157,
            "unit": "B/op",
            "extra": "12214 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - allocs/op",
            "value": 300,
            "unit": "allocs/op",
            "extra": "12214 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader",
            "value": 57884,
            "unit": "ns/op\t   13080 B/op\t     242 allocs/op",
            "extra": "20911 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - ns/op",
            "value": 57884,
            "unit": "ns/op",
            "extra": "20911 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - B/op",
            "value": 13080,
            "unit": "B/op",
            "extra": "20911 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - allocs/op",
            "value": 242,
            "unit": "allocs/op",
            "extra": "20911 times\n4 procs"
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
          "id": "ac1b03f1444eaf3af93c8b86e1d95a9198d38c47",
          "message": "Merge pull request #44 from ucan-wg/policy-filtering\n\npolicy: add a way to filter policies with a path",
          "timestamp": "2024-10-16T14:29:57+02:00",
          "tree_id": "948cc83b040d756898427494ea3470c2b262e58f",
          "url": "https://github.com/ucan-wg/go-ucan/commit/ac1b03f1444eaf3af93c8b86e1d95a9198d38c47"
        },
        "date": 1729081846942,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkContainerSerialisation/car_write",
            "value": 5614,
            "unit": "ns/op\t   17896 B/op\t      58 allocs/op",
            "extra": "195351 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - ns/op",
            "value": 5614,
            "unit": "ns/op",
            "extra": "195351 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - B/op",
            "value": 17896,
            "unit": "B/op",
            "extra": "195351 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "195351 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read",
            "value": 959235,
            "unit": "ns/op\t  148842 B/op\t    2621 allocs/op",
            "extra": "1224 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - ns/op",
            "value": 959235,
            "unit": "ns/op",
            "extra": "1224 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - B/op",
            "value": 148842,
            "unit": "B/op",
            "extra": "1224 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - allocs/op",
            "value": 2621,
            "unit": "allocs/op",
            "extra": "1224 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write",
            "value": 11037,
            "unit": "ns/op\t   24200 B/op\t      60 allocs/op",
            "extra": "108208 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - ns/op",
            "value": 11037,
            "unit": "ns/op",
            "extra": "108208 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - B/op",
            "value": 24200,
            "unit": "B/op",
            "extra": "108208 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - allocs/op",
            "value": 60,
            "unit": "allocs/op",
            "extra": "108208 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read",
            "value": 963119,
            "unit": "ns/op\t  150906 B/op\t    2623 allocs/op",
            "extra": "1221 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - ns/op",
            "value": 963119,
            "unit": "ns/op",
            "extra": "1221 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - B/op",
            "value": 150906,
            "unit": "B/op",
            "extra": "1221 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - allocs/op",
            "value": 2623,
            "unit": "allocs/op",
            "extra": "1221 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write",
            "value": 4394,
            "unit": "ns/op\t   16368 B/op\t      25 allocs/op",
            "extra": "268683 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - ns/op",
            "value": 4394,
            "unit": "ns/op",
            "extra": "268683 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - B/op",
            "value": 16368,
            "unit": "B/op",
            "extra": "268683 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - allocs/op",
            "value": 25,
            "unit": "allocs/op",
            "extra": "268683 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read",
            "value": 947706,
            "unit": "ns/op\t  138303 B/op\t    2524 allocs/op",
            "extra": "1263 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - ns/op",
            "value": 947706,
            "unit": "ns/op",
            "extra": "1263 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - B/op",
            "value": 138303,
            "unit": "B/op",
            "extra": "1263 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - allocs/op",
            "value": 2524,
            "unit": "allocs/op",
            "extra": "1263 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write",
            "value": 9153,
            "unit": "ns/op\t   20720 B/op\t      26 allocs/op",
            "extra": "136086 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - ns/op",
            "value": 9153,
            "unit": "ns/op",
            "extra": "136086 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - B/op",
            "value": 20720,
            "unit": "B/op",
            "extra": "136086 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "136086 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read",
            "value": 957519,
            "unit": "ns/op\t  140365 B/op\t    2526 allocs/op",
            "extra": "1234 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - ns/op",
            "value": 957519,
            "unit": "ns/op",
            "extra": "1234 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - B/op",
            "value": 140365,
            "unit": "B/op",
            "extra": "1234 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - allocs/op",
            "value": 2526,
            "unit": "allocs/op",
            "extra": "1234 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob",
            "value": 49.97,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "24123796 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - ns/op",
            "value": 49.97,
            "unit": "ns/op",
            "extra": "24123796 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "24123796 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "24123796 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad",
            "value": 14707,
            "unit": "ns/op\t   14569 B/op\t     162 allocs/op",
            "extra": "80958 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - ns/op",
            "value": 14707,
            "unit": "ns/op",
            "extra": "80958 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - B/op",
            "value": 14569,
            "unit": "B/op",
            "extra": "80958 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - allocs/op",
            "value": 162,
            "unit": "allocs/op",
            "extra": "80958 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson",
            "value": 107604,
            "unit": "ns/op\t   15102 B/op\t     369 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - ns/op",
            "value": 107604,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - B/op",
            "value": 15102,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - allocs/op",
            "value": 369,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal",
            "value": 53869,
            "unit": "ns/op\t    9600 B/op\t     208 allocs/op",
            "extra": "22772 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - ns/op",
            "value": 53869,
            "unit": "ns/op",
            "extra": "22772 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - B/op",
            "value": 9600,
            "unit": "B/op",
            "extra": "22772 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - allocs/op",
            "value": 208,
            "unit": "allocs/op",
            "extra": "22772 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal",
            "value": 95874,
            "unit": "ns/op\t   14157 B/op\t     300 allocs/op",
            "extra": "12496 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - ns/op",
            "value": 95874,
            "unit": "ns/op",
            "extra": "12496 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - B/op",
            "value": 14157,
            "unit": "B/op",
            "extra": "12496 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - allocs/op",
            "value": 300,
            "unit": "allocs/op",
            "extra": "12496 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson",
            "value": 59756,
            "unit": "ns/op\t   15112 B/op\t     248 allocs/op",
            "extra": "20464 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - ns/op",
            "value": 59756,
            "unit": "ns/op",
            "extra": "20464 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - B/op",
            "value": 15112,
            "unit": "B/op",
            "extra": "20464 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - allocs/op",
            "value": 248,
            "unit": "allocs/op",
            "extra": "20464 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader",
            "value": 107983,
            "unit": "ns/op\t   15053 B/op\t     368 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - ns/op",
            "value": 107983,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - B/op",
            "value": 15053,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - allocs/op",
            "value": 368,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter",
            "value": 56414,
            "unit": "ns/op\t    9120 B/op\t     239 allocs/op",
            "extra": "22028 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - ns/op",
            "value": 56414,
            "unit": "ns/op",
            "extra": "22028 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - B/op",
            "value": 9120,
            "unit": "B/op",
            "extra": "22028 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - allocs/op",
            "value": 239,
            "unit": "allocs/op",
            "extra": "22028 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader",
            "value": 96853,
            "unit": "ns/op\t   14157 B/op\t     300 allocs/op",
            "extra": "12345 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - ns/op",
            "value": 96853,
            "unit": "ns/op",
            "extra": "12345 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - B/op",
            "value": 14157,
            "unit": "B/op",
            "extra": "12345 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - allocs/op",
            "value": 300,
            "unit": "allocs/op",
            "extra": "12345 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader",
            "value": 57020,
            "unit": "ns/op\t   13080 B/op\t     242 allocs/op",
            "extra": "21172 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - ns/op",
            "value": 57020,
            "unit": "ns/op",
            "extra": "21172 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - B/op",
            "value": 13080,
            "unit": "B/op",
            "extra": "21172 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - allocs/op",
            "value": 242,
            "unit": "allocs/op",
            "extra": "21172 times\n4 procs"
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
          "id": "deaf9c4fe9e0734d66393949b592c8327abfcefc",
          "message": "Merge pull request #46 from ucan-wg/rework-policies\n\nselector: rework to match the spec, cleanup lots of edge cases",
          "timestamp": "2024-10-24T13:01:10+02:00",
          "tree_id": "b5b06019217dbadeea9e486b5e923e33a085fef7",
          "url": "https://github.com/ucan-wg/go-ucan/commit/deaf9c4fe9e0734d66393949b592c8327abfcefc"
        },
        "date": 1729767727464,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkContainerSerialisation/car_write",
            "value": 5632,
            "unit": "ns/op\t   17896 B/op\t      58 allocs/op",
            "extra": "214245 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - ns/op",
            "value": 5632,
            "unit": "ns/op",
            "extra": "214245 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - B/op",
            "value": 17896,
            "unit": "B/op",
            "extra": "214245 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "214245 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read",
            "value": 952301,
            "unit": "ns/op\t  148842 B/op\t    2621 allocs/op",
            "extra": "1240 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - ns/op",
            "value": 952301,
            "unit": "ns/op",
            "extra": "1240 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - B/op",
            "value": 148842,
            "unit": "B/op",
            "extra": "1240 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - allocs/op",
            "value": 2621,
            "unit": "allocs/op",
            "extra": "1240 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write",
            "value": 10997,
            "unit": "ns/op\t   24200 B/op\t      60 allocs/op",
            "extra": "109269 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - ns/op",
            "value": 10997,
            "unit": "ns/op",
            "extra": "109269 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - B/op",
            "value": 24200,
            "unit": "B/op",
            "extra": "109269 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - allocs/op",
            "value": 60,
            "unit": "allocs/op",
            "extra": "109269 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read",
            "value": 985050,
            "unit": "ns/op\t  150904 B/op\t    2623 allocs/op",
            "extra": "1216 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - ns/op",
            "value": 985050,
            "unit": "ns/op",
            "extra": "1216 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - B/op",
            "value": 150904,
            "unit": "B/op",
            "extra": "1216 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - allocs/op",
            "value": 2623,
            "unit": "allocs/op",
            "extra": "1216 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write",
            "value": 4707,
            "unit": "ns/op\t   16368 B/op\t      25 allocs/op",
            "extra": "217352 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - ns/op",
            "value": 4707,
            "unit": "ns/op",
            "extra": "217352 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - B/op",
            "value": 16368,
            "unit": "B/op",
            "extra": "217352 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - allocs/op",
            "value": 25,
            "unit": "allocs/op",
            "extra": "217352 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read",
            "value": 933504,
            "unit": "ns/op\t  138298 B/op\t    2524 allocs/op",
            "extra": "1254 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - ns/op",
            "value": 933504,
            "unit": "ns/op",
            "extra": "1254 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - B/op",
            "value": 138298,
            "unit": "B/op",
            "extra": "1254 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - allocs/op",
            "value": 2524,
            "unit": "allocs/op",
            "extra": "1254 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write",
            "value": 8678,
            "unit": "ns/op\t   20720 B/op\t      26 allocs/op",
            "extra": "141087 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - ns/op",
            "value": 8678,
            "unit": "ns/op",
            "extra": "141087 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - B/op",
            "value": 20720,
            "unit": "B/op",
            "extra": "141087 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "141087 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read",
            "value": 946096,
            "unit": "ns/op\t  140365 B/op\t    2526 allocs/op",
            "extra": "1238 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - ns/op",
            "value": 946096,
            "unit": "ns/op",
            "extra": "1238 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - B/op",
            "value": 140365,
            "unit": "B/op",
            "extra": "1238 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - allocs/op",
            "value": 2526,
            "unit": "allocs/op",
            "extra": "1238 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob",
            "value": 51.49,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "23832345 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - ns/op",
            "value": 51.49,
            "unit": "ns/op",
            "extra": "23832345 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "23832345 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "23832345 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad",
            "value": 14663,
            "unit": "ns/op\t   14569 B/op\t     162 allocs/op",
            "extra": "81810 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - ns/op",
            "value": 14663,
            "unit": "ns/op",
            "extra": "81810 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - B/op",
            "value": 14569,
            "unit": "B/op",
            "extra": "81810 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - allocs/op",
            "value": 162,
            "unit": "allocs/op",
            "extra": "81810 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson",
            "value": 107494,
            "unit": "ns/op\t   15102 B/op\t     369 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - ns/op",
            "value": 107494,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - B/op",
            "value": 15102,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - allocs/op",
            "value": 369,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal",
            "value": 52664,
            "unit": "ns/op\t    9600 B/op\t     208 allocs/op",
            "extra": "22936 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - ns/op",
            "value": 52664,
            "unit": "ns/op",
            "extra": "22936 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - B/op",
            "value": 9600,
            "unit": "B/op",
            "extra": "22936 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - allocs/op",
            "value": 208,
            "unit": "allocs/op",
            "extra": "22936 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal",
            "value": 95247,
            "unit": "ns/op\t   14157 B/op\t     300 allocs/op",
            "extra": "12493 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - ns/op",
            "value": 95247,
            "unit": "ns/op",
            "extra": "12493 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - B/op",
            "value": 14157,
            "unit": "B/op",
            "extra": "12493 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - allocs/op",
            "value": 300,
            "unit": "allocs/op",
            "extra": "12493 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson",
            "value": 56965,
            "unit": "ns/op\t   15112 B/op\t     248 allocs/op",
            "extra": "20934 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - ns/op",
            "value": 56965,
            "unit": "ns/op",
            "extra": "20934 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - B/op",
            "value": 15112,
            "unit": "B/op",
            "extra": "20934 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - allocs/op",
            "value": 248,
            "unit": "allocs/op",
            "extra": "20934 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader",
            "value": 106933,
            "unit": "ns/op\t   15053 B/op\t     368 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - ns/op",
            "value": 106933,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - B/op",
            "value": 15053,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - allocs/op",
            "value": 368,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter",
            "value": 53582,
            "unit": "ns/op\t    9120 B/op\t     239 allocs/op",
            "extra": "22224 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - ns/op",
            "value": 53582,
            "unit": "ns/op",
            "extra": "22224 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - B/op",
            "value": 9120,
            "unit": "B/op",
            "extra": "22224 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - allocs/op",
            "value": 239,
            "unit": "allocs/op",
            "extra": "22224 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader",
            "value": 96193,
            "unit": "ns/op\t   14157 B/op\t     300 allocs/op",
            "extra": "12472 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - ns/op",
            "value": 96193,
            "unit": "ns/op",
            "extra": "12472 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - B/op",
            "value": 14157,
            "unit": "B/op",
            "extra": "12472 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - allocs/op",
            "value": 300,
            "unit": "allocs/op",
            "extra": "12472 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader",
            "value": 56403,
            "unit": "ns/op\t   13080 B/op\t     242 allocs/op",
            "extra": "21259 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - ns/op",
            "value": 56403,
            "unit": "ns/op",
            "extra": "21259 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - B/op",
            "value": 13080,
            "unit": "B/op",
            "extra": "21259 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - allocs/op",
            "value": 242,
            "unit": "allocs/op",
            "extra": "21259 times\n4 procs"
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
          "id": "e76354fb0af6f16dc52eb59bba633553a5708600",
          "message": "Merge pull request #47 from ucan-wg/simplify-literal\n\nliteral: simplify package with built-in functions",
          "timestamp": "2024-10-24T14:33:33+02:00",
          "tree_id": "90c5fc08c871717a1ecf701938729c7c6b7f99cc",
          "url": "https://github.com/ucan-wg/go-ucan/commit/e76354fb0af6f16dc52eb59bba633553a5708600"
        },
        "date": 1729773257322,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkContainerSerialisation/car_write",
            "value": 5518,
            "unit": "ns/op\t   17896 B/op\t      58 allocs/op",
            "extra": "203031 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - ns/op",
            "value": 5518,
            "unit": "ns/op",
            "extra": "203031 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - B/op",
            "value": 17896,
            "unit": "B/op",
            "extra": "203031 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "203031 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read",
            "value": 957058,
            "unit": "ns/op\t  148841 B/op\t    2621 allocs/op",
            "extra": "1216 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - ns/op",
            "value": 957058,
            "unit": "ns/op",
            "extra": "1216 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - B/op",
            "value": 148841,
            "unit": "B/op",
            "extra": "1216 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - allocs/op",
            "value": 2621,
            "unit": "allocs/op",
            "extra": "1216 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write",
            "value": 10900,
            "unit": "ns/op\t   24200 B/op\t      60 allocs/op",
            "extra": "109063 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - ns/op",
            "value": 10900,
            "unit": "ns/op",
            "extra": "109063 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - B/op",
            "value": 24200,
            "unit": "B/op",
            "extra": "109063 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - allocs/op",
            "value": 60,
            "unit": "allocs/op",
            "extra": "109063 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read",
            "value": 968791,
            "unit": "ns/op\t  150905 B/op\t    2623 allocs/op",
            "extra": "1212 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - ns/op",
            "value": 968791,
            "unit": "ns/op",
            "extra": "1212 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - B/op",
            "value": 150905,
            "unit": "B/op",
            "extra": "1212 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - allocs/op",
            "value": 2623,
            "unit": "allocs/op",
            "extra": "1212 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write",
            "value": 4372,
            "unit": "ns/op\t   16368 B/op\t      25 allocs/op",
            "extra": "261961 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - ns/op",
            "value": 4372,
            "unit": "ns/op",
            "extra": "261961 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - B/op",
            "value": 16368,
            "unit": "B/op",
            "extra": "261961 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - allocs/op",
            "value": 25,
            "unit": "allocs/op",
            "extra": "261961 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read",
            "value": 942197,
            "unit": "ns/op\t  138299 B/op\t    2524 allocs/op",
            "extra": "1256 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - ns/op",
            "value": 942197,
            "unit": "ns/op",
            "extra": "1256 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - B/op",
            "value": 138299,
            "unit": "B/op",
            "extra": "1256 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - allocs/op",
            "value": 2524,
            "unit": "allocs/op",
            "extra": "1256 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write",
            "value": 8578,
            "unit": "ns/op\t   20720 B/op\t      26 allocs/op",
            "extra": "136495 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - ns/op",
            "value": 8578,
            "unit": "ns/op",
            "extra": "136495 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - B/op",
            "value": 20720,
            "unit": "B/op",
            "extra": "136495 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "136495 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read",
            "value": 986028,
            "unit": "ns/op\t  140364 B/op\t    2526 allocs/op",
            "extra": "1238 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - ns/op",
            "value": 986028,
            "unit": "ns/op",
            "extra": "1238 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - B/op",
            "value": 140364,
            "unit": "B/op",
            "extra": "1238 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - allocs/op",
            "value": 2526,
            "unit": "allocs/op",
            "extra": "1238 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob",
            "value": 49.61,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "24736586 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - ns/op",
            "value": 49.61,
            "unit": "ns/op",
            "extra": "24736586 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "24736586 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "24736586 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad",
            "value": 14815,
            "unit": "ns/op\t   14569 B/op\t     162 allocs/op",
            "extra": "80598 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - ns/op",
            "value": 14815,
            "unit": "ns/op",
            "extra": "80598 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - B/op",
            "value": 14569,
            "unit": "B/op",
            "extra": "80598 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - allocs/op",
            "value": 162,
            "unit": "allocs/op",
            "extra": "80598 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson",
            "value": 106841,
            "unit": "ns/op\t   15102 B/op\t     369 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - ns/op",
            "value": 106841,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - B/op",
            "value": 15102,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - allocs/op",
            "value": 369,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal",
            "value": 51852,
            "unit": "ns/op\t    9600 B/op\t     208 allocs/op",
            "extra": "22982 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - ns/op",
            "value": 51852,
            "unit": "ns/op",
            "extra": "22982 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - B/op",
            "value": 9600,
            "unit": "B/op",
            "extra": "22982 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - allocs/op",
            "value": 208,
            "unit": "allocs/op",
            "extra": "22982 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal",
            "value": 95579,
            "unit": "ns/op\t   14157 B/op\t     300 allocs/op",
            "extra": "12550 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - ns/op",
            "value": 95579,
            "unit": "ns/op",
            "extra": "12550 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - B/op",
            "value": 14157,
            "unit": "B/op",
            "extra": "12550 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - allocs/op",
            "value": 300,
            "unit": "allocs/op",
            "extra": "12550 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson",
            "value": 57033,
            "unit": "ns/op\t   15112 B/op\t     248 allocs/op",
            "extra": "21040 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - ns/op",
            "value": 57033,
            "unit": "ns/op",
            "extra": "21040 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - B/op",
            "value": 15112,
            "unit": "B/op",
            "extra": "21040 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - allocs/op",
            "value": 248,
            "unit": "allocs/op",
            "extra": "21040 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader",
            "value": 113419,
            "unit": "ns/op\t   15054 B/op\t     368 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - ns/op",
            "value": 113419,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - B/op",
            "value": 15054,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - allocs/op",
            "value": 368,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter",
            "value": 53821,
            "unit": "ns/op\t    9120 B/op\t     239 allocs/op",
            "extra": "22233 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - ns/op",
            "value": 53821,
            "unit": "ns/op",
            "extra": "22233 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - B/op",
            "value": 9120,
            "unit": "B/op",
            "extra": "22233 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - allocs/op",
            "value": 239,
            "unit": "allocs/op",
            "extra": "22233 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader",
            "value": 96271,
            "unit": "ns/op\t   14157 B/op\t     300 allocs/op",
            "extra": "12384 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - ns/op",
            "value": 96271,
            "unit": "ns/op",
            "extra": "12384 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - B/op",
            "value": 14157,
            "unit": "B/op",
            "extra": "12384 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - allocs/op",
            "value": 300,
            "unit": "allocs/op",
            "extra": "12384 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader",
            "value": 56893,
            "unit": "ns/op\t   13080 B/op\t     242 allocs/op",
            "extra": "21272 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - ns/op",
            "value": 56893,
            "unit": "ns/op",
            "extra": "21272 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - B/op",
            "value": 13080,
            "unit": "B/op",
            "extra": "21272 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - allocs/op",
            "value": 242,
            "unit": "allocs/op",
            "extra": "21272 times\n4 procs"
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
            "email": "batolettre@gmail.com",
            "name": "Michael Muré",
            "username": "MichaelMure"
          },
          "distinct": true,
          "id": "3b2ea8d435c96b5ff658632eda66ea891e45bf79",
          "message": "minor cleanups",
          "timestamp": "2024-10-24T14:39:39+02:00",
          "tree_id": "ded3bfe341aa722d216b350809779eb703a5b24b",
          "url": "https://github.com/ucan-wg/go-ucan/commit/3b2ea8d435c96b5ff658632eda66ea891e45bf79"
        },
        "date": 1729773628258,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkContainerSerialisation/car_write",
            "value": 5651,
            "unit": "ns/op\t   17896 B/op\t      58 allocs/op",
            "extra": "207939 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - ns/op",
            "value": 5651,
            "unit": "ns/op",
            "extra": "207939 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - B/op",
            "value": 17896,
            "unit": "B/op",
            "extra": "207939 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "207939 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read",
            "value": 969208,
            "unit": "ns/op\t  148841 B/op\t    2621 allocs/op",
            "extra": "1209 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - ns/op",
            "value": 969208,
            "unit": "ns/op",
            "extra": "1209 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - B/op",
            "value": 148841,
            "unit": "B/op",
            "extra": "1209 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - allocs/op",
            "value": 2621,
            "unit": "allocs/op",
            "extra": "1209 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write",
            "value": 10980,
            "unit": "ns/op\t   24200 B/op\t      60 allocs/op",
            "extra": "107181 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - ns/op",
            "value": 10980,
            "unit": "ns/op",
            "extra": "107181 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - B/op",
            "value": 24200,
            "unit": "B/op",
            "extra": "107181 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - allocs/op",
            "value": 60,
            "unit": "allocs/op",
            "extra": "107181 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read",
            "value": 998789,
            "unit": "ns/op\t  150911 B/op\t    2623 allocs/op",
            "extra": "1086 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - ns/op",
            "value": 998789,
            "unit": "ns/op",
            "extra": "1086 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - B/op",
            "value": 150911,
            "unit": "B/op",
            "extra": "1086 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - allocs/op",
            "value": 2623,
            "unit": "allocs/op",
            "extra": "1086 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write",
            "value": 4467,
            "unit": "ns/op\t   16368 B/op\t      25 allocs/op",
            "extra": "261300 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - ns/op",
            "value": 4467,
            "unit": "ns/op",
            "extra": "261300 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - B/op",
            "value": 16368,
            "unit": "B/op",
            "extra": "261300 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - allocs/op",
            "value": 25,
            "unit": "allocs/op",
            "extra": "261300 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read",
            "value": 952555,
            "unit": "ns/op\t  138300 B/op\t    2524 allocs/op",
            "extra": "1245 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - ns/op",
            "value": 952555,
            "unit": "ns/op",
            "extra": "1245 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - B/op",
            "value": 138300,
            "unit": "B/op",
            "extra": "1245 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - allocs/op",
            "value": 2524,
            "unit": "allocs/op",
            "extra": "1245 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write",
            "value": 8826,
            "unit": "ns/op\t   20720 B/op\t      26 allocs/op",
            "extra": "135027 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - ns/op",
            "value": 8826,
            "unit": "ns/op",
            "extra": "135027 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - B/op",
            "value": 20720,
            "unit": "B/op",
            "extra": "135027 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "135027 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read",
            "value": 971687,
            "unit": "ns/op\t  140364 B/op\t    2526 allocs/op",
            "extra": "1219 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - ns/op",
            "value": 971687,
            "unit": "ns/op",
            "extra": "1219 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - B/op",
            "value": 140364,
            "unit": "B/op",
            "extra": "1219 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - allocs/op",
            "value": 2526,
            "unit": "allocs/op",
            "extra": "1219 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob",
            "value": 49.73,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "24697796 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - ns/op",
            "value": 49.73,
            "unit": "ns/op",
            "extra": "24697796 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "24697796 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "24697796 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad",
            "value": 15048,
            "unit": "ns/op\t   14569 B/op\t     162 allocs/op",
            "extra": "78664 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - ns/op",
            "value": 15048,
            "unit": "ns/op",
            "extra": "78664 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - B/op",
            "value": 14569,
            "unit": "B/op",
            "extra": "78664 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - allocs/op",
            "value": 162,
            "unit": "allocs/op",
            "extra": "78664 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson",
            "value": 107942,
            "unit": "ns/op\t   15102 B/op\t     369 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - ns/op",
            "value": 107942,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - B/op",
            "value": 15102,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - allocs/op",
            "value": 369,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal",
            "value": 53135,
            "unit": "ns/op\t    9600 B/op\t     208 allocs/op",
            "extra": "22254 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - ns/op",
            "value": 53135,
            "unit": "ns/op",
            "extra": "22254 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - B/op",
            "value": 9600,
            "unit": "B/op",
            "extra": "22254 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - allocs/op",
            "value": 208,
            "unit": "allocs/op",
            "extra": "22254 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal",
            "value": 97377,
            "unit": "ns/op\t   14157 B/op\t     300 allocs/op",
            "extra": "12336 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - ns/op",
            "value": 97377,
            "unit": "ns/op",
            "extra": "12336 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - B/op",
            "value": 14157,
            "unit": "B/op",
            "extra": "12336 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - allocs/op",
            "value": 300,
            "unit": "allocs/op",
            "extra": "12336 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson",
            "value": 59199,
            "unit": "ns/op\t   15112 B/op\t     248 allocs/op",
            "extra": "20510 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - ns/op",
            "value": 59199,
            "unit": "ns/op",
            "extra": "20510 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - B/op",
            "value": 15112,
            "unit": "B/op",
            "extra": "20510 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - allocs/op",
            "value": 248,
            "unit": "allocs/op",
            "extra": "20510 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader",
            "value": 114890,
            "unit": "ns/op\t   15054 B/op\t     368 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - ns/op",
            "value": 114890,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - B/op",
            "value": 15054,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - allocs/op",
            "value": 368,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter",
            "value": 55697,
            "unit": "ns/op\t    9120 B/op\t     239 allocs/op",
            "extra": "21560 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - ns/op",
            "value": 55697,
            "unit": "ns/op",
            "extra": "21560 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - B/op",
            "value": 9120,
            "unit": "B/op",
            "extra": "21560 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - allocs/op",
            "value": 239,
            "unit": "allocs/op",
            "extra": "21560 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader",
            "value": 99755,
            "unit": "ns/op\t   14157 B/op\t     300 allocs/op",
            "extra": "12118 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - ns/op",
            "value": 99755,
            "unit": "ns/op",
            "extra": "12118 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - B/op",
            "value": 14157,
            "unit": "B/op",
            "extra": "12118 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - allocs/op",
            "value": 300,
            "unit": "allocs/op",
            "extra": "12118 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader",
            "value": 58381,
            "unit": "ns/op\t   13080 B/op\t     242 allocs/op",
            "extra": "20480 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - ns/op",
            "value": 58381,
            "unit": "ns/op",
            "extra": "20480 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - B/op",
            "value": 13080,
            "unit": "B/op",
            "extra": "20480 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - allocs/op",
            "value": 242,
            "unit": "allocs/op",
            "extra": "20480 times\n4 procs"
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
            "email": "batolettre@gmail.com",
            "name": "Michael Muré",
            "username": "MichaelMure"
          },
          "distinct": true,
          "id": "02c6ca24c2e2cd33f49caa7222f44479add4a55e",
          "message": "minor cleanups",
          "timestamp": "2024-10-24T14:40:43+02:00",
          "tree_id": "cdc0233761e38aac7d8745bbe321d8f322d4199d",
          "url": "https://github.com/ucan-wg/go-ucan/commit/02c6ca24c2e2cd33f49caa7222f44479add4a55e"
        },
        "date": 1729773694781,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkContainerSerialisation/car_write",
            "value": 5679,
            "unit": "ns/op\t   17896 B/op\t      58 allocs/op",
            "extra": "206404 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - ns/op",
            "value": 5679,
            "unit": "ns/op",
            "extra": "206404 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - B/op",
            "value": 17896,
            "unit": "B/op",
            "extra": "206404 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "206404 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read",
            "value": 1008470,
            "unit": "ns/op\t  148842 B/op\t    2621 allocs/op",
            "extra": "1208 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - ns/op",
            "value": 1008470,
            "unit": "ns/op",
            "extra": "1208 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - B/op",
            "value": 148842,
            "unit": "B/op",
            "extra": "1208 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - allocs/op",
            "value": 2621,
            "unit": "allocs/op",
            "extra": "1208 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write",
            "value": 11236,
            "unit": "ns/op\t   24200 B/op\t      60 allocs/op",
            "extra": "107012 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - ns/op",
            "value": 11236,
            "unit": "ns/op",
            "extra": "107012 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - B/op",
            "value": 24200,
            "unit": "B/op",
            "extra": "107012 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - allocs/op",
            "value": 60,
            "unit": "allocs/op",
            "extra": "107012 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read",
            "value": 989735,
            "unit": "ns/op\t  150906 B/op\t    2623 allocs/op",
            "extra": "1194 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - ns/op",
            "value": 989735,
            "unit": "ns/op",
            "extra": "1194 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - B/op",
            "value": 150906,
            "unit": "B/op",
            "extra": "1194 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - allocs/op",
            "value": 2623,
            "unit": "allocs/op",
            "extra": "1194 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write",
            "value": 4628,
            "unit": "ns/op\t   16368 B/op\t      25 allocs/op",
            "extra": "262526 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - ns/op",
            "value": 4628,
            "unit": "ns/op",
            "extra": "262526 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - B/op",
            "value": 16368,
            "unit": "B/op",
            "extra": "262526 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - allocs/op",
            "value": 25,
            "unit": "allocs/op",
            "extra": "262526 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read",
            "value": 962737,
            "unit": "ns/op\t  138301 B/op\t    2524 allocs/op",
            "extra": "1113 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - ns/op",
            "value": 962737,
            "unit": "ns/op",
            "extra": "1113 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - B/op",
            "value": 138301,
            "unit": "B/op",
            "extra": "1113 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - allocs/op",
            "value": 2524,
            "unit": "allocs/op",
            "extra": "1113 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write",
            "value": 8763,
            "unit": "ns/op\t   20720 B/op\t      26 allocs/op",
            "extra": "137292 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - ns/op",
            "value": 8763,
            "unit": "ns/op",
            "extra": "137292 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - B/op",
            "value": 20720,
            "unit": "B/op",
            "extra": "137292 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "137292 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read",
            "value": 972654,
            "unit": "ns/op\t  140364 B/op\t    2526 allocs/op",
            "extra": "1218 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - ns/op",
            "value": 972654,
            "unit": "ns/op",
            "extra": "1218 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - B/op",
            "value": 140364,
            "unit": "B/op",
            "extra": "1218 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - allocs/op",
            "value": 2526,
            "unit": "allocs/op",
            "extra": "1218 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob",
            "value": 49.57,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "23967138 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - ns/op",
            "value": 49.57,
            "unit": "ns/op",
            "extra": "23967138 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "23967138 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "23967138 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad",
            "value": 14838,
            "unit": "ns/op\t   14569 B/op\t     162 allocs/op",
            "extra": "80924 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - ns/op",
            "value": 14838,
            "unit": "ns/op",
            "extra": "80924 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - B/op",
            "value": 14569,
            "unit": "B/op",
            "extra": "80924 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - allocs/op",
            "value": 162,
            "unit": "allocs/op",
            "extra": "80924 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson",
            "value": 107788,
            "unit": "ns/op\t   15102 B/op\t     369 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - ns/op",
            "value": 107788,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - B/op",
            "value": 15102,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - allocs/op",
            "value": 369,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal",
            "value": 53024,
            "unit": "ns/op\t    9600 B/op\t     208 allocs/op",
            "extra": "22640 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - ns/op",
            "value": 53024,
            "unit": "ns/op",
            "extra": "22640 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - B/op",
            "value": 9600,
            "unit": "B/op",
            "extra": "22640 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - allocs/op",
            "value": 208,
            "unit": "allocs/op",
            "extra": "22640 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal",
            "value": 98729,
            "unit": "ns/op\t   14157 B/op\t     300 allocs/op",
            "extra": "12140 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - ns/op",
            "value": 98729,
            "unit": "ns/op",
            "extra": "12140 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - B/op",
            "value": 14157,
            "unit": "B/op",
            "extra": "12140 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - allocs/op",
            "value": 300,
            "unit": "allocs/op",
            "extra": "12140 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson",
            "value": 58767,
            "unit": "ns/op\t   15112 B/op\t     248 allocs/op",
            "extra": "20163 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - ns/op",
            "value": 58767,
            "unit": "ns/op",
            "extra": "20163 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - B/op",
            "value": 15112,
            "unit": "B/op",
            "extra": "20163 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - allocs/op",
            "value": 248,
            "unit": "allocs/op",
            "extra": "20163 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader",
            "value": 109070,
            "unit": "ns/op\t   15054 B/op\t     368 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - ns/op",
            "value": 109070,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - B/op",
            "value": 15054,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - allocs/op",
            "value": 368,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter",
            "value": 54997,
            "unit": "ns/op\t    9120 B/op\t     239 allocs/op",
            "extra": "21842 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - ns/op",
            "value": 54997,
            "unit": "ns/op",
            "extra": "21842 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - B/op",
            "value": 9120,
            "unit": "B/op",
            "extra": "21842 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - allocs/op",
            "value": 239,
            "unit": "allocs/op",
            "extra": "21842 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader",
            "value": 98109,
            "unit": "ns/op\t   14157 B/op\t     300 allocs/op",
            "extra": "12247 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - ns/op",
            "value": 98109,
            "unit": "ns/op",
            "extra": "12247 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - B/op",
            "value": 14157,
            "unit": "B/op",
            "extra": "12247 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - allocs/op",
            "value": 300,
            "unit": "allocs/op",
            "extra": "12247 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader",
            "value": 57987,
            "unit": "ns/op\t   13080 B/op\t     242 allocs/op",
            "extra": "20814 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - ns/op",
            "value": 57987,
            "unit": "ns/op",
            "extra": "20814 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - B/op",
            "value": 13080,
            "unit": "B/op",
            "extra": "20814 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - allocs/op",
            "value": 242,
            "unit": "allocs/op",
            "extra": "20814 times\n4 procs"
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
            "email": "batolettre@gmail.com",
            "name": "Michael Muré",
            "username": "MichaelMure"
          },
          "distinct": true,
          "id": "6d0fbd4d5a7fc79f4fb6b0b1d269605d20b96d41",
          "message": "minor cleanups",
          "timestamp": "2024-10-24T14:46:01+02:00",
          "tree_id": "e004fe55b7d123bc41ded35b971ebf2e589f6bf7",
          "url": "https://github.com/ucan-wg/go-ucan/commit/6d0fbd4d5a7fc79f4fb6b0b1d269605d20b96d41"
        },
        "date": 1729774008849,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkContainerSerialisation/car_write",
            "value": 5628,
            "unit": "ns/op\t   17896 B/op\t      58 allocs/op",
            "extra": "210698 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - ns/op",
            "value": 5628,
            "unit": "ns/op",
            "extra": "210698 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - B/op",
            "value": 17896,
            "unit": "B/op",
            "extra": "210698 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "210698 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read",
            "value": 977863,
            "unit": "ns/op\t  148836 B/op\t    2621 allocs/op",
            "extra": "1204 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - ns/op",
            "value": 977863,
            "unit": "ns/op",
            "extra": "1204 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - B/op",
            "value": 148836,
            "unit": "B/op",
            "extra": "1204 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - allocs/op",
            "value": 2621,
            "unit": "allocs/op",
            "extra": "1204 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write",
            "value": 11221,
            "unit": "ns/op\t   24200 B/op\t      60 allocs/op",
            "extra": "104548 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - ns/op",
            "value": 11221,
            "unit": "ns/op",
            "extra": "104548 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - B/op",
            "value": 24200,
            "unit": "B/op",
            "extra": "104548 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - allocs/op",
            "value": 60,
            "unit": "allocs/op",
            "extra": "104548 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read",
            "value": 993765,
            "unit": "ns/op\t  150903 B/op\t    2623 allocs/op",
            "extra": "1184 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - ns/op",
            "value": 993765,
            "unit": "ns/op",
            "extra": "1184 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - B/op",
            "value": 150903,
            "unit": "B/op",
            "extra": "1184 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - allocs/op",
            "value": 2623,
            "unit": "allocs/op",
            "extra": "1184 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write",
            "value": 4633,
            "unit": "ns/op\t   16368 B/op\t      25 allocs/op",
            "extra": "259533 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - ns/op",
            "value": 4633,
            "unit": "ns/op",
            "extra": "259533 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - B/op",
            "value": 16368,
            "unit": "B/op",
            "extra": "259533 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - allocs/op",
            "value": 25,
            "unit": "allocs/op",
            "extra": "259533 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read",
            "value": 958853,
            "unit": "ns/op\t  138303 B/op\t    2524 allocs/op",
            "extra": "1225 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - ns/op",
            "value": 958853,
            "unit": "ns/op",
            "extra": "1225 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - B/op",
            "value": 138303,
            "unit": "B/op",
            "extra": "1225 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - allocs/op",
            "value": 2524,
            "unit": "allocs/op",
            "extra": "1225 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write",
            "value": 8809,
            "unit": "ns/op\t   20720 B/op\t      26 allocs/op",
            "extra": "136501 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - ns/op",
            "value": 8809,
            "unit": "ns/op",
            "extra": "136501 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - B/op",
            "value": 20720,
            "unit": "B/op",
            "extra": "136501 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "136501 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read",
            "value": 975863,
            "unit": "ns/op\t  140365 B/op\t    2526 allocs/op",
            "extra": "1204 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - ns/op",
            "value": 975863,
            "unit": "ns/op",
            "extra": "1204 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - B/op",
            "value": 140365,
            "unit": "B/op",
            "extra": "1204 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - allocs/op",
            "value": 2526,
            "unit": "allocs/op",
            "extra": "1204 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob",
            "value": 49.77,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "24023499 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - ns/op",
            "value": 49.77,
            "unit": "ns/op",
            "extra": "24023499 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "24023499 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "24023499 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad",
            "value": 14920,
            "unit": "ns/op\t   14569 B/op\t     162 allocs/op",
            "extra": "80725 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - ns/op",
            "value": 14920,
            "unit": "ns/op",
            "extra": "80725 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - B/op",
            "value": 14569,
            "unit": "B/op",
            "extra": "80725 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - allocs/op",
            "value": 162,
            "unit": "allocs/op",
            "extra": "80725 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson",
            "value": 108993,
            "unit": "ns/op\t   15102 B/op\t     369 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - ns/op",
            "value": 108993,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - B/op",
            "value": 15102,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - allocs/op",
            "value": 369,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal",
            "value": 52934,
            "unit": "ns/op\t    9600 B/op\t     208 allocs/op",
            "extra": "22512 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - ns/op",
            "value": 52934,
            "unit": "ns/op",
            "extra": "22512 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - B/op",
            "value": 9600,
            "unit": "B/op",
            "extra": "22512 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - allocs/op",
            "value": 208,
            "unit": "allocs/op",
            "extra": "22512 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal",
            "value": 97380,
            "unit": "ns/op\t   14157 B/op\t     300 allocs/op",
            "extra": "12321 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - ns/op",
            "value": 97380,
            "unit": "ns/op",
            "extra": "12321 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - B/op",
            "value": 14157,
            "unit": "B/op",
            "extra": "12321 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - allocs/op",
            "value": 300,
            "unit": "allocs/op",
            "extra": "12321 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson",
            "value": 58022,
            "unit": "ns/op\t   15112 B/op\t     248 allocs/op",
            "extra": "20635 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - ns/op",
            "value": 58022,
            "unit": "ns/op",
            "extra": "20635 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - B/op",
            "value": 15112,
            "unit": "B/op",
            "extra": "20635 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - allocs/op",
            "value": 248,
            "unit": "allocs/op",
            "extra": "20635 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader",
            "value": 112049,
            "unit": "ns/op\t   15054 B/op\t     368 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - ns/op",
            "value": 112049,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - B/op",
            "value": 15054,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - allocs/op",
            "value": 368,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter",
            "value": 55029,
            "unit": "ns/op\t    9120 B/op\t     239 allocs/op",
            "extra": "21404 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - ns/op",
            "value": 55029,
            "unit": "ns/op",
            "extra": "21404 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - B/op",
            "value": 9120,
            "unit": "B/op",
            "extra": "21404 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - allocs/op",
            "value": 239,
            "unit": "allocs/op",
            "extra": "21404 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader",
            "value": 99160,
            "unit": "ns/op\t   14157 B/op\t     300 allocs/op",
            "extra": "12127 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - ns/op",
            "value": 99160,
            "unit": "ns/op",
            "extra": "12127 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - B/op",
            "value": 14157,
            "unit": "B/op",
            "extra": "12127 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - allocs/op",
            "value": 300,
            "unit": "allocs/op",
            "extra": "12127 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader",
            "value": 57986,
            "unit": "ns/op\t   13080 B/op\t     242 allocs/op",
            "extra": "20545 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - ns/op",
            "value": 57986,
            "unit": "ns/op",
            "extra": "20545 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - B/op",
            "value": 13080,
            "unit": "B/op",
            "extra": "20545 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - allocs/op",
            "value": 242,
            "unit": "allocs/op",
            "extra": "20545 times\n4 procs"
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
          "id": "6b72799818563bb1cb9b3bf33804acfd28ac2e43",
          "message": "Merge pull request #41 from ucan-wg/streamline-did\n\ndid: simplify public API, add missing required algorithms",
          "timestamp": "2024-10-24T17:09:10+02:00",
          "tree_id": "58e1c2c45cf6801639fc31f9b75d00669f7f034d",
          "url": "https://github.com/ucan-wg/go-ucan/commit/6b72799818563bb1cb9b3bf33804acfd28ac2e43"
        },
        "date": 1729782609110,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkContainerSerialisation/car_write",
            "value": 5549,
            "unit": "ns/op\t   17896 B/op\t      58 allocs/op",
            "extra": "213980 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - ns/op",
            "value": 5549,
            "unit": "ns/op",
            "extra": "213980 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - B/op",
            "value": 17896,
            "unit": "B/op",
            "extra": "213980 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "213980 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read",
            "value": 956066,
            "unit": "ns/op\t  149240 B/op\t    2651 allocs/op",
            "extra": "1239 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - ns/op",
            "value": 956066,
            "unit": "ns/op",
            "extra": "1239 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - B/op",
            "value": 149240,
            "unit": "B/op",
            "extra": "1239 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - allocs/op",
            "value": 2651,
            "unit": "allocs/op",
            "extra": "1239 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write",
            "value": 11006,
            "unit": "ns/op\t   24200 B/op\t      60 allocs/op",
            "extra": "107748 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - ns/op",
            "value": 11006,
            "unit": "ns/op",
            "extra": "107748 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - B/op",
            "value": 24200,
            "unit": "B/op",
            "extra": "107748 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - allocs/op",
            "value": 60,
            "unit": "allocs/op",
            "extra": "107748 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read",
            "value": 1021168,
            "unit": "ns/op\t  151303 B/op\t    2653 allocs/op",
            "extra": "1234 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - ns/op",
            "value": 1021168,
            "unit": "ns/op",
            "extra": "1234 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - B/op",
            "value": 151303,
            "unit": "B/op",
            "extra": "1234 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - allocs/op",
            "value": 2653,
            "unit": "allocs/op",
            "extra": "1234 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write",
            "value": 4405,
            "unit": "ns/op\t   16368 B/op\t      25 allocs/op",
            "extra": "268356 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - ns/op",
            "value": 4405,
            "unit": "ns/op",
            "extra": "268356 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - B/op",
            "value": 16368,
            "unit": "B/op",
            "extra": "268356 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - allocs/op",
            "value": 25,
            "unit": "allocs/op",
            "extra": "268356 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read",
            "value": 953307,
            "unit": "ns/op\t  138699 B/op\t    2554 allocs/op",
            "extra": "1260 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - ns/op",
            "value": 953307,
            "unit": "ns/op",
            "extra": "1260 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - B/op",
            "value": 138699,
            "unit": "B/op",
            "extra": "1260 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - allocs/op",
            "value": 2554,
            "unit": "allocs/op",
            "extra": "1260 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write",
            "value": 8565,
            "unit": "ns/op\t   20720 B/op\t      26 allocs/op",
            "extra": "137944 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - ns/op",
            "value": 8565,
            "unit": "ns/op",
            "extra": "137944 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - B/op",
            "value": 20720,
            "unit": "B/op",
            "extra": "137944 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "137944 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read",
            "value": 950713,
            "unit": "ns/op\t  140765 B/op\t    2556 allocs/op",
            "extra": "1242 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - ns/op",
            "value": 950713,
            "unit": "ns/op",
            "extra": "1242 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - B/op",
            "value": 140765,
            "unit": "B/op",
            "extra": "1242 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - allocs/op",
            "value": 2556,
            "unit": "allocs/op",
            "extra": "1242 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob",
            "value": 49.42,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "24464232 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - ns/op",
            "value": 49.42,
            "unit": "ns/op",
            "extra": "24464232 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "24464232 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "24464232 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad",
            "value": 14540,
            "unit": "ns/op\t   14570 B/op\t     162 allocs/op",
            "extra": "79800 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - ns/op",
            "value": 14540,
            "unit": "ns/op",
            "extra": "79800 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - B/op",
            "value": 14570,
            "unit": "B/op",
            "extra": "79800 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - allocs/op",
            "value": 162,
            "unit": "allocs/op",
            "extra": "79800 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson",
            "value": 107116,
            "unit": "ns/op\t   15141 B/op\t     372 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - ns/op",
            "value": 107116,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - B/op",
            "value": 15141,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - allocs/op",
            "value": 372,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal",
            "value": 53403,
            "unit": "ns/op\t    9600 B/op\t     208 allocs/op",
            "extra": "22831 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - ns/op",
            "value": 53403,
            "unit": "ns/op",
            "extra": "22831 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - B/op",
            "value": 9600,
            "unit": "B/op",
            "extra": "22831 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - allocs/op",
            "value": 208,
            "unit": "allocs/op",
            "extra": "22831 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal",
            "value": 95577,
            "unit": "ns/op\t   14197 B/op\t     303 allocs/op",
            "extra": "12555 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - ns/op",
            "value": 95577,
            "unit": "ns/op",
            "extra": "12555 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - B/op",
            "value": 14197,
            "unit": "B/op",
            "extra": "12555 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - allocs/op",
            "value": 303,
            "unit": "allocs/op",
            "extra": "12555 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson",
            "value": 57505,
            "unit": "ns/op\t   15112 B/op\t     248 allocs/op",
            "extra": "20800 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - ns/op",
            "value": 57505,
            "unit": "ns/op",
            "extra": "20800 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - B/op",
            "value": 15112,
            "unit": "B/op",
            "extra": "20800 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - allocs/op",
            "value": 248,
            "unit": "allocs/op",
            "extra": "20800 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader",
            "value": 108297,
            "unit": "ns/op\t   15094 B/op\t     371 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - ns/op",
            "value": 108297,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - B/op",
            "value": 15094,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - allocs/op",
            "value": 371,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter",
            "value": 53944,
            "unit": "ns/op\t    9120 B/op\t     239 allocs/op",
            "extra": "22044 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - ns/op",
            "value": 53944,
            "unit": "ns/op",
            "extra": "22044 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - B/op",
            "value": 9120,
            "unit": "B/op",
            "extra": "22044 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - allocs/op",
            "value": 239,
            "unit": "allocs/op",
            "extra": "22044 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader",
            "value": 97081,
            "unit": "ns/op\t   14197 B/op\t     303 allocs/op",
            "extra": "12289 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - ns/op",
            "value": 97081,
            "unit": "ns/op",
            "extra": "12289 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - B/op",
            "value": 14197,
            "unit": "B/op",
            "extra": "12289 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - allocs/op",
            "value": 303,
            "unit": "allocs/op",
            "extra": "12289 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader",
            "value": 56892,
            "unit": "ns/op\t   13080 B/op\t     242 allocs/op",
            "extra": "20676 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - ns/op",
            "value": 56892,
            "unit": "ns/op",
            "extra": "20676 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - B/op",
            "value": 13080,
            "unit": "B/op",
            "extra": "20676 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - allocs/op",
            "value": 242,
            "unit": "allocs/op",
            "extra": "20676 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "smoyer64@gmail.com",
            "name": "Steve Moyer",
            "username": "smoyer64"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "fcb527cc5264c3bca1d51f3459bceaa0ec5266f6",
          "message": "Merge pull request #50 from ucan-wg/fix/meta-optional-in-delegation\n\nfix: change meta optional in `delegation` `Token`, model and schema",
          "timestamp": "2024-10-24T18:10:30-04:00",
          "tree_id": "128e19731ee69bd3b3a8d0a9873e37817f94a956",
          "url": "https://github.com/ucan-wg/go-ucan/commit/fcb527cc5264c3bca1d51f3459bceaa0ec5266f6"
        },
        "date": 1729807875014,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkContainerSerialisation/car_write",
            "value": 5685,
            "unit": "ns/op\t   17896 B/op\t      58 allocs/op",
            "extra": "197557 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - ns/op",
            "value": 5685,
            "unit": "ns/op",
            "extra": "197557 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - B/op",
            "value": 17896,
            "unit": "B/op",
            "extra": "197557 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_write - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "197557 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read",
            "value": 1000665,
            "unit": "ns/op\t  147963 B/op\t    2651 allocs/op",
            "extra": "1183 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - ns/op",
            "value": 1000665,
            "unit": "ns/op",
            "extra": "1183 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - B/op",
            "value": 147963,
            "unit": "B/op",
            "extra": "1183 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/car_read - allocs/op",
            "value": 2651,
            "unit": "allocs/op",
            "extra": "1183 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write",
            "value": 12179,
            "unit": "ns/op\t   24200 B/op\t      60 allocs/op",
            "extra": "100128 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - ns/op",
            "value": 12179,
            "unit": "ns/op",
            "extra": "100128 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - B/op",
            "value": 24200,
            "unit": "B/op",
            "extra": "100128 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_write - allocs/op",
            "value": 60,
            "unit": "allocs/op",
            "extra": "100128 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read",
            "value": 1028881,
            "unit": "ns/op\t  150024 B/op\t    2653 allocs/op",
            "extra": "1106 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - ns/op",
            "value": 1028881,
            "unit": "ns/op",
            "extra": "1106 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - B/op",
            "value": 150024,
            "unit": "B/op",
            "extra": "1106 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/carBase64_read - allocs/op",
            "value": 2653,
            "unit": "allocs/op",
            "extra": "1106 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write",
            "value": 4947,
            "unit": "ns/op\t   16368 B/op\t      25 allocs/op",
            "extra": "254372 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - ns/op",
            "value": 4947,
            "unit": "ns/op",
            "extra": "254372 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - B/op",
            "value": 16368,
            "unit": "B/op",
            "extra": "254372 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_write - allocs/op",
            "value": 25,
            "unit": "allocs/op",
            "extra": "254372 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read",
            "value": 983043,
            "unit": "ns/op\t  137421 B/op\t    2554 allocs/op",
            "extra": "1197 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - ns/op",
            "value": 983043,
            "unit": "ns/op",
            "extra": "1197 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - B/op",
            "value": 137421,
            "unit": "B/op",
            "extra": "1197 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cbor_read - allocs/op",
            "value": 2554,
            "unit": "allocs/op",
            "extra": "1197 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write",
            "value": 8826,
            "unit": "ns/op\t   20720 B/op\t      26 allocs/op",
            "extra": "135708 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - ns/op",
            "value": 8826,
            "unit": "ns/op",
            "extra": "135708 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - B/op",
            "value": 20720,
            "unit": "B/op",
            "extra": "135708 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_write - allocs/op",
            "value": 26,
            "unit": "allocs/op",
            "extra": "135708 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read",
            "value": 996012,
            "unit": "ns/op\t  139487 B/op\t    2556 allocs/op",
            "extra": "1194 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - ns/op",
            "value": 996012,
            "unit": "ns/op",
            "extra": "1194 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - B/op",
            "value": 139487,
            "unit": "B/op",
            "extra": "1194 times\n4 procs"
          },
          {
            "name": "BenchmarkContainerSerialisation/cborBase64_read - allocs/op",
            "value": 2556,
            "unit": "allocs/op",
            "extra": "1194 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob",
            "value": 49.82,
            "unit": "ns/op\t       0 B/op\t       0 allocs/op",
            "extra": "24415770 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - ns/op",
            "value": 49.82,
            "unit": "ns/op",
            "extra": "24415770 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - B/op",
            "value": 0,
            "unit": "B/op",
            "extra": "24415770 times\n4 procs"
          },
          {
            "name": "BenchmarkGlob - allocs/op",
            "value": 0,
            "unit": "allocs/op",
            "extra": "24415770 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad",
            "value": 15286,
            "unit": "ns/op\t   14577 B/op\t     163 allocs/op",
            "extra": "77932 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - ns/op",
            "value": 15286,
            "unit": "ns/op",
            "extra": "77932 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - B/op",
            "value": 14577,
            "unit": "B/op",
            "extra": "77932 times\n4 procs"
          },
          {
            "name": "BenchmarkSchemaLoad - allocs/op",
            "value": 163,
            "unit": "allocs/op",
            "extra": "77932 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson",
            "value": 112039,
            "unit": "ns/op\t   15014 B/op\t     372 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - ns/op",
            "value": 112039,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - B/op",
            "value": 15014,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/FromDagJson - allocs/op",
            "value": 372,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal",
            "value": 54962,
            "unit": "ns/op\t    9584 B/op\t     208 allocs/op",
            "extra": "21936 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - ns/op",
            "value": 54962,
            "unit": "ns/op",
            "extra": "21936 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - B/op",
            "value": 9584,
            "unit": "B/op",
            "extra": "21936 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Seal - allocs/op",
            "value": 208,
            "unit": "allocs/op",
            "extra": "21936 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal",
            "value": 102751,
            "unit": "ns/op\t   14069 B/op\t     303 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - ns/op",
            "value": 102751,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - B/op",
            "value": 14069,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/Unseal - allocs/op",
            "value": 303,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson",
            "value": 61268,
            "unit": "ns/op\t   15096 B/op\t     248 allocs/op",
            "extra": "19262 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - ns/op",
            "value": 61268,
            "unit": "ns/op",
            "extra": "19262 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - B/op",
            "value": 15096,
            "unit": "B/op",
            "extra": "19262 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_buffers/ToDagJson - allocs/op",
            "value": 248,
            "unit": "allocs/op",
            "extra": "19262 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader",
            "value": 112745,
            "unit": "ns/op\t   14966 B/op\t     371 allocs/op",
            "extra": "9945 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - ns/op",
            "value": 112745,
            "unit": "ns/op",
            "extra": "9945 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - B/op",
            "value": 14966,
            "unit": "B/op",
            "extra": "9945 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/FromDagJsonReader - allocs/op",
            "value": 371,
            "unit": "allocs/op",
            "extra": "9945 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter",
            "value": 56948,
            "unit": "ns/op\t    9104 B/op\t     239 allocs/op",
            "extra": "21054 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - ns/op",
            "value": 56948,
            "unit": "ns/op",
            "extra": "21054 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - B/op",
            "value": 9104,
            "unit": "B/op",
            "extra": "21054 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/SealWriter - allocs/op",
            "value": 239,
            "unit": "allocs/op",
            "extra": "21054 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader",
            "value": 102304,
            "unit": "ns/op\t   14069 B/op\t     303 allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - ns/op",
            "value": 102304,
            "unit": "ns/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - B/op",
            "value": 14069,
            "unit": "B/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/UnsealReader - allocs/op",
            "value": 303,
            "unit": "allocs/op",
            "extra": "10000 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader",
            "value": 59726,
            "unit": "ns/op\t   13064 B/op\t     242 allocs/op",
            "extra": "19818 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - ns/op",
            "value": 59726,
            "unit": "ns/op",
            "extra": "19818 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - B/op",
            "value": 13064,
            "unit": "B/op",
            "extra": "19818 times\n4 procs"
          },
          {
            "name": "BenchmarkRoundTrip/via_streaming/ToDagJsonReader - allocs/op",
            "value": 242,
            "unit": "allocs/op",
            "extra": "19818 times\n4 procs"
          }
        ]
      }
    ]
  }
}