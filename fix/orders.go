package fix

// Orders is a sample of backup orders sorted by unix nano seconds. All orders
// are grouped together by their buffer length of 5 seconds each. The 5 seconds
// buffer length is used for unit tests in the typ/orders/buffer package.
const Orders = `
{
  "2022-05-02T00:00:00Z": {
    "AK": [
      {
        "PR": 1325.8,
        "SI": 9.55
      },
      {
        "PR": 1326,
        "SI": 2
      }
    ],
    "BD": [
      {
        "PR": 0.2,
        "SI": 10
      },
      {
        "PR": 0.7,
        "SI": 1
      }
    ],
    "TS": {
      "seconds": 1651449602,
      "nanos": 716740000
    }
  },
  "2022-05-02T00:00:05Z": {
    "AK": [
      {
        "PR": 1327.3,
        "SI": 11.983
      },
      {
        "PR": 1327.4,
        "SI": 29.483
      },
      {
        "PR": 1327.5,
        "SI": 34.456
      },
      {
        "PR": 1327.6,
        "SI": 35.963
      }
    ],
    "BD": [
      {
        "PR": 0.2,
        "SI": 10
      },
      {
        "PR": 0.7,
        "SI": 1
      },
      {
        "PR": 1,
        "SI": 505
      },
      {
        "PR": 35,
        "SI": 60
      }
    ],
    "TS": {
      "seconds": 1651449608,
      "nanos": 101664000
    }
  },
  "2022-05-02T00:00:10Z": {
    "AK": [
      {
        "PR": 1327,
        "SI": 15.494
      }
    ],
    "BD": [
      {
        "PR": 0.2,
        "SI": 10
      }
    ],
    "TS": {
      "seconds": 1651449612,
      "nanos": 403523000
    }
  },
  "2022-05-02T00:00:15Z": {
    "AK": [
      {
        "PR": 1325.7,
        "SI": 7.395
      },
      {
        "PR": 1325.8,
        "SI": 64.175
      },
      {
        "PR": 1325.9,
        "SI": 48.386
      },
      {
        "PR": 1326,
        "SI": 49.207
      },
      {
        "PR": 1326.1,
        "SI": 63.457
      },
      {
        "PR": 1326.2,
        "SI": 64.019
      }
    ],
    "BD": [
      {
        "PR": 0.2,
        "SI": 10
      },
      {
        "PR": 0.7,
        "SI": 1
      },
      {
        "PR": 1,
        "SI": 505
      },
      {
        "PR": 35,
        "SI": 60
      },
      {
        "PR": 98,
        "SI": 486
      },
      {
        "PR": 100,
        "SI": 0.01
      }
    ],
    "TS": {
      "seconds": 1651449616,
      "nanos": 725104000
    }
  },
  "2022-05-02T00:00:20Z": {
    "AK": [
      {
        "PR": 1324.7,
        "SI": 24.167
      },
      {
        "PR": 1324.8,
        "SI": 53.255
      },
      {
        "PR": 1324.9,
        "SI": 39.477
      },
      {
        "PR": 1325,
        "SI": 66.192
      }
    ],
    "BD": [
      {
        "PR": 0.2,
        "SI": 10
      },
      {
        "PR": 0.7,
        "SI": 1
      },
      {
        "PR": 1,
        "SI": 505
      },
      {
        "PR": 35,
        "SI": 60
      }
    ],
    "TS": {
      "seconds": 1651449623,
      "nanos": 76650000
    }
  },
  "2022-05-02T00:00:25Z": {
    "AK": [
      {
        "PR": 1325.4,
        "SI": 29.359
      },
      {
        "PR": 1325.5,
        "SI": 50.188
      },
      {
        "PR": 1325.6,
        "SI": 74.527
      }
    ],
    "BD": [
      {
        "PR": 0.2,
        "SI": 10
      },
      {
        "PR": 0.7,
        "SI": 1
      },
      {
        "PR": 1,
        "SI": 505
      }
    ],
    "TS": {
      "seconds": 1651449629,
      "nanos": 389851000
    }
  }
}
`
