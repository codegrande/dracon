package cyclonedx

import (
	"encoding/json"
	"testing"

	v1 "github.com/ocurity/dracon/api/proto/v1"
	"github.com/stretchr/testify/assert"
)

func TestToDraconLibrary(t *testing.T) {
	input := libraryBom
	issues, err := ToDracon([]byte(input), "json")
	assert.Nil(t, err)
	expectedIssues := []*v1.Issue{
		{
			Target:        "pkg:npm/juice-shop@11.1.2",
			Type:          "SBOM",
			Title:         "SBOM for pkg:npm/juice-shop@11.1.2",
			Severity:      v1.Severity_SEVERITY_INFO,
			CycloneDXSBOM: &input,
		},
	}
	assert.Equal(t, expectedIssues[0].Target, issues[0].Target)
	assert.Equal(t, expectedIssues[0].Type, issues[0].Type)
	assert.Equal(t, expectedIssues[0].Title, issues[0].Title)
	assert.Equal(t, expectedIssues[0].Severity, issues[0].Severity)
	var sbom1, sbom2 map[string]interface{}
	json.Unmarshal([]byte(*expectedIssues[0].CycloneDXSBOM), &sbom1)
	json.Unmarshal([]byte(*issues[0].CycloneDXSBOM), &sbom2)
	assert.Equal(t, sbom1, sbom2)
}

func TestToDraconSaaSInfra(t *testing.T) {
	input := saasBOM
	issues, err := ToDracon([]byte(input), "json")
	assert.Nil(t, err)
	expectedIssues := []*v1.Issue{
		{
			Target:        "acme-application",
			Type:          "SBOM",
			Title:         "SBOM for acme-application",
			Severity:      v1.Severity_SEVERITY_INFO,
			CycloneDXSBOM: &input,
		},
	}
	assert.Equal(t, expectedIssues[0].Target, issues[0].Target)
	assert.Equal(t, expectedIssues[0].Type, issues[0].Type)
	assert.Equal(t, expectedIssues[0].Title, issues[0].Title)
	assert.Equal(t, expectedIssues[0].Severity, issues[0].Severity)
	var sbom1, sbom2 map[string]interface{}
	json.Unmarshal([]byte(*expectedIssues[0].CycloneDXSBOM), &sbom1)
	json.Unmarshal([]byte(*issues[0].CycloneDXSBOM), &sbom2)
	assert.Equal(t, sbom1, sbom2)
}

const saasBOM = `{
	"bomFormat": "CycloneDX",
	"specVersion": "1.4",
	"serialNumber": "urn:uuid:3e671687-395b-41f5-a30f-a58921a69b79",
	"version": 1,
	"metadata": {
	  "timestamp": "2021-01-10T12:00:00Z",
	  "component": {
		"bom-ref": "acme-application",
		"type": "application",
		"name": "Acme Cloud Example",
		"version": "2022-1"
	  }
	},
	"services": [
	  {
		"bom-ref": "api-gateway",
		"provider": {
		  "name": "Acme Inc",
		  "url": [ "https://example.com" ]
		},
		"group": "com.example",
		"name": "API Gateway",
		"version": "2022-1",
		"description": "Example API Gateway",
		"endpoints": [
		  "https://example.com/",
		  "https://example.com/app"
		],
		"authenticated": false,
		"x-trust-boundary": true,
		"data": [
		  {
			"classification": "PII",
			"flow": "bi-directional"
		  },
		  {
			"classification": "PIFI",
			"flow": "bi-directional"
		  },
		  {
			"classification": "Public",
			"flow": "bi-directional"
		  }
		],
		"externalReferences": [
		  {
			"type": "documentation",
			"url": "http://example.com/app/swagger"
		  }
		],
		"services": [
		  {
			"bom-ref": "ms-1.example.com",
			"provider": {
			  "name": "Acme Inc",
			  "url": [ "https://example.com" ]
			},
			"group": "com.example",
			"name": "Microservice 1",
			"version": "2022-1",
			"description": "Example Microservice",
			"endpoints": [
			  "https://ms-1.example.com"
			],
			"authenticated": true,
			"x-trust-boundary": false,
			"data": [
			  {
				"classification": "PII",
				"flow": "bi-directional"
			  }
			],
			"externalReferences": [
			  {
				"type": "documentation",
				"url": "https://ms-1.example.com/swagger"
			  }
			]
		  },
		  {
			"bom-ref": "ms-2.example.com",
			"provider": {
			  "name": "Acme Inc",
			  "url": [ "https://example.com" ]
			},
			"group": "com.example",
			"name": "Microservice 2",
			"version": "2022-1",
			"description": "Example Microservice",
			"endpoints": [
			  "https://ms-2.example.com"
			],
			"authenticated": true,
			"x-trust-boundary": false,
			"data": [
			  {
				"classification": "PIFI",
				"flow": "bi-directional"
			  }
			],
			"externalReferences": [
			  {
				"type": "documentation",
				"url": "https://ms-2.example.com/swagger"
			  }
			]
		  },
		  {
			"bom-ref": "ms-3.example.com",
			"provider": {
			  "name": "Acme Inc",
			  "url": [ "https://example.com" ]
			},
			"group": "com.example",
			"name": "Microservice 3",
			"version": "2022-1",
			"description": "Example Microservice",
			"endpoints": [
			  "https://ms-3.example.com"
			],
			"authenticated": true,
			"x-trust-boundary": false,
			"data": [
			  {
				"classification": "Public",
				"flow": "bi-directional"
			  }
			],
			"externalReferences": [
			  {
				"type": "documentation",
				"url": "https://ms-3.example.com/swagger"
			  }
			]
		  },
		  {
			"bom-ref": "ms-1-pgsql.example.com",
			"group": "org.postgresql",
			"name": "Postgres",
			"version": "14.1",
			"description": "Postgres database for Microservice #1",
			"endpoints": [
			  "https://ms-1-pgsql.example.com:5432"
			],
			"authenticated": true,
			"x-trust-boundary": false,
			"data": [
			  {
				"classification": "PII",
				"flow": "bi-directional"
			  }
			]
		  },
		  {
			"bom-ref": "s3-example.amazon.com",
			"group": "com.amazon",
			"name": "S3",
			"description": "S3 bucket",
			"endpoints": [
			  "https://s3-example.amazon.com"
			],
			"authenticated": true,
			"x-trust-boundary": true,
			"data": [
			  {
				"classification": "Public",
				"flow": "bi-directional"
			  }
			]
		  }
		]
	  }
	],
	"dependencies": [
	  {
		"ref": "acme-application",
		"dependsOn": [ "api-gateway" ]
	  },
	  {
		"ref": "api-gateway",
		"dependsOn": [
		  "ms-1.example.com",
		  "ms-2.example.com",
		  "ms-3.example.com"
		]
	  },
	  {
		"ref": "ms-1.example.com",
		"dependsOn": [ "ms-1-pgsql.example.com" ]
	  },
	  {
		"ref": "ms-2.example.com",
		"dependsOn": [ ]
	  },
	  {
		"ref": "ms-3.example.com",
		"dependsOn": [ "s3-example.amazon.com" ]
	  }
	]
  }
  `

const libraryBom = `{
	"bomFormat": "CycloneDX",
	"specVersion": "1.2",
	"serialNumber": "urn:uuid:1f860713-54b9-4253-ba5a-9554851904af",
	"version": 1,
	"metadata": {
	  "timestamp": "2020-08-03T03:20:53.771Z",
	  "tools": [
		{
		  "vendor": "CycloneDX",
		  "name": "Node.js module",
		  "version": "2.0.0"
		}
	  ],
	  "component": {
		"type": "library",
		"bom-ref": "pkg:npm/juice-shop@11.1.2",
		"name": "juice-shop",
		"version": "11.1.2",
		"description": "Probably the most modern and sophisticated insecure web application",
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/juice-shop@11.1.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://owasp-juice.shop"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/bkimminich/juice-shop/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/bkimminich/juice-shop.git"
		  }
		]
	  }
	},
	"components": [
	  {
		"type": "library",
		"bom-ref": "pkg:npm/body-parser@1.19.0",
		"name": "body-parser",
		"version": "1.19.0",
		"description": "Node.js body parsing middleware",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "96b2709e57c9c4e09a6fd66a8fd979844f69f08a"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/body-parser@1.19.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/expressjs/body-parser#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/expressjs/body-parser/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/expressjs/body-parser.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/bytes@3.1.0",
		"name": "bytes",
		"version": "3.1.0",
		"description": "Utility to parse a string bytes to bytes and vice-versa",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "f6cf7933a360e0588fa9fde85651cdc7f805d1f6"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/bytes@3.1.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/visionmedia/bytes.js#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/visionmedia/bytes.js/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/visionmedia/bytes.js.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/content-type@1.0.4",
		"name": "content-type",
		"version": "1.0.4",
		"description": "Create and parse HTTP Content-Type header",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "e138cc75e040c727b1966fe5e5f8c9aee256fe3b"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/content-type@1.0.4",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jshttp/content-type#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jshttp/content-type/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jshttp/content-type.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/debug@2.6.9",
		"name": "debug",
		"version": "2.6.9",
		"description": "small debugging utility",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "5d128515df134ff327e90a4c93f4e077a536341f"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/debug@2.6.9",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/visionmedia/debug#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/visionmedia/debug/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/visionmedia/debug.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/ms@2.0.0",
		"name": "ms",
		"version": "2.0.0",
		"description": "Tiny milisecond conversion utility",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "5608aeadfc00be6c2901df5f9861788de0d597c8"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/ms@2.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/zeit/ms#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/zeit/ms/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/zeit/ms.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/depd@1.1.2",
		"name": "depd",
		"version": "1.1.2",
		"description": "Deprecate all the things",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "9bcd52e14c097763e749b274c4346ed2e560b5a9"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/depd@1.1.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/dougwilson/nodejs-depd#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/dougwilson/nodejs-depd/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/dougwilson/nodejs-depd.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/http-errors@1.7.2",
		"name": "http-errors",
		"version": "1.7.2",
		"description": "Create HTTP error objects",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "4f5029cf13239f31036e5b2e55292bcfbcc85c8f"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/http-errors@1.7.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jshttp/http-errors#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jshttp/http-errors/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jshttp/http-errors.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/inherits@2.0.3",
		"name": "inherits",
		"version": "2.0.3",
		"description": "Browser-friendly inheritance fully compatible with standard node.js inherits()",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "633c2c83e3da42a502f52466022480f4208261de"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "ISC"
			}
		  }
		],
		"purl": "pkg:npm/inherits@2.0.3",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/isaacs/inherits#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/isaacs/inherits/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/isaacs/inherits.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/setprototypeof@1.1.1",
		"name": "setprototypeof",
		"version": "1.1.1",
		"description": "A small polyfill for Object.setprototypeof",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "7e95acb24aa92f5885e0abef5ba131330d4ae683"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "ISC"
			}
		  }
		],
		"purl": "pkg:npm/setprototypeof@1.1.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/wesleytodd/setprototypeof"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/wesleytodd/setprototypeof/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/wesleytodd/setprototypeof.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/statuses@1.5.0",
		"name": "statuses",
		"version": "1.5.0",
		"description": "HTTP status utility",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "161c7dac177659fd9811f43771fa99381478628c"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/statuses@1.5.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jshttp/statuses#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jshttp/statuses/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jshttp/statuses.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/toidentifier@1.0.0",
		"name": "toidentifier",
		"version": "1.0.0",
		"description": "Convert a string of words to a JavaScript identifier",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "7e1be3470f1e77948bc43d94a3c8f4d7752ba553"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/toidentifier@1.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/component/toidentifier#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/component/toidentifier/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/component/toidentifier.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/iconv-lite@0.4.24",
		"name": "iconv-lite",
		"version": "0.4.24",
		"description": "Convert character encodings in pure javascript.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "2022b4b25fbddc21d2f524974a474aafe733908b"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/iconv-lite@0.4.24",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/ashtuchkin/iconv-lite"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/ashtuchkin/iconv-lite/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/ashtuchkin/iconv-lite.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/safer-buffer@2.1.2",
		"name": "safer-buffer",
		"version": "2.1.2",
		"description": "Modern Buffer API polyfill without footguns",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "44fa161b0187b9549dd84bb91802f9bd8385cd6a"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/safer-buffer@2.1.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/ChALkeR/safer-buffer#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/ChALkeR/safer-buffer/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/ChALkeR/safer-buffer.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/on-finished@2.3.0",
		"name": "on-finished",
		"version": "2.3.0",
		"description": "Execute a callback when a request closes, finishes, or errors",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "20f1336481b083cd75337992a16971aa2d906947"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/on-finished@2.3.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jshttp/on-finished#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jshttp/on-finished/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jshttp/on-finished.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/ee-first@1.1.1",
		"name": "ee-first",
		"version": "1.1.1",
		"description": "return the first event in a set of ee/event pairs",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "590c61156b0ae2f4f0255732a158b266bc56b21d"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/ee-first@1.1.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonathanong/ee-first#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonathanong/ee-first/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonathanong/ee-first.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/qs@6.7.0",
		"name": "qs",
		"version": "6.7.0",
		"description": "A querystring parser that supports nesting and arrays, with a depth limit",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "41dc1a015e3d581f1621776be31afb2876a9b1bc"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "BSD-3-Clause"
			}
		  }
		],
		"purl": "pkg:npm/qs@6.7.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/ljharb/qs"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/ljharb/qs/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/ljharb/qs.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/raw-body@2.4.0",
		"name": "raw-body",
		"version": "2.4.0",
		"description": "Get and validate the raw body of a readable stream.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "a1ce6fb9c9bc356ca52e89256ab59059e13d0332"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/raw-body@2.4.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/stream-utils/raw-body#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/stream-utils/raw-body/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/stream-utils/raw-body.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/unpipe@1.0.0",
		"name": "unpipe",
		"version": "1.0.0",
		"description": "Unpipe a stream from all destinations",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "b2bf4ee8514aae6165b4817829d21b2ef49904ec"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/unpipe@1.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/stream-utils/unpipe#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/stream-utils/unpipe/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/stream-utils/unpipe.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/type-is@1.6.18",
		"name": "type-is",
		"version": "1.6.18",
		"description": "Infer the content-type of a request.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "4e552cd05df09467dcbc4ef739de89f2cf37c131"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/type-is@1.6.18",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jshttp/type-is#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jshttp/type-is/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jshttp/type-is.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/media-typer@0.3.0",
		"name": "media-typer",
		"version": "0.3.0",
		"description": "Simple RFC 6838 media type parser and formatter",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "8710d7af0aa626f8fffa1ce00168545263255748"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/media-typer@0.3.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jshttp/media-typer#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jshttp/media-typer/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jshttp/media-typer.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/mime-types@2.1.27",
		"name": "mime-types",
		"version": "2.1.27",
		"description": "The ultimate javascript content-type utility.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "47949f98e279ea53119f5722e0f34e529bec009f"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/mime-types@2.1.27",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jshttp/mime-types#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jshttp/mime-types/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jshttp/mime-types.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/mime-db@1.44.0",
		"name": "mime-db",
		"version": "1.44.0",
		"description": "Media Type Database",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "fa11c5eb0aca1334b4233cb4d52f10c5a6272f92"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/mime-db@1.44.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jshttp/mime-db#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jshttp/mime-db/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jshttp/mime-db.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/check-dependencies@1.1.0",
		"name": "check-dependencies",
		"version": "1.1.0",
		"description": "Checks if currently installed npm/bower dependencies are installed in the exact same versions that are specified in package.json/bower.json",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "3aa2df4061770179d8e88e8bf9315c53722ddff4"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/check-dependencies@1.1.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/mgol/check-dependencies"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/mgol/check-dependencies/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/mgol/check-dependencies.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/bower-config@1.4.3",
		"name": "bower-config",
		"version": "1.4.3",
		"description": "The Bower config reader and writer.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "3454fecdc5f08e7aa9cc6d556e492be0669689ae"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/bower-config@1.4.3",
		"externalReferences": [
		  {
			"type": "website",
			"url": "http://bower.io"
		  },
		  {
			"type": "vcs",
			"url": "https://github.com/bower/bower/tree/master/packages/bower-config"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/graceful-fs@4.2.4",
		"name": "graceful-fs",
		"version": "4.2.4",
		"description": "A drop-in replacement for fs, making various improvements.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "2256bde14d3632958c465ebc96dc467ca07a29fb"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "ISC"
			}
		  }
		],
		"purl": "pkg:npm/graceful-fs@4.2.4",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/isaacs/node-graceful-fs#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/isaacs/node-graceful-fs/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/isaacs/node-graceful-fs.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/minimist@0.2.1",
		"name": "minimist",
		"version": "0.2.1",
		"description": "parse argument options",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "827ba4e7593464e7c221e8c5bed930904ee2c455"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/minimist@0.2.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/substack/minimist"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/substack/minimist/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/substack/minimist.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/mout@1.2.2",
		"name": "mout",
		"version": "1.2.2",
		"description": "Modular Utilities",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "c9b718a499806a0632cede178e80f436259e777d"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/mout@1.2.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "http://moutjs.com/"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/mout/mout/issues/"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/mout/mout.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/osenv@0.1.5",
		"name": "osenv",
		"version": "0.1.5",
		"description": "Look up environment settings specific to different operating systems",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "85cdfafaeb28e8677f416e287592b5f3f49ea410"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "ISC"
			}
		  }
		],
		"purl": "pkg:npm/osenv@0.1.5",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/npm/osenv#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/npm/osenv/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/npm/osenv.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/os-homedir@1.0.2",
		"name": "os-homedir",
		"version": "1.0.2",
		"description": "Node.js 4 os.homedir() ponyfill",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "ffbc4988336e0e833de0c168c7ef152121aa7fb3"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/os-homedir@1.0.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/sindresorhus/os-homedir#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/sindresorhus/os-homedir/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/sindresorhus/os-homedir.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/os-tmpdir@1.0.2",
		"name": "os-tmpdir",
		"version": "1.0.2",
		"description": "Node.js os.tmpdir() ponyfill",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "bbe67406c79aa85c5cfec766fe5734555dfa1274"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/os-tmpdir@1.0.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/sindresorhus/os-tmpdir#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/sindresorhus/os-tmpdir/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/sindresorhus/os-tmpdir.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/untildify@2.1.0",
		"name": "untildify",
		"version": "2.1.0",
		"description": "Convert a tilde path to an absolute path: ~/dev => /Users/sindresorhus/dev",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "17eb2807987f76952e9c0485fc311d06a826a2e0"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/untildify@2.1.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/sindresorhus/untildify#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/sindresorhus/untildify/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/sindresorhus/untildify.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/wordwrap@0.0.3",
		"name": "wordwrap",
		"version": "0.0.3",
		"description": "Wrap those words. Show them at what columns to start and stop.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "a3d5da6cd5c0bc0008d37234bbaf1bed63059107"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/wordwrap@0.0.3",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/substack/node-wordwrap#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/substack/node-wordwrap/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/substack/node-wordwrap.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/chalk@2.4.2",
		"name": "chalk",
		"version": "2.4.2",
		"description": "Terminal string styling done right",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "cd42541677a54333cf541a49108c1432b44c9424"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/chalk@2.4.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/chalk/chalk#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/chalk/chalk/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/chalk/chalk.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/ansi-styles@3.2.1",
		"name": "ansi-styles",
		"version": "3.2.1",
		"description": "ANSI escape codes for styling strings in the terminal",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "41fbb20243e50b12be0f04b8dedbf07520ce841d"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/ansi-styles@3.2.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/chalk/ansi-styles#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/chalk/ansi-styles/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/chalk/ansi-styles.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/color-convert@1.9.3",
		"name": "color-convert",
		"version": "1.9.3",
		"description": "Plain color conversion functions",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "bb71850690e1f136567de629d2d5471deda4c1e8"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/color-convert@1.9.3",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/Qix-/color-convert#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/Qix-/color-convert/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/Qix-/color-convert.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/color-name@1.1.3",
		"name": "color-name",
		"version": "1.1.3",
		"description": "A list of color names and its values",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "a7d0558bd89c42f795dd42328f740831ca53bc25"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/color-name@1.1.3",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/dfcreative/color-name"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/dfcreative/color-name/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+ssh://git@github.com/dfcreative/color-name.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/escape-string-regexp@1.0.5",
		"name": "escape-string-regexp",
		"version": "1.0.5",
		"description": "Escape RegExp special characters",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "1b61c0562190a8dff6ae3bb2cf0200ca130b86d4"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/escape-string-regexp@1.0.5",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/sindresorhus/escape-string-regexp#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/sindresorhus/escape-string-regexp/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/sindresorhus/escape-string-regexp.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/supports-color@5.5.0",
		"name": "supports-color",
		"version": "5.5.0",
		"description": "Detect whether a terminal supports color",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "e2e69a44ac8772f78a1ec0b35b689df6530efc8f"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/supports-color@5.5.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/chalk/supports-color#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/chalk/supports-color/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/chalk/supports-color.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/has-flag@3.0.0",
		"name": "has-flag",
		"version": "3.0.0",
		"description": "Check if argv has a specific flag",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "b5d454dc2199ae225699f3467e5a07f3b955bafd"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/has-flag@3.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/sindresorhus/has-flag#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/sindresorhus/has-flag/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/sindresorhus/has-flag.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/findup-sync@2.0.0",
		"name": "findup-sync",
		"version": "2.0.0",
		"description": "Find the first file matching a given pattern in the current directory or the nearest ancestor directory.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "9326b1488c22d1a6088650a86901b2d9a90a2cbc"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/findup-sync@2.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/js-cli/node-findup-sync#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/js-cli/node-findup-sync/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/js-cli/node-findup-sync.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/detect-file@1.0.0",
		"name": "detect-file",
		"version": "1.0.0",
		"description": "Detects if a file exists and returns the resolved filepath.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "f0d66d03672a825cb1b73bdb3fe62310c8e552b7"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/detect-file@1.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/doowb/detect-file"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/doowb/detect-file/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/doowb/detect-file.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/is-glob@3.1.0",
		"name": "is-glob",
		"version": "3.1.0",
		"description": "Returns true if the given string looks like a glob pattern or an extglob pattern. This makes it easy to create code that only uses external modules like node-glob when necessary, resulting in much faster code execution and initialization time, and a better user experience.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "7ba5ae24217804ac70707b96922567486cc3e84a"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/is-glob@3.1.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/is-glob"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/is-glob/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/is-glob.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/is-extglob@2.1.1",
		"name": "is-extglob",
		"version": "2.1.1",
		"description": "Returns true if a string has an extglob.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "a88c02535791f02ed37c76a1b9ea9773c833f8c2"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/is-extglob@2.1.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/is-extglob"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/is-extglob/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/is-extglob.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/micromatch@3.1.10",
		"name": "micromatch",
		"version": "3.1.10",
		"description": "Glob matching for javascript/node.js. A drop-in replacement and faster alternative to minimatch and multimatch.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "70859bc95c9840952f359a068a3fc49f9ecfac23"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/micromatch@3.1.10",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/micromatch/micromatch"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/micromatch/micromatch/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/micromatch/micromatch.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/arr-diff@4.0.0",
		"name": "arr-diff",
		"version": "4.0.0",
		"description": "Returns an array with only the unique values from the first array, by excluding all values from additional arrays using strict equality for comparisons.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "d6461074febfec71e7e15235761a329a5dc7c520"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/arr-diff@4.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/arr-diff"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/arr-diff/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/arr-diff.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/array-unique@0.3.2",
		"name": "array-unique",
		"version": "0.3.2",
		"description": "Remove duplicate values from an array. Fastest ES5 implementation.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "a894b75d4bc4f6cd679ef3244a9fd8f46ae2d428"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/array-unique@0.3.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/array-unique"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/array-unique/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/array-unique.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/braces@2.3.2",
		"name": "braces",
		"version": "2.3.2",
		"description": "Bash-like brace expansion, implemented in JavaScript. Safer than other brace expansion libs, with complete support for the Bash 4.3 braces specification, without sacrificing speed.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "5979fd3f14cd531565e5fa2df1abfff1dfaee729"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/braces@2.3.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/micromatch/braces"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/micromatch/braces/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/micromatch/braces.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/arr-flatten@1.1.0",
		"name": "arr-flatten",
		"version": "1.1.0",
		"description": "Recursively flatten an array or arrays.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "36048bbff4e7b47e136644316c99669ea5ae91f1"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/arr-flatten@1.1.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/arr-flatten"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/arr-flatten/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/arr-flatten.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/extend-shallow@2.0.1",
		"name": "extend-shallow",
		"version": "2.0.1",
		"description": "Extend an object with the properties of additional objects. node.js/javascript util.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "51af7d614ad9a9f610ea1bafbb989d6b1c56890f"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/extend-shallow@2.0.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/extend-shallow"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/extend-shallow/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/extend-shallow.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/is-extendable@0.1.1",
		"name": "is-extendable",
		"version": "0.1.1",
		"description": "Returns true if a value is any of the object types: array, regexp, plain object, function or date. This is useful for determining if a value can be extended, e.g. \"can the value have keys?\"",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "62b110e289a471418e3ec36a617d472e301dfc89"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/is-extendable@0.1.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/is-extendable"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/is-extendable/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/is-extendable.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/fill-range@4.0.0",
		"name": "fill-range",
		"version": "4.0.0",
		"description": "Fill in a range of numbers or letters, optionally passing an increment or step to use, or create a regex-compatible range with options.toRegex",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "d544811d428f98eb06a63dc402d2403c328c38f7"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/fill-range@4.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/fill-range"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/fill-range/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/fill-range.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/is-number@3.0.0",
		"name": "is-number",
		"version": "3.0.0",
		"description": "Returns true if the value is a number. comprehensive tests.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "24fd6201a4782cf50561c810276afc7d12d71195"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/is-number@3.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/is-number"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/is-number/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/is-number.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/kind-of@3.2.2",
		"name": "kind-of",
		"version": "3.2.2",
		"description": "Get the native type of a value.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "31ea21a734bab9bbb0f32466d893aea51e4a3c64"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/kind-of@3.2.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/kind-of"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/kind-of/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/kind-of.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/is-buffer@1.1.6",
		"name": "is-buffer",
		"version": "1.1.6",
		"description": "Determine if an object is a Buffer",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "efaa2ea9daa0d7ab2ea13a97b2b8ad51fefbe8be"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/is-buffer@1.1.6",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/feross/is-buffer#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/feross/is-buffer/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/feross/is-buffer.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/repeat-string@1.6.1",
		"name": "repeat-string",
		"version": "1.6.1",
		"description": "Repeat the given string n times. Fastest implementation for repeating a string.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "8dcae470e1c88abc2d600fff4a776286da75e637"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/repeat-string@1.6.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/repeat-string"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/repeat-string/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/repeat-string.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/to-regex-range@2.1.1",
		"name": "to-regex-range",
		"version": "2.1.1",
		"description": "Pass two numbers, get a regex-compatible source string for matching ranges. Validated against more than 2.78 million test assertions.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "7c80c17b9dfebe599e27367e0d4dd5590141db38"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/to-regex-range@2.1.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/micromatch/to-regex-range"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/micromatch/to-regex-range/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/micromatch/to-regex-range.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/isobject@3.0.1",
		"name": "isobject",
		"version": "3.0.1",
		"description": "Returns true if the value is an object and not an array or null.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "4e431e92b11a9731636aa1f9c8d1ccbcfdab78df"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/isobject@3.0.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/isobject"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/isobject/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/isobject.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/repeat-element@1.1.3",
		"name": "repeat-element",
		"version": "1.1.3",
		"description": "Create an array by repeating the given value n times.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "782e0d825c0c5a3bb39731f84efee6b742e6b1ce"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/repeat-element@1.1.3",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/repeat-element"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/repeat-element/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/repeat-element.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/snapdragon@0.8.2",
		"name": "snapdragon",
		"version": "0.8.2",
		"description": "Fast, pluggable and easy-to-use parser-renderer factory.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "64922e7c565b0e14204ba1aa7d6964278d25182d"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/snapdragon@0.8.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/snapdragon"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/snapdragon/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/snapdragon.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/base@0.11.2",
		"name": "base",
		"version": "0.11.2",
		"description": "base is the foundation for creating modular, unit testable and highly pluggable node.js applications, starting with a handful of common methods, like set, get, del and use.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "7bde5ced145b6d551a90db87f83c558b4eb48a8f"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/base@0.11.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/node-base/base"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/node-base/base/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/node-base/base.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/cache-base@1.0.1",
		"name": "cache-base",
		"version": "1.0.1",
		"description": "Basic object cache with get, set, del, and has methods for node.js/javascript projects.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "0a7f46416831c8b662ee36fe4e7c59d76f666ab2"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/cache-base@1.0.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/cache-base"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/cache-base/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/cache-base.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/collection-visit@1.0.0",
		"name": "collection-visit",
		"version": "1.0.0",
		"description": "Visit a method over the items in an object, or map visit over the objects in an array.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "4bc0373c164bc3291b4d368c829cf1a80a59dca0"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/collection-visit@1.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/collection-visit"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/collection-visit/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/collection-visit.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/map-visit@1.0.0",
		"name": "map-visit",
		"version": "1.0.0",
		"description": "Map visit over an array of objects.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "ecdca8f13144e660f1b5bd41f12f3479d98dfb8f"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/map-visit@1.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/map-visit"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/map-visit/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/map-visit.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/object-visit@1.0.1",
		"name": "object-visit",
		"version": "1.0.1",
		"description": "Call a specified method on each value in the given object.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "f79c4493af0c5377b59fe39d395e41042dd045bb"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/object-visit@1.0.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/object-visit"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/object-visit/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/object-visit.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/component-emitter@1.3.0",
		"name": "component-emitter",
		"version": "1.3.0",
		"description": "Event emitter",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "16e4070fba8ae29b679f2215853ee181ab2eabc0"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/component-emitter@1.3.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/component/emitter#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/component/emitter/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/component/emitter.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/get-value@2.0.6",
		"name": "get-value",
		"version": "2.0.6",
		"description": "Use property paths (a.b.c) to get a nested value from an object.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "dc15ca1c672387ca76bd37ac0a395ba2042a2c28"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/get-value@2.0.6",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/get-value"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/get-value/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/get-value.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/has-value@1.0.0",
		"name": "has-value",
		"version": "1.0.0",
		"description": "Returns true if a value exists, false if empty. Works with deeply nested values using object paths.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "18b281da585b1c5c51def24c930ed29a0be6b177"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/has-value@1.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/has-value"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/has-value/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/has-value.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/has-values@1.0.0",
		"name": "has-values",
		"version": "1.0.0",
		"description": "Returns true if any values exist, false if empty. Works for booleans, functions, numbers, strings, nulls, objects and arrays. ",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "95b0b63fec2146619a6fe57fe75628d5a39efe4f"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/has-values@1.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/has-values"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/has-values/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/has-values.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/kind-of@4.0.0",
		"name": "kind-of",
		"version": "4.0.0",
		"description": "Get the native type of a value.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "20813df3d712928b207378691a45066fae72dd57"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/kind-of@4.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/kind-of"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/kind-of/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/kind-of.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/set-value@2.0.1",
		"name": "set-value",
		"version": "2.0.1",
		"description": "Create nested values and any intermediaries using dot notation ('a.b.c') paths.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "a18d40530e6f07de4228c7defe4227af8cad005b"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/set-value@2.0.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/set-value"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/set-value/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/set-value.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/is-plain-object@2.0.4",
		"name": "is-plain-object",
		"version": "2.0.4",
		"description": "Returns true if an object was created by the Object constructor.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "2c163b3fafb1b606d9d17928f05c2a1c38e07677"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/is-plain-object@2.0.4",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/is-plain-object"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/is-plain-object/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/is-plain-object.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/split-string@3.1.0",
		"name": "split-string",
		"version": "3.1.0",
		"description": "Split a string on a character except when the character is escaped.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "7cb09dda3a86585705c64b39a6466038682e8fe2"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/split-string@3.1.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/split-string"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/split-string/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/split-string.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/extend-shallow@3.0.2",
		"name": "extend-shallow",
		"version": "3.0.2",
		"description": "Extend an object with the properties of additional objects. node.js/javascript util.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "26a71aaf073b39fb2127172746131c2704028db8"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/extend-shallow@3.0.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/extend-shallow"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/extend-shallow/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/extend-shallow.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/assign-symbols@1.0.0",
		"name": "assign-symbols",
		"version": "1.0.0",
		"description": "Assign the enumerable es6 Symbol properties from an object (or objects) to the first object passed on the arguments. Can be used as a supplement to other extend, assign or merge methods as a polyfill for the Symbols part of the es6 Object.assign method.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "59667f41fadd4f20ccbc2bb96b8d4f7f78ec0367"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/assign-symbols@1.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/assign-symbols"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/assign-symbols/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/assign-symbols.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/is-extendable@1.0.1",
		"name": "is-extendable",
		"version": "1.0.1",
		"description": "Returns true if a value is a plain object, array or function.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "a7470f9e426733d81bd81e1155264e3a3507cab4"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/is-extendable@1.0.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/is-extendable"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/is-extendable/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/is-extendable.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/to-object-path@0.3.0",
		"name": "to-object-path",
		"version": "0.3.0",
		"description": "Create an object path from a list or array of strings.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "297588b7b0e7e0ac08e04e672f85c1f4999e17af"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/to-object-path@0.3.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/to-object-path"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/to-object-path/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/to-object-path.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/union-value@1.0.1",
		"name": "union-value",
		"version": "1.0.1",
		"description": "Set an array of unique values as the property of an object. Supports setting deeply nested properties using using object-paths/dot notation.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "0b6fe7b835aecda61c6ea4d4f02c14221e109847"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/union-value@1.0.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/union-value"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/union-value/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/union-value.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/arr-union@3.1.0",
		"name": "arr-union",
		"version": "3.1.0",
		"description": "Combines a list of arrays, returning a single array with unique values, using strict equality for comparisons.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "e39b09aea9def866a8f206e288af63919bae39c4"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/arr-union@3.1.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/arr-union"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/arr-union/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/arr-union.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/unset-value@1.0.0",
		"name": "unset-value",
		"version": "1.0.0",
		"description": "Delete nested properties from an object using dot notation.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "8376873f7d2335179ffb1e6fc3a8ed0dfc8ab559"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/unset-value@1.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/unset-value"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/unset-value/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/unset-value.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/has-value@0.3.1",
		"name": "has-value",
		"version": "0.3.1",
		"description": "Returns true if a value exists, false if empty. Works with deeply nested values using object paths.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "7b1f58bada62ca827ec0a2078025654845995e1f"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/has-value@0.3.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/has-value"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/has-value/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/has-value.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/has-values@0.1.4",
		"name": "has-values",
		"version": "0.1.4",
		"description": "Returns true if any values exist, false if empty. Works for booleans, functions, numbers, strings, nulls, objects and arrays. ",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "6d61de95d91dfca9b9a02089ad384bff8f62b771"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/has-values@0.1.4",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/has-values"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/has-values/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/has-values.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/isobject@2.1.0",
		"name": "isobject",
		"version": "2.1.0",
		"description": "Returns true if the value is an object and not an array or null.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "f065561096a3f1da2ef46272f815c840d87e0c89"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/isobject@2.1.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/isobject"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/isobject/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/isobject.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/isarray@1.0.0",
		"name": "isarray",
		"version": "1.0.0",
		"description": "Array#isArray for older browsers",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "bb935d48582cba168c06834957a54a3e07124f11"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/isarray@1.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/juliangruber/isarray"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/juliangruber/isarray/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/juliangruber/isarray.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/class-utils@0.3.6",
		"name": "class-utils",
		"version": "0.3.6",
		"description": "Utils for working with JavaScript classes and prototype methods.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "f93369ae8b9a7ce02fd41faad0ca83033190c463"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/class-utils@0.3.6",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/class-utils"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/class-utils/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/class-utils.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/define-property@0.2.5",
		"name": "define-property",
		"version": "0.2.5",
		"description": "Define a non-enumerable property on an object.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "c35b1ef918ec3c990f9a5bc57be04aacec5c8116"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/define-property@0.2.5",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/define-property"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/define-property/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/define-property.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/is-descriptor@0.1.6",
		"name": "is-descriptor",
		"version": "0.1.6",
		"description": "Returns true if a value has the characteristics of a valid JavaScript descriptor. Works for data descriptors and accessor descriptors.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "366d8240dde487ca51823b1ab9f07a10a78251ca"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/is-descriptor@0.1.6",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/is-descriptor"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/is-descriptor/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/is-descriptor.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/is-accessor-descriptor@0.1.6",
		"name": "is-accessor-descriptor",
		"version": "0.1.6",
		"description": "Returns true if a value has the characteristics of a valid JavaScript accessor descriptor.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "a9e12cb3ae8d876727eeef3843f8a0897b5c98d6"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/is-accessor-descriptor@0.1.6",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/is-accessor-descriptor"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/is-accessor-descriptor/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/is-accessor-descriptor.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/is-data-descriptor@0.1.4",
		"name": "is-data-descriptor",
		"version": "0.1.4",
		"description": "Returns true if a value has the characteristics of a valid JavaScript data descriptor.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "0b5ee648388e2c860282e793f1856fec3f301b56"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/is-data-descriptor@0.1.4",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/is-data-descriptor"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/is-data-descriptor/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/is-data-descriptor.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/kind-of@5.1.0",
		"name": "kind-of",
		"version": "5.1.0",
		"description": "Get the native type of a value.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "729c91e2d857b7a419a1f9aa65685c4c33f5845d"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/kind-of@5.1.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/kind-of"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/kind-of/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/kind-of.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/static-extend@0.1.2",
		"name": "static-extend",
		"version": "0.1.2",
		"description": "Adds a static extend method to a class, to simplify inheritance. Extends the static properties, prototype properties, and descriptors from a Parent constructor onto Child constructors.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "60809c39cbff55337226fd5e0b520f341f1fb5c6"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/static-extend@0.1.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/static-extend"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/static-extend/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/static-extend.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/object-copy@0.1.0",
		"name": "object-copy",
		"version": "0.1.0",
		"description": "Copy static properties, prototype properties, and descriptors from one object to another.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "7e7d858b781bd7c991a41ba975ed3812754e998c"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/object-copy@0.1.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/object-copy"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/object-copy/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/object-copy.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/copy-descriptor@0.1.1",
		"name": "copy-descriptor",
		"version": "0.1.1",
		"description": "Copy a descriptor from object A to object B",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "676f6eb3c39997c2ee1ac3a924fd6124748f578d"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/copy-descriptor@0.1.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/copy-descriptor"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/copy-descriptor/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/copy-descriptor.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/define-property@1.0.0",
		"name": "define-property",
		"version": "1.0.0",
		"description": "Define a non-enumerable property on an object.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "769ebaaf3f4a63aad3af9e8d304c9bbe79bfb0e6"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/define-property@1.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/define-property"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/define-property/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/define-property.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/is-descriptor@1.0.2",
		"name": "is-descriptor",
		"version": "1.0.2",
		"description": "Returns true if a value has the characteristics of a valid JavaScript descriptor. Works for data descriptors and accessor descriptors.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "3b159746a66604b04f8c81524ba365c5f14d86ec"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/is-descriptor@1.0.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/is-descriptor"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/is-descriptor/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/is-descriptor.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/is-accessor-descriptor@1.0.0",
		"name": "is-accessor-descriptor",
		"version": "1.0.0",
		"description": "Returns true if a value has the characteristics of a valid JavaScript accessor descriptor.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "169c2f6d3df1f992618072365c9b0ea1f6878656"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/is-accessor-descriptor@1.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/is-accessor-descriptor"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/is-accessor-descriptor/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/is-accessor-descriptor.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/kind-of@6.0.3",
		"name": "kind-of",
		"version": "6.0.3",
		"description": "Get the native type of a value.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "07c05034a6c349fa06e24fa35aa76db4580ce4dd"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/kind-of@6.0.3",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/kind-of"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/kind-of/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/kind-of.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/is-data-descriptor@1.0.0",
		"name": "is-data-descriptor",
		"version": "1.0.0",
		"description": "Returns true if a value has the characteristics of a valid JavaScript data descriptor.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "d84876321d0e7add03990406abbbbd36ba9268c7"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/is-data-descriptor@1.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/is-data-descriptor"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/is-data-descriptor/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/is-data-descriptor.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/mixin-deep@1.3.2",
		"name": "mixin-deep",
		"version": "1.3.2",
		"description": "Deeply mix the properties of objects into the first object. Like merge-deep, but doesn't clone.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "1120b43dc359a785dce65b55b82e257ccf479566"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/mixin-deep@1.3.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/mixin-deep"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/mixin-deep/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/mixin-deep.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/for-in@1.0.2",
		"name": "for-in",
		"version": "1.0.2",
		"description": "Iterate over the own and inherited enumerable properties of an object, and return an object with properties that evaluate to true from the callback. Exit early by returning false. JavaScript/Node.js",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "81068d295a8142ec0ac726c6e2200c30fb6d5e80"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/for-in@1.0.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/for-in"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/for-in/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/for-in.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/pascalcase@0.1.1",
		"name": "pascalcase",
		"version": "0.1.1",
		"description": "Convert a string to pascal-case.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "b363e55e8006ca6fe21784d2db22bd15d7917f14"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/pascalcase@0.1.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/pascalcase"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/pascalcase/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/pascalcase.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/map-cache@0.2.2",
		"name": "map-cache",
		"version": "0.2.2",
		"description": "Basic cache object for storing key-value pairs.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "c32abd0bd6525d9b051645bb4f26ac5dc98a0dbf"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/map-cache@0.2.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/map-cache"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/map-cache/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/map-cache.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/source-map@0.5.7",
		"name": "source-map",
		"version": "0.5.7",
		"description": "Generates and consumes source maps",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "8a039d2d1021d22d1ea14c80d8ea468ba2ef3fcc"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "BSD-3-Clause"
			}
		  }
		],
		"purl": "pkg:npm/source-map@0.5.7",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/mozilla/source-map"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/mozilla/source-map/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+ssh://git@github.com/mozilla/source-map.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/source-map-resolve@0.5.3",
		"name": "source-map-resolve",
		"version": "0.5.3",
		"description": "Resolve the source map and/or sources for a generated file.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "190866bece7553e1f8f267a2ee82c606b5509a1a"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/source-map-resolve@0.5.3",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/lydell/source-map-resolve#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/lydell/source-map-resolve/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/lydell/source-map-resolve.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/atob@2.1.2",
		"name": "atob",
		"version": "2.1.2",
		"description": "atob for Node.JS and Linux / Mac / Windows CLI (it's a one-liner)",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "6d9517eb9e030d2436666651e86bd9f6f13533c9"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "name": "(MIT OR Apache-2.0)"
			}
		  }
		],
		"purl": "pkg:npm/atob@2.1.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://git.coolaj86.com/coolaj86/atob.js.git"
		  },
		  {
			"type": "vcs",
			"url": "git://git.coolaj86.com/coolaj86/atob.js.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/decode-uri-component@0.2.0",
		"name": "decode-uri-component",
		"version": "0.2.0",
		"description": "A better decodeURIComponent",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "eb3913333458775cb84cd1a1fae062106bb87545"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/decode-uri-component@0.2.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/SamVerschueren/decode-uri-component#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/SamVerschueren/decode-uri-component/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/SamVerschueren/decode-uri-component.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/resolve-url@0.2.1",
		"name": "resolve-url",
		"version": "0.2.1",
		"description": "Like Node.js’ path.resolve/url.resolve for the browser.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "2c637fe77c893afd2a663fe21aa9080068e2052a"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/resolve-url@0.2.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/lydell/resolve-url#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/lydell/resolve-url/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/lydell/resolve-url.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/source-map-url@0.4.0",
		"name": "source-map-url",
		"version": "0.4.0",
		"description": "Tools for working with sourceMappingURL comments.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "3e935d7ddd73631b97659956d55128e87b5084a3"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/source-map-url@0.4.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/lydell/source-map-url#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/lydell/source-map-url/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/lydell/source-map-url.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/urix@0.1.0",
		"name": "urix",
		"version": "0.1.0",
		"description": "Makes Windows-style paths more unix and URI friendly.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "da937f7a62e21fec1fd18d49b35c2935067a6c72"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/urix@0.1.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/lydell/urix#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/lydell/urix/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/lydell/urix.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/use@3.1.1",
		"name": "use",
		"version": "3.1.1",
		"description": "Easily add plugin support to your node.js application.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "d50c8cac79a19fbc20f2911f56eb973f4e10070f"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/use@3.1.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/use"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/use/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/use.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/snapdragon-node@2.1.1",
		"name": "snapdragon-node",
		"version": "2.1.1",
		"description": "Snapdragon utility for creating a new AST node in custom code, such as plugins.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "6c175f86ff14bdb0724563e8f3c1b021a286853b"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/snapdragon-node@2.1.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/snapdragon-node"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/snapdragon-node/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/snapdragon-node.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/snapdragon-util@3.0.1",
		"name": "snapdragon-util",
		"version": "3.0.1",
		"description": "Utilities for the snapdragon parser/compiler.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "f956479486f2acd79700693f6f7b805e45ab56e2"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/snapdragon-util@3.0.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/snapdragon-util"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/snapdragon-util/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/snapdragon-util.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/to-regex@3.0.2",
		"name": "to-regex",
		"version": "3.0.2",
		"description": "Generate a regex from a string or array of strings.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "13cfdd9b336552f30b51f33a8ae1b42a7a7599ce"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/to-regex@3.0.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/to-regex"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/to-regex/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/to-regex.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/define-property@2.0.2",
		"name": "define-property",
		"version": "2.0.2",
		"description": "Define a non-enumerable property on an object. Uses Reflect.defineProperty when available, otherwise Object.defineProperty.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "d459689e8d654ba77e02a817f8710d702cb16e9d"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/define-property@2.0.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/define-property"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/define-property/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/define-property.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/regex-not@1.0.2",
		"name": "regex-not",
		"version": "1.0.2",
		"description": "Create a javascript regular expression for matching everything except for the given string.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "1f4ece27e00b0b65e0247a6810e6a85d83a5752c"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/regex-not@1.0.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/regex-not"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/regex-not/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/regex-not.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/safe-regex@1.1.0",
		"name": "safe-regex",
		"version": "1.1.0",
		"description": "detect possibly catastrophic, exponential-time regular expressions",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "40a3669f3b077d1e943d44629e157dd48023bf2e"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/safe-regex@1.1.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/substack/safe-regex"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/substack/safe-regex/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/substack/safe-regex.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/ret@0.1.15",
		"name": "ret",
		"version": "0.1.15",
		"description": "Tokenizes a string that represents a regular expression.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "b8a4825d5bdb1fc3f6f53c2bc33f81388681c7bc"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/ret@0.1.15",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/fent/ret.js#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/fent/ret.js/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/fent/ret.js.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/extglob@2.0.4",
		"name": "extglob",
		"version": "2.0.4",
		"description": "Extended glob support for JavaScript. Adds (almost) the expressive power of regular expressions to glob patterns.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "ad00fe4dc612a9232e8718711dc5cb5ab0285543"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/extglob@2.0.4",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/micromatch/extglob"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/micromatch/extglob/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/micromatch/extglob.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/expand-brackets@2.1.4",
		"name": "expand-brackets",
		"version": "2.1.4",
		"description": "Expand POSIX bracket expressions (character classes) in glob patterns.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "b77735e315ce30f6b6eff0f83b04151a22449622"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/expand-brackets@2.1.4",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/expand-brackets"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/expand-brackets/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/expand-brackets.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/posix-character-classes@0.1.1",
		"name": "posix-character-classes",
		"version": "0.1.1",
		"description": "POSIX character classes for creating regular expressions.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "01eac0fe3b5af71a2a6c02feabb8c1fef7e00eab"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/posix-character-classes@0.1.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/posix-character-classes"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/posix-character-classes/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/posix-character-classes.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/fragment-cache@0.2.1",
		"name": "fragment-cache",
		"version": "0.2.1",
		"description": "A cache for managing namespaced sub-caches",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "4290fad27f13e89be7f33799c6bc5a0abfff0d19"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/fragment-cache@0.2.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/fragment-cache"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/fragment-cache/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/fragment-cache.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/nanomatch@1.2.13",
		"name": "nanomatch",
		"version": "1.2.13",
		"description": "Fast, minimal glob matcher for node.js. Similar to micromatch, minimatch and multimatch, but complete Bash 4.3 wildcard support only (no support for exglobs, posix brackets or braces)",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "b87a8aa4fc0de8fe6be88895b38983ff265bd119"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/nanomatch@1.2.13",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/micromatch/nanomatch"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/micromatch/nanomatch/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/micromatch/nanomatch.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/is-windows@1.0.2",
		"name": "is-windows",
		"version": "1.0.2",
		"description": "Returns true if the platform is windows. UMD module, works with node.js, commonjs, browser, AMD, electron, etc.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "d1850eb9791ecd18e6182ce12a30f396634bb19d"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/is-windows@1.0.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/is-windows"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/is-windows/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/is-windows.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/object.pick@1.3.0",
		"name": "object.pick",
		"version": "1.3.0",
		"description": "Returns a filtered copy of an object with only the specified keys, similar to _.pick from lodash / underscore.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "87a10ac4c1694bd2e1cbf53591a66141fb5dd747"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/object.pick@1.3.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/object.pick"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/object.pick/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/object.pick.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/resolve-dir@1.0.1",
		"name": "resolve-dir",
		"version": "1.0.1",
		"description": "Resolve a directory that is either local, global or in the user's home directory.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "79a40644c362be82f26effe739c9bb5382046f43"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/resolve-dir@1.0.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/resolve-dir"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/resolve-dir/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/resolve-dir.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/expand-tilde@2.0.2",
		"name": "expand-tilde",
		"version": "2.0.2",
		"description": "Bash-like tilde expansion for node.js. Expands a leading tilde in a file path to the user home directory, or ~+ to the cwd.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "97e801aa052df02454de46b02bf621642cdc8502"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/expand-tilde@2.0.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/expand-tilde"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/expand-tilde/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/expand-tilde.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/homedir-polyfill@1.0.3",
		"name": "homedir-polyfill",
		"version": "1.0.3",
		"description": "Node.js os.homedir polyfill for older versions of node.js.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "743298cef4e5af3e194161fbadcc2151d3a058e8"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/homedir-polyfill@1.0.3",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/doowb/homedir-polyfill"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/doowb/homedir-polyfill/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/doowb/homedir-polyfill.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/parse-passwd@1.0.0",
		"name": "parse-passwd",
		"version": "1.0.0",
		"description": "Parse a passwd file into a list of users.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "6d5b934a456993b23d37f40a382d6f1666a8e5c6"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/parse-passwd@1.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/doowb/parse-passwd"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/doowb/parse-passwd/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/doowb/parse-passwd.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/global-modules@1.0.0",
		"name": "global-modules",
		"version": "1.0.0",
		"description": "The directory used by npm for globally installed npm modules.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "6d770f0eb523ac78164d72b5e71a8877265cc3ea"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/global-modules@1.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/global-modules"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/global-modules/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/global-modules.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/global-prefix@1.0.2",
		"name": "global-prefix",
		"version": "1.0.2",
		"description": "Get the npm global path prefix.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "dbf743c6c14992593c655568cb66ed32c0122ebe"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/global-prefix@1.0.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/global-prefix"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/global-prefix/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/global-prefix.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/ini@1.3.5",
		"name": "ini",
		"version": "1.3.5",
		"description": "An ini encoder/decoder for node",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "eee25f56db1c9ec6085e0c22778083f596abf927"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "ISC"
			}
		  }
		],
		"purl": "pkg:npm/ini@1.3.5",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/isaacs/ini#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/isaacs/ini/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/isaacs/ini.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/which@1.3.1",
		"name": "which",
		"version": "1.3.1",
		"description": "Like which(1) unix command. Find the first instance of an executable in the PATH.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "a45043d54f5805316da8d62f9f50918d3da70b0a"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "ISC"
			}
		  }
		],
		"purl": "pkg:npm/which@1.3.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/isaacs/node-which#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/isaacs/node-which/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/isaacs/node-which.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/isexe@2.0.0",
		"name": "isexe",
		"version": "2.0.0",
		"description": "Minimal module to check if a file is executable.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "e8fbf374dc556ff8947a10dcb0572d633f2cfa10"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "ISC"
			}
		  }
		],
		"purl": "pkg:npm/isexe@2.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/isaacs/isexe#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/isaacs/isexe/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/isaacs/isexe.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/lodash.camelcase@4.3.0",
		"name": "lodash.camelcase",
		"version": "4.3.0",
		"description": "The lodash method _.camelCase exported as a module.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "b28aa6288a2b9fc651035c7711f65ab6190331a6"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/lodash.camelcase@4.3.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://lodash.com/"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/lodash/lodash/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/lodash/lodash.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/minimist@1.2.5",
		"name": "minimist",
		"version": "1.2.5",
		"description": "parse argument options",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "67d66014b66a6a8aaa0c083c5fd58df4e4e97602"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/minimist@1.2.5",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/substack/minimist"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/substack/minimist/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/substack/minimist.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/semver@5.7.1",
		"name": "semver",
		"version": "5.7.1",
		"description": "The semantic version parser used by npm.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "a954f931aeba508d307bbf069eff0c01c96116f7"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "ISC"
			}
		  }
		],
		"purl": "pkg:npm/semver@5.7.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/npm/node-semver#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/npm/node-semver/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/npm/node-semver.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/clarinet@0.12.4",
		"name": "clarinet",
		"version": "0.12.4",
		"description": "SAX based evented streaming JSON parser in JavaScript (browser and node)",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "5d7196a2b2347ff283db2e2bf1ef615c0aa6afdb"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "BSD-2-Clause"
			}
		  }
		],
		"purl": "pkg:npm/clarinet@0.12.4",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/dscape/clarinet"
		  },
		  {
			"type": "issue-tracker",
			"url": "http://github.com/dscape/clarinet/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+ssh://git@github.com/dscape/clarinet.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/colors@1.4.0",
		"name": "colors",
		"version": "1.4.0",
		"description": "get colors in your node.js console",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "c50491479d4c1bdaed2c9ced32cf7c7dc2360f78"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/colors@1.4.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/Marak/colors.js"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/Marak/colors.js/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+ssh://git@github.com/Marak/colors.js.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/compression@1.7.4",
		"name": "compression",
		"version": "1.7.4",
		"description": "Node.js compression middleware",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "95523eff170ca57c29a0ca41e6fe131f41e5bb8f"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/compression@1.7.4",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/expressjs/compression#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/expressjs/compression/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/expressjs/compression.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/accepts@1.3.7",
		"name": "accepts",
		"version": "1.3.7",
		"description": "Higher-level content negotiation",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "531bc726517a3b2b41f850021c6cc15eaab507cd"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/accepts@1.3.7",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jshttp/accepts#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jshttp/accepts/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jshttp/accepts.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/negotiator@0.6.2",
		"name": "negotiator",
		"version": "0.6.2",
		"description": "HTTP content negotiation",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "feacf7ccf525a77ae9634436a64883ffeca346fb"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/negotiator@0.6.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jshttp/negotiator#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jshttp/negotiator/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jshttp/negotiator.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/bytes@3.0.0",
		"name": "bytes",
		"version": "3.0.0",
		"description": "Utility to parse a string bytes to bytes and vice-versa",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "d32815404d689699f85a4ea4fa8755dd13a96048"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/bytes@3.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/visionmedia/bytes.js#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/visionmedia/bytes.js/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/visionmedia/bytes.js.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/compressible@2.0.18",
		"name": "compressible",
		"version": "2.0.18",
		"description": "Compressible Content-Type / mime checking",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "af53cca6b070d4c3c0750fbd77286a6d7cc46fba"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/compressible@2.0.18",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jshttp/compressible#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jshttp/compressible/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jshttp/compressible.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/on-headers@1.0.2",
		"name": "on-headers",
		"version": "1.0.2",
		"description": "Execute a listener when a response is about to write headers",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "772b0ae6aaa525c399e489adfad90c403eb3c28f"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/on-headers@1.0.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jshttp/on-headers#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jshttp/on-headers/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jshttp/on-headers.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/safe-buffer@5.1.2",
		"name": "safe-buffer",
		"version": "5.1.2",
		"description": "Safer Node.js Buffer API",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "991ec69d296e0313747d59bdfd2b745c35f8828d"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/safe-buffer@5.1.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/feross/safe-buffer"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/feross/safe-buffer/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/feross/safe-buffer.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/vary@1.1.2",
		"name": "vary",
		"version": "1.1.2",
		"description": "Manipulate the HTTP Vary header",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "2299f02c6ded30d4a5961b0b9f74524a18f634fc"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/vary@1.1.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jshttp/vary#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jshttp/vary/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jshttp/vary.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/concurrently@5.2.0",
		"name": "concurrently",
		"version": "5.2.0",
		"description": "Run commands concurrently",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "ead55121d08a0fc817085584c123cedec2e08975"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/concurrently@5.2.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/kimmobrunfeldt/concurrently#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/kimmobrunfeldt/concurrently/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/kimmobrunfeldt/concurrently.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/date-fns@2.14.0",
		"name": "date-fns",
		"version": "2.14.0",
		"description": "Modern JavaScript date utility library",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "359a87a265bb34ef2e38f93ecf63ac453f9bc7ba"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/date-fns@2.14.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/date-fns/date-fns#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/date-fns/date-fns/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/date-fns/date-fns.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/lodash@4.17.19",
		"name": "lodash",
		"version": "4.17.19",
		"description": "Lodash modular utilities.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "e48ddedbe30b3321783c5b4301fbd353bc1e4a4b"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/lodash@4.17.19",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://lodash.com/"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/lodash/lodash/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/lodash/lodash.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/read-pkg@4.0.1",
		"name": "read-pkg",
		"version": "4.0.1",
		"description": "Read a package.json file",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "963625378f3e1c4d48c85872b5a6ec7d5d093237"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/read-pkg@4.0.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/sindresorhus/read-pkg#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/sindresorhus/read-pkg/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/sindresorhus/read-pkg.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/normalize-package-data@2.5.0",
		"name": "normalize-package-data",
		"version": "2.5.0",
		"description": "Normalizes data that can be found in package.json files.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "e66db1838b200c1dfc233225d12cb36520e234a8"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "BSD-2-Clause"
			}
		  }
		],
		"purl": "pkg:npm/normalize-package-data@2.5.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/npm/normalize-package-data#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/npm/normalize-package-data/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/npm/normalize-package-data.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/hosted-git-info@2.8.8",
		"name": "hosted-git-info",
		"version": "2.8.8",
		"description": "Provides metadata and conversions from repository urls for Github, Bitbucket and Gitlab",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "7539bd4bc1e0e0a895815a2e0262420b12858488"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "ISC"
			}
		  }
		],
		"purl": "pkg:npm/hosted-git-info@2.8.8",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/npm/hosted-git-info"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/npm/hosted-git-info/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/npm/hosted-git-info.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/resolve@1.17.0",
		"name": "resolve",
		"version": "1.17.0",
		"description": "resolve like require.resolve() on behalf of files asynchronously and synchronously",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "b25941b54968231cc2d1bb76a79cb7f2c0bf8444"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/resolve@1.17.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/browserify/resolve#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/browserify/resolve/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/browserify/resolve.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/path-parse@1.0.6",
		"name": "path-parse",
		"version": "1.0.6",
		"description": "Node.js path.parse() ponyfill",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "d62dbb5679405d72c4737ec58600e9ddcf06d24c"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/path-parse@1.0.6",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jbgutierrez/path-parse#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jbgutierrez/path-parse/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jbgutierrez/path-parse.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/validate-npm-package-license@3.0.4",
		"name": "validate-npm-package-license",
		"version": "3.0.4",
		"description": "Give me a string and I'll tell you if it's a valid npm package license string",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "fc91f6b9c7ba15c857f4cb2c5defeec39d4f410a"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "Apache-2.0"
			}
		  }
		],
		"purl": "pkg:npm/validate-npm-package-license@3.0.4",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/kemitchell/validate-npm-package-license.js#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/kemitchell/validate-npm-package-license.js/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/kemitchell/validate-npm-package-license.js.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/spdx-correct@3.1.1",
		"name": "spdx-correct",
		"version": "3.1.1",
		"description": "correct invalid SPDX expressions",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "dece81ac9c1e6713e5f7d1b6f17d468fa53d89a9"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "Apache-2.0"
			}
		  }
		],
		"purl": "pkg:npm/spdx-correct@3.1.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jslicense/spdx-correct.js#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jslicense/spdx-correct.js/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jslicense/spdx-correct.js.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/spdx-expression-parse@3.0.1",
		"name": "spdx-expression-parse",
		"version": "3.0.1",
		"description": "parse SPDX license expressions",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "cf70f50482eefdc98e3ce0a6833e4a53ceeba679"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/spdx-expression-parse@3.0.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jslicense/spdx-expression-parse.js#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jslicense/spdx-expression-parse.js/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jslicense/spdx-expression-parse.js.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/spdx-exceptions@2.3.0",
		"name": "spdx-exceptions",
		"version": "2.3.0",
		"description": "list of SPDX standard license exceptions",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "3f28ce1a77a00372683eade4a433183527a2163d"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "CC-BY-3.0"
			}
		  }
		],
		"purl": "pkg:npm/spdx-exceptions@2.3.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/kemitchell/spdx-exceptions.json#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/kemitchell/spdx-exceptions.json/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/kemitchell/spdx-exceptions.json.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/spdx-license-ids@3.0.5",
		"name": "spdx-license-ids",
		"version": "3.0.5",
		"description": "A list of SPDX license identifiers",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "3694b5804567a458d3c8045842a6358632f62654"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "CC0-1.0"
			}
		  }
		],
		"purl": "pkg:npm/spdx-license-ids@3.0.5",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/shinnn/spdx-license-ids#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/shinnn/spdx-license-ids/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/shinnn/spdx-license-ids.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/parse-json@4.0.0",
		"name": "parse-json",
		"version": "4.0.0",
		"description": "Parse JSON with more helpful errors",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "be35f5425be1f7f6c747184f98a788cb99477ee0"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/parse-json@4.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/sindresorhus/parse-json#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/sindresorhus/parse-json/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/sindresorhus/parse-json.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/error-ex@1.3.2",
		"name": "error-ex",
		"version": "1.3.2",
		"description": "Easy error subclassing and stack customization",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "b4ac40648107fdcdcfae242f428bea8a14d4f1bf"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/error-ex@1.3.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/qix-/node-error-ex#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/qix-/node-error-ex/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/qix-/node-error-ex.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/is-arrayish@0.2.1",
		"name": "is-arrayish",
		"version": "0.2.1",
		"description": "Determines if an object can be used as an array",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "77c99840527aa8ecb1a8ba697b80645a7a926a9d"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/is-arrayish@0.2.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/qix-/node-is-arrayish#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/qix-/node-is-arrayish/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/qix-/node-is-arrayish.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/json-parse-better-errors@1.0.2",
		"name": "json-parse-better-errors",
		"version": "1.0.2",
		"description": "JSON.parse with context information on error",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "bb867cfb3450e69107c131d1c514bab3dc8bcaa9"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/json-parse-better-errors@1.0.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/zkat/json-parse-better-errors#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/zkat/json-parse-better-errors/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/zkat/json-parse-better-errors.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/pify@3.0.0",
		"name": "pify",
		"version": "3.0.0",
		"description": "Promisify a callback-style function",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "e5a4acd2c101fdf3d9a4d07f0dbc4db49dd28176"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/pify@3.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/sindresorhus/pify#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/sindresorhus/pify/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/sindresorhus/pify.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/rxjs@6.6.0",
		"name": "rxjs",
		"version": "6.6.0",
		"description": "Reactive Extensions for modern JavaScript",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "af2901eedf02e3a83ffa7f886240ff9018bbec84"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "Apache-2.0"
			}
		  }
		],
		"purl": "pkg:npm/rxjs@6.6.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/ReactiveX/RxJS"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/ReactiveX/RxJS/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/reactivex/rxjs.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/tslib@1.13.0",
		"name": "tslib",
		"version": "1.13.0",
		"description": "Runtime library for TypeScript helper functions",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "c881e13cc7015894ed914862d276436fa9a47043"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "0BSD"
			}
		  }
		],
		"purl": "pkg:npm/tslib@1.13.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://www.typescriptlang.org/"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/Microsoft/TypeScript/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/Microsoft/tslib.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/spawn-command@0.0.2-1",
		"name": "spawn-command",
		"version": "0.0.2-1",
		"description": "Spawn commands like child_process.exec does but return a ChildProcess",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "62f5e9466981c1b796dc5929937e11c9c6921bd0"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/spawn-command@0.0.2-1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/mmalecki/spawn-command#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/mmalecki/spawn-command/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/mmalecki/spawn-command.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/supports-color@6.1.0",
		"name": "supports-color",
		"version": "6.1.0",
		"description": "Detect whether a terminal supports color",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "0764abc69c63d5ac842dd4867e8d025e880df8f3"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/supports-color@6.1.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/chalk/supports-color#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/chalk/supports-color/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/chalk/supports-color.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/tree-kill@1.2.2",
		"name": "tree-kill",
		"version": "1.2.2",
		"description": "kill trees of processes",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "4ca09a9092c88b73a7cdc5e8a01b507b0790a0cc"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/tree-kill@1.2.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/pkrumins/node-tree-kill"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/pkrumins/node-tree-kill/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/pkrumins/node-tree-kill.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/yargs@13.3.2",
		"name": "yargs",
		"version": "13.3.2",
		"description": "yargs the modern, pirate-themed, successor to optimist.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "ad7ffefec1aa59565ac915f82dccb38a9c31a2dd"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/yargs@13.3.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://yargs.js.org/"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/yargs/yargs/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/yargs/yargs.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/cliui@5.0.0",
		"name": "cliui",
		"version": "5.0.0",
		"description": "easily create complex multi-column command-line-interfaces",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "deefcfdb2e800784aa34f46fa08e06851c7bbbc5"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "ISC"
			}
		  }
		],
		"purl": "pkg:npm/cliui@5.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/yargs/cliui#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/yargs/cliui/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+ssh://git@github.com/yargs/cliui.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/string-width@3.1.0",
		"name": "string-width",
		"version": "3.1.0",
		"description": "Get the visual width of a string - the number of columns required to display it",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "22767be21b62af1081574306f69ac51b62203961"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/string-width@3.1.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/sindresorhus/string-width#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/sindresorhus/string-width/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/sindresorhus/string-width.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/emoji-regex@7.0.3",
		"name": "emoji-regex",
		"version": "7.0.3",
		"description": "A regular expression to match all Emoji-only symbols as per the Unicode Standard.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "933a04052860c85e83c122479c4748a8e4c72156"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/emoji-regex@7.0.3",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://mths.be/emoji-regex"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/mathiasbynens/emoji-regex/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/mathiasbynens/emoji-regex.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/is-fullwidth-code-point@2.0.0",
		"name": "is-fullwidth-code-point",
		"version": "2.0.0",
		"description": "Check if the character represented by a given Unicode code point is fullwidth",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "a3b30a5c4f199183167aaab93beefae3ddfb654f"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/is-fullwidth-code-point@2.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/sindresorhus/is-fullwidth-code-point#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/sindresorhus/is-fullwidth-code-point/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/sindresorhus/is-fullwidth-code-point.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/strip-ansi@5.2.0",
		"name": "strip-ansi",
		"version": "5.2.0",
		"description": "Strip ANSI escape codes from a string",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "8c9a536feb6afc962bdfa5b104a5091c1ad9c0ae"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/strip-ansi@5.2.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/chalk/strip-ansi#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/chalk/strip-ansi/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/chalk/strip-ansi.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/ansi-regex@4.1.0",
		"name": "ansi-regex",
		"version": "4.1.0",
		"description": "Regular expression for matching ANSI escape codes",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "8b9f8f08cf1acb843756a839ca8c7e3168c51997"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/ansi-regex@4.1.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/chalk/ansi-regex#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/chalk/ansi-regex/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/chalk/ansi-regex.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/wrap-ansi@5.1.0",
		"name": "wrap-ansi",
		"version": "5.1.0",
		"description": "Wordwrap a string with ANSI escape codes",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "1fd1f67235d5b6d0fee781056001bfb694c03b09"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/wrap-ansi@5.1.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/chalk/wrap-ansi#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/chalk/wrap-ansi/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/chalk/wrap-ansi.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/find-up@3.0.0",
		"name": "find-up",
		"version": "3.0.0",
		"description": "Find a file or directory by walking up parent directories",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "49169f1d7993430646da61ecc5ae355c21c97b73"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/find-up@3.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/sindresorhus/find-up#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/sindresorhus/find-up/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/sindresorhus/find-up.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/locate-path@3.0.0",
		"name": "locate-path",
		"version": "3.0.0",
		"description": "Get the first path that exists on disk of multiple paths",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "dbec3b3ab759758071b58fe59fc41871af21400e"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/locate-path@3.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/sindresorhus/locate-path#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/sindresorhus/locate-path/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/sindresorhus/locate-path.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/p-locate@3.0.0",
		"name": "p-locate",
		"version": "3.0.0",
		"description": "Get the first fulfilled promise that satisfies the provided testing function",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "322d69a05c0264b25997d9f40cd8a891ab0064a4"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/p-locate@3.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/sindresorhus/p-locate#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/sindresorhus/p-locate/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/sindresorhus/p-locate.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/p-limit@2.3.0",
		"name": "p-limit",
		"version": "2.3.0",
		"description": "Run multiple promise-returning & async functions with limited concurrency",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "3dd33c647a214fdfffd835933eb086da0dc21db1"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/p-limit@2.3.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/sindresorhus/p-limit#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/sindresorhus/p-limit/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/sindresorhus/p-limit.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/p-try@2.2.0",
		"name": "p-try",
		"version": "2.2.0",
		"description": "Start a promise chain",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "cb2868540e313d61de58fafbe35ce9004d5540e6"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/p-try@2.2.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/sindresorhus/p-try#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/sindresorhus/p-try/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/sindresorhus/p-try.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/path-exists@3.0.0",
		"name": "path-exists",
		"version": "3.0.0",
		"description": "Check if a path exists",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "ce0ebeaa5f78cb18925ea7d810d7b59b010fd515"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/path-exists@3.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/sindresorhus/path-exists#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/sindresorhus/path-exists/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/sindresorhus/path-exists.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/get-caller-file@2.0.5",
		"name": "get-caller-file",
		"version": "2.0.5",
		"description": "[![Build Status](https://travis-ci.org/stefanpenner/get-caller-file.svg?branch=master)](https://travis-ci.org/stefanpenner/get-caller-file) [![Build status](https://ci.appveyor.com/api/projects/status/ol2q94g1932cy14a/branch/master?svg=true)](https://ci.appveyor.com/project/embercli/get-caller-file/branch/master)",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "4f94412a82db32f36e3b0b9741f8a97feb031f7e"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "ISC"
			}
		  }
		],
		"purl": "pkg:npm/get-caller-file@2.0.5",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/stefanpenner/get-caller-file#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/stefanpenner/get-caller-file/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/stefanpenner/get-caller-file.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/require-directory@2.1.1",
		"name": "require-directory",
		"version": "2.1.1",
		"description": "Recursively iterates over specified directory, require()'ing each file, and returning a nested hash structure containing those modules.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "8c64ad5fd30dab1c976e2344ffe7f792a6a6df42"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/require-directory@2.1.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/troygoode/node-require-directory/"
		  },
		  {
			"type": "issue-tracker",
			"url": "http://github.com/troygoode/node-require-directory/issues/"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/troygoode/node-require-directory.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/require-main-filename@2.0.0",
		"name": "require-main-filename",
		"version": "2.0.0",
		"description": "shim for require.main.filename() that works in as many environments as possible",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "d0b329ecc7cc0f61649f62215be69af54aa8989b"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "ISC"
			}
		  }
		],
		"purl": "pkg:npm/require-main-filename@2.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/yargs/require-main-filename#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/yargs/require-main-filename/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+ssh://git@github.com/yargs/require-main-filename.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/set-blocking@2.0.0",
		"name": "set-blocking",
		"version": "2.0.0",
		"description": "set blocking stdio and stderr ensuring that terminal output does not truncate",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "045f9782d011ae9a6803ddd382b24392b3d890f7"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "ISC"
			}
		  }
		],
		"purl": "pkg:npm/set-blocking@2.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/yargs/set-blocking#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/yargs/set-blocking/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/yargs/set-blocking.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/which-module@2.0.0",
		"name": "which-module",
		"version": "2.0.0",
		"description": "Find the module object for something that was require()d",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "d9ef07dce77b9902b8a3a8fa4b31c3e3f7e6e87a"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "ISC"
			}
		  }
		],
		"purl": "pkg:npm/which-module@2.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/nexdrew/which-module#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/nexdrew/which-module/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/nexdrew/which-module.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/y18n@4.0.0",
		"name": "y18n",
		"version": "4.0.0",
		"description": "the bare-bones internationalization library used by yargs",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "95ef94f85ecc81d007c264e190a120f0a3c8566b"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "ISC"
			}
		  }
		],
		"purl": "pkg:npm/y18n@4.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/yargs/y18n"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/yargs/y18n/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+ssh://git@github.com/yargs/y18n.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/yargs-parser@13.1.2",
		"name": "yargs-parser",
		"version": "13.1.2",
		"description": "the mighty option parser used by yargs",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "130f09702ebaeef2650d54ce6e3e5706f7a4fb38"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "ISC"
			}
		  }
		],
		"purl": "pkg:npm/yargs-parser@13.1.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/yargs/yargs-parser#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/yargs/yargs-parser/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+ssh://git@github.com/yargs/yargs-parser.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/camelcase@5.3.1",
		"name": "camelcase",
		"version": "5.3.1",
		"description": "Convert a dash/dot/underscore/space separated string to camelCase or PascalCase: foo-bar → fooBar",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "e3c9b31569e106811df242f715725a1f4c494320"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/camelcase@5.3.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/sindresorhus/camelcase#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/sindresorhus/camelcase/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/sindresorhus/camelcase.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/decamelize@1.2.0",
		"name": "decamelize",
		"version": "1.2.0",
		"description": "Convert a camelized string into a lowercased one with a custom separator: unicornRainbow → unicorn_rainbow",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "f6534d15148269b20352e7bee26f501f9a191290"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/decamelize@1.2.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/sindresorhus/decamelize#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/sindresorhus/decamelize/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/sindresorhus/decamelize.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/config@3.3.1",
		"name": "config",
		"version": "3.3.1",
		"description": "Configuration control for production node deployments",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "b6a70e2908a43b98ed20be7e367edf0cc8ed5a19"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/config@3.3.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "http://lorenwest.github.com/node-config"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/lorenwest/node-config/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+ssh://git@github.com/lorenwest/node-config.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/json5@2.1.3",
		"name": "json5",
		"version": "2.1.3",
		"description": "JSON for humans.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "c9b0f7fa9233bfe5807fe66fcf3a5617ed597d43"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/json5@2.1.3",
		"externalReferences": [
		  {
			"type": "website",
			"url": "http://json5.org/"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/json5/json5/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/json5/json5.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/cookie-parser@1.4.5",
		"name": "cookie-parser",
		"version": "1.4.5",
		"description": "Parse HTTP request cookies",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "3e572d4b7c0c80f9c61daf604e4336831b5d1d49"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/cookie-parser@1.4.5",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/expressjs/cookie-parser#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/expressjs/cookie-parser/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/expressjs/cookie-parser.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/cookie@0.4.0",
		"name": "cookie",
		"version": "0.4.0",
		"description": "HTTP server cookie parsing and serialization",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "beb437e7022b3b6d49019d088665303ebe9c14ba"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/cookie@0.4.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jshttp/cookie#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jshttp/cookie/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jshttp/cookie.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/cookie-signature@1.0.6",
		"name": "cookie-signature",
		"version": "1.0.6",
		"description": "Sign and unsign cookies",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "e303a882b342cc3ee8ca513a79999734dab3ae2c"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/cookie-signature@1.0.6",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/visionmedia/node-cookie-signature#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/visionmedia/node-cookie-signature/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/visionmedia/node-cookie-signature.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/cors@2.8.5",
		"name": "cors",
		"version": "2.8.5",
		"description": "Node.js CORS middleware",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "eac11da51592dd86b9f06f6e7ac293b3df875d29"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/cors@2.8.5",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/expressjs/cors#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/expressjs/cors/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/expressjs/cors.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/object-assign@4.1.1",
		"name": "object-assign",
		"version": "4.1.1",
		"description": "ES2015 Object.assign() ponyfill",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "2109adc7965887cfc05cbbd442cac8bfbb360863"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/object-assign@4.1.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/sindresorhus/object-assign#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/sindresorhus/object-assign/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/sindresorhus/object-assign.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/dottie@2.0.2",
		"name": "dottie",
		"version": "2.0.2",
		"description": "Fast and safe nested object access and manipulation in JavaScript",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "cc91c0726ce3a054ebf11c55fbc92a7f266dd154"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/dottie@2.0.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/mickhansen/dottie.js#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/mickhansen/dottie.js/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/mickhansen/dottie.js.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/download@7.1.0",
		"name": "download",
		"version": "7.1.0",
		"description": "Download and extract files",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "9059aa9d70b503ee76a132897be6dec8e5587233"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/download@7.1.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/kevva/download#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/kevva/download/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/kevva/download.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/archive-type@4.0.0",
		"name": "archive-type",
		"version": "4.0.0",
		"description": "Detect the archive type of a Buffer/Uint8Array",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "f92e72233056dfc6969472749c267bdb046b1d70"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/archive-type@4.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/kevva/archive-type#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/kevva/archive-type/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/kevva/archive-type.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/file-type@4.4.0",
		"name": "file-type",
		"version": "4.4.0",
		"description": "Detect the file type of a Buffer/Uint8Array",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "1b600e5fca1fbdc6e80c0a70c71c8dba5f7906c5"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/file-type@4.4.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/sindresorhus/file-type#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/sindresorhus/file-type/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/sindresorhus/file-type.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/caw@2.0.1",
		"name": "caw",
		"version": "2.0.1",
		"description": "Construct HTTP/HTTPS agents for tunneling proxies",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "6c3ca071fc194720883c2dc5da9b074bfc7e9e95"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/caw@2.0.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/kevva/caw#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/kevva/caw/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/kevva/caw.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/get-proxy@2.1.0",
		"name": "get-proxy",
		"version": "2.1.0",
		"description": "Get configured proxy",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "349f2b4d91d44c4d4d4e9cba2ad90143fac5ef93"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/get-proxy@2.1.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/kevva/get-proxy#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/kevva/get-proxy/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/kevva/get-proxy.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/npm-conf@1.1.3",
		"name": "npm-conf",
		"version": "1.1.3",
		"description": "Get the npm config",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "256cc47bd0e218c259c4e9550bf413bc2192aff9"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/npm-conf@1.1.3",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/kevva/npm-conf#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/kevva/npm-conf/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/kevva/npm-conf.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/config-chain@1.1.12",
		"name": "config-chain",
		"version": "1.1.12",
		"description": "HANDLE CONFIGURATION ONCE AND FOR ALL",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "0fde8d091200eb5e808caf25fe618c02f48e4efa"
		  }
		],
		"purl": "pkg:npm/config-chain@1.1.12",
		"externalReferences": [
		  {
			"type": "website",
			"url": "http://github.com/dominictarr/config-chain"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/dominictarr/config-chain/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/dominictarr/config-chain.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/proto-list@1.2.4",
		"name": "proto-list",
		"version": "1.2.4",
		"description": "A utility for managing a prototype chain",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "212d5bfe1318306a420f6402b8e26ff39647a849"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "ISC"
			}
		  }
		],
		"purl": "pkg:npm/proto-list@1.2.4",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/isaacs/proto-list#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/isaacs/proto-list/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/isaacs/proto-list.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/isurl@1.0.0",
		"name": "isurl",
		"version": "1.0.0",
		"description": "Checks whether a value is a WHATWG URL.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "b27f4f49f3cdaa3ea44a0a5b7f3462e6edc39d67"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/isurl@1.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/stevenvachon/isurl#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/stevenvachon/isurl/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/stevenvachon/isurl.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/has-to-string-tag-x@1.4.1",
		"name": "has-to-string-tag-x",
		"version": "1.4.1",
		"description": "Tests if ES6 @@toStringTag is supported.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "a045ab383d7b4b2012a00148ab0aa5f290044d4d"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/has-to-string-tag-x@1.4.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/Xotic750/has-to-string-tag-x"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/Xotic750/has-to-string-tag-x/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/Xotic750/has-to-string-tag-x.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/has-symbol-support-x@1.4.2",
		"name": "has-symbol-support-x",
		"version": "1.4.2",
		"description": "Tests if ES6 Symbol is supported.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "1409f98bc00247da45da67cee0a36f282ff26455"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/has-symbol-support-x@1.4.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/Xotic750/has-symbol-support-x"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/Xotic750/has-symbol-support-x/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/Xotic750/has-symbol-support-x.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/is-object@1.0.1",
		"name": "is-object",
		"version": "1.0.1",
		"description": "Checks whether a value is an object",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "8952688c5ec2ffd6b03ecc85e769e02903083470"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/is-object@1.0.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/ljharb/is-object"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/ljharb/is-object/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/ljharb/is-object.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/tunnel-agent@0.6.0",
		"name": "tunnel-agent",
		"version": "0.6.0",
		"description": "HTTP proxy tunneling agent. Formerly part of mikeal/request, now a standalone module.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "27a5dea06b36b04a0a9966774b290868f0fc40fd"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "Apache-2.0"
			}
		  }
		],
		"purl": "pkg:npm/tunnel-agent@0.6.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/mikeal/tunnel-agent#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/mikeal/tunnel-agent/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/mikeal/tunnel-agent.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/url-to-options@1.0.1",
		"name": "url-to-options",
		"version": "1.0.1",
		"description": "Convert a WHATWG URL to an http(s).request options object.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "1505a03a289a48cbd7a434efbaeec5055f5633a9"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/url-to-options@1.0.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/stevenvachon/url-to-options#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/stevenvachon/url-to-options/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/stevenvachon/url-to-options.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/content-disposition@0.5.3",
		"name": "content-disposition",
		"version": "0.5.3",
		"description": "Create and parse Content-Disposition header",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "e130caf7e7279087c5616c2007d0485698984fbd"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/content-disposition@0.5.3",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jshttp/content-disposition#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jshttp/content-disposition/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jshttp/content-disposition.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/decompress@4.2.1",
		"name": "decompress",
		"version": "4.2.1",
		"description": "Extracting archives made easy",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "007f55cc6a62c055afa37c07eb6a4ee1b773f118"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/decompress@4.2.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/kevva/decompress#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/kevva/decompress/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/kevva/decompress.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/decompress-tar@4.1.1",
		"name": "decompress-tar",
		"version": "4.1.1",
		"description": "decompress tar plugin",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "718cbd3fcb16209716e70a26b84e7ba4592e5af1"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/decompress-tar@4.1.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/kevva/decompress-tar#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/kevva/decompress-tar/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/kevva/decompress-tar.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/file-type@5.2.0",
		"name": "file-type",
		"version": "5.2.0",
		"description": "Detect the file type of a Buffer/Uint8Array",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "2ddbea7c73ffe36368dfae49dc338c058c2b8ad6"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/file-type@5.2.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/sindresorhus/file-type#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/sindresorhus/file-type/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/sindresorhus/file-type.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/is-stream@1.1.0",
		"name": "is-stream",
		"version": "1.1.0",
		"description": "Check if something is a Node.js stream",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "12d4a3dd4e68e0b79ceb8dbc84173ae80d91ca44"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/is-stream@1.1.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/sindresorhus/is-stream#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/sindresorhus/is-stream/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/sindresorhus/is-stream.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/tar-stream@1.6.2",
		"name": "tar-stream",
		"version": "1.6.2",
		"description": "tar-stream is a streaming tar parser and generator and nothing else. It is streams2 and operates purely using streams which means you can easily extract/parse tarballs without ever hitting the file system.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "8ea55dab37972253d9a9af90fdcd559ae435c555"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/tar-stream@1.6.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/mafintosh/tar-stream"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/mafintosh/tar-stream/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/mafintosh/tar-stream.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/bl@1.2.2",
		"name": "bl",
		"version": "1.2.2",
		"description": "Buffer List: collect buffers and access with a standard readable Buffer interface, streamable too!",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "a160911717103c07410cef63ef51b397c025af9c"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/bl@1.2.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/rvagg/bl"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/rvagg/bl/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/rvagg/bl.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/readable-stream@2.3.7",
		"name": "readable-stream",
		"version": "2.3.7",
		"description": "Streams3, a user-land copy of the stream library from Node.js",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "1eca1cf711aef814c04f62252a36a62f6cb23b57"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/readable-stream@2.3.7",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/nodejs/readable-stream#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/nodejs/readable-stream/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/nodejs/readable-stream.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/core-util-is@1.0.2",
		"name": "core-util-is",
		"version": "1.0.2",
		"description": "The util.is* functions introduced in Node v0.12.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "b5fd54220aa2bc5ab57aab7140c940754503c1a7"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/core-util-is@1.0.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/isaacs/core-util-is#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/isaacs/core-util-is/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/isaacs/core-util-is.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/process-nextick-args@2.0.1",
		"name": "process-nextick-args",
		"version": "2.0.1",
		"description": "process.nextTick but always with args",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "7820d9b16120cc55ca9ae7792680ae7dba6d7fe2"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/process-nextick-args@2.0.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/calvinmetcalf/process-nextick-args"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/calvinmetcalf/process-nextick-args/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/calvinmetcalf/process-nextick-args.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/string_decoder@1.1.1",
		"name": "string_decoder",
		"version": "1.1.1",
		"description": "The string_decoder module from Node core",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "9cf1611ba62685d7030ae9e4ba34149c3af03fc8"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/string_decoder@1.1.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/nodejs/string_decoder"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/nodejs/string_decoder/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/nodejs/string_decoder.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/util-deprecate@1.0.2",
		"name": "util-deprecate",
		"version": "1.0.2",
		"description": "The Node.js util.deprecate() function with browser support",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "450d4dc9fa70de732762fbd2d4a28981419a0ccf"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/util-deprecate@1.0.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/TooTallNate/util-deprecate"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/TooTallNate/util-deprecate/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/TooTallNate/util-deprecate.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/buffer-alloc@1.2.0",
		"name": "buffer-alloc",
		"version": "1.2.0",
		"description": "A [ponyfill](https://ponyfill.com) for Buffer.alloc.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "890dd90d923a873e08e10e5fd51a57e5b7cce0ec"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/buffer-alloc@1.2.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/LinusU/buffer-alloc#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/LinusU/buffer-alloc/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/LinusU/buffer-alloc.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/buffer-alloc-unsafe@1.1.0",
		"name": "buffer-alloc-unsafe",
		"version": "1.1.0",
		"description": "A [ponyfill](https://ponyfill.com) for Buffer.allocUnsafe.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "bd7dc26ae2972d0eda253be061dba992349c19f0"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/buffer-alloc-unsafe@1.1.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/LinusU/buffer-alloc-unsafe#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/LinusU/buffer-alloc-unsafe/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/LinusU/buffer-alloc-unsafe.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/buffer-fill@1.0.0",
		"name": "buffer-fill",
		"version": "1.0.0",
		"description": "A [ponyfill](https://ponyfill.com) for Buffer.fill.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "f8f78b76789888ef39f205cd637f68e702122b2c"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/buffer-fill@1.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/LinusU/buffer-fill#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/LinusU/buffer-fill/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/LinusU/buffer-fill.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/end-of-stream@1.4.4",
		"name": "end-of-stream",
		"version": "1.4.4",
		"description": "Call a callback when a readable/writable/duplex stream has completed or failed.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "5ae64a5f45057baf3626ec14da0ca5e4b2431eb0"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/end-of-stream@1.4.4",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/mafintosh/end-of-stream"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/mafintosh/end-of-stream/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/mafintosh/end-of-stream.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/once@1.4.0",
		"name": "once",
		"version": "1.4.0",
		"description": "Run a function exactly one time",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "583b1aa775961d4b113ac17d9c50baef9dd76bd1"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "ISC"
			}
		  }
		],
		"purl": "pkg:npm/once@1.4.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/isaacs/once#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/isaacs/once/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/isaacs/once.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/wrappy@1.0.2",
		"name": "wrappy",
		"version": "1.0.2",
		"description": "Callback wrapping utility",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "b5243d8f3ec1aa35f1364605bc0d1036e30ab69f"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "ISC"
			}
		  }
		],
		"purl": "pkg:npm/wrappy@1.0.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/npm/wrappy"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/npm/wrappy/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/npm/wrappy.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/fs-constants@1.0.0",
		"name": "fs-constants",
		"version": "1.0.0",
		"description": "Require constants across node and the browser",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "6be0de9be998ce16af8afc24497b9ee9b7ccd9ad"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/fs-constants@1.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/mafintosh/fs-constants"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/mafintosh/fs-constants/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/mafintosh/fs-constants.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/to-buffer@1.1.1",
		"name": "to-buffer",
		"version": "1.1.1",
		"description": "Pass in a string, get a buffer back. Pass in a buffer, get the same buffer back",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "493bd48f62d7c43fcded313a03dcadb2e1213a80"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/to-buffer@1.1.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/mafintosh/to-buffer"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/mafintosh/to-buffer/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/mafintosh/to-buffer.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/xtend@4.0.2",
		"name": "xtend",
		"version": "4.0.2",
		"description": "extend like a boss",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "bb72779f5fa465186b1f438f674fa347fdb5db54"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/xtend@4.0.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/Raynos/xtend"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/Raynos/xtend/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/Raynos/xtend.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/decompress-tarbz2@4.1.1",
		"name": "decompress-tarbz2",
		"version": "4.1.1",
		"description": "decompress tar.bz2 plugin",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "3082a5b880ea4043816349f378b56c516be1a39b"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/decompress-tarbz2@4.1.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/kevva/decompress-tarbz2#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/kevva/decompress-tarbz2/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/kevva/decompress-tarbz2.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/file-type@6.2.0",
		"name": "file-type",
		"version": "6.2.0",
		"description": "Detect the file type of a Buffer/Uint8Array",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "e50cd75d356ffed4e306dc4f5bcf52a79903a919"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/file-type@6.2.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/sindresorhus/file-type#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/sindresorhus/file-type/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/sindresorhus/file-type.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/seek-bzip@1.0.5",
		"name": "seek-bzip",
		"version": "1.0.5",
		"description": "a pure-JavaScript Node.JS module for random-access decoding bzip2 data",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "cfe917cb3d274bcffac792758af53173eb1fabdc"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/seek-bzip@1.0.5",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/cscott/seek-bzip#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/cscott/seek-bzip/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/cscott/seek-bzip.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/commander@2.8.1",
		"name": "commander",
		"version": "2.8.1",
		"description": "the complete solution for node.js command-line programs",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "06be367febfda0c330aa1e2a072d3dc9762425d4"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/commander@2.8.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/tj/commander.js#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/tj/commander.js/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/tj/commander.js.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/graceful-readlink@1.0.1",
		"name": "graceful-readlink",
		"version": "1.0.1",
		"description": "graceful fs.readlink",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "4cafad76bc62f02fa039b2f94e9a3dd3a391a725"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/graceful-readlink@1.0.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/zhiyelee/graceful-readlink"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/zhiyelee/graceful-readlink/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/zhiyelee/graceful-readlink.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/unbzip2-stream@1.4.3",
		"name": "unbzip2-stream",
		"version": "1.4.3",
		"description": "streaming unbzip2 implementation in pure javascript for node and browsers",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "b0da04c4371311df771cdc215e87f2130991ace7"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/unbzip2-stream@1.4.3",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/regular/unbzip2-stream#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/regular/unbzip2-stream/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/regular/unbzip2-stream.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/buffer@5.6.0",
		"name": "buffer",
		"version": "5.6.0",
		"description": "Node.js Buffer API, for the browser",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "a31749dc7d81d84db08abf937b6b8c4033f62786"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/buffer@5.6.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/feross/buffer"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/feross/buffer/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/feross/buffer.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/base64-js@1.3.1",
		"name": "base64-js",
		"version": "1.3.1",
		"description": "Base64 encoding/decoding in pure JS",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "58ece8cb75dd07e71ed08c736abc5fac4dbf8df1"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/base64-js@1.3.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/beatgammit/base64-js"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/beatgammit/base64-js/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/beatgammit/base64-js.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/ieee754@1.1.13",
		"name": "ieee754",
		"version": "1.1.13",
		"description": "Read/write IEEE754 floating point numbers from/to a Buffer or array-like object",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "ec168558e95aa181fd87d37f55c32bbcb6708b84"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "BSD-3-Clause"
			}
		  }
		],
		"purl": "pkg:npm/ieee754@1.1.13",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/feross/ieee754#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/feross/ieee754/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/feross/ieee754.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/through@2.3.8",
		"name": "through",
		"version": "2.3.8",
		"description": "simplified stream construction",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "0dd4c9ffaabc357960b1b724115d7e0e86a2e1f5"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/through@2.3.8",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/dominictarr/through"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/dominictarr/through/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/dominictarr/through.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/decompress-targz@4.1.1",
		"name": "decompress-targz",
		"version": "4.1.1",
		"description": "decompress tar.gz plugin",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "c09bc35c4d11f3de09f2d2da53e9de23e7ce1eee"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/decompress-targz@4.1.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/kevva/decompress-targz#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/kevva/decompress-targz/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/kevva/decompress-targz.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/decompress-unzip@4.0.1",
		"name": "decompress-unzip",
		"version": "4.0.1",
		"description": "decompress zip plugin",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "deaaccdfd14aeaf85578f733ae8210f9b4848f69"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/decompress-unzip@4.0.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/kevva/decompress-unzip#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/kevva/decompress-unzip/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/kevva/decompress-unzip.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/file-type@3.9.0",
		"name": "file-type",
		"version": "3.9.0",
		"description": "Detect the file type of a Buffer/Uint8Array",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "257a078384d1db8087bc449d107d52a52672b9e9"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/file-type@3.9.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/sindresorhus/file-type#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/sindresorhus/file-type/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/sindresorhus/file-type.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/get-stream@2.3.1",
		"name": "get-stream",
		"version": "2.3.1",
		"description": "Get a stream as a string, buffer, or array",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "5f38f93f346009666ee0150a054167f91bdd95de"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/get-stream@2.3.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/sindresorhus/get-stream#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/sindresorhus/get-stream/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/sindresorhus/get-stream.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/pinkie-promise@2.0.1",
		"name": "pinkie-promise",
		"version": "2.0.1",
		"description": "ES2015 Promise ponyfill",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "2135d6dfa7a358c069ac9b178776288228450ffa"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/pinkie-promise@2.0.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/floatdrop/pinkie-promise#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/floatdrop/pinkie-promise/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/floatdrop/pinkie-promise.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/pinkie@2.0.4",
		"name": "pinkie",
		"version": "2.0.4",
		"description": "Itty bitty little widdle twinkie pinkie ES2015 Promise implementation",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "72556b80cfa0d48a974e80e77248e80ed4f7f870"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/pinkie@2.0.4",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/floatdrop/pinkie#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/floatdrop/pinkie/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/floatdrop/pinkie.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/pify@2.3.0",
		"name": "pify",
		"version": "2.3.0",
		"description": "Promisify a callback-style function",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "ed141a6ac043a849ea588498e7dca8b15330e90c"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/pify@2.3.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/sindresorhus/pify#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/sindresorhus/pify/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/sindresorhus/pify.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/yauzl@2.10.0",
		"name": "yauzl",
		"version": "2.10.0",
		"description": "yet another unzip library for node",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "c7eb17c93e112cb1086fa6d8e51fb0667b79a5f9"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/yauzl@2.10.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/thejoshwolfe/yauzl"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/thejoshwolfe/yauzl/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/thejoshwolfe/yauzl.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/buffer-crc32@0.2.13",
		"name": "buffer-crc32",
		"version": "0.2.13",
		"description": "A pure javascript CRC32 algorithm that plays nice with binary data",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "0d333e3f00eac50aa1454abd30ef8c2a5d9a7242"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/buffer-crc32@0.2.13",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/brianloveswords/buffer-crc32"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/brianloveswords/buffer-crc32/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/brianloveswords/buffer-crc32.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/fd-slicer@1.1.0",
		"name": "fd-slicer",
		"version": "1.1.0",
		"description": "safely create multiple ReadStream or WriteStream objects from the same file descriptor",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "25c7c89cb1f9077f8891bbe61d8f390eae256f1e"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/fd-slicer@1.1.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/andrewrk/node-fd-slicer#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/andrewrk/node-fd-slicer/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/andrewrk/node-fd-slicer.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/pend@1.2.0",
		"name": "pend",
		"version": "1.2.0",
		"description": "dead-simple optimistic async helper",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "7a57eb550a6783f9115331fcf4663d5c8e007a50"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/pend@1.2.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/andrewrk/node-pend#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/andrewrk/node-pend/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/andrewrk/node-pend.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/make-dir@1.3.0",
		"name": "make-dir",
		"version": "1.3.0",
		"description": "Make a directory and its parents if needed - Think mkdir -p",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "79c1033b80515bd6d24ec9933e860ca75ee27f0c"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/make-dir@1.3.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/sindresorhus/make-dir#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/sindresorhus/make-dir/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/sindresorhus/make-dir.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/strip-dirs@2.1.0",
		"name": "strip-dirs",
		"version": "2.1.0",
		"description": "Remove leading directory components from a path, like tar's --strip-components option",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "4987736264fc344cf20f6c34aca9d13d1d4ed6c5"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/strip-dirs@2.1.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/shinnn/node-strip-dirs#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/shinnn/node-strip-dirs/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/shinnn/node-strip-dirs.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/is-natural-number@4.0.1",
		"name": "is-natural-number",
		"version": "4.0.1",
		"description": "Check if a value is a natural number",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "ab9d76e1db4ced51e35de0c72ebecf09f734cde8"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/is-natural-number@4.0.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/shinnn/is-natural-number.js#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/shinnn/is-natural-number.js/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/shinnn/is-natural-number.js.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/ext-name@5.0.0",
		"name": "ext-name",
		"version": "5.0.0",
		"description": "Get the file extension and MIME type from a file",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "70781981d183ee15d13993c8822045c506c8f0a6"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/ext-name@5.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/kevva/ext-name#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/kevva/ext-name/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/kevva/ext-name.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/ext-list@2.2.2",
		"name": "ext-list",
		"version": "2.2.2",
		"description": "List of known file extensions and their MIME types",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "0b98e64ed82f5acf0f2931babf69212ef52ddd37"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/ext-list@2.2.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/kevva/ext-list#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/kevva/ext-list/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/kevva/ext-list.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/sort-keys-length@1.0.1",
		"name": "sort-keys-length",
		"version": "1.0.1",
		"description": "Sort objecy keys by length",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "9cb6f4f4e9e48155a6aa0671edd336ff1479a188"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/sort-keys-length@1.0.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/kevva/sort-keys-length#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/kevva/sort-keys-length/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/kevva/sort-keys-length.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/sort-keys@1.1.2",
		"name": "sort-keys",
		"version": "1.1.2",
		"description": "Sort the keys of an object",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "441b6d4d346798f1b4e49e8920adfba0e543f9ad"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/sort-keys@1.1.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/sindresorhus/sort-keys#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/sindresorhus/sort-keys/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/sindresorhus/sort-keys.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/is-plain-obj@1.1.0",
		"name": "is-plain-obj",
		"version": "1.1.0",
		"description": "Check if a value is a plain object",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "71a50c8429dfca773c92a390a4a03b39fcd51d3e"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/is-plain-obj@1.1.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/sindresorhus/is-plain-obj#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/sindresorhus/is-plain-obj/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/sindresorhus/is-plain-obj.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/file-type@8.1.0",
		"name": "file-type",
		"version": "8.1.0",
		"description": "Detect the file type of a Buffer/Uint8Array",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "244f3b7ef641bbe0cca196c7276e4b332399f68c"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/file-type@8.1.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/sindresorhus/file-type#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/sindresorhus/file-type/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/sindresorhus/file-type.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/filenamify@2.1.0",
		"name": "filenamify",
		"version": "2.1.0",
		"description": "Convert a string to a valid safe filename",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "88faf495fb1b47abfd612300002a16228c677ee9"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/filenamify@2.1.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/sindresorhus/filenamify#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/sindresorhus/filenamify/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/sindresorhus/filenamify.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/filename-reserved-regex@2.0.0",
		"name": "filename-reserved-regex",
		"version": "2.0.0",
		"description": "Regular expression for matching reserved filename characters",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "abf73dfab735d045440abfea2d91f389ebbfa229"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/filename-reserved-regex@2.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/sindresorhus/filename-reserved-regex#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/sindresorhus/filename-reserved-regex/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/sindresorhus/filename-reserved-regex.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/strip-outer@1.0.1",
		"name": "strip-outer",
		"version": "1.0.1",
		"description": "Strip a substring from the start/end of a string",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "b2fd2abf6604b9d1e6013057195df836b8a9d631"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/strip-outer@1.0.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/sindresorhus/strip-outer#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/sindresorhus/strip-outer/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/sindresorhus/strip-outer.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/trim-repeated@1.0.0",
		"name": "trim-repeated",
		"version": "1.0.0",
		"description": "Trim a consecutively repeated substring: foo--bar---baz → foo-bar-baz",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "e3646a2ea4e891312bf7eace6cfb05380bc01c21"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/trim-repeated@1.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/sindresorhus/trim-repeated#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/sindresorhus/trim-repeated/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/sindresorhus/trim-repeated.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/get-stream@3.0.0",
		"name": "get-stream",
		"version": "3.0.0",
		"description": "Get a stream as a string, buffer, or array",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "8e943d1358dc37555054ecbe2edb05aa174ede14"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/get-stream@3.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/sindresorhus/get-stream#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/sindresorhus/get-stream/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/sindresorhus/get-stream.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/got@8.3.2",
		"name": "got",
		"version": "8.3.2",
		"description": "Simplified HTTP requests",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "1d23f64390e97f776cac52e5b936e5f514d2e937"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/got@8.3.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/sindresorhus/got#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/sindresorhus/got/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/sindresorhus/got.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/%40sindresorhus/is@0.7.0",
		"group": "@sindresorhus",
		"name": "is",
		"version": "0.7.0",
		"description": "Type check values: is.string('🦄') //=> true",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "9a06f4f137ee84d7df0460c1fdb1135ffa6c50fd"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/%40sindresorhus/is@0.7.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/sindresorhus/is#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/sindresorhus/is/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/sindresorhus/is.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/cacheable-request@2.1.4",
		"name": "cacheable-request",
		"version": "2.1.4",
		"description": "Wrap native HTTP requests with RFC compliant cache support",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "0d808801b6342ad33c91df9d0b44dc09b91e5c3d"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/cacheable-request@2.1.4",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/lukechilds/cacheable-request"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/lukechilds/cacheable-request/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/lukechilds/cacheable-request.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/clone-response@1.0.2",
		"name": "clone-response",
		"version": "1.0.2",
		"description": "Clone a Node.js HTTP response stream",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "d1dc973920314df67fbeb94223b4ee350239e96b"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/clone-response@1.0.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/lukechilds/clone-response"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/lukechilds/clone-response/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/lukechilds/clone-response.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/mimic-response@1.0.1",
		"name": "mimic-response",
		"version": "1.0.1",
		"description": "Mimic a Node.js HTTP response stream",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "4923538878eef42063cb8a3e3b0798781487ab1b"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/mimic-response@1.0.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/sindresorhus/mimic-response#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/sindresorhus/mimic-response/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/sindresorhus/mimic-response.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/http-cache-semantics@3.8.1",
		"name": "http-cache-semantics",
		"version": "3.8.1",
		"description": "Parses Cache-Control and other headers. Helps building correct HTTP caches and proxies",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "39b0e16add9b605bf0a9ef3d9daaf4843b4cacd2"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "BSD-2-Clause"
			}
		  }
		],
		"purl": "pkg:npm/http-cache-semantics@3.8.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/pornel/http-cache-semantics#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/pornel/http-cache-semantics/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/pornel/http-cache-semantics.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/keyv@3.0.0",
		"name": "keyv",
		"version": "3.0.0",
		"description": "Simple key-value storage with support for multiple backends",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "44923ba39e68b12a7cec7df6c3268c031f2ef373"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/keyv@3.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/lukechilds/keyv"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/lukechilds/keyv/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/lukechilds/keyv.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/json-buffer@3.0.0",
		"name": "json-buffer",
		"version": "3.0.0",
		"description": "JSON parse & stringify that supports binary via bops & base64",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "5b1f397afc75d677bde8bcfc0e47e1f9a3d9a898"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/json-buffer@3.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/dominictarr/json-buffer"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/dominictarr/json-buffer/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/dominictarr/json-buffer.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/lowercase-keys@1.0.0",
		"name": "lowercase-keys",
		"version": "1.0.0",
		"description": "Lowercase the keys of an object",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "4e3366b39e7f5457e35f1324bdf6f88d0bfc7306"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/lowercase-keys@1.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/sindresorhus/lowercase-keys#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/sindresorhus/lowercase-keys/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/sindresorhus/lowercase-keys.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/normalize-url@2.0.1",
		"name": "normalize-url",
		"version": "2.0.1",
		"description": "Normalize a URL",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "835a9da1551fa26f70e92329069a23aa6574d7e6"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/normalize-url@2.0.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/sindresorhus/normalize-url#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/sindresorhus/normalize-url/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/sindresorhus/normalize-url.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/prepend-http@2.0.0",
		"name": "prepend-http",
		"version": "2.0.0",
		"description": "Prepend http:// to humanized URLs like todomvc.com and localhost",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "e92434bfa5ea8c19f41cdfd401d741a3c819d897"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/prepend-http@2.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/sindresorhus/prepend-http#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/sindresorhus/prepend-http/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/sindresorhus/prepend-http.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/query-string@5.1.1",
		"name": "query-string",
		"version": "5.1.1",
		"description": "Parse and stringify URL query strings",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "a78c012b71c17e05f2e3fa2319dd330682efb3cb"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/query-string@5.1.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/sindresorhus/query-string#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/sindresorhus/query-string/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/sindresorhus/query-string.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/strict-uri-encode@1.1.0",
		"name": "strict-uri-encode",
		"version": "1.1.0",
		"description": "A stricter URI encode adhering to RFC 3986",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "279b225df1d582b1f54e65addd4352e18faa0713"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/strict-uri-encode@1.1.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/kevva/strict-uri-encode#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/kevva/strict-uri-encode/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/kevva/strict-uri-encode.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/sort-keys@2.0.0",
		"name": "sort-keys",
		"version": "2.0.0",
		"description": "Sort the keys of an object",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "658535584861ec97d730d6cf41822e1f56684128"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/sort-keys@2.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/sindresorhus/sort-keys#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/sindresorhus/sort-keys/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/sindresorhus/sort-keys.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/responselike@1.0.2",
		"name": "responselike",
		"version": "1.0.2",
		"description": "A response-like object for mocking a Node.js HTTP response stream",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "918720ef3b631c5642be068f15ade5a46f4ba1e7"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/responselike@1.0.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/lukechilds/responselike#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/lukechilds/responselike/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/lukechilds/responselike.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/lowercase-keys@1.0.1",
		"name": "lowercase-keys",
		"version": "1.0.1",
		"description": "Lowercase the keys of an object",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "6f9e30b47084d971a7c820ff15a6c5167b74c26f"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/lowercase-keys@1.0.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/sindresorhus/lowercase-keys#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/sindresorhus/lowercase-keys/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/sindresorhus/lowercase-keys.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/decompress-response@3.3.0",
		"name": "decompress-response",
		"version": "3.3.0",
		"description": "Decompress a HTTP response if needed",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "80a4dd323748384bfa248083622aedec982adff3"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/decompress-response@3.3.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/sindresorhus/decompress-response#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/sindresorhus/decompress-response/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/sindresorhus/decompress-response.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/duplexer3@0.1.4",
		"name": "duplexer3",
		"version": "0.1.4",
		"description": "Like duplexer but using streams3",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "ee01dd1cac0ed3cbc7fdbea37dc0a8f1ce002ce2"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "BSD-3-Clause"
			}
		  }
		],
		"purl": "pkg:npm/duplexer3@0.1.4",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/floatdrop/duplexer3#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/floatdrop/duplexer3/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/floatdrop/duplexer3.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/into-stream@3.1.0",
		"name": "into-stream",
		"version": "3.1.0",
		"description": "Convert a buffer/string/array/object/iterable/promise into a stream",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "96fb0a936c12babd6ff1752a17d05616abd094c6"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/into-stream@3.1.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/sindresorhus/into-stream#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/sindresorhus/into-stream/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/sindresorhus/into-stream.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/from2@2.3.0",
		"name": "from2",
		"version": "2.3.0",
		"description": "Convenience wrapper for ReadableStream, with an API lifted from \"from\" and \"through2\"",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "8bfb5502bde4a4d36cfdeea007fcca21d7e382af"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/from2@2.3.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/hughsk/from2"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/hughsk/from2/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/hughsk/from2.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/p-is-promise@1.1.0",
		"name": "p-is-promise",
		"version": "1.1.0",
		"description": "Check if something is a promise",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "9c9456989e9f6588017b0434d56097675c3da05e"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/p-is-promise@1.1.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/sindresorhus/p-is-promise#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/sindresorhus/p-is-promise/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/sindresorhus/p-is-promise.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/is-retry-allowed@1.2.0",
		"name": "is-retry-allowed",
		"version": "1.2.0",
		"description": "Is retry allowed for Error?",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "d778488bd0a4666a3be8a1482b9f2baafedea8b4"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/is-retry-allowed@1.2.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/floatdrop/is-retry-allowed#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/floatdrop/is-retry-allowed/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/floatdrop/is-retry-allowed.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/p-cancelable@0.4.1",
		"name": "p-cancelable",
		"version": "0.4.1",
		"description": "Create a promise that can be canceled",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "35f363d67d52081c8d9585e37bcceb7e0bbcb2a0"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/p-cancelable@0.4.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/sindresorhus/p-cancelable#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/sindresorhus/p-cancelable/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/sindresorhus/p-cancelable.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/p-timeout@2.0.1",
		"name": "p-timeout",
		"version": "2.0.1",
		"description": "Timeout a promise after a specified amount of time",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "d8dd1979595d2dc0139e1fe46b8b646cb3cdf038"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/p-timeout@2.0.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/sindresorhus/p-timeout#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/sindresorhus/p-timeout/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/sindresorhus/p-timeout.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/p-finally@1.0.0",
		"name": "p-finally",
		"version": "1.0.0",
		"description": "Promise#finally() ponyfill - Invoked when the promise is settled regardless of outcome",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "3fbcfb15b899a44123b34b6dcc18b724336a2cae"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/p-finally@1.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/sindresorhus/p-finally#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/sindresorhus/p-finally/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/sindresorhus/p-finally.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/timed-out@4.0.1",
		"name": "timed-out",
		"version": "4.0.1",
		"description": "Emit ETIMEDOUT or ESOCKETTIMEDOUT when ClientRequest is hanged",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "f32eacac5a175bea25d7fab565ab3ed8741ef56f"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/timed-out@4.0.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/floatdrop/timed-out#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/floatdrop/timed-out/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/floatdrop/timed-out.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/url-parse-lax@3.0.0",
		"name": "url-parse-lax",
		"version": "3.0.0",
		"description": "Lax url.parse() with support for protocol-less URLs & IPs",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "16b5cafc07dbe3676c1b1999177823d6503acb0c"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/url-parse-lax@3.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/sindresorhus/url-parse-lax#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/sindresorhus/url-parse-lax/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/sindresorhus/url-parse-lax.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/p-event@2.3.1",
		"name": "p-event",
		"version": "2.3.1",
		"description": "Promisify an event by waiting for it to be emitted",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "596279ef169ab2c3e0cae88c1cfbb08079993ef6"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/p-event@2.3.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/sindresorhus/p-event#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/sindresorhus/p-event/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/sindresorhus/p-event.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/errorhandler@1.5.1",
		"name": "errorhandler",
		"version": "1.5.1",
		"description": "Development-only error handler middleware",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "b9ba5d17cf90744cd1e851357a6e75bf806a9a91"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/errorhandler@1.5.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/expressjs/errorhandler#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/expressjs/errorhandler/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/expressjs/errorhandler.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/escape-html@1.0.3",
		"name": "escape-html",
		"version": "1.0.3",
		"description": "Escape string for use in HTML",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "0258eae4d3d0c0974de1c169188ef0051d1d1988"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/escape-html@1.0.3",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/component/escape-html#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/component/escape-html/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/component/escape-html.git"
		  }
		]
	  },
	  {
		"type": "framework",
		"bom-ref": "pkg:npm/express@4.17.1",
		"name": "express",
		"version": "4.17.1",
		"description": "Fast, unopinionated, minimalist web framework",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "4491fc38605cf51f8629d39c2b5d026f98a4c134"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/express@4.17.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "http://expressjs.com/"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/expressjs/express/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/expressjs/express.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/array-flatten@1.1.1",
		"name": "array-flatten",
		"version": "1.1.1",
		"description": "Flatten an array of nested arrays into a single flat array",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "9a5f699051b1e7073328f2a008968b64ea2955d2"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/array-flatten@1.1.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/blakeembrey/array-flatten"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/blakeembrey/array-flatten/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/blakeembrey/array-flatten.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/encodeurl@1.0.2",
		"name": "encodeurl",
		"version": "1.0.2",
		"description": "Encode a URL to a percent-encoded form, excluding already-encoded sequences",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "ad3ff4c86ec2d029322f5a02c3a9a606c95b3f59"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/encodeurl@1.0.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/pillarjs/encodeurl#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/pillarjs/encodeurl/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/pillarjs/encodeurl.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/etag@1.8.1",
		"name": "etag",
		"version": "1.8.1",
		"description": "Create simple HTTP ETags",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "41ae2eeb65efa62268aebfea83ac7d79299b0887"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/etag@1.8.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jshttp/etag#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jshttp/etag/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jshttp/etag.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/finalhandler@1.1.2",
		"name": "finalhandler",
		"version": "1.1.2",
		"description": "Node.js final http responder",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "b7e7d000ffd11938d0fdb053506f6ebabe9f587d"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/finalhandler@1.1.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/pillarjs/finalhandler#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/pillarjs/finalhandler/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/pillarjs/finalhandler.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/parseurl@1.3.3",
		"name": "parseurl",
		"version": "1.3.3",
		"description": "parse a url with memoization",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "9da19e7bee8d12dff0513ed5b76957793bc2e8d4"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/parseurl@1.3.3",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/pillarjs/parseurl#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/pillarjs/parseurl/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/pillarjs/parseurl.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/fresh@0.5.2",
		"name": "fresh",
		"version": "0.5.2",
		"description": "HTTP response freshness testing",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "3d8cadd90d976569fa835ab1f8e4b23a105605a7"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/fresh@0.5.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jshttp/fresh#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jshttp/fresh/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jshttp/fresh.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/merge-descriptors@1.0.1",
		"name": "merge-descriptors",
		"version": "1.0.1",
		"description": "Merge objects using descriptors",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "b00aaa556dd8b44568150ec9d1b953f3f90cbb61"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/merge-descriptors@1.0.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/component/merge-descriptors#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/component/merge-descriptors/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/component/merge-descriptors.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/methods@1.1.2",
		"name": "methods",
		"version": "1.1.2",
		"description": "HTTP methods that node supports",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "5529a4d67654134edcc5266656835b0f851afcee"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/methods@1.1.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jshttp/methods#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jshttp/methods/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jshttp/methods.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/path-to-regexp@0.1.7",
		"name": "path-to-regexp",
		"version": "0.1.7",
		"description": "Express style path to RegExp utility",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "df604178005f522f15eb4490e7247a1bfaa67f8c"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/path-to-regexp@0.1.7",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/component/path-to-regexp#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/component/path-to-regexp/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/component/path-to-regexp.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/proxy-addr@2.0.6",
		"name": "proxy-addr",
		"version": "2.0.6",
		"description": "Determine address of proxied request",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "fdc2336505447d3f2f2c638ed272caf614bbb2bf"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/proxy-addr@2.0.6",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jshttp/proxy-addr#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jshttp/proxy-addr/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jshttp/proxy-addr.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/forwarded@0.1.2",
		"name": "forwarded",
		"version": "0.1.2",
		"description": "Parse HTTP X-Forwarded-For header",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "98c23dab1175657b8c0573e8ceccd91b0ff18c84"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/forwarded@0.1.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jshttp/forwarded#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jshttp/forwarded/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jshttp/forwarded.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/ipaddr.js@1.9.1",
		"name": "ipaddr.js",
		"version": "1.9.1",
		"description": "A library for manipulating IPv4 and IPv6 addresses in JavaScript.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "bff38543eeb8984825079ff3a2a8e6cbd46781b3"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/ipaddr.js@1.9.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/whitequark/ipaddr.js#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/whitequark/ipaddr.js/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/whitequark/ipaddr.js.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/range-parser@1.2.1",
		"name": "range-parser",
		"version": "1.2.1",
		"description": "Range header field string parser",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "3cf37023d199e1c24d1a55b84800c2f3e6468031"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/range-parser@1.2.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jshttp/range-parser#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jshttp/range-parser/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jshttp/range-parser.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/send@0.17.1",
		"name": "send",
		"version": "0.17.1",
		"description": "Better streaming static file server with Range and conditional-GET support",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "c1d8b059f7900f7466dd4938bdc44e11ddb376c8"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/send@0.17.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/pillarjs/send#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/pillarjs/send/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/pillarjs/send.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/destroy@1.0.4",
		"name": "destroy",
		"version": "1.0.4",
		"description": "destroy a stream if possible",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "978857442c44749e4206613e37946205826abd80"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/destroy@1.0.4",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/stream-utils/destroy#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/stream-utils/destroy/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/stream-utils/destroy.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/mime@1.6.0",
		"name": "mime",
		"version": "1.6.0",
		"description": "A comprehensive library for mime-type mapping",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "32cd9e5c64553bd58d19a568af452acff04981b1"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/mime@1.6.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/broofa/node-mime#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/broofa/node-mime/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/broofa/node-mime.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/ms@2.1.1",
		"name": "ms",
		"version": "2.1.1",
		"description": "Tiny millisecond conversion utility",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "30a5864eb3ebb0a66f2ebe6d727af06a09d86e0a"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/ms@2.1.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/zeit/ms#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/zeit/ms/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/zeit/ms.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/serve-static@1.14.1",
		"name": "serve-static",
		"version": "1.14.1",
		"description": "Serve static files",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "666e636dc4f010f7ef29970a88a674320898b2f9"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/serve-static@1.14.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/expressjs/serve-static#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/expressjs/serve-static/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/expressjs/serve-static.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/utils-merge@1.0.1",
		"name": "utils-merge",
		"version": "1.0.1",
		"description": "merge() utility function",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "9f95710f50a267947b2ccc124741c1028427e713"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/utils-merge@1.0.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jaredhanson/utils-merge#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "http://github.com/jaredhanson/utils-merge/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/jaredhanson/utils-merge.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/express-jwt@0.1.3",
		"name": "express-jwt",
		"version": "0.1.3",
		"description": "JWT authentication middleware.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "7c78221f8b9d72106aff556a8a5b8e852d41b12f"
		  }
		],
		"purl": "pkg:npm/express-jwt@0.1.3",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/auth0/express-jwt#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "http://github.com/auth0/express-jwt/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/auth0/express-jwt.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/jsonwebtoken@0.1.0",
		"name": "jsonwebtoken",
		"version": "0.1.0",
		"description": "JSON Web Token implementation (symmetric and asymmetric)",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "505628492092fe35d08b600fa6768cd06711aaa2"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/jsonwebtoken@0.1.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/auth0/node-jsonwebtoken#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/auth0/node-jsonwebtoken/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/auth0/node-jsonwebtoken.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/jws@0.2.6",
		"name": "jws",
		"version": "0.2.6",
		"description": "Implementation of JSON Web Signatures",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "e9b7e9ac8d2ac1067413233bc6c20fbd8868e9ba"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/jws@0.2.6",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/brianloveswords/node-jws#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/brianloveswords/node-jws/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/brianloveswords/node-jws.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/base64url@0.0.6",
		"name": "base64url",
		"version": "0.0.6",
		"description": "For encoding to/from base64urls",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "9597b36b330db1c42477322ea87ea8027499b82b"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/base64url@0.0.6",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/brianloveswords/base64url#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/brianloveswords/base64url/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/brianloveswords/base64url.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/jwa@0.0.1",
		"name": "jwa",
		"version": "0.0.1",
		"description": "JWA implementation (supports all JWS algorithms)",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "2d05f54d68f170648c30fe45944731a388cd07cc"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/jwa@0.0.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/brianloveswords/node-jwa#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/brianloveswords/node-jwa/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/brianloveswords/node-jwa.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/moment@2.0.0",
		"name": "moment",
		"version": "2.0.0",
		"description": "Parse, manipulate, and display dates.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "2bbc5b44c321837693ab6efcadbd46ed946211fe"
		  }
		],
		"purl": "pkg:npm/moment@2.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "http://momentjs.com"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/timrwood/moment/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/timrwood/moment.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/express-rate-limit@5.1.3",
		"name": "express-rate-limit",
		"version": "5.1.3",
		"description": "Basic IP rate-limiting middleware for Express. Use to limit repeated requests to public APIs and/or endpoints such as password reset.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "656bacce3f093034976346958a0f0199902c9174"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/express-rate-limit@5.1.3",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/nfriedly/express-rate-limit"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/nfriedly/express-rate-limit/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/nfriedly/express-rate-limit.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/express-robots-txt@0.4.1",
		"name": "express-robots-txt",
		"version": "0.4.1",
		"description": "Express middleware to serve and generate robots.txt",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "f3123a9875fd885d3c11cf4a7348b89a20f40ffc"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/express-robots-txt@0.4.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/modosc/express-robots-txt"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/modosc/express-robots-txt/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/modosc/express-robots-txt.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/express-security.txt@2.0.0",
		"name": "express-security.txt",
		"version": "2.0.0",
		"description": "[![Build Status](https://travis-ci.org/gergelyke/express-security.txt.svg?branch=master)](https://travis-ci.org/gergelyke/express-security.txt)",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "e5b825109ea88ccfb3001c1558a4739528d1fde0"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "ISC"
			}
		  }
		],
		"purl": "pkg:npm/express-security.txt@2.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/gergelyke/express-security.txt#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/gergelyke/express-security.txt/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/gergelyke/express-security.txt.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/file-stream-rotator@0.5.7",
		"name": "file-stream-rotator",
		"version": "0.5.7",
		"description": "Automated stream rotation useful for log files",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "868a2e5966f7640a17dd86eda0e4467c089f6286"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/file-stream-rotator@0.5.7",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/rogerc/file-stream-rotator#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/rogerc/file-stream-rotator/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/rogerc/file-stream-rotator.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/moment@2.27.0",
		"name": "moment",
		"version": "2.27.0",
		"description": "Parse, validate, manipulate, and display dates",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "8bff4e3e26a236220dfe3e36de756b6ebaa0105d"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/moment@2.27.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://momentjs.com"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/moment/moment/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/moment/moment.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/file-type@12.4.2",
		"name": "file-type",
		"version": "12.4.2",
		"description": "Detect the file type of a Buffer/Uint8Array/ArrayBuffer",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "a344ea5664a1d01447ee7fb1b635f72feb6169d9"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/file-type@12.4.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/sindresorhus/file-type#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/sindresorhus/file-type/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/sindresorhus/file-type.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/finale-rest@1.1.1",
		"name": "finale-rest",
		"version": "1.1.1",
		"description": "Create REST resources and controllers with Sequelize and Express or Restify",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "74dc49fb1655e938cc84210acf8c349887090086"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/finale-rest@1.1.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/tommybananas/finale#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/tommybananas/finale/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/tommybananas/finale.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/bluebird@3.7.2",
		"name": "bluebird",
		"version": "3.7.2",
		"description": "Full featured Promises/A+ implementation with exceptionally good performance",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "9f229c15be272454ffa973ace0dbee79a1b0c36f"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/bluebird@3.7.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/petkaantonov/bluebird"
		  },
		  {
			"type": "issue-tracker",
			"url": "http://github.com/petkaantonov/bluebird/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/petkaantonov/bluebird.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/inflection@1.12.0",
		"name": "inflection",
		"version": "1.12.0",
		"description": "A port of inflection-js to node.js module",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "a200935656d6f5f6bc4dc7502e1aecb703228416"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/inflection@1.12.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/dreamerslab/node.inflection#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/dreamerslab/node.inflection/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/dreamerslab/node.inflection.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/fs-extra@8.1.0",
		"name": "fs-extra",
		"version": "8.1.0",
		"description": "fs-extra contains methods that aren't included in the vanilla Node.js fs package. Such as mkdir -p, cp -r, and rm -rf.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "49d43c45a88cd9677668cb7be1b46efdb8d2e1c0"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/fs-extra@8.1.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jprichardson/node-fs-extra"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jprichardson/node-fs-extra/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jprichardson/node-fs-extra.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/jsonfile@4.0.0",
		"name": "jsonfile",
		"version": "4.0.0",
		"description": "Easily read/write JSON files.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "8771aae0799b64076b76640fca058f9c10e33ecb"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/jsonfile@4.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jprichardson/node-jsonfile#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jprichardson/node-jsonfile/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+ssh://git@github.com/jprichardson/node-jsonfile.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/universalify@0.1.2",
		"name": "universalify",
		"version": "0.1.2",
		"description": "Make a callback- or promise-based function support both promises and callbacks.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "b646f69be3942dabcecc9d6639c80dc105efaa66"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/universalify@0.1.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/RyanZim/universalify#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/RyanZim/universalify/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/RyanZim/universalify.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/glob@7.1.6",
		"name": "glob",
		"version": "7.1.6",
		"description": "a little globber",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "141f33b81a7c2492e125594307480c46679278a6"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "ISC"
			}
		  }
		],
		"purl": "pkg:npm/glob@7.1.6",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/isaacs/node-glob#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/isaacs/node-glob/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/isaacs/node-glob.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/fs.realpath@1.0.0",
		"name": "fs.realpath",
		"version": "1.0.0",
		"description": "Use node's fs.realpath, but fall back to the JS implementation if the native one fails",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "1504ad2523158caa40db4a2787cb01411994ea4f"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "ISC"
			}
		  }
		],
		"purl": "pkg:npm/fs.realpath@1.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/isaacs/fs.realpath#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/isaacs/fs.realpath/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/isaacs/fs.realpath.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/inflight@1.0.6",
		"name": "inflight",
		"version": "1.0.6",
		"description": "Add callbacks to requests in flight to avoid async duplication",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "49bd6331d7d02d0c09bc910a1075ba8165b56df9"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "ISC"
			}
		  }
		],
		"purl": "pkg:npm/inflight@1.0.6",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/isaacs/inflight"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/isaacs/inflight/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/npm/inflight.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/minimatch@3.0.4",
		"name": "minimatch",
		"version": "3.0.4",
		"description": "a glob matcher in javascript",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "5166e286457f03306064be5497e8dbb0c3d32083"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "ISC"
			}
		  }
		],
		"purl": "pkg:npm/minimatch@3.0.4",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/isaacs/minimatch#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/isaacs/minimatch/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/isaacs/minimatch.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/brace-expansion@1.1.11",
		"name": "brace-expansion",
		"version": "1.1.11",
		"description": "Brace expansion as known from sh/bash",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "3c7fcbf529d87226f3d2f52b966ff5271eb441dd"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/brace-expansion@1.1.11",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/juliangruber/brace-expansion"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/juliangruber/brace-expansion/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/juliangruber/brace-expansion.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/balanced-match@1.0.0",
		"name": "balanced-match",
		"version": "1.0.0",
		"description": "Match balanced character pairs, like \"{\" and \"}\"",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "89b4d199ab2bee49de164ea02b89ce462d71b767"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/balanced-match@1.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/juliangruber/balanced-match"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/juliangruber/balanced-match/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/juliangruber/balanced-match.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/concat-map@0.0.1",
		"name": "concat-map",
		"version": "0.0.1",
		"description": "concatenative mapdashery",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "d8a96bd77fd68df7793a73036a3ba0d5405d477b"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/concat-map@0.0.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/substack/node-concat-map#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/substack/node-concat-map/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/substack/node-concat-map.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/path-is-absolute@1.0.1",
		"name": "path-is-absolute",
		"version": "1.0.1",
		"description": "Node.js 0.12 path.isAbsolute() ponyfill",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "174b9268735534ffbc7ace6bf53a5a9e1b5c5f5f"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/path-is-absolute@1.0.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/sindresorhus/path-is-absolute#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/sindresorhus/path-is-absolute/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/sindresorhus/path-is-absolute.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/grunt@1.2.1",
		"name": "grunt",
		"version": "1.2.1",
		"description": "The JavaScript Task Runner",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "5a1fcdfc222841108893e4e50c1a46f413a564ab"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/grunt@1.2.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://gruntjs.com/"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/gruntjs/grunt/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/gruntjs/grunt.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/dateformat@3.0.3",
		"name": "dateformat",
		"version": "3.0.3",
		"description": "A node.js package for Steven Levithan's excellent dateFormat() function.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "a6e37499a4d9a9cf85ef5872044d62901c9889ae"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/dateformat@3.0.3",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/felixge/node-dateformat"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/felixge/node-dateformat/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/felixge/node-dateformat.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/eventemitter2@0.4.14",
		"name": "eventemitter2",
		"version": "0.4.14",
		"description": "A Node.js event emitter implementation with namespaces, wildcards, TTL and browser support.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "8f61b75cde012b2e9eb284d4545583b5643b61ab"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/eventemitter2@0.4.14",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/hij1nx/EventEmitter2#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/hij1nx/EventEmitter2/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/hij1nx/EventEmitter2.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/exit@0.1.2",
		"name": "exit",
		"version": "0.1.2",
		"description": "A replacement for process.exit that ensures stdio are fully drained before exiting.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "0632638f8d877cc82107d30a0fff1a17cba1cd0c"
		  }
		],
		"purl": "pkg:npm/exit@0.1.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/cowboy/node-exit"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/cowboy/node-exit/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/cowboy/node-exit.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/findup-sync@0.3.0",
		"name": "findup-sync",
		"version": "0.3.0",
		"description": "Find the first file matching a given pattern in the current directory or the nearest ancestor directory.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "37930aa5d816b777c03445e1966cc6790a4c0b16"
		  }
		],
		"purl": "pkg:npm/findup-sync@0.3.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/cowboy/node-findup-sync"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/cowboy/node-findup-sync/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/cowboy/node-findup-sync.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/glob@5.0.15",
		"name": "glob",
		"version": "5.0.15",
		"description": "a little globber",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "1bc936b9e02f4a603fcc222ecf7633d30b8b93b1"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "ISC"
			}
		  }
		],
		"purl": "pkg:npm/glob@5.0.15",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/isaacs/node-glob#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/isaacs/node-glob/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/isaacs/node-glob.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/grunt-cli@1.3.2",
		"name": "grunt-cli",
		"version": "1.3.2",
		"description": "The grunt command line interface",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "60f12d12c1b5aae94ae3469c6b5fe24e960014e8"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/grunt-cli@1.3.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/gruntjs/grunt-cli#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/gruntjs/grunt-cli/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/gruntjs/grunt-cli.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/grunt-known-options@1.1.1",
		"name": "grunt-known-options",
		"version": "1.1.1",
		"description": "The known options used in Grunt",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "6cc088107bd0219dc5d3e57d91923f469059804d"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/grunt-known-options@1.1.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "http://gruntjs.com/"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/gruntjs/grunt-known-options/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/gruntjs/grunt-known-options.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/interpret@1.1.0",
		"name": "interpret",
		"version": "1.1.0",
		"description": "A dictionary of file extensions and associated module loaders.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "7ed1b1410c6a0e0f78cf95d3b8440c63f78b8614"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/interpret@1.1.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/tkellen/node-interpret"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/tkellen/node-interpret/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/tkellen/node-interpret.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/liftoff@2.5.0",
		"name": "liftoff",
		"version": "2.5.0",
		"description": "Launch your command line tool with ease.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "2009291bb31cea861bbf10a7c15a28caf75c31ec"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/liftoff@2.5.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/js-cli/js-liftoff#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/js-cli/js-liftoff/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/js-cli/js-liftoff.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/extend@3.0.2",
		"name": "extend",
		"version": "3.0.2",
		"description": "Port of jQuery.extend for node.js and the browser",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "f8b1136b4071fbd8eb140aff858b1019ec2915fa"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/extend@3.0.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/justmoon/node-extend#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/justmoon/node-extend/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/justmoon/node-extend.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/fined@1.2.0",
		"name": "fined",
		"version": "1.2.0",
		"description": "Find a file given a declaration of locations.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "d00beccf1aa2b475d16d423b0238b713a2c4a37b"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/fined@1.2.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/gulpjs/fined#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/gulpjs/fined/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/gulpjs/fined.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/object.defaults@1.1.0",
		"name": "object.defaults",
		"version": "1.1.0",
		"description": "Like extend but only copies missing properties/values to the target object.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "3a7f868334b407dea06da16d88d5cd29e435fecf"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/object.defaults@1.1.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/object.defaults"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/object.defaults/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/object.defaults.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/array-each@1.0.1",
		"name": "array-each",
		"version": "1.0.1",
		"description": "Loop over each item in an array and call the given function on every element.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "a794af0c05ab1752846ee753a1f211a05ba0c44f"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/array-each@1.0.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/array-each"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/array-each/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/array-each.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/array-slice@1.1.0",
		"name": "array-slice",
		"version": "1.1.0",
		"description": "Array-slice method. Slices array from the start index up to, but not including, the end index.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "e368ea15f89bc7069f7ffb89aec3a6c7d4ac22d4"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/array-slice@1.1.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/array-slice"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/array-slice/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/array-slice.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/for-own@1.0.0",
		"name": "for-own",
		"version": "1.0.0",
		"description": "Iterate over the own enumerable properties of an object, and return an object with properties that evaluate to true from the callback. Exit early by returning false. JavaScript/Node.js.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "c63332f415cedc4b04dbfe70cf836494c53cb44b"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/for-own@1.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/for-own"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/for-own/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/for-own.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/parse-filepath@1.0.2",
		"name": "parse-filepath",
		"version": "1.0.2",
		"description": "Pollyfill for node.js path.parse, parses a filepath into an object.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "a632127f53aaf3d15876f5872f3ffac763d6c891"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/parse-filepath@1.0.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/parse-filepath"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/parse-filepath/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/parse-filepath.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/is-absolute@1.0.0",
		"name": "is-absolute",
		"version": "1.0.0",
		"description": "Returns true if a file path is absolute. Does not rely on the path module and can be used as a polyfill for node.js native path.isAbolute.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "395e1ae84b11f26ad1795e73c17378e48a301576"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/is-absolute@1.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/is-absolute"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/is-absolute/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/is-absolute.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/is-relative@1.0.0",
		"name": "is-relative",
		"version": "1.0.0",
		"description": "Returns true if the path appears to be relative.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "a1bb6935ce8c5dba1e8b9754b9b2dcc020e2260d"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/is-relative@1.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/is-relative"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/is-relative/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/is-relative.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/is-unc-path@1.0.0",
		"name": "is-unc-path",
		"version": "1.0.0",
		"description": "Returns true if a filepath is a windows UNC file path.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "d731e8898ed090a12c352ad2eaed5095ad322c9d"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/is-unc-path@1.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/is-unc-path"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/is-unc-path/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/is-unc-path.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/unc-path-regex@0.1.2",
		"name": "unc-path-regex",
		"version": "0.1.2",
		"description": "Regular expression for testing if a file path is a windows UNC file path. Can also be used as a component of another regexp via the .source property.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "e73dd3d7b0d7c5ed86fbac6b0ae7d8c6a69d50fa"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/unc-path-regex@0.1.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/regexhq/unc-path-regex"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/regexhq/unc-path-regex/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/regexhq/unc-path-regex.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/path-root@0.1.1",
		"name": "path-root",
		"version": "0.1.1",
		"description": "Get the root of a posix or windows filepath.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "9a4a6814cac1c0cd73360a95f32083c8ea4745b7"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/path-root@0.1.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/path-root"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/path-root/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/path-root.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/path-root-regex@0.1.2",
		"name": "path-root-regex",
		"version": "0.1.2",
		"description": "Regular expression for getting the root of a posix or windows filepath.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "bfccdc8df5b12dc52c8b43ec38d18d72c04ba96d"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/path-root-regex@0.1.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/regexhq/path-root-regex"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/regexhq/path-root-regex/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/regexhq/path-root-regex.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/flagged-respawn@1.0.1",
		"name": "flagged-respawn",
		"version": "1.0.1",
		"description": "A tool for respawning node binaries when special flags are present.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "e7de6f1279ddd9ca9aac8a5971d618606b3aab41"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/flagged-respawn@1.0.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/gulpjs/flagged-respawn#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/gulpjs/flagged-respawn/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/gulpjs/flagged-respawn.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/object.map@1.0.1",
		"name": "object.map",
		"version": "1.0.1",
		"description": "Similar to map for arrays, this creates a new object by calling the callback on each property of the original object.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "cf83e59dc8fcc0ad5f4250e1f78b3b81bd801d37"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/object.map@1.0.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/object.map"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/object.map/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/object.map.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/make-iterator@1.0.1",
		"name": "make-iterator",
		"version": "1.0.1",
		"description": "Convert an argument into a valid iterator. Based on the .makeIterator() implementation in mout https://github.com/mout/mout.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "29b33f312aa8f547c4a5e490f56afcec99133ad6"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/make-iterator@1.0.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/make-iterator"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/make-iterator/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/make-iterator.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/rechoir@0.6.2",
		"name": "rechoir",
		"version": "0.6.2",
		"description": "Require any supported file as a node module.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "85204b54dba82d5742e28c96756ef43af50e3384"
		  }
		],
		"purl": "pkg:npm/rechoir@0.6.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/tkellen/node-rechoir"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/tkellen/node-rechoir/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/tkellen/node-rechoir.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/nopt@4.0.3",
		"name": "nopt",
		"version": "4.0.3",
		"description": "Option parsing for Node, supporting types, shorthands, etc. Used by npm.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "a375cad9d02fd921278d954c2254d5aa57e15e48"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "ISC"
			}
		  }
		],
		"purl": "pkg:npm/nopt@4.0.3",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/npm/nopt#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/npm/nopt/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/npm/nopt.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/abbrev@1.1.1",
		"name": "abbrev",
		"version": "1.1.1",
		"description": "Like ruby's abbrev module, but in js",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "f8f2c887ad10bf67f634f005b6987fed3179aac8"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "ISC"
			}
		  }
		],
		"purl": "pkg:npm/abbrev@1.1.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/isaacs/abbrev-js#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/isaacs/abbrev-js/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+ssh://git@github.com/isaacs/abbrev-js.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/v8flags@3.1.3",
		"name": "v8flags",
		"version": "3.1.3",
		"description": "Get available v8 flags.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "fc9dc23521ca20c5433f81cc4eb9b3033bb105d8"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/v8flags@3.1.3",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/gulpjs/v8flags#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/gulpjs/v8flags/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/gulpjs/v8flags.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/grunt-legacy-log@2.0.0",
		"name": "grunt-legacy-log",
		"version": "2.0.0",
		"description": "The Grunt 0.4.x logger.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "c8cd2c6c81a4465b9bbf2d874d963fef7a59ffb9"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/grunt-legacy-log@2.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "http://gruntjs.com/"
		  },
		  {
			"type": "issue-tracker",
			"url": "http://github.com/gruntjs/grunt-legacy-log/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/gruntjs/grunt-legacy-log.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/colors@1.1.2",
		"name": "colors",
		"version": "1.1.2",
		"description": "get colors in your node.js console",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "168a4701756b6a7f51a12ce0c97bfa28c084ed63"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/colors@1.1.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/Marak/colors.js"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/Marak/colors.js/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+ssh://git@github.com/Marak/colors.js.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/grunt-legacy-log-utils@2.0.1",
		"name": "grunt-legacy-log-utils",
		"version": "2.0.1",
		"description": "Static methods for the Grunt 0.4.x logger.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "d2f442c7c0150065d9004b08fd7410d37519194e"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/grunt-legacy-log-utils@2.0.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "http://gruntjs.com/"
		  },
		  {
			"type": "issue-tracker",
			"url": "http://github.com/gruntjs/grunt-legacy-log-utils/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/gruntjs/grunt-legacy-log-utils.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/hooker@0.2.3",
		"name": "hooker",
		"version": "0.2.3",
		"description": "Monkey-patch (hook) functions for debugging and stuff.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "b834f723cc4a242aa65963459df6d984c5d3d959"
		  }
		],
		"purl": "pkg:npm/hooker@0.2.3",
		"externalReferences": [
		  {
			"type": "website",
			"url": "http://github.com/cowboy/javascript-hooker"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/cowboy/javascript-hooker/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/cowboy/javascript-hooker.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/grunt-legacy-util@1.1.1",
		"name": "grunt-legacy-util",
		"version": "1.1.1",
		"description": "Some old grunt utils provided for backwards compatibility.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "e10624e7c86034e5b870c8a8616743f0a0845e42"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/grunt-legacy-util@1.1.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "http://gruntjs.com/"
		  },
		  {
			"type": "issue-tracker",
			"url": "http://github.com/gruntjs/grunt-legacy-util/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/gruntjs/grunt-legacy-util.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/async@1.5.2",
		"name": "async",
		"version": "1.5.2",
		"description": "Higher-order functions and common patterns for asynchronous code",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "ec6a61ae56480c0c3cb241c95618e20892f9672a"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/async@1.5.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/caolan/async#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/caolan/async/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/caolan/async.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/getobject@0.1.0",
		"name": "getobject",
		"version": "0.1.0",
		"description": "get.and.set.deep.objects.easily = true",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "047a449789fa160d018f5486ed91320b6ec7885c"
		  }
		],
		"purl": "pkg:npm/getobject@0.1.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/cowboy/node-getobject"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/cowboy/node-getobject/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/cowboy/node-getobject.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/underscore.string@3.3.5",
		"name": "underscore.string",
		"version": "3.3.5",
		"description": "String manipulation extensions for Underscore.js javascript library.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "fc2ad255b8bd309e239cbc5816fd23a9b7ea4023"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/underscore.string@3.3.5",
		"externalReferences": [
		  {
			"type": "website",
			"url": "http://epeli.github.com/underscore.string/"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/epeli/underscore.string/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/epeli/underscore.string.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/sprintf-js@1.1.2",
		"name": "sprintf-js",
		"version": "1.1.2",
		"description": "JavaScript sprintf implementation",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "da1765262bf8c0f571749f2ad6c26300207ae673"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "BSD-3-Clause"
			}
		  }
		],
		"purl": "pkg:npm/sprintf-js@1.1.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/alexei/sprintf.js#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/alexei/sprintf.js/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/alexei/sprintf.js.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/js-yaml@3.14.0",
		"name": "js-yaml",
		"version": "3.14.0",
		"description": "YAML 1.2 parser and serializer",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "a7a34170f26a21bb162424d8adacb4113a69e482"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/js-yaml@3.14.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/nodeca/js-yaml"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/nodeca/js-yaml/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/nodeca/js-yaml.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/argparse@1.0.10",
		"name": "argparse",
		"version": "1.0.10",
		"description": "Very powerful CLI arguments parser. Native port of argparse - python's options parsing library",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "bcd6791ea5ae09725e17e5ad988134cd40b3d911"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/argparse@1.0.10",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/nodeca/argparse#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/nodeca/argparse/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/nodeca/argparse.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/sprintf-js@1.0.3",
		"name": "sprintf-js",
		"version": "1.0.3",
		"description": "JavaScript sprintf implementation",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "04e6926f662895354f3dd015203633b857297e2c"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "BSD-3-Clause"
			}
		  }
		],
		"purl": "pkg:npm/sprintf-js@1.0.3",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/alexei/sprintf.js#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/alexei/sprintf.js/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/alexei/sprintf.js.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/esprima@4.0.1",
		"name": "esprima",
		"version": "4.0.1",
		"description": "ECMAScript parsing infrastructure for multipurpose analysis",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "13b04cdb3e6c5d19df91ab6987a8695619b0aa71"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "BSD-2-Clause"
			}
		  }
		],
		"purl": "pkg:npm/esprima@4.0.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "http://esprima.org"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jquery/esprima/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jquery/esprima.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/mkdirp@1.0.4",
		"name": "mkdirp",
		"version": "1.0.4",
		"description": "Recursively mkdir, like mkdir -p",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "3eb5ed62622756d79a5f0e2a221dfebad75c2f7e"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/mkdirp@1.0.4",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/isaacs/node-mkdirp#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/isaacs/node-mkdirp/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/isaacs/node-mkdirp.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/nopt@3.0.6",
		"name": "nopt",
		"version": "3.0.6",
		"description": "Option parsing for Node, supporting types, shorthands, etc. Used by npm.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "c6465dbf08abcd4db359317f79ac68a646b28ff9"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "ISC"
			}
		  }
		],
		"purl": "pkg:npm/nopt@3.0.6",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/npm/nopt#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/npm/nopt/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/npm/nopt.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/rimraf@3.0.2",
		"name": "rimraf",
		"version": "3.0.2",
		"description": "A deep deletion module for node (like rm -rf)",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "f1a5402ba6220ad52cc1282bac1ae3aa49fd061a"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "ISC"
			}
		  }
		],
		"purl": "pkg:npm/rimraf@3.0.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/isaacs/rimraf#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/isaacs/rimraf/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/isaacs/rimraf.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/grunt-contrib-compress@1.6.0",
		"name": "grunt-contrib-compress",
		"version": "1.6.0",
		"description": "Compress files and folders",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "9708885c738a97a12c5f3072dc97dbc31b4121db"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/grunt-contrib-compress@1.6.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/gruntjs/grunt-contrib-compress#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/gruntjs/grunt-contrib-compress/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/gruntjs/grunt-contrib-compress.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/archiver@1.3.0",
		"name": "archiver",
		"version": "1.3.0",
		"description": "a streaming interface for archive generation",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "4f2194d6d8f99df3f531e6881f14f15d55faaf22"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/archiver@1.3.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/archiverjs/node-archiver"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/archiverjs/node-archiver/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/archiverjs/node-archiver.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/archiver-utils@1.3.0",
		"name": "archiver-utils",
		"version": "1.3.0",
		"description": "utility functions for archiver",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "e50b4c09c70bf3d680e32ff1b7994e9f9d895174"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/archiver-utils@1.3.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/archiverjs/archiver-utils#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/archiverjs/archiver-utils/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/archiverjs/archiver-utils.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/lazystream@1.0.0",
		"name": "lazystream",
		"version": "1.0.0",
		"description": "Open Node Streams on demand.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "f6995fe0f820392f61396be89462407bb77168e4"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/lazystream@1.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jpommerening/node-lazystream"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jpommerening/node-lazystream/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jpommerening/node-lazystream.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/normalize-path@2.1.1",
		"name": "normalize-path",
		"version": "2.1.1",
		"description": "Normalize file path slashes to be unix-like forward slashes. Also condenses repeat slashes to a single slash and removes and trailing slashes unless disabled.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "1ab28b556e198363a8c1a6f7e6fa20137fe6aed9"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/normalize-path@2.1.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/normalize-path"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/normalize-path/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/normalize-path.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/remove-trailing-separator@1.1.0",
		"name": "remove-trailing-separator",
		"version": "1.1.0",
		"description": "Removes separators from the end of the string.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "c24bce2a283adad5bc3f58e0d48249b92379d8ef"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "ISC"
			}
		  }
		],
		"purl": "pkg:npm/remove-trailing-separator@1.1.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/darsain/remove-trailing-separator#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/darsain/remove-trailing-separator/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/darsain/remove-trailing-separator.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/async@2.6.3",
		"name": "async",
		"version": "2.6.3",
		"description": "Higher-order functions and common patterns for asynchronous code",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "d72625e2344a3656e3a3ad4fa749fa83299d82ff"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/async@2.6.3",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://caolan.github.io/async/"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/caolan/async/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/caolan/async.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/walkdir@0.0.11",
		"name": "walkdir",
		"version": "0.0.11",
		"description": "Find files simply. Walks a directory tree emitting events based on what it finds. Presents a familiar callback/emitter/a+sync interface. Walk a tree of any depth.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "a16d025eb931bd03b52f308caed0f40fcebe9532"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/walkdir@0.0.11",
		"externalReferences": [
		  {
			"type": "website",
			"url": "http://github.com/soldair/node-walkdir"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/soldair/node-walkdir/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/soldair/node-walkdir.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/zip-stream@1.2.0",
		"name": "zip-stream",
		"version": "1.2.0",
		"description": "a streaming zip archive generator.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "a8bc45f4c1b49699c6b90198baacaacdbcd4ba04"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/zip-stream@1.2.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/archiverjs/node-zip-stream"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/archiverjs/node-zip-stream/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/archiverjs/node-zip-stream.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/compress-commons@1.2.2",
		"name": "compress-commons",
		"version": "1.2.2",
		"description": "a library that defines a common interface for working with archive formats within node",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "524a9f10903f3a813389b0225d27c48bb751890f"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/compress-commons@1.2.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/archiverjs/node-compress-commons"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/archiverjs/node-compress-commons/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/archiverjs/node-compress-commons.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/crc32-stream@2.0.0",
		"name": "crc32-stream",
		"version": "2.0.0",
		"description": "a streaming CRC32 checksumer",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "e3cdd3b4df3168dd74e3de3fbbcb7b297fe908f4"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/crc32-stream@2.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/archiverjs/node-crc32-stream"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/archiverjs/node-crc32-stream/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/archiverjs/node-crc32-stream.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/crc@3.8.0",
		"name": "crc",
		"version": "3.8.0",
		"description": "Module for calculating Cyclic Redundancy Check (CRC) for Node.js and the Browser.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "ad60269c2c856f8c299e2c4cc0de4556914056c6"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/crc@3.8.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/alexgorbatchev/node-crc"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/alexgorbatchev/node-crc/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/alexgorbatchev/node-crc.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/chalk@1.1.3",
		"name": "chalk",
		"version": "1.1.3",
		"description": "Terminal string styling done right. Much color.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "a8115c55e4a702fe4d150abd3872822a7e09fc98"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/chalk@1.1.3",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/chalk/chalk#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/chalk/chalk/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/chalk/chalk.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/ansi-styles@2.2.1",
		"name": "ansi-styles",
		"version": "2.2.1",
		"description": "ANSI escape codes for styling strings in the terminal",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "b432dd3358b634cf75e1e4664368240533c1ddbe"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/ansi-styles@2.2.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/chalk/ansi-styles#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/chalk/ansi-styles/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/chalk/ansi-styles.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/has-ansi@2.0.0",
		"name": "has-ansi",
		"version": "2.0.0",
		"description": "Check if a string has ANSI escape codes",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "34f5049ce1ecdf2b0649af3ef24e45ed35416d91"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/has-ansi@2.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/sindresorhus/has-ansi#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/sindresorhus/has-ansi/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/sindresorhus/has-ansi.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/ansi-regex@2.1.1",
		"name": "ansi-regex",
		"version": "2.1.1",
		"description": "Regular expression for matching ANSI escape codes",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "c3b33ab5ee360d86e0e628f0468ae7ef27d654df"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/ansi-regex@2.1.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/chalk/ansi-regex#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/chalk/ansi-regex/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/chalk/ansi-regex.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/strip-ansi@3.0.1",
		"name": "strip-ansi",
		"version": "3.0.1",
		"description": "Strip ANSI escape codes",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "6a385fb8853d952d5ff05d0e8aaf94278dc63dcf"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/strip-ansi@3.0.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/chalk/strip-ansi#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/chalk/strip-ansi/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/chalk/strip-ansi.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/supports-color@2.0.0",
		"name": "supports-color",
		"version": "2.0.0",
		"description": "Detect whether a terminal supports color",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "535d045ce6b6363fa40117084629995e9df324c7"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/supports-color@2.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/chalk/supports-color#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/chalk/supports-color/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/chalk/supports-color.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/pretty-bytes@4.0.2",
		"name": "pretty-bytes",
		"version": "4.0.2",
		"description": "Convert bytes to a human readable string: 1337 → 1.34 kB",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "b2bf82e7350d65c6c33aa95aaa5a4f6327f61cd9"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/pretty-bytes@4.0.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/sindresorhus/pretty-bytes#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/sindresorhus/pretty-bytes/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/sindresorhus/pretty-bytes.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/stream-buffers@2.2.0",
		"name": "stream-buffers",
		"version": "2.2.0",
		"description": "Buffer-backed Streams for reading and writing.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "91d5f5130d1cef96dcfa7f726945188741d09ee4"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "Unlicense"
			}
		  }
		],
		"purl": "pkg:npm/stream-buffers@2.2.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/samcday/node-stream-buffer#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/samcday/node-stream-buffer/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/samcday/node-stream-buffer.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/grunt-replace-json@0.1.0",
		"name": "grunt-replace-json",
		"version": "0.1.0",
		"description": "Updates attributes of json files.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "2e58602249181718f744147a9365e4d383ca15af"
		  }
		],
		"purl": "pkg:npm/grunt-replace-json@0.1.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/exo-dev/grunt-replace-json"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/exo-dev/grunt-replace-json/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/exo-dev/grunt-replace-json.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/lodash.set@4.3.2",
		"name": "lodash.set",
		"version": "4.3.2",
		"description": "The lodash method _.set exported as a module.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "d8757b1da807dde24816b0d6a84bea1a76230b23"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/lodash.set@4.3.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://lodash.com/"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/lodash/lodash/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/lodash/lodash.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/hashids@2.2.1",
		"name": "hashids",
		"version": "2.2.1",
		"description": "Generate YouTube-like ids from numbers. Use Hashids when you do not want to expose your database ids to the user.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "ad0c600f0083aa0df7451dfd184e53db34f71289"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/hashids@2.2.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "http://hashids.org/javascript"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/niieani/hashids.js/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/niieani/hashids.js.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/helmet@3.23.3",
		"name": "helmet",
		"version": "3.23.3",
		"description": "help secure Express/Connect apps with various HTTP headers",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "5ba30209c5f73ded4ab65746a3a11bedd4579ab7"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/helmet@3.23.3",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://helmetjs.github.io/"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/helmetjs/helmet/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/helmetjs/helmet.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/depd@2.0.0",
		"name": "depd",
		"version": "2.0.0",
		"description": "Deprecate all the things",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "b696163cc757560d09cf22cc8fad1571b79e76df"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/depd@2.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/dougwilson/nodejs-depd#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/dougwilson/nodejs-depd/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/dougwilson/nodejs-depd.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/dont-sniff-mimetype@1.1.0",
		"name": "dont-sniff-mimetype",
		"version": "1.1.0",
		"description": "Middleware to prevent mimetype from being sniffed",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "c7d0427f8bcb095762751252af59d148b0a623b2"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/dont-sniff-mimetype@1.1.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://helmetjs.github.io/docs/dont-sniff-mimetype"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/helmetjs/dont-sniff-mimetype/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/helmetjs/dont-sniff-mimetype.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/feature-policy@0.3.0",
		"name": "feature-policy",
		"version": "0.3.0",
		"description": "Middleware to set the Feature-Policy HTTP header",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "7430e8e54a40da01156ca30aaec1a381ce536069"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/feature-policy@0.3.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://helmetjs.github.io/docs/feature-policy/"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/helmetjs/feature-policy/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/helmetjs/feature-policy.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/helmet-crossdomain@0.4.0",
		"name": "helmet-crossdomain",
		"version": "0.4.0",
		"description": "Set the X-Permitted-Cross-Domain-Policies header in Express apps",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "5f1fe5a836d0325f1da0a78eaa5fd8429078894e"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/helmet-crossdomain@0.4.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://helmetjs.github.io/docs/crossdomain/"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/helmetjs/crossdomain/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/helmetjs/crossdomain.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/helmet-csp@2.10.0",
		"name": "helmet-csp",
		"version": "2.10.0",
		"description": "Content Security Policy middleware.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "685dde1747bc16c5e28ad9d91e229a69f0a85e84"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/helmet-csp@2.10.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://helmetjs.github.io/docs/csp/"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/helmetjs/csp/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/helmetjs/csp.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/bowser@2.9.0",
		"name": "bowser",
		"version": "2.9.0",
		"description": "Lightweight browser detector",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "3bed854233b419b9a7422d9ee3e85504373821c9"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/bowser@2.9.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/lancedikson/bowser"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/lancedikson/bowser/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/lancedikson/bowser.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/camelize@1.0.0",
		"name": "camelize",
		"version": "1.0.0",
		"description": "recursively transform key strings to camel-case",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "164a5483e630fa4321e5af07020e531831b2609b"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/camelize@1.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/substack/camelize"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/substack/camelize/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/substack/camelize.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/content-security-policy-builder@2.1.0",
		"name": "content-security-policy-builder",
		"version": "2.1.0",
		"description": "Build Content Security Policy directives.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "0a2364d769a3d7014eec79ff7699804deb8cfcbb"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/content-security-policy-builder@2.1.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/helmetjs/content-security-policy-builder"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/helmetjs/content-security-policy-builder/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/helmetjs/content-security-policy-builder.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/dasherize@2.0.0",
		"name": "dasherize",
		"version": "2.0.0",
		"description": "recursively transform key strings to dash-case",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "6d809c9cd0cf7bb8952d80fc84fa13d47ddb1308"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/dasherize@2.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/shahata/dasherize"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/shahata/dasherize/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/shahata/dasherize.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/hide-powered-by@1.1.0",
		"name": "hide-powered-by",
		"version": "1.1.0",
		"description": "Middleware to remove the X-Powered-By header",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "be3ea9cab4bdb16f8744be873755ca663383fa7a"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/hide-powered-by@1.1.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://helmetjs.github.io/docs/hide-powered-by/"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/helmetjs/hide-powered-by/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/helmetjs/hide-powered-by.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/hpkp@2.0.0",
		"name": "hpkp",
		"version": "2.0.0",
		"description": "HTTP Public Key Pinning (HPKP) middleware",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "10e142264e76215a5d30c44ec43de64dee6d1672"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/hpkp@2.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/helmetjs/hpkp#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/helmetjs/hpkp/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/helmetjs/hpkp.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/hsts@2.2.0",
		"name": "hsts",
		"version": "2.2.0",
		"description": "HTTP Strict Transport Security middleware.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "09119d42f7a8587035d027dda4522366fe75d964"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/hsts@2.2.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://helmetjs.github.io/docs/hsts/"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/helmetjs/hsts/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/helmetjs/hsts.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/nocache@2.1.0",
		"name": "nocache",
		"version": "2.1.0",
		"description": "Middleware to destroy caching",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "120c9ffec43b5729b1d5de88cd71aa75a0ba491f"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/nocache@2.1.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://helmetjs.github.io/docs/nocache/"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/helmetjs/nocache/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/helmetjs/nocache.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/referrer-policy@1.2.0",
		"name": "referrer-policy",
		"version": "1.2.0",
		"description": "Middleware to set the Referrer-Policy HTTP header",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "b99cfb8b57090dc454895ef897a4cc35ef67a98e"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/referrer-policy@1.2.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://helmetjs.github.io/docs/referrer-policy/"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/helmetjs/referrer-policy/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/helmetjs/referrer-policy.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/x-xss-protection@1.3.0",
		"name": "x-xss-protection",
		"version": "1.3.0",
		"description": "Middleware to set the X-XSS-Protection header",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "3e3a8dd638da80421b0e9fff11a2dbe168f6d52c"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/x-xss-protection@1.3.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://helmetjs.github.io/docs/xss-filter/"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/helmetjs/x-xss-protection/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/helmetjs/x-xss-protection.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/html-entities@1.3.1",
		"name": "html-entities",
		"version": "1.3.1",
		"description": "Faster HTML entities encode/decode library.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "fb9a1a4b5b14c5daba82d3e34c6ae4fe701a0e44"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/html-entities@1.3.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/mdevils/node-html-entities#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/mdevils/node-html-entities/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/mdevils/node-html-entities.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/i18n@0.8.6",
		"name": "i18n",
		"version": "0.8.6",
		"description": "lightweight translation module with dynamic json storage",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "9608c58bfb8e29db34aa9a09e37dccb9b5666e01"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/i18n@0.8.6",
		"externalReferences": [
		  {
			"type": "website",
			"url": "http://github.com/mashpie/i18n-node"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/mashpie/i18n-node/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+ssh://git@github.com/mashpie/i18n-node.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/make-plural@6.2.1",
		"name": "make-plural",
		"version": "6.2.1",
		"description": "Unicode CLDR pluralization rules as JavaScript functions",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "2790af1d05fb2fc35a111ce759ffdb0aca1339a3"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "ISC"
			}
		  }
		],
		"purl": "pkg:npm/make-plural@6.2.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/eemeli/make-plural#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/eemeli/make-plural/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/eemeli/make-plural.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/math-interval-parser@2.0.1",
		"name": "math-interval-parser",
		"version": "2.0.1",
		"description": "Parse math interval",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "e22cd6d15a0a7f4c03aec560db76513da615bed4"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/math-interval-parser@2.0.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/Semigradsky/math-interval-parser#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/Semigradsky/math-interval-parser/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/Semigradsky/math-interval-parser.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/messageformat@2.3.0",
		"name": "messageformat",
		"version": "2.3.0",
		"description": "PluralFormat and SelectFormat Message and i18n Tool - A JavaScript Implemenation of the ICU standards.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "de263c49029d5eae65d7ee25e0754f57f425ad91"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/messageformat@2.3.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://messageformat.github.io/messageformat/"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/messageformat/messageformat/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/messageformat/messageformat.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/make-plural@4.3.0",
		"name": "make-plural",
		"version": "4.3.0",
		"description": "Translates Unicode CLDR pluralization rules to executable JavaScript",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "f23de08efdb0cac2e0c9ba9f315b0dff6b4c2735"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "ISC"
			}
		  }
		],
		"purl": "pkg:npm/make-plural@4.3.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/eemeli/make-plural#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/eemeli/make-plural/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/eemeli/make-plural.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/messageformat-formatters@2.0.1",
		"name": "messageformat-formatters",
		"version": "2.0.1",
		"description": "Formatters for messageformat",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "0492c1402a48775f751c9b17c0354e92be012b08"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/messageformat-formatters@2.0.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://messageformat.github.io/"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/messageformat/messageformat/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/messageformat/messageformat.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/messageformat-parser@4.1.3",
		"name": "messageformat-parser",
		"version": "4.1.3",
		"description": "A PEG.js parser for ICU MessageFormat strings",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "b824787f57fcda7d50769f5b63e8d4fda68f5b9e"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/messageformat-parser@4.1.3",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://messageformat.github.io/"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/messageformat/messageformat/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/messageformat/messageformat.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/mustache@4.0.1",
		"name": "mustache",
		"version": "4.0.1",
		"description": "Logic-less {{mustache}} templates with JavaScript",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "d99beb031701ad433338e7ea65e0489416c854a2"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/mustache@4.0.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/janl/mustache.js"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/janl/mustache.js/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/janl/mustache.js.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/is-docker@2.0.0",
		"name": "is-docker",
		"version": "2.0.0",
		"description": "Check if the process is running inside a Docker container",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "2cb0df0e75e2d064fe1864c37cdeacb7b2dcf25b"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/is-docker@2.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/sindresorhus/is-docker#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/sindresorhus/is-docker/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/sindresorhus/is-docker.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/is-heroku@2.0.0",
		"name": "is-heroku",
		"version": "2.0.0",
		"description": "Check if your code is running on Heroku",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "6482d1e861435103ae9f69e66f9bd28eb4ea0bca"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/is-heroku@2.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/sindresorhus/is-heroku#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/sindresorhus/is-heroku/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/sindresorhus/is-heroku.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/jsonwebtoken@0.4.0",
		"name": "jsonwebtoken",
		"version": "0.4.0",
		"description": "JSON Web Token implementation (symmetric and asymmetric)",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "7dfa44ac8a588e16e0453c81f11ab6addd0742fe"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/jsonwebtoken@0.4.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/auth0/node-jsonwebtoken#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/auth0/node-jsonwebtoken/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/auth0/node-jsonwebtoken.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/jssha@2.4.2",
		"name": "jssha",
		"version": "2.4.2",
		"description": "jsSHA is a JavaScript implementation of the complete Secure Hash Standard family (SHA-1, SHA-224, SHA3-224, SHA-256, SHA3-256, SHA-384, SHA3-384, SHA-512, SHA3-512, SHAKE128, and SHAKE256) as well as HMAC",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "d950b095634928bd6b2bda1d42da9a3a762d65e9"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "BSD-3-Clause"
			}
		  }
		],
		"purl": "pkg:npm/jssha@2.4.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/Caligatio/jsSHA"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/Caligatio/jsSHA/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/Caligatio/jsSHA.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/libxmljs2@0.25.5",
		"name": "libxmljs2",
		"version": "0.25.5",
		"description": "libxml bindings for v8 javascript engine",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "71f491af11a9df29e04648ea46b14a41b256fe66"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/libxmljs2@0.25.5",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/marudor/libxmljs2#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "http://github.com/marudor/libxmljs2/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+ssh://git@github.com/marudor/libxmljs2.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/bindings@1.5.0",
		"name": "bindings",
		"version": "1.5.0",
		"description": "Helper module for loading your native module's .node file",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "10353c9e945334bc0511a6d90b38fbc7c9c504df"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/bindings@1.5.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/TooTallNate/node-bindings"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/TooTallNate/node-bindings/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/TooTallNate/node-bindings.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/file-uri-to-path@1.0.0",
		"name": "file-uri-to-path",
		"version": "1.0.0",
		"description": "Convert a file: URI to a file path",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "553a7b8446ff6f684359c445f1e37a05dacc33dd"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/file-uri-to-path@1.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/TooTallNate/file-uri-to-path"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/TooTallNate/file-uri-to-path/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/TooTallNate/file-uri-to-path.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/nan@2.14.1",
		"name": "nan",
		"version": "2.14.1",
		"description": "Native Abstractions for Node.js: C++ header for Node 0.8 -> 14 compatibility",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "d7be34dfa3105b91494c3147089315eff8874b01"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/nan@2.14.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/nodejs/nan#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/nodejs/nan/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/nodejs/nan.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/node-pre-gyp@0.15.0",
		"name": "node-pre-gyp",
		"version": "0.15.0",
		"description": "Node.js native addon binary install tool",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "c2fc383276b74c7ffa842925241553e8b40f1087"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "BSD-3-Clause"
			}
		  }
		],
		"purl": "pkg:npm/node-pre-gyp@0.15.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/mapbox/node-pre-gyp#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/mapbox/node-pre-gyp/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/mapbox/node-pre-gyp.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/detect-libc@1.0.3",
		"name": "detect-libc",
		"version": "1.0.3",
		"description": "Node.js module to detect the C standard library (libc) implementation family and version",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "fa137c4bd698edf55cd5cd02ac559f91a4c4ba9b"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "Apache-2.0"
			}
		  }
		],
		"purl": "pkg:npm/detect-libc@1.0.3",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/lovell/detect-libc#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/lovell/detect-libc/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/lovell/detect-libc.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/mkdirp@0.5.5",
		"name": "mkdirp",
		"version": "0.5.5",
		"description": "Recursively mkdir, like mkdir -p",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "d91cefd62d1436ca0f41620e251288d420099def"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/mkdirp@0.5.5",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/substack/node-mkdirp#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/substack/node-mkdirp/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/substack/node-mkdirp.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/needle@2.5.0",
		"name": "needle",
		"version": "2.5.0",
		"description": "The leanest and most handsome HTTP client in the Nodelands.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "e6fc4b3cc6c25caed7554bd613a5cf0bac8c31c0"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/needle@2.5.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/tomas/needle#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/tomas/needle/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/tomas/needle.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/debug@3.2.6",
		"name": "debug",
		"version": "3.2.6",
		"description": "small debugging utility",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "e83d17de16d8a7efb7717edbe5fb10135eee629b"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/debug@3.2.6",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/visionmedia/debug#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/visionmedia/debug/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/visionmedia/debug.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/ms@2.1.2",
		"name": "ms",
		"version": "2.1.2",
		"description": "Tiny millisecond conversion utility",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "d09d1f357b443f493382a8eb3ccd183872ae6009"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/ms@2.1.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/zeit/ms#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/zeit/ms/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/zeit/ms.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/sax@1.2.4",
		"name": "sax",
		"version": "1.2.4",
		"description": "An evented streaming XML parser in JavaScript",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "2816234e2378bddc4e5354fab5caa895df7100d9"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "ISC"
			}
		  }
		],
		"purl": "pkg:npm/sax@1.2.4",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/isaacs/sax-js#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/isaacs/sax-js/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/isaacs/sax-js.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/npm-packlist@1.4.8",
		"name": "npm-packlist",
		"version": "1.4.8",
		"description": "Get a list of the files to add from a folder into an npm package",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "56ee6cc135b9f98ad3d51c1c95da22bbb9b2ef3e"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "ISC"
			}
		  }
		],
		"purl": "pkg:npm/npm-packlist@1.4.8",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://www.npmjs.com/package/npm-packlist"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/npm/npm-packlist/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/npm/npm-packlist.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/ignore-walk@3.0.3",
		"name": "ignore-walk",
		"version": "3.0.3",
		"description": "Nested/recursive .gitignore/.npmignore parsing and filtering.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "017e2447184bfeade7c238e4aefdd1e8f95b1e37"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "ISC"
			}
		  }
		],
		"purl": "pkg:npm/ignore-walk@3.0.3",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/isaacs/ignore-walk#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/isaacs/ignore-walk/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/isaacs/ignore-walk.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/npm-bundled@1.1.1",
		"name": "npm-bundled",
		"version": "1.1.1",
		"description": "list things in node_modules that are bundledDependencies, or transitive dependencies thereof",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "1edd570865a94cdb1bc8220775e29466c9fb234b"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "ISC"
			}
		  }
		],
		"purl": "pkg:npm/npm-bundled@1.1.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/npm/npm-bundled#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/npm/npm-bundled/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/npm/npm-bundled.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/npm-normalize-package-bin@1.0.1",
		"name": "npm-normalize-package-bin",
		"version": "1.0.1",
		"description": "Turn any flavor of allowable package.json bin into a normalized object",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "6e79a41f23fd235c0623218228da7d9c23b8f6e2"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "ISC"
			}
		  }
		],
		"purl": "pkg:npm/npm-normalize-package-bin@1.0.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/npm/npm-normalize-package-bin#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/npm/npm-normalize-package-bin/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/npm/npm-normalize-package-bin.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/npmlog@4.1.2",
		"name": "npmlog",
		"version": "4.1.2",
		"description": "logger for npm",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "08a7f2a8bf734604779a9efa4ad5cc717abb954b"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "ISC"
			}
		  }
		],
		"purl": "pkg:npm/npmlog@4.1.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/npm/npmlog#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/npm/npmlog/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/npm/npmlog.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/are-we-there-yet@1.1.5",
		"name": "are-we-there-yet",
		"version": "1.1.5",
		"description": "Keep track of the overall completion of many disparate processes",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "4b35c2944f062a8bfcda66410760350fe9ddfc21"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "ISC"
			}
		  }
		],
		"purl": "pkg:npm/are-we-there-yet@1.1.5",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/iarna/are-we-there-yet"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/iarna/are-we-there-yet/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/iarna/are-we-there-yet.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/delegates@1.0.0",
		"name": "delegates",
		"version": "1.0.0",
		"description": "delegate methods and accessors to another property",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "84c6e159b81904fdca59a0ef44cd870d31250f9a"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/delegates@1.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/visionmedia/node-delegates#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/visionmedia/node-delegates/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/visionmedia/node-delegates.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/console-control-strings@1.1.0",
		"name": "console-control-strings",
		"version": "1.1.0",
		"description": "A library of cross-platform tested terminal/console command strings for doing things like color and cursor positioning.  This is a subset of both ansi and vt100.  All control codes included work on both Windows & Unix-like OSes, except where noted.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "3d7cf4464db6446ea644bf4b39507f9851008e8e"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "ISC"
			}
		  }
		],
		"purl": "pkg:npm/console-control-strings@1.1.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/iarna/console-control-strings#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/iarna/console-control-strings/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/iarna/console-control-strings.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/gauge@2.7.4",
		"name": "gauge",
		"version": "2.7.4",
		"description": "A terminal based horizontal guage",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "2c03405c7538c39d7eb37b317022e325fb018bf7"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "ISC"
			}
		  }
		],
		"purl": "pkg:npm/gauge@2.7.4",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/iarna/gauge"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/iarna/gauge/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/iarna/gauge.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/aproba@1.2.0",
		"name": "aproba",
		"version": "1.2.0",
		"description": "A ridiculously light-weight argument validator (now browser friendly)",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "6802e6264efd18c790a1b0d517f0f2627bf2c94a"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "ISC"
			}
		  }
		],
		"purl": "pkg:npm/aproba@1.2.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/iarna/aproba"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/iarna/aproba/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/iarna/aproba.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/has-unicode@2.0.1",
		"name": "has-unicode",
		"version": "2.0.1",
		"description": "Try to guess if your terminal supports unicode",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "e0e6fe6a28cf51138855e086d1691e771de2a8b9"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "ISC"
			}
		  }
		],
		"purl": "pkg:npm/has-unicode@2.0.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/iarna/has-unicode"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/iarna/has-unicode/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/iarna/has-unicode.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/signal-exit@3.0.3",
		"name": "signal-exit",
		"version": "3.0.3",
		"description": "when you want to fire an event no matter how a process exits.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "a1410c2edd8f077b08b4e253c8eacfcaf057461c"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "ISC"
			}
		  }
		],
		"purl": "pkg:npm/signal-exit@3.0.3",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/tapjs/signal-exit"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/tapjs/signal-exit/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/tapjs/signal-exit.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/string-width@1.0.2",
		"name": "string-width",
		"version": "1.0.2",
		"description": "Get the visual width of a string - the number of columns required to display it",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "118bdf5b8cdc51a2a7e70d211e07e2b0b9b107d3"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/string-width@1.0.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/sindresorhus/string-width#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/sindresorhus/string-width/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/sindresorhus/string-width.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/code-point-at@1.1.0",
		"name": "code-point-at",
		"version": "1.1.0",
		"description": "ES2015 String#codePointAt() ponyfill",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "0d070b4d043a5bea33a2f1a40e2edb3d9a4ccf77"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/code-point-at@1.1.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/sindresorhus/code-point-at#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/sindresorhus/code-point-at/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/sindresorhus/code-point-at.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/is-fullwidth-code-point@1.0.0",
		"name": "is-fullwidth-code-point",
		"version": "1.0.0",
		"description": "Check if the character represented by a given Unicode code point is fullwidth",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "ef9e31386f031a7f0d643af82fde50c457ef00cb"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/is-fullwidth-code-point@1.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/sindresorhus/is-fullwidth-code-point#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/sindresorhus/is-fullwidth-code-point/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/sindresorhus/is-fullwidth-code-point.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/number-is-nan@1.0.1",
		"name": "number-is-nan",
		"version": "1.0.1",
		"description": "ES2015 Number.isNaN() ponyfill",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "097b602b53422a522c1afb8790318336941a011d"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/number-is-nan@1.0.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/sindresorhus/number-is-nan#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/sindresorhus/number-is-nan/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/sindresorhus/number-is-nan.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/wide-align@1.1.3",
		"name": "wide-align",
		"version": "1.1.3",
		"description": "A wide-character aware text alignment function for use on the console or with fixed width fonts.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "ae074e6bdc0c14a431e804e624549c633b000457"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "ISC"
			}
		  }
		],
		"purl": "pkg:npm/wide-align@1.1.3",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/iarna/wide-align#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/iarna/wide-align/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/iarna/wide-align.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/string-width@2.1.1",
		"name": "string-width",
		"version": "2.1.1",
		"description": "Get the visual width of a string - the number of columns required to display it",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "ab93f27a8dc13d28cac815c462143a6d9012ae9e"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/string-width@2.1.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/sindresorhus/string-width#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/sindresorhus/string-width/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/sindresorhus/string-width.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/strip-ansi@4.0.0",
		"name": "strip-ansi",
		"version": "4.0.0",
		"description": "Strip ANSI escape codes",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "a8479022eb1ac368a871389b635262c505ee368f"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/strip-ansi@4.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/chalk/strip-ansi#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/chalk/strip-ansi/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/chalk/strip-ansi.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/ansi-regex@3.0.0",
		"name": "ansi-regex",
		"version": "3.0.0",
		"description": "Regular expression for matching ANSI escape codes",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "ed0317c322064f79466c02966bddb605ab37d998"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/ansi-regex@3.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/chalk/ansi-regex#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/chalk/ansi-regex/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/chalk/ansi-regex.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/rc@1.2.8",
		"name": "rc",
		"version": "1.2.8",
		"description": "hardwired configuration loader",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "cd924bf5200a075b83c188cd6b9e211b7fc0d3ed"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "name": "(BSD-2-Clause OR MIT OR Apache-2.0)"
			}
		  }
		],
		"purl": "pkg:npm/rc@1.2.8",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/dominictarr/rc#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/dominictarr/rc/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/dominictarr/rc.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/deep-extend@0.6.0",
		"name": "deep-extend",
		"version": "0.6.0",
		"description": "Recursive object extending",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "c4fa7c95404a17a9c3e8ca7e1537312b736330ac"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/deep-extend@0.6.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/unclechu/node-deep-extend"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/unclechu/node-deep-extend/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/unclechu/node-deep-extend.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/strip-json-comments@2.0.1",
		"name": "strip-json-comments",
		"version": "2.0.1",
		"description": "Strip comments from JSON. Lets you use comments in your JSON files!",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "3c531942e908c2697c0ec344858c286c7ca0a60a"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/strip-json-comments@2.0.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/sindresorhus/strip-json-comments#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/sindresorhus/strip-json-comments/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/sindresorhus/strip-json-comments.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/rimraf@2.7.1",
		"name": "rimraf",
		"version": "2.7.1",
		"description": "A deep deletion module for node (like rm -rf)",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "35797f13a7fdadc566142c29d4f07ccad483e3ec"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "ISC"
			}
		  }
		],
		"purl": "pkg:npm/rimraf@2.7.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/isaacs/rimraf#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/isaacs/rimraf/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/isaacs/rimraf.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/tar@4.4.13",
		"name": "tar",
		"version": "4.4.13",
		"description": "tar for node",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "43b364bc52888d555298637b10d60790254ab525"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "ISC"
			}
		  }
		],
		"purl": "pkg:npm/tar@4.4.13",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/npm/node-tar#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/npm/node-tar/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/npm/node-tar.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/chownr@1.1.4",
		"name": "chownr",
		"version": "1.1.4",
		"description": "like chown -R",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "6fc9d7b42d32a583596337666e7d08084da2cc6b"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "ISC"
			}
		  }
		],
		"purl": "pkg:npm/chownr@1.1.4",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/isaacs/chownr#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/isaacs/chownr/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/isaacs/chownr.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/fs-minipass@1.2.7",
		"name": "fs-minipass",
		"version": "1.2.7",
		"description": "fs read and write streams based on minipass",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "ccff8570841e7fe4265693da88936c55aed7f7c7"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "ISC"
			}
		  }
		],
		"purl": "pkg:npm/fs-minipass@1.2.7",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/npm/fs-minipass#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/npm/fs-minipass/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/npm/fs-minipass.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/minipass@2.9.0",
		"name": "minipass",
		"version": "2.9.0",
		"description": "minimal implementation of a PassThrough stream",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "e713762e7d3e32fed803115cf93e04bca9fcc9a6"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "ISC"
			}
		  }
		],
		"purl": "pkg:npm/minipass@2.9.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/isaacs/minipass#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/isaacs/minipass/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/isaacs/minipass.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/yallist@3.1.1",
		"name": "yallist",
		"version": "3.1.1",
		"description": "Yet Another Linked List",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "dbb7daf9bfd8bac9ab45ebf602b8cbad0d5d08fd"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "ISC"
			}
		  }
		],
		"purl": "pkg:npm/yallist@3.1.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/isaacs/yallist#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/isaacs/yallist/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/isaacs/yallist.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/minizlib@1.3.3",
		"name": "minizlib",
		"version": "1.3.3",
		"description": "A small fast zlib stream built on [minipass](http://npm.im/minipass) and Node.js's zlib binding.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "2290de96818a34c29551c8a8d301216bd65a861d"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/minizlib@1.3.3",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/isaacs/minizlib#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/isaacs/minizlib/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/isaacs/minizlib.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/marsdb@0.6.11",
		"name": "marsdb",
		"version": "0.6.11",
		"description": "MarsDB is a lightweight client-side MongoDB-like database, Promise based, written in ES6",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "005b4af67e41df4efe73efed577adf29135d9712"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/marsdb@0.6.11",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/c58/marsdb"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/c58/marsdb/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+ssh://git@github.com/c58/marsdb.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/check-types@6.0.0",
		"name": "check-types",
		"version": "6.0.0",
		"description": "A little library for asserting types and values.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "de12a2ffc295df40b3afca2bcfeb831bc5f5edf0"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/check-types@6.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/philbooth/check-types.js"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/philbooth/check-types.js/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/philbooth/check-types.js.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/double-ended-queue@0.9.7",
		"name": "double-ended-queue",
		"version": "0.9.7",
		"description": "Extremely fast double-ended queue implementation",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "8ae0a7265df66cdc3f07dce558e9716adb586ab8"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/double-ended-queue@0.9.7",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/petkaantonov/deque"
		  },
		  {
			"type": "issue-tracker",
			"url": "http://github.com/petkaantonov/deque/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/petkaantonov/deque.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/eventemitter3@1.1.1",
		"name": "eventemitter3",
		"version": "1.1.1",
		"description": "EventEmitter3 focuses on performance while maintaining a Node.js AND browser compatible interface.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "47786bdaa087caf7b1b75e73abc5c7d540158cd0"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/eventemitter3@1.1.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/primus/eventemitter3#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/primus/eventemitter3/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/primus/eventemitter3.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/fast.js@0.1.1",
		"name": "fast.js",
		"version": "0.1.1",
		"description": "Faster user-land reimplementations of native functions with extra helpers.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "7c024d55ae144882fbcee44b79005fe2dcabd9fe"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/fast.js@0.1.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/codemix/fast.js"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/codemix/fast.js/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+ssh://git@github.com/codemix/fast.js.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/geojson-utils@1.1.0",
		"name": "geojson-utils",
		"version": "1.1.0",
		"description": "GeoJSON Utilities for JavaScript",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "e8ffb4c81c0a75b3e306f5187265d6f23040f50b"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/geojson-utils@1.1.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/maxogden/geojson-js-utils#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/maxogden/geojson-js-utils/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/maxogden/geojson-js-utils.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/invariant@2.2.4",
		"name": "invariant",
		"version": "2.2.4",
		"description": "invariant",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "610f3c92c9359ce1db616e538008d23ff35158e6"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/invariant@2.2.4",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/zertosh/invariant#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/zertosh/invariant/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/zertosh/invariant.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/loose-envify@1.4.0",
		"name": "loose-envify",
		"version": "1.4.0",
		"description": "Fast (and loose) selective process.env replacer using js-tokens instead of an AST",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "71ee51fa7be4caec1a63839f7e682d8132d30caf"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/loose-envify@1.4.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/zertosh/loose-envify"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/zertosh/loose-envify/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/zertosh/loose-envify.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/js-tokens@4.0.0",
		"name": "js-tokens",
		"version": "4.0.0",
		"description": "A regex that tokenizes JavaScript.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "19203fb59991df98e3a287050d4647cdeaf32499"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/js-tokens@4.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/lydell/js-tokens#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/lydell/js-tokens/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/lydell/js-tokens.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/morgan@1.10.0",
		"name": "morgan",
		"version": "1.10.0",
		"description": "HTTP request logger middleware for node.js",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "091778abc1fc47cd3509824653dae1faab6b17d7"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/morgan@1.10.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/expressjs/morgan#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/expressjs/morgan/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/expressjs/morgan.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/basic-auth@2.0.1",
		"name": "basic-auth",
		"version": "2.0.1",
		"description": "node.js basic auth parser",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "b998279bf47ce38344b4f3cf916d4679bbf51e3a"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/basic-auth@2.0.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jshttp/basic-auth#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jshttp/basic-auth/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jshttp/basic-auth.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/multer@1.4.2",
		"name": "multer",
		"version": "1.4.2",
		"description": "Middleware for handling multipart/form-data.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "2f1f4d12dbaeeba74cb37e623f234bf4d3d2057a"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/multer@1.4.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/expressjs/multer#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/expressjs/multer/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/expressjs/multer.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/append-field@1.0.0",
		"name": "append-field",
		"version": "1.0.0",
		"description": "A [W3C HTML JSON forms spec](http://www.w3.org/TR/html-json-forms/) compliant field appender (for lack of a better name). Useful for people implementing application/x-www-form-urlencoded and multipart/form-data parsers.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "1e3440e915f0b1203d23748e78edd7b9b5b43e56"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/append-field@1.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/LinusU/node-append-field#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/LinusU/node-append-field/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+ssh://git@github.com/LinusU/node-append-field.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/busboy@0.2.14",
		"name": "busboy",
		"version": "0.2.14",
		"description": "A streaming parser for HTML form data for node.js",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "6c2a622efcf47c57bbbe1e2a9c37ad36c7925453"
		  }
		],
		"purl": "pkg:npm/busboy@0.2.14",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/mscdex/busboy#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/mscdex/busboy/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+ssh://git@github.com/mscdex/busboy.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/dicer@0.2.5",
		"name": "dicer",
		"version": "0.2.5",
		"description": "A very fast streaming multipart parser for node.js",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "5996c086bb33218c812c090bddc09cd12facb70f"
		  }
		],
		"purl": "pkg:npm/dicer@0.2.5",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/mscdex/dicer#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/mscdex/dicer/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+ssh://git@github.com/mscdex/dicer.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/readable-stream@1.1.14",
		"name": "readable-stream",
		"version": "1.1.14",
		"description": "Streams3, a user-land copy of the stream library from Node.js v0.11.x",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "7cf4c54ef648e3813084c636dd2079e166c081d9"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/readable-stream@1.1.14",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/isaacs/readable-stream#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/isaacs/readable-stream/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/isaacs/readable-stream.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/isarray@0.0.1",
		"name": "isarray",
		"version": "0.0.1",
		"description": "Array#isArray for older browsers",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "8a18acfca9a8f4177e09abfc6038939b05d1eedf"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/isarray@0.0.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/juliangruber/isarray"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/juliangruber/isarray/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/juliangruber/isarray.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/string_decoder@0.10.31",
		"name": "string_decoder",
		"version": "0.10.31",
		"description": "The string_decoder module from Node core",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "62e203bc41766c6c28c9fc84301dab1c5310fa94"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/string_decoder@0.10.31",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/rvagg/string_decoder"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/rvagg/string_decoder/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/rvagg/string_decoder.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/streamsearch@0.1.2",
		"name": "streamsearch",
		"version": "0.1.2",
		"description": "Streaming Boyer-Moore-Horspool searching for node.js",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "808b9d0e56fc273d809ba57338e929919a1a9f1a"
		  }
		],
		"purl": "pkg:npm/streamsearch@0.1.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/mscdex/streamsearch#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/mscdex/streamsearch/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+ssh://git@github.com/mscdex/streamsearch.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/concat-stream@1.6.2",
		"name": "concat-stream",
		"version": "1.6.2",
		"description": "writable stream that concatenates strings or binary data and calls a callback with the result",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "904bdf194cd3122fc675c77fc4ac3d4ff0fd1a34"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/concat-stream@1.6.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/maxogden/concat-stream#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "http://github.com/maxogden/concat-stream/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+ssh://git@github.com/maxogden/concat-stream.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/buffer-from@1.1.1",
		"name": "buffer-from",
		"version": "1.1.1",
		"description": "A [ponyfill](https://ponyfill.com) for Buffer.from, uses native implementation if available.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "32713bc028f75c02fdb710d7c7bcec1f2c6070ef"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/buffer-from@1.1.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/LinusU/buffer-from#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/LinusU/buffer-from/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/LinusU/buffer-from.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/typedarray@0.0.6",
		"name": "typedarray",
		"version": "0.0.6",
		"description": "TypedArray polyfill for old browsers",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "867ac74e3864187b1d3d47d996a78ec5c8830777"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/typedarray@0.0.6",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/substack/typedarray"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/substack/typedarray/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/substack/typedarray.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/node-pre-gyp@0.14.0",
		"name": "node-pre-gyp",
		"version": "0.14.0",
		"description": "Node.js native addon binary install tool",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "9a0596533b877289bcad4e143982ca3d904ddc83"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "BSD-3-Clause"
			}
		  }
		],
		"purl": "pkg:npm/node-pre-gyp@0.14.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/mapbox/node-pre-gyp#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/mapbox/node-pre-gyp/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/mapbox/node-pre-gyp.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/notevil@1.3.3",
		"name": "notevil",
		"version": "1.3.3",
		"description": "Evalulate javascript like the built-in eval() method but safely",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "56b8a935d8978e0c000749621aca3928b823cb01"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/notevil@1.3.3",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/mmckegg/notevil#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/mmckegg/notevil/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/mmckegg/notevil.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/esprima@1.0.4",
		"name": "esprima",
		"version": "1.0.4",
		"description": "ECMAScript parsing infrastructure for multipurpose analysis",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "9f557e08fc3b4d26ece9dd34f8fbf476b62585ad"
		  }
		],
		"purl": "pkg:npm/esprima@1.0.4",
		"externalReferences": [
		  {
			"type": "website",
			"url": "http://esprima.org"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/ariya/esprima/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+ssh://git@github.com/ariya/esprima.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/hoister@0.0.2",
		"name": "hoister",
		"version": "0.0.2",
		"description": "Put all function and variable declarations at the top of the scope in an AST",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "0d0d8b1ce0f191553e61afec654f6b180eb96e5d"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/hoister@0.0.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/mmckegg/hoister#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/mmckegg/hoister/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/mmckegg/hoister.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/otplib@11.0.1",
		"name": "otplib",
		"version": "11.0.1",
		"description": "HMAC-based (HOTP) and Time-based (TOTP) One-Time Password library",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "7d64aa87029f07c99c7f96819fb10cdb67dea886"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/otplib@11.0.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://yeojz.github.io/otplib"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/yeojz/otplib/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/yeojz/otplib.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/thirty-two@1.0.2",
		"name": "thirty-two",
		"version": "1.0.2",
		"description": "Implementation RFC 3548 Base32 encoding/decoding for node.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "4ca2fffc02a51290d2744b9e3f557693ca6b627a"
		  }
		],
		"purl": "pkg:npm/thirty-two@1.0.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/chrisumbel/thirty-two#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/chrisumbel/thirty-two/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/chrisumbel/thirty-two.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/pdfkit@0.11.0",
		"name": "pdfkit",
		"version": "0.11.0",
		"description": "A PDF generation library for Node.js",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "9cdb2fc42bd2913587fe3ddf48cc5bbb3c36f7de"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/pdfkit@0.11.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "http://pdfkit.org/"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/foliojs/pdfkit/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/foliojs/pdfkit.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/crypto-js@3.3.0",
		"name": "crypto-js",
		"version": "3.3.0",
		"description": "JavaScript library of crypto standards.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "846dd1cce2f68aacfa156c8578f926a609b7976b"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/crypto-js@3.3.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "http://github.com/brix/crypto-js"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/brix/crypto-js/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+ssh://git@github.com/brix/crypto-js.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/fontkit@1.8.1",
		"name": "fontkit",
		"version": "1.8.1",
		"description": "An advanced font engine for Node and the browser",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "ae77485376f1096b45548bf6ced9a07af62a7846"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/fontkit@1.8.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/foliojs/fontkit#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/foliojs/fontkit/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/foliojs/fontkit.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/babel-runtime@6.26.0",
		"name": "babel-runtime",
		"version": "6.26.0",
		"description": "babel selfContained runtime",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "965c7058668e82b55d7bfe04ff2337bc8b5647fe"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/babel-runtime@6.26.0",
		"externalReferences": [
		  {
			"type": "vcs",
			"url": "https://github.com/babel/babel/tree/master/packages/babel-runtime"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/core-js@2.6.11",
		"name": "core-js",
		"version": "2.6.11",
		"description": "Standard library",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "38831469f9922bded8ee21c9dc46985e0399308c"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/core-js@2.6.11",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/zloirock/core-js#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/zloirock/core-js/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/zloirock/core-js.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/regenerator-runtime@0.11.1",
		"name": "regenerator-runtime",
		"version": "0.11.1",
		"description": "Runtime for Regenerator-compiled generator and async functions.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "be05ad7f9bf7d22e056f9726cee5017fbf19e2e9"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/regenerator-runtime@0.11.1",
		"externalReferences": [
		  {
			"type": "vcs",
			"url": "https://github.com/facebook/regenerator/tree/master/packages/regenerator-runtime"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/brfs@2.0.2",
		"name": "brfs",
		"version": "2.0.2",
		"description": "browserify fs.readFileSync() static asset inliner",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "44237878fa82aa479ce4f5fe2c1796ec69f07845"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/brfs@2.0.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/substack/brfs"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/substack/brfs/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/substack/brfs.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/quote-stream@1.0.2",
		"name": "quote-stream",
		"version": "1.0.2",
		"description": "transform a stream into a quoted string",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "84963f8c9c26b942e153feeb53aae74652b7e0b2"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/quote-stream@1.0.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/substack/quote-stream"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/substack/quote-stream/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/substack/quote-stream.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/buffer-equal@0.0.1",
		"name": "buffer-equal",
		"version": "0.0.1",
		"description": "return whether two buffers are equal",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "91bc74b11ea405bc916bc6aa908faafa5b4aac4b"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/buffer-equal@0.0.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/substack/node-buffer-equal#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/substack/node-buffer-equal/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/substack/node-buffer-equal.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/through2@2.0.5",
		"name": "through2",
		"version": "2.0.5",
		"description": "A tiny wrapper around Node streams2 Transform to avoid explicit subclassing noise",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "01c1e39eb31d07cb7d03a96a70823260b23132cd"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/through2@2.0.5",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/rvagg/through2#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/rvagg/through2/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/rvagg/through2.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/static-module@3.0.4",
		"name": "static-module",
		"version": "3.0.4",
		"description": "convert module usage to inline expressions",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "bfbd1d1c38dd1fbbf0bb4af0c1b3ae18a93a2b68"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/static-module@3.0.4",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/substack/static-module"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/substack/static-module/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/substack/static-module.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/acorn-node@1.8.2",
		"name": "acorn-node",
		"version": "1.8.2",
		"description": "the acorn javascript parser, preloaded with plugins for syntax parity with recent node versions",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "114c95d64539e53dede23de8b9d96df7c7ae2af8"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "Apache-2.0"
			}
		  }
		],
		"purl": "pkg:npm/acorn-node@1.8.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/browserify/acorn-node"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/browserify/acorn-node/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/browserify/acorn-node.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/acorn@7.3.1",
		"name": "acorn",
		"version": "7.3.1",
		"description": "ECMAScript parser",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "85010754db53c3fbaf3b9ea3e083aa5c5d147ffd"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/acorn@7.3.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/acornjs/acorn"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/acornjs/acorn/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/acornjs/acorn.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/acorn-walk@7.2.0",
		"name": "acorn-walk",
		"version": "7.2.0",
		"description": "ECMAScript (ESTree) AST walker",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "0de889a601203909b0fbe07b8938dc21d2e967bc"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/acorn-walk@7.2.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/acornjs/acorn"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/acornjs/acorn/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/acornjs/acorn.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/convert-source-map@1.7.0",
		"name": "convert-source-map",
		"version": "1.7.0",
		"description": "Converts a source-map from/to  different formats and allows adding/changing properties.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "17a2cb882d7f77d3490585e2ce6c524424a3a442"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/convert-source-map@1.7.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/thlorenz/convert-source-map"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/thlorenz/convert-source-map/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/thlorenz/convert-source-map.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/duplexer2@0.1.4",
		"name": "duplexer2",
		"version": "0.1.4",
		"description": "Like duplexer but using streams3",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "8b12dab878c0d69e3e7891051662a32fc6bddcc1"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "BSD-3-Clause"
			}
		  }
		],
		"purl": "pkg:npm/duplexer2@0.1.4",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/deoxxa/duplexer2#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/deoxxa/duplexer2/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/deoxxa/duplexer2.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/escodegen@1.14.3",
		"name": "escodegen",
		"version": "1.14.3",
		"description": "ECMAScript code generator",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "4e7b81fba61581dc97582ed78cab7f0e8d63f503"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "BSD-2-Clause"
			}
		  }
		],
		"purl": "pkg:npm/escodegen@1.14.3",
		"externalReferences": [
		  {
			"type": "website",
			"url": "http://github.com/estools/escodegen"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/estools/escodegen/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+ssh://git@github.com/estools/escodegen.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/estraverse@4.3.0",
		"name": "estraverse",
		"version": "4.3.0",
		"description": "ECMAScript JS AST traversal functions",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "398ad3f3c5a24948be7725e83d11a7de28cdbd1d"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "BSD-2-Clause"
			}
		  }
		],
		"purl": "pkg:npm/estraverse@4.3.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/estools/estraverse"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/estools/estraverse/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+ssh://git@github.com/estools/estraverse.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/esutils@2.0.3",
		"name": "esutils",
		"version": "2.0.3",
		"description": "utility box for ECMAScript language tools",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "74d2eb4de0b8da1293711910d50775b9b710ef64"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "BSD-2-Clause"
			}
		  }
		],
		"purl": "pkg:npm/esutils@2.0.3",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/estools/esutils"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/estools/esutils/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+ssh://git@github.com/estools/esutils.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/optionator@0.8.3",
		"name": "optionator",
		"version": "0.8.3",
		"description": "option parsing and help generation",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "84fa1d036fe9d3c7e21d99884b601167ec8fb495"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/optionator@0.8.3",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/gkz/optionator"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/gkz/optionator/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/gkz/optionator.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/deep-is@0.1.3",
		"name": "deep-is",
		"version": "0.1.3",
		"description": "node's assert.deepEqual algorithm except for NaN being equal to NaN",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "b369d6fb5dbc13eecf524f91b070feedc357cf34"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/deep-is@0.1.3",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/thlorenz/deep-is#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/thlorenz/deep-is/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+ssh://git@github.com/thlorenz/deep-is.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/fast-levenshtein@2.0.6",
		"name": "fast-levenshtein",
		"version": "2.0.6",
		"description": "Efficient implementation of Levenshtein algorithm  with locale-specific collator support.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "3d8a5c66883a16a30ca8643e851f19baa7797917"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/fast-levenshtein@2.0.6",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/hiddentao/fast-levenshtein#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/hiddentao/fast-levenshtein/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/hiddentao/fast-levenshtein.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/levn@0.3.0",
		"name": "levn",
		"version": "0.3.0",
		"description": "Light ECMAScript (JavaScript) Value Notation - human written, concise, typed, flexible",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "3b09924edf9f083c0490fdd4c0bc4421e04764ee"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/levn@0.3.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/gkz/levn"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/gkz/levn/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/gkz/levn.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/prelude-ls@1.1.2",
		"name": "prelude-ls",
		"version": "1.1.2",
		"description": "prelude.ls is a functionally oriented utility library. It is powerful and flexible. Almost all of its functions are curried. It is written in, and is the recommended base library for, LiveScript.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "21932a549f5e52ffd9a827f570e04be62a97da54"
		  }
		],
		"purl": "pkg:npm/prelude-ls@1.1.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "http://preludels.com"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/gkz/prelude-ls/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/gkz/prelude-ls.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/type-check@0.3.2",
		"name": "type-check",
		"version": "0.3.2",
		"description": "type-check allows you to check the types of JavaScript values at runtime with a Haskell like type syntax.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "5884cab512cf1d355e3fb784f30804b2b520db72"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/type-check@0.3.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/gkz/type-check"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/gkz/type-check/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/gkz/type-check.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/word-wrap@1.2.3",
		"name": "word-wrap",
		"version": "1.2.3",
		"description": "Wrap words to a specified length.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "610636f6b1f703891bd34771ccb17fb93b47079c"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/word-wrap@1.2.3",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/word-wrap"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/word-wrap/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/word-wrap.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/source-map@0.6.1",
		"name": "source-map",
		"version": "0.6.1",
		"description": "Generates and consumes source maps",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "74722af32e9614e9c287a8d0bbde48b5e2f1a263"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "BSD-3-Clause"
			}
		  }
		],
		"purl": "pkg:npm/source-map@0.6.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/mozilla/source-map"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/mozilla/source-map/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+ssh://git@github.com/mozilla/source-map.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/has@1.0.3",
		"name": "has",
		"version": "1.0.3",
		"description": "Object.prototype.hasOwnProperty.call shortcut",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "722d7cbfc1f6aa8241f16dd814e011e1f41e8796"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/has@1.0.3",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/tarruda/has"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/tarruda/has/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/tarruda/has.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/function-bind@1.1.1",
		"name": "function-bind",
		"version": "1.1.1",
		"description": "Implementation of Function.prototype.bind",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "a56899d3ea3c9bab874bb9773b7c5ede92f4895d"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/function-bind@1.1.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/Raynos/function-bind"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/Raynos/function-bind/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/Raynos/function-bind.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/magic-string@0.25.1",
		"name": "magic-string",
		"version": "0.25.1",
		"description": "Modify strings, generate sourcemaps",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "b1c248b399cd7485da0fe7385c2fc7011843266e"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/magic-string@0.25.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/rich-harris/magic-string#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/rich-harris/magic-string/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/rich-harris/magic-string.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/sourcemap-codec@1.4.8",
		"name": "sourcemap-codec",
		"version": "1.4.8",
		"description": "Encode/decode sourcemap mappings",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "ea804bd94857402e6992d05a38ef1ae35a9ab4c4"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/sourcemap-codec@1.4.8",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/Rich-Harris/sourcemap-codec"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/Rich-Harris/sourcemap-codec/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/Rich-Harris/sourcemap-codec.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/merge-source-map@1.0.4",
		"name": "merge-source-map",
		"version": "1.0.4",
		"description": "Merge old source map and new source map in multi-transform flow",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "a5de46538dae84d4114cc5ea02b4772a6346701f"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/merge-source-map@1.0.4",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/keik/merge-source-map#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/keik/merge-source-map/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/keik/merge-source-map.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/object-inspect@1.8.0",
		"name": "object-inspect",
		"version": "1.8.0",
		"description": "string representations of objects in node and the browser",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "df807e5ecf53a609cc6bfe93eac3cc7be5b3a9d0"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/object-inspect@1.8.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/inspect-js/object-inspect"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/inspect-js/object-inspect/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/inspect-js/object-inspect.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/scope-analyzer@2.1.1",
		"name": "scope-analyzer",
		"version": "2.1.1",
		"description": "simple scope analysis for javascript ASTs",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "5156c27de084d74bf75af9e9506aaf95c6e73dd6"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "Apache-2.0"
			}
		  }
		],
		"purl": "pkg:npm/scope-analyzer@2.1.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/goto-bus-stop/scope-analyzer"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/goto-bus-stop/scope-analyzer/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/goto-bus-stop/scope-analyzer.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/array-from@2.1.1",
		"name": "array-from",
		"version": "2.1.1",
		"description": "A ponyfill for the ES 2015 (ES6) Array.from().",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "cfe9d8c26628b9dc5aecc62a9f5d8f1f352c1195"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/array-from@2.1.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/studio-b12/array-from#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/studio-b12/array-from/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+ssh://git@github.com/studio-b12/array-from.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/dash-ast@1.0.0",
		"name": "dash-ast",
		"version": "1.0.0",
		"description": "walk an AST, quickly",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "12029ba5fb2f8aa6f0a861795b23c1b4b6c27d37"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "Apache-2.0"
			}
		  }
		],
		"purl": "pkg:npm/dash-ast@1.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/goto-bus-stop/dash-ast"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/goto-bus-stop/dash-ast/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/goto-bus-stop/dash-ast.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/es6-map@0.1.5",
		"name": "es6-map",
		"version": "0.1.5",
		"description": "ECMAScript6 Map polyfill",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "9136e0503dcc06a301690f0bb14ff4e364e949f0"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/es6-map@0.1.5",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/medikoo/es6-map#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/medikoo/es6-map/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/medikoo/es6-map.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/d@1.0.1",
		"name": "d",
		"version": "1.0.1",
		"description": "Property descriptor factory",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "8698095372d58dbee346ffd0c7093f99f8f9eb5a"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "ISC"
			}
		  }
		],
		"purl": "pkg:npm/d@1.0.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/medikoo/d#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/medikoo/d/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/medikoo/d.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/es5-ext@0.10.53",
		"name": "es5-ext",
		"version": "0.10.53",
		"description": "ECMAScript extensions and shims",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "93c5a3acfdbef275220ad72644ad02ee18368de1"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "ISC"
			}
		  }
		],
		"purl": "pkg:npm/es5-ext@0.10.53",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/medikoo/es5-ext#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/medikoo/es5-ext/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/medikoo/es5-ext.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/es6-iterator@2.0.3",
		"name": "es6-iterator",
		"version": "2.0.3",
		"description": "Iterator abstraction based on ES6 specification",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "a7de889141a05a94b0854403b2d0a0fbfa98f3b7"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/es6-iterator@2.0.3",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/medikoo/es6-iterator#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/medikoo/es6-iterator/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/medikoo/es6-iterator.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/es6-symbol@3.1.3",
		"name": "es6-symbol",
		"version": "3.1.3",
		"description": "ECMAScript 6 Symbol polyfill",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "bad5d3c1bcdac28269f4cb331e431c78ac705d18"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "ISC"
			}
		  }
		],
		"purl": "pkg:npm/es6-symbol@3.1.3",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/medikoo/es6-symbol#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/medikoo/es6-symbol/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/medikoo/es6-symbol.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/ext@1.4.0",
		"name": "ext",
		"version": "1.4.0",
		"description": "JavaScript utilities with respect to emerging standard",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "89ae7a07158f79d35517882904324077e4379244"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "ISC"
			}
		  }
		],
		"purl": "pkg:npm/ext@1.4.0",
		"externalReferences": [
		  {
			"type": "vcs",
			"url": "https://github.com/medikoo/es5-ext/tree/ext"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/type@2.0.0",
		"name": "type",
		"version": "2.0.0",
		"description": "Runtime validation and processing of JavaScript types",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "5f16ff6ef2eb44f260494dae271033b29c09a9c3"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "ISC"
			}
		  }
		],
		"purl": "pkg:npm/type@2.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/medikoo/type#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/medikoo/type/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/medikoo/type.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/next-tick@1.0.0",
		"name": "next-tick",
		"version": "1.0.0",
		"description": "Environment agnostic nextTick polyfill",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "ca86d1fe8828169b0120208e3dc8424b9db8342c"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/next-tick@1.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/medikoo/next-tick#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/medikoo/next-tick/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/medikoo/next-tick.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/type@1.2.0",
		"name": "type",
		"version": "1.2.0",
		"description": "Runtime validation and processing of JavaScript types",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "848dd7698dafa3e54a6c479e759c4bc3f18847a0"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "ISC"
			}
		  }
		],
		"purl": "pkg:npm/type@1.2.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/medikoo/type#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/medikoo/type/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/medikoo/type.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/es6-set@0.1.5",
		"name": "es6-set",
		"version": "0.1.5",
		"description": "ECMAScript6 Set polyfill",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "d2b3ec5d4d800ced818db538d28974db0a73ccb1"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/es6-set@0.1.5",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/medikoo/es6-set#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/medikoo/es6-set/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/medikoo/es6-set.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/es6-symbol@3.1.1",
		"name": "es6-symbol",
		"version": "3.1.1",
		"description": "ECMAScript 6 Symbol polyfill",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "bf00ef4fdab6ba1b46ecb7b629b4c7ed5715cc77"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/es6-symbol@3.1.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/medikoo/es6-symbol#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/medikoo/es6-symbol/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/medikoo/es6-symbol.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/event-emitter@0.3.5",
		"name": "event-emitter",
		"version": "0.3.5",
		"description": "Environment agnostic event emitter",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "df8c69eef1647923c7157b9ce83840610b02cc39"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/event-emitter@0.3.5",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/medikoo/event-emitter#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/medikoo/event-emitter/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/medikoo/event-emitter.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/estree-is-function@1.0.0",
		"name": "estree-is-function",
		"version": "1.0.0",
		"description": "check if an AST node is a function of some sort",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "c0adc29806d7f18a74db7df0f3b2666702e37ad2"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "Apache-2.0"
			}
		  }
		],
		"purl": "pkg:npm/estree-is-function@1.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/goto-bus-stop/estree-is-function"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/goto-bus-stop/estree-is-function/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/goto-bus-stop/estree-is-function.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/get-assigned-identifiers@1.2.0",
		"name": "get-assigned-identifiers",
		"version": "1.2.0",
		"description": "get a list of identifiers that are initialised by a JavaScript AST node.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "6dbf411de648cbaf8d9169ebb0d2d576191e2ff1"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "Apache-2.0"
			}
		  }
		],
		"purl": "pkg:npm/get-assigned-identifiers@1.2.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/goto-bus-stop/get-assigned-identifiers"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/goto-bus-stop/get-assigned-identifiers/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/goto-bus-stop/get-assigned-identifiers.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/shallow-copy@0.0.1",
		"name": "shallow-copy",
		"version": "0.0.1",
		"description": "make a shallow copy of an object or array",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "415f42702d73d810330292cc5ee86eae1a11a170"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/shallow-copy@0.0.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/substack/shallow-copy"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/substack/shallow-copy/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/substack/shallow-copy.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/static-eval@2.1.0",
		"name": "static-eval",
		"version": "2.1.0",
		"description": "evaluate statically-analyzable expressions",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "a16dbe54522d7fa5ef1389129d813fd47b148014"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/static-eval@2.1.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/browserify/static-eval"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/browserify/static-eval/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/browserify/static-eval.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/brotli@1.3.2",
		"name": "brotli",
		"version": "1.3.2",
		"description": "A port of the Brotli compression algorithm as used in WOFF2",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "525a9cad4fcba96475d7d388f6aecb13eed52f46"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/brotli@1.3.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/devongovett/brotli.js"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/devongovett/brotli.js/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/devongovett/brotli.js.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/browserify-optional@1.0.1",
		"name": "browserify-optional",
		"version": "1.0.1",
		"description": "A browserify transform that allows optional dependencies in try..catch blocks",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "1e13722cfde0d85f121676c2a72ced533a018869"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/browserify-optional@1.0.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/devongovett/browserify-optional"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/devongovett/browserify-optional/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/devongovett/browserify-optional.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/ast-transform@0.0.0",
		"name": "ast-transform",
		"version": "0.0.0",
		"description": "Convenience wrapper for performing AST transformations with browserify transform streams",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "74944058887d8283e189d954600947bc98fe0062"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/ast-transform@0.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/hughsk/ast-transform"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/hughsk/ast-transform/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/hughsk/ast-transform.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/escodegen@1.2.0",
		"name": "escodegen",
		"version": "1.2.0",
		"description": "ECMAScript code generator",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "09de7967791cc958b7f89a2ddb6d23451af327e1"
		  }
		],
		"purl": "pkg:npm/escodegen@1.2.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "http://github.com/Constellation/escodegen"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/Constellation/escodegen/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+ssh://git@github.com/Constellation/escodegen.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/estraverse@1.5.1",
		"name": "estraverse",
		"version": "1.5.1",
		"description": "ECMAScript JS AST traversal functions",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "867a3e8e58a9f84618afb6c2ddbcd916b7cbaf71"
		  }
		],
		"purl": "pkg:npm/estraverse@1.5.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/Constellation/estraverse"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/Constellation/estraverse/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+ssh://git@github.com/Constellation/estraverse.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/esutils@1.0.0",
		"name": "esutils",
		"version": "1.0.0",
		"description": "utility box for ECMAScript language tools",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "8151d358e20c8acc7fb745e7472c0025fe496570"
		  }
		],
		"purl": "pkg:npm/esutils@1.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/Constellation/esutils"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/Constellation/esutils/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+ssh://git@github.com/Constellation/esutils.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/source-map@0.1.43",
		"name": "source-map",
		"version": "0.1.43",
		"description": "Generates and consumes source maps",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "c24bc146ca517c1471f5dacbe2571b2b7f9e3346"
		  }
		],
		"purl": "pkg:npm/source-map@0.1.43",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/mozilla/source-map"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/mozilla/source-map/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+ssh://git@github.com/mozilla/source-map.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/amdefine@1.0.1",
		"name": "amdefine",
		"version": "1.0.1",
		"description": "Provide AMD's define() API for declaring modules in the AMD format",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "4a5282ac164729e93619bcfd3ad151f817ce91f5"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "name": "BSD-3-Clause OR MIT"
			}
		  }
		],
		"purl": "pkg:npm/amdefine@1.0.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "http://github.com/jrburke/amdefine"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jrburke/amdefine/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jrburke/amdefine.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/ast-types@0.7.8",
		"name": "ast-types",
		"version": "0.7.8",
		"description": "Esprima-compatible implementation of the Mozilla JS Parser API",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "902d2e0d60d071bdcd46dc115e1809ed11c138a9"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/ast-types@0.7.8",
		"externalReferences": [
		  {
			"type": "website",
			"url": "http://github.com/benjamn/ast-types"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/benjamn/ast-types/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/benjamn/ast-types.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/browser-resolve@1.11.3",
		"name": "browser-resolve",
		"version": "1.11.3",
		"description": "resolve which handles browser field support in package.json",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "9b7cbb3d0f510e4cb86bdbd796124d28b5890af6"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/browser-resolve@1.11.3",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/shtylman/node-browser-resolve#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/shtylman/node-browser-resolve/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/shtylman/node-browser-resolve.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/resolve@1.1.7",
		"name": "resolve",
		"version": "1.1.7",
		"description": "resolve like require.resolve() on behalf of files asynchronously and synchronously",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "203114d82ad2c5ed9e8e0411b3932875e889e97b"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/resolve@1.1.7",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/substack/node-resolve#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/substack/node-resolve/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/substack/node-resolve.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/clone@1.0.4",
		"name": "clone",
		"version": "1.0.4",
		"description": "deep cloning of objects and arrays",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "da309cc263df15994c688ca902179ca3c7cd7c7e"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/clone@1.0.4",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/pvorb/node-clone#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/pvorb/node-clone/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/pvorb/node-clone.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/deep-equal@1.1.1",
		"name": "deep-equal",
		"version": "1.1.1",
		"description": "node's assert.deepEqual algorithm",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "b5c98c942ceffaf7cb051e24e1434a25a2e6076a"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/deep-equal@1.1.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/substack/node-deep-equal#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/substack/node-deep-equal/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+ssh://git@github.com/substack/node-deep-equal.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/is-arguments@1.0.4",
		"name": "is-arguments",
		"version": "1.0.4",
		"description": "Is this an arguments object? It's a harder question than you think.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "3faf966c7cba0ff437fb31f6250082fcf0448cf3"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/is-arguments@1.0.4",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/ljharb/is-arguments"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/ljharb/is-arguments/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/ljharb/is-arguments.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/is-date-object@1.0.2",
		"name": "is-date-object",
		"version": "1.0.2",
		"description": "Is this value a JS Date object? This module works cross-realm/iframe, and despite ES6 @@toStringTag.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "bda736f2cd8fd06d32844e7743bfa7494c3bfd7e"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/is-date-object@1.0.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/ljharb/is-date-object#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/ljharb/is-date-object/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/ljharb/is-date-object.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/is-regex@1.1.0",
		"name": "is-regex",
		"version": "1.1.0",
		"description": "Is this value a JS regex? Works cross-realm/iframe, and despite ES6 @@toStringTag",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "ece38e389e490df0dc21caea2bd596f987f767ff"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/is-regex@1.1.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/ljharb/is-regex"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/ljharb/is-regex/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/ljharb/is-regex.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/has-symbols@1.0.1",
		"name": "has-symbols",
		"version": "1.0.1",
		"description": "Determine if the JS environment has Symbol support. Supports spec, or shams.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "9f5214758a44196c406d9bd76cebf81ec2dd31e8"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/has-symbols@1.0.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/ljharb/has-symbols#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/ljharb/has-symbols/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/ljharb/has-symbols.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/object-is@1.1.2",
		"name": "object-is",
		"version": "1.1.2",
		"description": "ES2015-compliant shim for Object.is - differentiates between -0 and +0",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "c5d2e87ff9e119f78b7a088441519e2eec1573b6"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/object-is@1.1.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/es-shims/object-is"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/es-shims/object-is/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/es-shims/object-is.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/define-properties@1.1.3",
		"name": "define-properties",
		"version": "1.1.3",
		"description": "Define multiple non-enumerable properties at once. Uses Object.defineProperty when available; falls back to standard assignment in older engines.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "cf88da6cbee26fe6db7094f61d870cbd84cee9f1"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/define-properties@1.1.3",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/ljharb/define-properties#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/ljharb/define-properties/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/ljharb/define-properties.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/object-keys@1.1.1",
		"name": "object-keys",
		"version": "1.1.1",
		"description": "An Object.keys replacement, in case Object.keys is not available. From https://github.com/es-shims/es5-shim",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "1c47f272df277f3b1daf061677d9c82e2322c60e"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/object-keys@1.1.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/ljharb/object-keys#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/ljharb/object-keys/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/ljharb/object-keys.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/es-abstract@1.17.6",
		"name": "es-abstract",
		"version": "1.17.6",
		"description": "ECMAScript spec abstract operations.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "9142071707857b2cacc7b89ecb670316c3e2d52a"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/es-abstract@1.17.6",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/ljharb/es-abstract#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/ljharb/es-abstract/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/ljharb/es-abstract.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/es-to-primitive@1.2.1",
		"name": "es-to-primitive",
		"version": "1.2.1",
		"description": "ECMAScript “ToPrimitive” algorithm. Provides ES5 and ES2015 versions.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "e55cd4c9cdc188bcefb03b366c736323fc5c898a"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/es-to-primitive@1.2.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/ljharb/es-to-primitive#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/ljharb/es-to-primitive/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/ljharb/es-to-primitive.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/is-callable@1.2.0",
		"name": "is-callable",
		"version": "1.2.0",
		"description": "Is this JS value callable? Works with Functions and GeneratorFunctions, despite ES6 @@toStringTag.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "83336560b54a38e35e3a2df7afd0454d691468bb"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/is-callable@1.2.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/ljharb/is-callable#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/ljharb/is-callable/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/ljharb/is-callable.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/is-symbol@1.0.3",
		"name": "is-symbol",
		"version": "1.0.3",
		"description": "Determine if a value is an ES6 Symbol or not.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "38e1014b9e6329be0de9d24a414fd7441ec61937"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/is-symbol@1.0.3",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/inspect-js/is-symbol#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/inspect-js/is-symbol/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/inspect-js/is-symbol.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/object.assign@4.1.0",
		"name": "object.assign",
		"version": "4.1.0",
		"description": "ES6 spec-compliant Object.assign shim. From https://github.com/es-shims/es6-shim",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "968bf1100d7956bb3ca086f006f846b3bc4008da"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/object.assign@4.1.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/ljharb/object.assign#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/ljharb/object.assign/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/ljharb/object.assign.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/string.prototype.trimend@1.0.1",
		"name": "string.prototype.trimend",
		"version": "1.0.1",
		"description": "ES2019 spec-compliant String.prototype.trimEnd shim.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "85812a6b847ac002270f5808146064c995fb6913"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/string.prototype.trimend@1.0.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/es-shims/String.prototype.trimEnd#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/es-shims/String.prototype.trimEnd/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/es-shims/String.prototype.trimEnd.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/string.prototype.trimstart@1.0.1",
		"name": "string.prototype.trimstart",
		"version": "1.0.1",
		"description": "ES2019 spec-compliant String.prototype.trimStart shim.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "14af6d9f34b053f7cfc89b72f8f2ee14b9039a54"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/string.prototype.trimstart@1.0.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/es-shims/String.prototype.trimStart#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/es-shims/String.prototype.trimStart/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/es-shims/String.prototype.trimStart.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/regexp.prototype.flags@1.3.0",
		"name": "regexp.prototype.flags",
		"version": "1.3.0",
		"description": "ES6 spec-compliant RegExp.prototype.flags shim.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "7aba89b3c13a64509dabcf3ca8d9fbb9bdf5cb75"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/regexp.prototype.flags@1.3.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/es-shims/RegExp.prototype.flags#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/es-shims/RegExp.prototype.flags/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/es-shims/RegExp.prototype.flags.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/dfa@1.2.0",
		"name": "dfa",
		"version": "1.2.0",
		"description": "A state machine compiler",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "96ac3204e2d29c49ea5b57af8d92c2ae12790657"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/dfa@1.2.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/devongovett/dfa#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/devongovett/dfa/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+ssh://git@github.com/devongovett/dfa.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/restructure@0.5.4",
		"name": "restructure",
		"version": "0.5.4",
		"description": "Declaratively encode and decode binary data",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "f54e7dd563590fb34fd6bf55876109aeccb28de8"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/restructure@0.5.4",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/devongovett/restructure"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/devongovett/restructure/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/devongovett/restructure.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/tiny-inflate@1.0.3",
		"name": "tiny-inflate",
		"version": "1.0.3",
		"description": "A tiny inflate implementation",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "122715494913a1805166aaf7c93467933eea26c4"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/tiny-inflate@1.0.3",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/devongovett/tiny-inflate"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/devongovett/tiny-inflate/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/devongovett/tiny-inflate.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/unicode-properties@1.3.1",
		"name": "unicode-properties",
		"version": "1.3.1",
		"description": "Provides fast access to unicode character properties",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "cc642b6314bde2c691d65dd94cece09ed84f1282"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/unicode-properties@1.3.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/devongovett/unicode-properties"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/devongovett/unicode-properties/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/devongovett/unicode-properties.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/unicode-trie@2.0.0",
		"name": "unicode-trie",
		"version": "2.0.0",
		"description": "Unicode Trie data structure for fast character metadata lookup, ported from ICU",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "8fd8845696e2e14a8b67d78fa9e0dd2cad62fec8"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/unicode-trie@2.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/devongovett/unicode-trie"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/devongovett/unicode-trie/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/devongovett/unicode-trie.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/pako@0.2.9",
		"name": "pako",
		"version": "0.2.9",
		"description": "zlib port to javascript - fast, modularized, with browser support",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "f3f7522f4ef782348da8161bad9ecfd51bf83a75"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/pako@0.2.9",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/nodeca/pako"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/nodeca/pako/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/nodeca/pako.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/unicode-trie@0.3.1",
		"name": "unicode-trie",
		"version": "0.3.1",
		"description": "Unicode Trie data structure for fast character metadata lookup, ported from ICU",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "d671dddd89101a08bac37b6a5161010602052085"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/unicode-trie@0.3.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/devongovett/unicode-trie"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/devongovett/unicode-trie/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/devongovett/unicode-trie.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/linebreak@1.0.2",
		"name": "linebreak",
		"version": "1.0.2",
		"description": "An implementation of the Unicode Line Breaking Algorithm (UAX #14)",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "4b5781733e9a9eb2849dba2f963e47c887f8aa06"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/linebreak@1.0.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/devongovett/linebreaker"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/devongovett/linebreaker/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/devongovett/linebreaker.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/base64-js@0.0.8",
		"name": "base64-js",
		"version": "0.0.8",
		"description": "Base64 encoding/decoding in pure JS",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "1101e9544f4a76b1bc3b26d452ca96d7a35e7978"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/base64-js@0.0.8",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/beatgammit/base64-js#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/beatgammit/base64-js/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/beatgammit/base64-js.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/unicode-trie@1.0.0",
		"name": "unicode-trie",
		"version": "1.0.0",
		"description": "Unicode Trie data structure for fast character metadata lookup, ported from ICU",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "f649afdca127135edb55ca0ad7c8c60656d92ad1"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/unicode-trie@1.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/devongovett/unicode-trie"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/devongovett/unicode-trie/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/devongovett/unicode-trie.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/png-js@1.0.0",
		"name": "png-js",
		"version": "1.0.0",
		"description": "A PNG decoder in JavaScript",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "e5484f1e8156996e383aceebb3789fd75df1874d"
		  }
		],
		"purl": "pkg:npm/png-js@1.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/devongovett/png.js#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "http://github.com/devongovett/png.js/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/devongovett/png.js.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/portscanner@2.2.0",
		"name": "portscanner",
		"version": "2.2.0",
		"description": "Asynchronous port scanner for Node.js",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "6059189b3efa0965c9d96a56b958eb9508411cf1"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/portscanner@2.2.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/baalexander/node-portscanner"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/baalexander/node-portscanner/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/baalexander/node-portscanner.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/is-number-like@1.0.8",
		"name": "is-number-like",
		"version": "1.0.8",
		"description": "Checks whether provided parameter looks like a number",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "2e129620b50891042e44e9bbbb30593e75cfbbe3"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "ISC"
			}
		  }
		],
		"purl": "pkg:npm/is-number-like@1.0.8",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/vigour-io/is-number-like#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/vigour-io/is-number-like/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/vigour-io/is-number-like.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/lodash.isfinite@3.3.2",
		"name": "lodash.isfinite",
		"version": "3.3.2",
		"description": "The lodash method _.isFinite exported as a module.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "fb89b65a9a80281833f0b7478b3a5104f898ebb3"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/lodash.isfinite@3.3.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://lodash.com/"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/lodash/lodash/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/lodash/lodash.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/prom-client@11.5.3",
		"name": "prom-client",
		"version": "11.5.3",
		"description": "Client for prometheus",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "5fedfce1083bac6c2b223738e966d0e1643756f8"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "Apache-2.0"
			}
		  }
		],
		"purl": "pkg:npm/prom-client@11.5.3",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/siimon/prom-client"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/siimon/prom-client/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+ssh://git@github.com/siimon/prom-client.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/tdigest@0.1.1",
		"name": "tdigest",
		"version": "0.1.1",
		"description": "javascript implementation of Dunning's T-Digest for streaming quantile approximation",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "2e3cb2c39ea449e55d1e6cd91117accca4588021"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/tdigest@0.1.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/welch/tdigest"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/welch/tdigest/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/welch/tdigest.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/bintrees@1.0.1",
		"name": "bintrees",
		"version": "1.0.1",
		"description": "Binary Search Trees",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "0e655c9b9c2435eaab68bf4027226d2b55a34524"
		  }
		],
		"purl": "pkg:npm/bintrees@1.0.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/vadimg/js_bintrees#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/vadimg/js_bintrees/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/vadimg/js_bintrees.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/pug@2.0.4",
		"name": "pug",
		"version": "2.0.4",
		"description": "A clean, whitespace-sensitive template language for writing HTML",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "ee7682ec0a60494b38d48a88f05f3b0ac931377d"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/pug@2.0.4",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://pugjs.org"
		  },
		  {
			"type": "vcs",
			"url": "https://github.com/pugjs/pug/tree/master/packages/pug"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/pug-code-gen@2.0.2",
		"name": "pug-code-gen",
		"version": "2.0.2",
		"description": "Default code-generator for pug.  It generates HTML via a JavaScript template function.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "ad0967162aea077dcf787838d94ed14acb0217c2"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/pug-code-gen@2.0.2",
		"externalReferences": [
		  {
			"type": "vcs",
			"url": "https://github.com/pugjs/pug/tree/master/packages/pug-code-gen"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/constantinople@3.1.2",
		"name": "constantinople",
		"version": "3.1.2",
		"description": "Determine whether a JavaScript expression evaluates to a constant (using acorn)",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "d45ed724f57d3d10500017a7d3a889c1381ae647"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/constantinople@3.1.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/ForbesLindesay/constantinople#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/ForbesLindesay/constantinople/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/ForbesLindesay/constantinople.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/%40types/babel-types@7.0.8",
		"group": "@types",
		"name": "babel-types",
		"version": "7.0.8",
		"description": "TypeScript definitions for babel-types",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "267f405bda841ffae731e7714166b88254cc3e19"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/%40types/babel-types@7.0.8",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/DefinitelyTyped/DefinitelyTyped#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/DefinitelyTyped/DefinitelyTyped/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/DefinitelyTyped/DefinitelyTyped.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/%40types/babylon@6.16.5",
		"group": "@types",
		"name": "babylon",
		"version": "6.16.5",
		"description": "TypeScript definitions for babylon",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "1c5641db69eb8cdf378edd25b4be7754beeb48b4"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/%40types/babylon@6.16.5",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/DefinitelyTyped/DefinitelyTyped#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/DefinitelyTyped/DefinitelyTyped/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/DefinitelyTyped/DefinitelyTyped.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/babel-types@6.26.0",
		"name": "babel-types",
		"version": "6.26.0",
		"description": "Babel Types is a Lodash-esque utility library for AST nodes",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "a3b073f94ab49eb6fa55cd65227a334380632497"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/babel-types@6.26.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://babeljs.io/"
		  },
		  {
			"type": "vcs",
			"url": "https://github.com/babel/babel/tree/master/packages/babel-types"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/to-fast-properties@1.0.3",
		"name": "to-fast-properties",
		"version": "1.0.3",
		"description": "Force V8 to use fast properties for an object",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "b83571fa4d8c25b82e231b06e3a3055de4ca1a47"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/to-fast-properties@1.0.3",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/sindresorhus/to-fast-properties#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/sindresorhus/to-fast-properties/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/sindresorhus/to-fast-properties.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/babylon@6.18.0",
		"name": "babylon",
		"version": "6.18.0",
		"description": "A JavaScript parser",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "af2f3b88fa6f5c1e4c634d1a0f8eac4f55b395e3"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/babylon@6.18.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://babeljs.io/"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/babel/babylon/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/babel/babylon.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/doctypes@1.1.0",
		"name": "doctypes",
		"version": "1.1.0",
		"description": "Shorthands for commonly used doctypes",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "ea80b106a87538774e8a3a4a5afe293de489e0a9"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/doctypes@1.1.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/pugjs/doctypes#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/pugjs/doctypes/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/pugjs/doctypes.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/js-stringify@1.0.2",
		"name": "js-stringify",
		"version": "1.0.2",
		"description": "Stringify an object so it can be safely inlined in JavaScript code",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "1736fddfd9724f28a3682adc6230ae7e4e9679db"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/js-stringify@1.0.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jadejs/js-stringify#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jadejs/js-stringify/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jadejs/js-stringify.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/pug-attrs@2.0.4",
		"name": "pug-attrs",
		"version": "2.0.4",
		"description": "Generate code for Pug attributes",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "b2f44c439e4eb4ad5d4ef25cac20d18ad28cc336"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/pug-attrs@2.0.4",
		"externalReferences": [
		  {
			"type": "vcs",
			"url": "https://github.com/pugjs/pug/tree/master/packages/pug-attrs"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/pug-runtime@2.0.5",
		"name": "pug-runtime",
		"version": "2.0.5",
		"description": "The runtime components for the pug templating language",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "6da7976c36bf22f68e733c359240d8ae7a32953a"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/pug-runtime@2.0.5",
		"externalReferences": [
		  {
			"type": "vcs",
			"url": "https://github.com/pugjs/pug/tree/master/packages/pug-runtime"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/pug-error@1.3.3",
		"name": "pug-error",
		"version": "1.3.3",
		"description": "Standard error objects for pug",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "f342fb008752d58034c185de03602dd9ffe15fa6"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/pug-error@1.3.3",
		"externalReferences": [
		  {
			"type": "vcs",
			"url": "https://github.com/pugjs/pug/tree/master/packages/pug-error"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/void-elements@2.0.1",
		"name": "void-elements",
		"version": "2.0.1",
		"description": "Array of \"void elements\" defined by the HTML specification.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "c066afb582bb1cb4128d60ea92392e94d5e9dbec"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/void-elements@2.0.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/hemanth/void-elements"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/hemanth/void-elements/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/hemanth/void-elements.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/with@5.1.1",
		"name": "with",
		"version": "5.1.1",
		"description": "Compile time with for strict mode JavaScript",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "fa4daa92daf32c4ea94ed453c81f04686b575dfe"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/with@5.1.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/pugjs/with#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/pugjs/with/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/pugjs/with.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/acorn@3.3.0",
		"name": "acorn",
		"version": "3.3.0",
		"description": "ECMAScript parser",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "45e37fb39e8da3f25baee3ff5369e2bb5f22017a"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/acorn@3.3.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/ternjs/acorn"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/ternjs/acorn/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/ternjs/acorn.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/acorn-globals@3.1.0",
		"name": "acorn-globals",
		"version": "3.1.0",
		"description": "Detect global variables in JavaScript using acorn",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "fd8270f71fbb4996b004fa880ee5d46573a731bf"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/acorn-globals@3.1.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/ForbesLindesay/acorn-globals#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/ForbesLindesay/acorn-globals/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/ForbesLindesay/acorn-globals.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/acorn@4.0.13",
		"name": "acorn",
		"version": "4.0.13",
		"description": "ECMAScript parser",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "105495ae5361d697bd195c825192e1ad7f253787"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/acorn@4.0.13",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/ternjs/acorn"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/ternjs/acorn/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/ternjs/acorn.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/pug-filters@3.1.1",
		"name": "pug-filters",
		"version": "3.1.1",
		"description": "Code for processing filters in pug templates",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "ab2cc82db9eeccf578bda89130e252a0db026aa7"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/pug-filters@3.1.1",
		"externalReferences": [
		  {
			"type": "vcs",
			"url": "https://github.com/pugjs/pug/tree/master/packages/pug-filters"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/clean-css@4.2.3",
		"name": "clean-css",
		"version": "4.2.3",
		"description": "A well-tested CSS minifier",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "507b5de7d97b48ee53d84adb0160ff6216380f78"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/clean-css@4.2.3",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jakubpawlowicz/clean-css"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jakubpawlowicz/clean-css/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jakubpawlowicz/clean-css.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/jstransformer@1.0.0",
		"name": "jstransformer",
		"version": "1.0.0",
		"description": "Normalize the API of any jstransformer",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "ed8bf0921e2f3f1ed4d5c1a44f68709ed24722c3"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/jstransformer@1.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jstransformers/jstransformer#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jstransformers/jstransformer/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jstransformers/jstransformer.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/is-promise@2.2.2",
		"name": "is-promise",
		"version": "2.2.2",
		"description": "Test whether an object looks like a promises-a+ promise",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "39ab959ccbf9a774cf079f7b40c7a26f763135f1"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/is-promise@2.2.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/then/is-promise#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/then/is-promise/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/then/is-promise.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/promise@7.3.1",
		"name": "promise",
		"version": "7.3.1",
		"description": "Bare bones Promises/A+ implementation",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "064b72602b18f90f29192b8b1bc418ffd1ebd3bf"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/promise@7.3.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/then/promise#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/then/promise/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/then/promise.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/asap@2.0.6",
		"name": "asap",
		"version": "2.0.6",
		"description": "High-priority task queue for Node.js and browsers",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "e50347611d7e690943208bbdafebcbc2fb866d46"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/asap@2.0.6",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/kriskowal/asap#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/kriskowal/asap/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/kriskowal/asap.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/pug-walk@1.1.8",
		"name": "pug-walk",
		"version": "1.1.8",
		"description": "Walk and transform a pug AST",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "b408f67f27912f8c21da2f45b7230c4bd2a5ea7a"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/pug-walk@1.1.8",
		"externalReferences": [
		  {
			"type": "vcs",
			"url": "https://github.com/pugjs/pug/tree/master/packages/pug-walk"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/uglify-js@2.8.29",
		"name": "uglify-js",
		"version": "2.8.29",
		"description": "JavaScript parser, mangler/compressor and beautifier toolkit",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "29c5733148057bb4e1f75df35b7a9cb72e6a59dd"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "BSD-2-Clause"
			}
		  }
		],
		"purl": "pkg:npm/uglify-js@2.8.29",
		"externalReferences": [
		  {
			"type": "website",
			"url": "http://lisperator.net/uglifyjs"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/mishoo/UglifyJS2/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/mishoo/UglifyJS2.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/yargs@3.10.0",
		"name": "yargs",
		"version": "3.10.0",
		"description": "Light-weight option parsing with an argv hash. No optstrings attached.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "f7ee7bd857dd7c1d2d38c0e74efbd681d1431fd1"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/yargs@3.10.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/bcoe/yargs#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/bcoe/yargs/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+ssh://git@github.com/bcoe/yargs.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/camelcase@1.2.1",
		"name": "camelcase",
		"version": "1.2.1",
		"description": "Convert a dash/dot/underscore/space separated string to camelCase: foo-bar → fooBar",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "9bb5304d2e0b56698b2c758b08a3eaa9daa58a39"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/camelcase@1.2.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/sindresorhus/camelcase#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/sindresorhus/camelcase/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/sindresorhus/camelcase.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/cliui@2.1.0",
		"name": "cliui",
		"version": "2.1.0",
		"description": "easily create complex multi-column command-line-interfaces",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "4b475760ff80264c762c3a1719032e91c7fea0d1"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "ISC"
			}
		  }
		],
		"purl": "pkg:npm/cliui@2.1.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/bcoe/cliui#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/bcoe/cliui/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+ssh://git@github.com/bcoe/cliui.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/center-align@0.1.3",
		"name": "center-align",
		"version": "0.1.3",
		"description": "Center-align the text in a string.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "aa0d32629b6ee972200411cbd4461c907bc2b7ad"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/center-align@0.1.3",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/center-align"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/center-align/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/center-align.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/align-text@0.1.4",
		"name": "align-text",
		"version": "0.1.4",
		"description": "Align the text in a string.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "0cd90a561093f35d0a99256c22b7069433fad117"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/align-text@0.1.4",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/align-text"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/align-text/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/jonschlinkert/align-text.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/longest@1.0.1",
		"name": "longest",
		"version": "1.0.1",
		"description": "Get the longest item in an array.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "30a0b2da38f73770e8294a0d22e6625ed77d0097"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/longest@1.0.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/longest"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/longest/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/longest.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/lazy-cache@1.0.4",
		"name": "lazy-cache",
		"version": "1.0.4",
		"description": "Cache requires to be lazy-loaded when needed.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "a1d78fc3a50474cb80845d3b3b6e1da49a446e8e"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/lazy-cache@1.0.4",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/lazy-cache"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/lazy-cache/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/lazy-cache.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/right-align@0.1.3",
		"name": "right-align",
		"version": "0.1.3",
		"description": "Right-align the text in a string.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "61339b722fe6a3515689210d24e14c96148613ef"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/right-align@0.1.3",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/right-align"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/right-align/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/jonschlinkert/right-align.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/wordwrap@0.0.2",
		"name": "wordwrap",
		"version": "0.0.2",
		"description": "Wrap those words. Show them at what columns to start and stop.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "b79669bb42ecb409f83d583cad52ca17eaa1643f"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "name": "MIT/X11"
			}
		  }
		],
		"purl": "pkg:npm/wordwrap@0.0.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/substack/node-wordwrap#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/substack/node-wordwrap/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/substack/node-wordwrap.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/window-size@0.1.0",
		"name": "window-size",
		"version": "0.1.0",
		"description": "Reliable way to to get the height and width of the terminal/console in a node.js environment.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "5438cd2ea93b202efa3a19fe8887aee7c94f9c9d"
		  }
		],
		"purl": "pkg:npm/window-size@0.1.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jonschlinkert/window-size"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jonschlinkert/window-size/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jonschlinkert/window-size.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/pug-lexer@4.1.0",
		"name": "pug-lexer",
		"version": "4.1.0",
		"description": "The pug lexer (takes a string and converts it to an array of tokens)",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "531cde48c7c0b1fcbbc2b85485c8665e31489cfd"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/pug-lexer@4.1.0",
		"externalReferences": [
		  {
			"type": "vcs",
			"url": "https://github.com/pugjs/pug/tree/master/packages/pug-lexer"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/character-parser@2.2.0",
		"name": "character-parser",
		"version": "2.2.0",
		"description": "Parse JavaScript one character at a time to look for snippets in Templates.  This is not a validator, it's just designed to allow you to have sections of JavaScript delimited by brackets robustly.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "c7ce28f36d4bcd9744e5ffc2c5fcde1c73261fc0"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/character-parser@2.2.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/ForbesLindesay/character-parser#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/ForbesLindesay/character-parser/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/ForbesLindesay/character-parser.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/is-expression@3.0.0",
		"name": "is-expression",
		"version": "3.0.0",
		"description": "Check if a string is a valid JavaScript expression",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "39acaa6be7fd1f3471dc42c7416e61c24317ac9f"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/is-expression@3.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/pugjs/is-expression#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/pugjs/is-expression/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/pugjs/is-expression.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/pug-linker@3.0.6",
		"name": "pug-linker",
		"version": "3.0.6",
		"description": "Link multiple pug ASTs together using include/extends",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "f5bf218b0efd65ce6670f7afc51658d0f82989fb"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/pug-linker@3.0.6",
		"externalReferences": [
		  {
			"type": "vcs",
			"url": "https://github.com/pugjs/pug/tree/master/packages/pug-linker"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/pug-load@2.0.12",
		"name": "pug-load",
		"version": "2.0.12",
		"description": "The Pug loader is responsible for loading the depenendencies of a given Pug file.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "d38c85eb85f6e2f704dea14dcca94144d35d3e7b"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/pug-load@2.0.12",
		"externalReferences": [
		  {
			"type": "vcs",
			"url": "https://github.com/pugjs/pug/tree/master/packages/pug-load"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/pug-parser@5.0.1",
		"name": "pug-parser",
		"version": "5.0.1",
		"description": "The pug parser (takes an array of tokens and converts it to an abstract syntax tree)",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "03e7ada48b6840bd3822f867d7d90f842d0ffdc9"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/pug-parser@5.0.1",
		"externalReferences": [
		  {
			"type": "vcs",
			"url": "https://github.com/pugjs/pug/tree/master/packages/pug-parser"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/token-stream@0.0.1",
		"name": "token-stream",
		"version": "0.0.1",
		"description": "Take an array of token and produce a more useful API to give to a parser",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "ceeefc717a76c4316f126d0b9dbaa55d7e7df01a"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/token-stream@0.0.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jadejs/token-stream#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jadejs/token-stream/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jadejs/token-stream.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/pug-strip-comments@1.0.4",
		"name": "pug-strip-comments",
		"version": "1.0.4",
		"description": "Strip comments from a Pug token stream (from the lexer)",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "cc1b6de1f6e8f5931cf02ec66cdffd3f50eaf8a8"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/pug-strip-comments@1.0.4",
		"externalReferences": [
		  {
			"type": "vcs",
			"url": "https://github.com/pugjs/pug/tree/master/packages/pug-strip-comments"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/replace@1.2.0",
		"name": "replace",
		"version": "1.2.0",
		"description": "Command line search and replace utility",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "a25da288841aab22f0f7e95dc1d249dbd2ed6e26"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/replace@1.2.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/ALMaclaine/replace#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/ALMaclaine/replace/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/ALMaclaine/replace.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/yargs@15.4.1",
		"name": "yargs",
		"version": "15.4.1",
		"description": "yargs the modern, pirate-themed, successor to optimist.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "0d87a16de01aee9d8bec2bfbf74f67851730f4f8"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/yargs@15.4.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://yargs.js.org/"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/yargs/yargs/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/yargs/yargs.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/cliui@6.0.0",
		"name": "cliui",
		"version": "6.0.0",
		"description": "easily create complex multi-column command-line-interfaces",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "511d702c0c4e41ca156d7d0e96021f23e13225b1"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "ISC"
			}
		  }
		],
		"purl": "pkg:npm/cliui@6.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/yargs/cliui#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/yargs/cliui/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+ssh://git@github.com/yargs/cliui.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/string-width@4.2.0",
		"name": "string-width",
		"version": "4.2.0",
		"description": "Get the visual width of a string - the number of columns required to display it",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "952182c46cc7b2c313d1596e623992bd163b72b5"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/string-width@4.2.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/sindresorhus/string-width#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/sindresorhus/string-width/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/sindresorhus/string-width.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/emoji-regex@8.0.0",
		"name": "emoji-regex",
		"version": "8.0.0",
		"description": "A regular expression to match all Emoji-only symbols as per the Unicode Standard.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "e818fd69ce5ccfcb404594f842963bf53164cc37"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/emoji-regex@8.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://mths.be/emoji-regex"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/mathiasbynens/emoji-regex/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/mathiasbynens/emoji-regex.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/is-fullwidth-code-point@3.0.0",
		"name": "is-fullwidth-code-point",
		"version": "3.0.0",
		"description": "Check if the character represented by a given Unicode code point is fullwidth",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "f116f8064fe90b3f7844a38997c0b75051269f1d"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/is-fullwidth-code-point@3.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/sindresorhus/is-fullwidth-code-point#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/sindresorhus/is-fullwidth-code-point/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/sindresorhus/is-fullwidth-code-point.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/strip-ansi@6.0.0",
		"name": "strip-ansi",
		"version": "6.0.0",
		"description": "Strip ANSI escape codes from a string",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "0b1571dd7669ccd4f3e06e14ef1eed26225ae532"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/strip-ansi@6.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/chalk/strip-ansi#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/chalk/strip-ansi/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/chalk/strip-ansi.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/ansi-regex@5.0.0",
		"name": "ansi-regex",
		"version": "5.0.0",
		"description": "Regular expression for matching ANSI escape codes",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "388539f55179bf39339c81af30a654d69f87cb75"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/ansi-regex@5.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/chalk/ansi-regex#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/chalk/ansi-regex/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/chalk/ansi-regex.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/wrap-ansi@6.2.0",
		"name": "wrap-ansi",
		"version": "6.2.0",
		"description": "Wordwrap a string with ANSI escape codes",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "e9393ba07102e6c91a3b221478f0257cd2856e53"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/wrap-ansi@6.2.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/chalk/wrap-ansi#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/chalk/wrap-ansi/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/chalk/wrap-ansi.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/ansi-styles@4.2.1",
		"name": "ansi-styles",
		"version": "4.2.1",
		"description": "ANSI escape codes for styling strings in the terminal",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "90ae75c424d008d2624c5bf29ead3177ebfcf359"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/ansi-styles@4.2.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/chalk/ansi-styles#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/chalk/ansi-styles/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/chalk/ansi-styles.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/%40types/color-name@1.1.1",
		"group": "@types",
		"name": "color-name",
		"version": "1.1.1",
		"description": "TypeScript definitions for color-name",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "1c1261bbeaa10a8055bbc5d8ab84b7b2afc846a0"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/%40types/color-name@1.1.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/DefinitelyTyped/DefinitelyTyped#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/DefinitelyTyped/DefinitelyTyped/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/DefinitelyTyped/DefinitelyTyped.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/color-convert@2.0.1",
		"name": "color-convert",
		"version": "2.0.1",
		"description": "Plain color conversion functions",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "72d3a68d598c9bdb3af2ad1e84f21d896abd4de3"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/color-convert@2.0.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/Qix-/color-convert#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/Qix-/color-convert/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/Qix-/color-convert.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/color-name@1.1.4",
		"name": "color-name",
		"version": "1.1.4",
		"description": "A list of color names and its values",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "c2a09a87acbde69543de6f63fa3995c826c536a2"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/color-name@1.1.4",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/colorjs/color-name"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/colorjs/color-name/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+ssh://git@github.com/colorjs/color-name.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/find-up@4.1.0",
		"name": "find-up",
		"version": "4.1.0",
		"description": "Find a file or directory by walking up parent directories",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "97afe7d6cdc0bc5928584b7c8d7b16e8a9aa5d19"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/find-up@4.1.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/sindresorhus/find-up#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/sindresorhus/find-up/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/sindresorhus/find-up.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/locate-path@5.0.0",
		"name": "locate-path",
		"version": "5.0.0",
		"description": "Get the first path that exists on disk of multiple paths",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "1afba396afd676a6d42504d0a67a3a7eb9f62aa0"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/locate-path@5.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/sindresorhus/locate-path#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/sindresorhus/locate-path/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/sindresorhus/locate-path.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/p-locate@4.1.0",
		"name": "p-locate",
		"version": "4.1.0",
		"description": "Get the first fulfilled promise that satisfies the provided testing function",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "a3428bb7088b3a60292f66919278b7c297ad4f07"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/p-locate@4.1.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/sindresorhus/p-locate#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/sindresorhus/p-locate/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/sindresorhus/p-locate.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/path-exists@4.0.0",
		"name": "path-exists",
		"version": "4.0.0",
		"description": "Check if a path exists",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "513bdbe2d3b95d7762e8c1137efa195c6c61b5b3"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/path-exists@4.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/sindresorhus/path-exists#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/sindresorhus/path-exists/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/sindresorhus/path-exists.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/yargs-parser@18.1.3",
		"name": "yargs-parser",
		"version": "18.1.3",
		"description": "the mighty option parser used by yargs",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "be68c4975c6b2abf469236b0c870362fab09a7b0"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "ISC"
			}
		  }
		],
		"purl": "pkg:npm/yargs-parser@18.1.3",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/yargs/yargs-parser#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/yargs/yargs-parser/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/yargs/yargs-parser.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/request@2.88.2",
		"name": "request",
		"version": "2.88.2",
		"description": "Simplified HTTP request client.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "d73c918731cb5a87da047e207234146f664d12b3"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "Apache-2.0"
			}
		  }
		],
		"purl": "pkg:npm/request@2.88.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/request/request#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "http://github.com/request/request/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/request/request.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/aws-sign2@0.7.0",
		"name": "aws-sign2",
		"version": "0.7.0",
		"description": "AWS signing. Originally pulled from LearnBoost/knox, maintained as vendor in request, now a standalone module.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "b46e890934a9591f2d2f6f86d7e6a9f1b3fe76a8"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "Apache-2.0"
			}
		  }
		],
		"purl": "pkg:npm/aws-sign2@0.7.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/mikeal/aws-sign#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/mikeal/aws-sign/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/mikeal/aws-sign.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/aws4@1.10.0",
		"name": "aws4",
		"version": "1.10.0",
		"description": "Signs and prepares requests using AWS Signature Version 4",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "a17b3a8ea811060e74d47d306122400ad4497ae2"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/aws4@1.10.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/mhart/aws4#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/mhart/aws4/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/mhart/aws4.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/caseless@0.12.0",
		"name": "caseless",
		"version": "0.12.0",
		"description": "Caseless object set/get/has, very useful when working with HTTP headers.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "1b681c21ff84033c826543090689420d187151dc"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "Apache-2.0"
			}
		  }
		],
		"purl": "pkg:npm/caseless@0.12.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/mikeal/caseless#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/mikeal/caseless/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/mikeal/caseless.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/combined-stream@1.0.8",
		"name": "combined-stream",
		"version": "1.0.8",
		"description": "A stream that emits multiple other streams one after another.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "c3d45a8b34fd730631a110a8a2520682b31d5a7f"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/combined-stream@1.0.8",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/felixge/node-combined-stream"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/felixge/node-combined-stream/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/felixge/node-combined-stream.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/delayed-stream@1.0.0",
		"name": "delayed-stream",
		"version": "1.0.0",
		"description": "Buffers events from a stream until you are ready to handle them.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "df3ae199acadfb7d440aaae0b29e2272b24ec619"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/delayed-stream@1.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/felixge/node-delayed-stream"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/felixge/node-delayed-stream/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/felixge/node-delayed-stream.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/forever-agent@0.6.1",
		"name": "forever-agent",
		"version": "0.6.1",
		"description": "HTTP Agent that keeps socket connections alive between keep-alive requests. Formerly part of mikeal/request, now a standalone module.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "fbc71f0c41adeb37f96c577ad1ed42d8fdacca91"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "Apache-2.0"
			}
		  }
		],
		"purl": "pkg:npm/forever-agent@0.6.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/mikeal/forever-agent#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/mikeal/forever-agent/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/mikeal/forever-agent.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/form-data@2.3.3",
		"name": "form-data",
		"version": "2.3.3",
		"description": "A library to create readable \"multipart/form-data\" streams. Can be used to submit forms and file uploads to other web applications.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "dcce52c05f644f298c6a7ab936bd724ceffbf3a6"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/form-data@2.3.3",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/form-data/form-data#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/form-data/form-data/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/form-data/form-data.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/asynckit@0.4.0",
		"name": "asynckit",
		"version": "0.4.0",
		"description": "Minimal async jobs utility library, with streams support",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "c79ed97f7f34cb8f2ba1bc9790bcc366474b4b79"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/asynckit@0.4.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/alexindigo/asynckit#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/alexindigo/asynckit/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/alexindigo/asynckit.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/har-validator@5.1.3",
		"name": "har-validator",
		"version": "5.1.3",
		"description": "Extremely fast HTTP Archive (HAR) validator using JSON Schema",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "1ef89ebd3e4996557675eed9893110dc350fa080"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/har-validator@5.1.3",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/ahmadnassri/node-har-validator"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/ahmadnassri/node-har-validator/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/ahmadnassri/node-har-validator.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/ajv@6.12.3",
		"name": "ajv",
		"version": "6.12.3",
		"description": "Another JSON Schema Validator",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "18c5af38a111ddeb4f2697bd78d68abc1cabd706"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/ajv@6.12.3",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/ajv-validator/ajv"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/ajv-validator/ajv/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/ajv-validator/ajv.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/fast-deep-equal@3.1.3",
		"name": "fast-deep-equal",
		"version": "3.1.3",
		"description": "Fast deep equal",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "3a7d56b559d6cbc3eb512325244e619a65c6c525"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/fast-deep-equal@3.1.3",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/epoberezkin/fast-deep-equal#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/epoberezkin/fast-deep-equal/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/epoberezkin/fast-deep-equal.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/fast-json-stable-stringify@2.1.0",
		"name": "fast-json-stable-stringify",
		"version": "2.1.0",
		"description": "deterministic JSON.stringify() - a faster version of substack's json-stable-strigify without jsonify",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "874bf69c6f404c2b5d99c481341399fd55892633"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/fast-json-stable-stringify@2.1.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/epoberezkin/fast-json-stable-stringify"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/epoberezkin/fast-json-stable-stringify/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/epoberezkin/fast-json-stable-stringify.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/json-schema-traverse@0.4.1",
		"name": "json-schema-traverse",
		"version": "0.4.1",
		"description": "Traverse JSON Schema passing each schema object to callback",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "69f6a87d9513ab8bb8fe63bdb0979c448e684660"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/json-schema-traverse@0.4.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/epoberezkin/json-schema-traverse#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/epoberezkin/json-schema-traverse/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/epoberezkin/json-schema-traverse.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/uri-js@4.2.2",
		"name": "uri-js",
		"version": "4.2.2",
		"description": "An RFC 3986/3987 compliant, scheme extendable URI/IRI parsing/validating/resolving library for JavaScript.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "94c540e1ff772956e2299507c010aea6c8838eb0"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "BSD-2-Clause"
			}
		  }
		],
		"purl": "pkg:npm/uri-js@4.2.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/garycourt/uri-js"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/garycourt/uri-js/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+ssh://git@github.com/garycourt/uri-js.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/punycode@2.1.1",
		"name": "punycode",
		"version": "2.1.1",
		"description": "A robust Punycode converter that fully complies to RFC 3492 and RFC 5891, and works on nearly all JavaScript platforms.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "b58b010ac40c22c5657616c8d2c2c02c7bf479ec"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/punycode@2.1.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://mths.be/punycode"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/bestiejs/punycode.js/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/bestiejs/punycode.js.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/har-schema@2.0.0",
		"name": "har-schema",
		"version": "2.0.0",
		"description": "JSON Schema for HTTP Archive (HAR)",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "a94c2224ebcac04782a0d9035521f24735b7ec92"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "ISC"
			}
		  }
		],
		"purl": "pkg:npm/har-schema@2.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/ahmadnassri/har-schema"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/ahmadnassri/har-schema/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/ahmadnassri/har-schema.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/http-signature@1.2.0",
		"name": "http-signature",
		"version": "1.2.0",
		"description": "Reference implementation of Joyent's HTTP Signature scheme.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "9aecd925114772f3d95b65a60abb8f7c18fbace1"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/http-signature@1.2.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/joyent/node-http-signature/"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/joyent/node-http-signature/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/joyent/node-http-signature.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/assert-plus@1.0.0",
		"name": "assert-plus",
		"version": "1.0.0",
		"description": "Extra assertions on top of node's assert module",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "f12e0f3c5d77b0b1cdd9146942e4e96c1e4dd525"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/assert-plus@1.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/mcavage/node-assert-plus#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/mcavage/node-assert-plus/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/mcavage/node-assert-plus.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/jsprim@1.4.1",
		"name": "jsprim",
		"version": "1.4.1",
		"description": "utilities for primitive JavaScript types",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "313e66bc1e5cc06e438bc1b7499c2e5c56acb6a2"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/jsprim@1.4.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/joyent/node-jsprim#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/joyent/node-jsprim/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/joyent/node-jsprim.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/extsprintf@1.3.0",
		"name": "extsprintf",
		"version": "1.3.0",
		"description": "extended POSIX-style sprintf",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "96918440e3041a7a414f8c52e3c574eb3c3e1e05"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/extsprintf@1.3.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/davepacheco/node-extsprintf#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/davepacheco/node-extsprintf/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/davepacheco/node-extsprintf.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/json-schema@0.2.3",
		"name": "json-schema",
		"version": "0.2.3",
		"description": "JSON Schema validation and specifications",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "b480c892e59a2f05954ce727bd3f2a4e882f9e13"
		  }
		],
		"purl": "pkg:npm/json-schema@0.2.3",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/kriszyp/json-schema#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/kriszyp/json-schema/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+ssh://git@github.com/kriszyp/json-schema.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/verror@1.10.0",
		"name": "verror",
		"version": "1.10.0",
		"description": "richer JavaScript errors",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "3a105ca17053af55d6e270c1f8288682e18da400"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/verror@1.10.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/davepacheco/node-verror#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/davepacheco/node-verror/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/davepacheco/node-verror.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/sshpk@1.16.1",
		"name": "sshpk",
		"version": "1.16.1",
		"description": "A library for finding and using SSH public keys",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "fb661c0bef29b39db40769ee39fa70093d6f6877"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/sshpk@1.16.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/arekinath/node-sshpk#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/arekinath/node-sshpk/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/joyent/node-sshpk.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/asn1@0.2.4",
		"name": "asn1",
		"version": "0.2.4",
		"description": "Contains parsers and serializers for ASN.1 (currently BER only)",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "8d2475dfab553bb33e77b54e59e880bb8ce23136"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/asn1@0.2.4",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/joyent/node-asn1#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/joyent/node-asn1/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/joyent/node-asn1.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/bcrypt-pbkdf@1.0.2",
		"name": "bcrypt-pbkdf",
		"version": "1.0.2",
		"description": "Port of the OpenBSD bcrypt_pbkdf function to pure JS",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "a4301d389b6a43f9b67ff3ca11a3f6637e360e9e"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "BSD-3-Clause"
			}
		  }
		],
		"purl": "pkg:npm/bcrypt-pbkdf@1.0.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/joyent/node-bcrypt-pbkdf#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/joyent/node-bcrypt-pbkdf/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/joyent/node-bcrypt-pbkdf.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/tweetnacl@0.14.5",
		"name": "tweetnacl",
		"version": "0.14.5",
		"description": "Port of TweetNaCl cryptographic library to JavaScript",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "5ae68177f192d4456269d108afa93ff8743f4f64"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "Unlicense"
			}
		  }
		],
		"purl": "pkg:npm/tweetnacl@0.14.5",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://tweetnacl.js.org"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/dchest/tweetnacl-js/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/dchest/tweetnacl-js.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/dashdash@1.14.1",
		"name": "dashdash",
		"version": "1.14.1",
		"description": "A light, featureful and explicit option parsing library.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "853cfa0f7cbe2fed5de20326b8dd581035f6e2f0"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/dashdash@1.14.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/trentm/node-dashdash#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/trentm/node-dashdash/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/trentm/node-dashdash.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/ecc-jsbn@0.1.2",
		"name": "ecc-jsbn",
		"version": "0.1.2",
		"description": "ECC JS code based on JSBN",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "3a83a904e54353287874c564b7549386849a98c9"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/ecc-jsbn@0.1.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/quartzjer/ecc-jsbn"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/quartzjer/ecc-jsbn/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/quartzjer/ecc-jsbn.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/jsbn@0.1.1",
		"name": "jsbn",
		"version": "0.1.1",
		"description": "The jsbn library is a fast, portable implementation of large-number math in pure JavaScript, enabling public-key crypto and other applications on desktop and mobile browsers.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "a5e654c2e5a2deb5f201d96cefbca80c0ef2f513"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/jsbn@0.1.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/andyperlitch/jsbn#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/andyperlitch/jsbn/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/andyperlitch/jsbn.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/getpass@0.1.7",
		"name": "getpass",
		"version": "0.1.7",
		"description": "getpass for node.js",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "5eff8e3e684d569ae4cb2b1282604e8ba62149fa"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/getpass@0.1.7",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/arekinath/node-getpass#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/arekinath/node-getpass/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/arekinath/node-getpass.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/is-typedarray@1.0.0",
		"name": "is-typedarray",
		"version": "1.0.0",
		"description": "Detect whether or not an object is a Typed Array",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "e479c80858df0c1b11ddda6940f96011fcda4a9a"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/is-typedarray@1.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/hughsk/is-typedarray"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/hughsk/is-typedarray/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/hughsk/is-typedarray.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/isstream@0.1.2",
		"name": "isstream",
		"version": "0.1.2",
		"description": "Determine if an object is a Stream",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "47e63f7af55afa6f92e1500e690eb8b8529c099a"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/isstream@0.1.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/rvagg/isstream"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/rvagg/isstream/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/rvagg/isstream.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/json-stringify-safe@5.0.1",
		"name": "json-stringify-safe",
		"version": "5.0.1",
		"description": "Like JSON.stringify, but doesn't blow up on circular refs.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "1296a2d58fd45f19a0f6ce01d65701e2c735b6eb"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "ISC"
			}
		  }
		],
		"purl": "pkg:npm/json-stringify-safe@5.0.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/isaacs/json-stringify-safe"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/isaacs/json-stringify-safe/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/isaacs/json-stringify-safe.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/oauth-sign@0.9.0",
		"name": "oauth-sign",
		"version": "0.9.0",
		"description": "OAuth 1 signing. Formerly a vendor lib in mikeal/request, now a standalone module.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "47a7b016baa68b5fa0ecf3dee08a85c679ac6455"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "Apache-2.0"
			}
		  }
		],
		"purl": "pkg:npm/oauth-sign@0.9.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/mikeal/oauth-sign#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/mikeal/oauth-sign/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/mikeal/oauth-sign.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/performance-now@2.1.0",
		"name": "performance-now",
		"version": "2.1.0",
		"description": "Implements performance.now (based on process.hrtime).",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "6309f4e0e5fa913ec1c69307ae364b4b377c9e7b"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/performance-now@2.1.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/braveg1rl/performance-now"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/braveg1rl/performance-now/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/braveg1rl/performance-now.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/qs@6.5.2",
		"name": "qs",
		"version": "6.5.2",
		"description": "A querystring parser that supports nesting and arrays, with a depth limit",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "cb3ae806e8740444584ef154ce8ee98d403f3e36"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "BSD-3-Clause"
			}
		  }
		],
		"purl": "pkg:npm/qs@6.5.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/ljharb/qs"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/ljharb/qs/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/ljharb/qs.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/tough-cookie@2.5.0",
		"name": "tough-cookie",
		"version": "2.5.0",
		"description": "RFC6265 Cookies and Cookie Jar for node.js",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "cd9fb2a0aa1d5a12b473bd9fb96fa3dcff65ade2"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "BSD-3-Clause"
			}
		  }
		],
		"purl": "pkg:npm/tough-cookie@2.5.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/salesforce/tough-cookie"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/salesforce/tough-cookie/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/salesforce/tough-cookie.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/psl@1.8.0",
		"name": "psl",
		"version": "1.8.0",
		"description": "Domain name parser based on the Public Suffix List",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "9326f8bcfb013adcc005fdff056acce020e51c24"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/psl@1.8.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/lupomontero/psl#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/lupomontero/psl/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+ssh://git@github.com/lupomontero/psl.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/uuid@3.4.0",
		"name": "uuid",
		"version": "3.4.0",
		"description": "RFC4122 (v1, v4, and v5) UUIDs",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "b23e4358afa8a202fe7a100af1f5f883f02007ee"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/uuid@3.4.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/uuidjs/uuid#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/uuidjs/uuid/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/uuidjs/uuid.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/sanitize-filename@1.6.3",
		"name": "sanitize-filename",
		"version": "1.6.3",
		"description": "Sanitize a string for use as a filename",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "755ebd752045931977e30b2025d340d7c9090378"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "name": "WTFPL OR ISC"
			}
		  }
		],
		"purl": "pkg:npm/sanitize-filename@1.6.3",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/parshap/node-sanitize-filename#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/parshap/node-sanitize-filename/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+ssh://git@github.com/parshap/node-sanitize-filename.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/truncate-utf8-bytes@1.0.2",
		"name": "truncate-utf8-bytes",
		"version": "1.0.2",
		"description": "Truncate string to given length in bytes",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "405923909592d56f78a5818434b0b78489ca5f2b"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "WTFPL"
			}
		  }
		],
		"purl": "pkg:npm/truncate-utf8-bytes@1.0.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/parshap/truncate-utf8-bytes#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/parshap/truncate-utf8-bytes/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/parshap/truncate-utf8-bytes.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/utf8-byte-length@1.0.4",
		"name": "utf8-byte-length",
		"version": "1.0.4",
		"description": "Get utf8 byte length of string",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "f45f150c4c66eee968186505ab93fcbb8ad6bf61"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "WTFPL"
			}
		  }
		],
		"purl": "pkg:npm/utf8-byte-length@1.0.4",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/parshap/utf8-byte-length#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/parshap/utf8-byte-length/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/parshap/utf8-byte-length.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/sanitize-html@1.4.2",
		"name": "sanitize-html",
		"version": "1.4.2",
		"description": "Clean up user-submitted HTML, preserving whitelisted elements and whitelisted attributes on a per-element basis",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "0bcc3dc92ba79d8b5dbea8b851c13d50d5ed3d58"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/sanitize-html@1.4.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/punkave/sanitize-html#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/punkave/sanitize-html/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/punkave/sanitize-html.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/he@0.4.1",
		"name": "he",
		"version": "0.4.1",
		"description": "A robust HTML entities encoder/decoder with full Unicode support.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "c86667614d2dd71bc737a197c760fb2eec8a1921"
		  }
		],
		"purl": "pkg:npm/he@0.4.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "http://mths.be/he"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/mathiasbynens/he/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/mathiasbynens/he.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/htmlparser2@3.3.0",
		"name": "htmlparser2",
		"version": "3.3.0",
		"description": "Fast & forgiving HTML/XML/RSS parser",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "cc70d05a59f6542e43f0e685c982e14c924a9efe"
		  }
		],
		"purl": "pkg:npm/htmlparser2@3.3.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/fb55/htmlparser2#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "http://github.com/fb55/htmlparser2/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/fb55/htmlparser2.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/domelementtype@1.3.1",
		"name": "domelementtype",
		"version": "1.3.1",
		"description": "all the types of nodes in htmlparser2's dom",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "d048c44b37b0d10a7f2a3d5fee3f4333d790481f"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "BSD-2-Clause"
			}
		  }
		],
		"purl": "pkg:npm/domelementtype@1.3.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/fb55/domelementtype#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/fb55/domelementtype/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/fb55/domelementtype.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/domhandler@2.1.0",
		"name": "domhandler",
		"version": "2.1.0",
		"description": "handler for htmlparser2 that turns pages into a dom",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "d2646f5e57f6c3bab11cf6cb05d3c0acf7412594"
		  }
		],
		"purl": "pkg:npm/domhandler@2.1.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/fb55/domhandler#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/fb55/domhandler/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/fb55/domhandler.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/domutils@1.1.6",
		"name": "domutils",
		"version": "1.1.6",
		"description": "utilities for working with htmlparser2's dom",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "bddc3de099b9a2efacc51c623f28f416ecc57485"
		  }
		],
		"purl": "pkg:npm/domutils@1.1.6",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/FB55/domutils#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/FB55/domutils/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/FB55/domutils.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/readable-stream@1.0.34",
		"name": "readable-stream",
		"version": "1.0.34",
		"description": "Streams2, a user-land copy of the stream library from Node.js v0.10.x",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "125820e34bc842d2f2aaafafe4c2916ee32c157c"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/readable-stream@1.0.34",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/isaacs/readable-stream#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/isaacs/readable-stream/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/isaacs/readable-stream.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/lodash@2.4.2",
		"name": "lodash",
		"version": "2.4.2",
		"description": "A utility library delivering consistency, customization, performance, & extras.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "fadd834b9683073da179b3eae6d9c0d15053f73e"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/lodash@2.4.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "http://lodash.com/"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/lodash/lodash/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/lodash/lodash.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/semver@7.3.2",
		"name": "semver",
		"version": "7.3.2",
		"description": "The semantic version parser used by npm.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "604962b052b81ed0786aae84389ffba70ffd3938"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "ISC"
			}
		  }
		],
		"purl": "pkg:npm/semver@7.3.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/npm/node-semver#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/npm/node-semver/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/npm/node-semver.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/sequelize@5.22.3",
		"name": "sequelize",
		"version": "5.22.3",
		"description": "Multi dialect ORM for Node.JS",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "7e7a92ddd355d883c9eb11cdb106d874d0d2636f"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/sequelize@5.22.3",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://sequelize.org/"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/sequelize/sequelize/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/sequelize/sequelize.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/cls-bluebird@2.1.0",
		"name": "cls-bluebird",
		"version": "2.1.0",
		"description": "Make bluebird work with the continuation-local-storage module.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "37ef1e080a8ffb55c2f4164f536f1919e7968aee"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "BSD-2-Clause"
			}
		  }
		],
		"purl": "pkg:npm/cls-bluebird@2.1.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/TimBeyer/cls-bluebird#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/TimBeyer/cls-bluebird/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/TimBeyer/cls-bluebird.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/is-bluebird@1.0.2",
		"name": "is-bluebird",
		"version": "1.0.2",
		"description": "Is this a bluebird promise I see before me?",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "096439060f4aa411abee19143a84d6a55346d6e2"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/is-bluebird@1.0.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/overlookmotel/is-bluebird#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/overlookmotel/is-bluebird/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/overlookmotel/is-bluebird.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/shimmer@1.2.1",
		"name": "shimmer",
		"version": "1.2.1",
		"description": "Safe(r) monkeypatching for JavaScript.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "610859f7de327b587efebf501fb43117f9aff337"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "BSD-2-Clause"
			}
		  }
		],
		"purl": "pkg:npm/shimmer@1.2.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/othiym23/shimmer#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/othiym23/shimmer/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/othiym23/shimmer.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/debug@4.1.1",
		"name": "debug",
		"version": "4.1.1",
		"description": "small debugging utility",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "3b72260255109c6b589cee050f1d516139664791"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/debug@4.1.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/visionmedia/debug#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/visionmedia/debug/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/visionmedia/debug.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/moment-timezone@0.5.31",
		"name": "moment-timezone",
		"version": "0.5.31",
		"description": "Parse and display moments in any timezone.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "9c40d8c5026f0c7ab46eda3d63e49c155148de05"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/moment-timezone@0.5.31",
		"externalReferences": [
		  {
			"type": "website",
			"url": "http://momentjs.com/timezone/"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/moment/moment-timezone/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/moment/moment-timezone.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/retry-as-promised@3.2.0",
		"name": "retry-as-promised",
		"version": "3.2.0",
		"description": "Retry a failed promise",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "769f63d536bec4783549db0777cb56dadd9d8543"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/retry-as-promised@3.2.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/mickhansen/retry-as-promised"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/mickhansen/retry-as-promised/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/mickhansen/retry-as-promised.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/any-promise@1.3.0",
		"name": "any-promise",
		"version": "1.3.0",
		"description": "Resolve any installed ES6 compatible promise",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "abc6afeedcea52e809cdc0376aed3ce39635d17f"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/any-promise@1.3.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "http://github.com/kevinbeaty/any-promise"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/kevinbeaty/any-promise/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/kevinbeaty/any-promise.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/semver@6.3.0",
		"name": "semver",
		"version": "6.3.0",
		"description": "The semantic version parser used by npm.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "ee0a64c8af5e8ceea67687b133761e1becbd1d3d"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "ISC"
			}
		  }
		],
		"purl": "pkg:npm/semver@6.3.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/npm/node-semver#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/npm/node-semver/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/npm/node-semver.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/sequelize-pool@2.3.0",
		"name": "sequelize-pool",
		"version": "2.3.0",
		"description": "Resource pooling for Node.JS",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "64f1fe8744228172c474f530604b6133be64993d"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/sequelize-pool@2.3.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/sushantdhiman/sequelize-pool#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/sushantdhiman/sequelize-pool/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+ssh://git@github.com/sushantdhiman/sequelize-pool.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/toposort-class@1.0.1",
		"name": "toposort-class",
		"version": "1.0.1",
		"description": "Topological sort of directed acyclic graphs (like dependecy lists)",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "7ffd1f78c8be28c3ba45cd4e1a3f5ee193bd9988"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/toposort-class@1.0.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/gustavohenke/toposort#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/gustavohenke/toposort/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/gustavohenke/toposort.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/validator@10.11.0",
		"name": "validator",
		"version": "10.11.0",
		"description": "String validation and sanitization",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "003108ea6e9a9874d31ccc9e5006856ccd76b228"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/validator@10.11.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/chriso/validator.js"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/chriso/validator.js/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/chriso/validator.js.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/wkx@0.4.8",
		"name": "wkx",
		"version": "0.4.8",
		"description": "A WKT/WKB/EWKT/EWKB/TWKB/GeoJSON parser and serializer",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "a092cf088d112683fdc7182fd31493b2c5820003"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/wkx@0.4.8",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/cschwarz/wkx#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/cschwarz/wkx/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+ssh://git@github.com/cschwarz/wkx.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/%40types/node@14.0.23",
		"group": "@types",
		"name": "node",
		"version": "14.0.23",
		"description": "TypeScript definitions for Node.js",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "676fa0883450ed9da0bb24156213636290892806"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/%40types/node@14.0.23",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/DefinitelyTyped/DefinitelyTyped#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/DefinitelyTyped/DefinitelyTyped/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/DefinitelyTyped/DefinitelyTyped.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/sequelize-noupdate-attributes@1.0.0",
		"name": "sequelize-noupdate-attributes",
		"version": "1.0.0",
		"description": "A very simple Sequelize plugin which adds no update and readonly attributes support.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "70ab724563742b2c6fbadc507c91c01041b5fb38"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/sequelize-noupdate-attributes@1.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/diosney/node-sequelize-noupdate-attributes"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/diosney/node-sequelize-noupdate-attributes/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/diosney/node-sequelize-noupdate-attributes.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/serve-index@1.9.1",
		"name": "serve-index",
		"version": "1.9.1",
		"description": "Serve directory listings",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "d3768d69b1e7d82e5ce050fff5b453bea12a9239"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/serve-index@1.9.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/expressjs/serve-index#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/expressjs/serve-index/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/expressjs/serve-index.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/batch@0.6.1",
		"name": "batch",
		"version": "0.6.1",
		"description": "Simple async batch with concurrency control and progress reporting.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "dc34314f4e679318093fc760272525f94bf25c16"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/batch@0.6.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/visionmedia/batch#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/visionmedia/batch/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/visionmedia/batch.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/http-errors@1.6.3",
		"name": "http-errors",
		"version": "1.6.3",
		"description": "Create HTTP error objects",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "8b55680bb4be283a0b5bf4ea2e38580be1d9320d"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/http-errors@1.6.3",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jshttp/http-errors#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jshttp/http-errors/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jshttp/http-errors.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/setprototypeof@1.1.0",
		"name": "setprototypeof",
		"version": "1.1.0",
		"description": "A small polyfill for Object.setprototypeof",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "d0bd85536887b6fe7c0d818cb962d9d91c54e656"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "ISC"
			}
		  }
		],
		"purl": "pkg:npm/setprototypeof@1.1.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/wesleytodd/setprototypeof"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/wesleytodd/setprototypeof/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/wesleytodd/setprototypeof.git"
		  }
		]
	  },
	  {
		"type": "framework",
		"bom-ref": "pkg:npm/socket.io@2.3.0",
		"name": "socket.io",
		"version": "2.3.0",
		"description": "node.js realtime framework server",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "cd762ed6a4faeca59bc1f3e243c0969311eb73fb"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/socket.io@2.3.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/socketio/socket.io#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/socketio/socket.io/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/socketio/socket.io.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/engine.io@3.4.2",
		"name": "engine.io",
		"version": "3.4.2",
		"description": "The realtime engine behind Socket.IO. Provides the foundation of a bidirectional connection between client and server",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "8fc84ee00388e3e228645e0a7d3dfaeed5bd122c"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/engine.io@3.4.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/socketio/engine.io"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/socketio/engine.io/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+ssh://git@github.com/socketio/engine.io.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/base64id@2.0.0",
		"name": "base64id",
		"version": "2.0.0",
		"description": "Generates a base64 id",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "2770ac6bc47d312af97a8bf9a634342e0cd25cb6"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/base64id@2.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/faeldt/base64id#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/faeldt/base64id/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/faeldt/base64id.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/cookie@0.3.1",
		"name": "cookie",
		"version": "0.3.1",
		"description": "HTTP server cookie parsing and serialization",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "e7e0a1f9ef43b4c8ba925c5c5a96e806d16873bb"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/cookie@0.3.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jshttp/cookie#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jshttp/cookie/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/jshttp/cookie.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/engine.io-parser@2.2.0",
		"name": "engine.io-parser",
		"version": "2.2.0",
		"description": "Parser for the client for the realtime Engine",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "312c4894f57d52a02b420868da7b5c1c84af80ed"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/engine.io-parser@2.2.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/socketio/engine.io-parser"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/socketio/engine.io-parser/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+ssh://git@github.com/socketio/engine.io-parser.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/after@0.8.2",
		"name": "after",
		"version": "0.8.2",
		"description": "after - tiny flow control",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "fedb394f9f0e02aa9768e702bda23b505fae7e1f"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/after@0.8.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/Raynos/after#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/Raynos/after/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/Raynos/after.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/arraybuffer.slice@0.0.7",
		"name": "arraybuffer.slice",
		"version": "0.0.7",
		"description": "Exports a function for slicing ArrayBuffers (no polyfilling)",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "3bbc4275dd584cc1b10809b89d4e8b63a69e7675"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/arraybuffer.slice@0.0.7",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/rase-/arraybuffer.slice"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/rase-/arraybuffer.slice/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+ssh://git@github.com/rase-/arraybuffer.slice.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/base64-arraybuffer@0.1.5",
		"name": "base64-arraybuffer",
		"version": "0.1.5",
		"description": "Encode/decode base64 data into ArrayBuffers",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "73926771923b5a19747ad666aa5cd4bf9c6e9ce8"
		  }
		],
		"purl": "pkg:npm/base64-arraybuffer@0.1.5",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/niklasvh/base64-arraybuffer"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/niklasvh/base64-arraybuffer/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/niklasvh/base64-arraybuffer.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/blob@0.0.5",
		"name": "blob",
		"version": "0.0.5",
		"description": "Abstracts out Blob and uses BlobBulder in cases where it is supported with any vendor prefix.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "d680eeef25f8cd91ad533f5b01eed48e64caf683"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/blob@0.0.5",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/webmodules/blob"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/webmodules/blob/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/webmodules/blob.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/has-binary2@1.0.3",
		"name": "has-binary2",
		"version": "1.0.3",
		"description": "A function that takes anything in javascript and returns true if its argument contains binary data.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "7776ac627f3ea77250cfc332dab7ddf5e4f5d11d"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/has-binary2@1.0.3"
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/isarray@2.0.1",
		"name": "isarray",
		"version": "2.0.1",
		"description": "Array#isArray for older browsers",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "a37d94ed9cda2d59865c9f76fe596ee1f338741e"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/isarray@2.0.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/juliangruber/isarray"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/juliangruber/isarray/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/juliangruber/isarray.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/ws@7.3.1",
		"name": "ws",
		"version": "7.3.1",
		"description": "Simple to use, blazing fast and thoroughly tested websocket client and server for Node.js",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "d0547bf67f7ce4f12a72dfe31262c68d7dc551c8"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/ws@7.3.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/websockets/ws"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/websockets/ws/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/websockets/ws.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/socket.io-adapter@1.1.2",
		"name": "socket.io-adapter",
		"version": "1.1.2",
		"description": "default socket.io in-memory adapter",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "ab3f0d6f66b8fc7fca3959ab5991f82221789be9"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/socket.io-adapter@1.1.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/socketio/socket.io-adapter#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/socketio/socket.io-adapter/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/socketio/socket.io-adapter.git"
		  }
		]
	  },
	  {
		"type": "framework",
		"bom-ref": "pkg:npm/socket.io-client@2.3.0",
		"name": "socket.io-client",
		"version": "2.3.0",
		"description": "[![Build Status](https://secure.travis-ci.org/socketio/socket.io-client.svg?branch=master)](http://travis-ci.org/socketio/socket.io-client) [![Dependency Status](https://david-dm.org/socketio/socket.io-client.svg)](https://david-dm.org/socketio/socket.io-client) [![devDependency Status](https://david-dm.org/socketio/socket.io-client/dev-status.svg)](https://david-dm.org/socketio/socket.io-client#info=devDependencies) [![NPM version](https://badge.fury.io/js/socket.io-client.svg)](https://www.npmjs.com/package/socket.io-client) ![Downloads](http://img.shields.io/npm/dm/socket.io-client.svg?style=flat) [![](http://slack.socket.io/badge.svg?)](http://slack.socket.io)",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "14d5ba2e00b9bcd145ae443ab96b3f86cbcc1bb4"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/socket.io-client@2.3.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/Automattic/socket.io-client#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/Automattic/socket.io-client/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/Automattic/socket.io-client.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/backo2@1.0.2",
		"name": "backo2",
		"version": "1.0.2",
		"description": "simple backoff based on segmentio/backo",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "31ab1ac8b129363463e35b3ebb69f4dfcfba7947"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/backo2@1.0.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/mokesmokes/backo#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/mokesmokes/backo/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/mokesmokes/backo.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/component-bind@1.0.0",
		"name": "component-bind",
		"version": "1.0.0",
		"description": "function binding utility",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "00c608ab7dcd93897c0009651b1d3a8e1e73bbd1"
		  }
		],
		"purl": "pkg:npm/component-bind@1.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/component/bind#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/component/bind/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/component/bind.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/component-emitter@1.2.1",
		"name": "component-emitter",
		"version": "1.2.1",
		"description": "Event emitter",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "137918d6d78283f7df7a6b7c5a63e140e69425e6"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/component-emitter@1.2.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/component/emitter#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/component/emitter/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/component/emitter.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/engine.io-client@3.4.3",
		"name": "engine.io-client",
		"version": "3.4.3",
		"description": "Client for the realtime Engine",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "192d09865403e3097e3575ebfeb3861c4d01a66c"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/engine.io-client@3.4.3",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/socketio/engine.io-client"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/socketio/engine.io-client/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/socketio/engine.io-client.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/component-inherit@0.0.3",
		"name": "component-inherit",
		"version": "0.0.3",
		"description": "Prototype inheritance utility",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "645fc4adf58b72b649d5cae65135619db26ff143"
		  }
		],
		"purl": "pkg:npm/component-inherit@0.0.3",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/component/inherit#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/component/inherit/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/component/inherit.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/has-cors@1.1.0",
		"name": "has-cors",
		"version": "1.1.0",
		"description": "Detects support for Cross-Origin Resource Sharing",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "5e474793f7ea9843d1bb99c23eef49ff126fff39"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/has-cors@1.1.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/component/has-cors#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/component/has-cors/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/component/has-cors.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/indexof@0.0.1",
		"name": "indexof",
		"version": "0.0.1",
		"description": "Microsoft sucks",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "82dc336d232b9062179d05ab3293a66059fd435d"
		  }
		],
		"purl": "pkg:npm/indexof@0.0.1"
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/parseqs@0.0.5",
		"name": "parseqs",
		"version": "0.0.5",
		"description": "Provides methods for parsing a query string into an object, and vice versa.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "d5208a3738e46766e291ba2ea173684921a8b89d"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/parseqs@0.0.5",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/get/querystring"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/get/querystring/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/get/querystring.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/better-assert@1.0.2",
		"name": "better-assert",
		"version": "1.0.2",
		"description": "Better assertions for node, reporting the expr, filename, lineno etc",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "40866b9e1b9e0b55b481894311e68faffaebc522"
		  }
		],
		"purl": "pkg:npm/better-assert@1.0.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/visionmedia/better-assert#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/visionmedia/better-assert/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/visionmedia/better-assert.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/callsite@1.0.0",
		"name": "callsite",
		"version": "1.0.0",
		"description": "access to v8's CallSites",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "280398e5d664bd74038b6f0905153e6e8af1bc20"
		  }
		],
		"purl": "pkg:npm/callsite@1.0.0"
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/parseuri@0.0.5",
		"name": "parseuri",
		"version": "0.0.5",
		"description": "Method that parses a URI and returns an array of its components",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "80204a50d4dbb779bfdc6ebe2778d90e4bce320a"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/parseuri@0.0.5",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/get/parseuri"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/get/parseuri/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/get/parseuri.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/ws@6.1.4",
		"name": "ws",
		"version": "6.1.4",
		"description": "Simple to use, blazing fast and thoroughly tested websocket client and server for Node.js",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "5b5c8800afab925e94ccb29d153c8d02c1776ef9"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/ws@6.1.4",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/websockets/ws"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/websockets/ws/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/websockets/ws.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/async-limiter@1.0.1",
		"name": "async-limiter",
		"version": "1.0.1",
		"description": "asynchronous function queue with adjustable concurrency",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "dd379e94f0db8310b08291f9d64c3209766617fd"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/async-limiter@1.0.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/strml/async-limiter#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/strml/async-limiter/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/strml/async-limiter.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/xmlhttprequest-ssl@1.5.5",
		"name": "xmlhttprequest-ssl",
		"version": "1.5.5",
		"description": "XMLHttpRequest for Node",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "c2876b06168aadc40e57d97e81191ac8f4398b3e"
		  }
		],
		"purl": "pkg:npm/xmlhttprequest-ssl@1.5.5",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/mjwwit/node-XMLHttpRequest#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "http://github.com/mjwwit/node-XMLHttpRequest/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/mjwwit/node-XMLHttpRequest.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/yeast@0.1.2",
		"name": "yeast",
		"version": "0.1.2",
		"description": "Tiny but linear growing unique id generator",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "008e06d8094320c372dbc2f8ed76a0ca6c8ac419"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/yeast@0.1.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/unshiftio/yeast"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/unshiftio/yeast/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/unshiftio/yeast.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/object-component@0.0.3",
		"name": "object-component",
		"version": "0.0.3",
		"description": "Object utils.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "f0c69aa50efc95b866c186f400a33769cb2f1291"
		  }
		],
		"purl": "pkg:npm/object-component@0.0.3"
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/socket.io-parser@3.3.0",
		"name": "socket.io-parser",
		"version": "3.3.0",
		"description": "socket.io protocol parser",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "2b52a96a509fdf31440ba40fed6094c7d4f1262f"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/socket.io-parser@3.3.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/Automattic/socket.io-parser#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/Automattic/socket.io-parser/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/Automattic/socket.io-parser.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/debug@3.1.0",
		"name": "debug",
		"version": "3.1.0",
		"description": "small debugging utility",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "5bb5a0672628b64149566ba16819e61518c67261"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/debug@3.1.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/visionmedia/debug#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/visionmedia/debug/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/visionmedia/debug.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/to-array@0.1.4",
		"name": "to-array",
		"version": "0.1.4",
		"description": "Turn an array like into an array",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "17e6c11f73dd4f3d74cda7a4ff3238e9ad9bf890"
		  }
		],
		"purl": "pkg:npm/to-array@0.1.4",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/Raynos/to-array"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/Raynos/to-array/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/Raynos/to-array.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/socket.io-parser@3.4.1",
		"name": "socket.io-parser",
		"version": "3.4.1",
		"description": "socket.io protocol parser",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "b06af838302975837eab2dc980037da24054d64a"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/socket.io-parser@3.4.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/socketio/socket.io-parser#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/socketio/socket.io-parser/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/socketio/socket.io-parser.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/sqlite3@4.2.0",
		"name": "sqlite3",
		"version": "4.2.0",
		"description": "Asynchronous, non-blocking SQLite3 bindings",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "49026d665e9fc4f922e56fb9711ba5b4c85c4901"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "BSD-3-Clause"
			}
		  }
		],
		"purl": "pkg:npm/sqlite3@4.2.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/mapbox/node-sqlite3"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/mapbox/node-sqlite3/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/mapbox/node-sqlite3.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/node-pre-gyp@0.11.0",
		"name": "node-pre-gyp",
		"version": "0.11.0",
		"description": "Node.js native addon binary install tool",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "db1f33215272f692cd38f03238e3e9b47c5dd054"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "BSD-3-Clause"
			}
		  }
		],
		"purl": "pkg:npm/node-pre-gyp@0.11.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/mapbox/node-pre-gyp#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/mapbox/node-pre-gyp/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/mapbox/node-pre-gyp.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/svg-captcha@1.4.0",
		"name": "svg-captcha",
		"version": "1.4.0",
		"description": "generate svg captcha in node.js or express.js",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "32ead3c6463936c218bb3bc9ed04fea4eeffe492"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/svg-captcha@1.4.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/steambap/svg-captcha#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/steambap/svg-captcha/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/steambap/svg-captcha.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/opentype.js@0.7.3",
		"name": "opentype.js",
		"version": "0.7.3",
		"description": "OpenType font parser",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "40fb8ce18bfd60e74448efdfe442834098397aab"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/opentype.js@0.7.3",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/nodebox/opentype.js#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/nodebox/opentype.js/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/nodebox/opentype.js.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/swagger-ui-express@4.1.4",
		"name": "swagger-ui-express",
		"version": "4.1.4",
		"description": "Swagger UI Express",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "8b814ad998b850a1cf90e71808d6d0a8a8daf742"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/swagger-ui-express@4.1.4",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/scottie1984/swagger-ui-express"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/scottie1984/swagger-ui-express/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+ssh://git@github.com/scottie1984/swagger-ui-express.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/swagger-ui-dist@3.28.0",
		"name": "swagger-ui-dist",
		"version": "3.28.0",
		"description": "[![NPM version](https://badge.fury.io/js/swagger-ui-dist.svg)](http://badge.fury.io/js/swagger-ui-dist)",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "7c30ece92f815c1f34de3d394e12983e97f3d421"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "Apache-2.0"
			}
		  }
		],
		"purl": "pkg:npm/swagger-ui-dist@3.28.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/swagger-api/swagger-ui#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/swagger-api/swagger-ui/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+ssh://git@github.com/swagger-api/swagger-ui.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/unzipper@0.9.15",
		"name": "unzipper",
		"version": "0.9.15",
		"description": "Unzip cross-platform streaming API ",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "97d99203dad17698ee39882483c14e4845c7549c"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/unzipper@0.9.15",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/ZJONSSON/node-unzipper#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/ZJONSSON/node-unzipper/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/ZJONSSON/node-unzipper.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/big-integer@1.6.48",
		"name": "big-integer",
		"version": "1.6.48",
		"description": "An arbitrary length integer library for Javascript",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "8fd88bd1632cba4a1c8c3e3d7159f08bb95b4b9e"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "Unlicense"
			}
		  }
		],
		"purl": "pkg:npm/big-integer@1.6.48",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/peterolson/BigInteger.js#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/peterolson/BigInteger.js/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+ssh://git@github.com/peterolson/BigInteger.js.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/binary@0.3.0",
		"name": "binary",
		"version": "0.3.0",
		"description": "Unpack multibyte binary values from buffers",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "9f60553bc5ce8c3386f3b553cff47462adecaa79"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/binary@0.3.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/substack/node-binary#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/substack/node-binary/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+ssh://git@github.com/substack/node-binary.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/buffers@0.1.1",
		"name": "buffers",
		"version": "0.1.1",
		"description": "Treat a collection of Buffers as a single contiguous partially mutable Buffer.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "b24579c3bed4d6d396aeee6d9a8ae7f5482ab7bb"
		  }
		],
		"purl": "pkg:npm/buffers@0.1.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/substack/node-buffers#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/substack/node-buffers/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+ssh://git@github.com/substack/node-buffers.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/chainsaw@0.1.0",
		"name": "chainsaw",
		"version": "0.1.0",
		"description": "Build chainable fluent interfaces the easy way... with a freakin' chainsaw!",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "5eab50b28afe58074d0d58291388828b5e5fbc98"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "name": "MIT/X11"
			}
		  }
		],
		"purl": "pkg:npm/chainsaw@0.1.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/substack/node-chainsaw#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/substack/node-chainsaw/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+ssh://git@github.com/substack/node-chainsaw.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/traverse@0.3.9",
		"name": "traverse",
		"version": "0.3.9",
		"description": "Traverse and transform objects by visiting every node on a recursive walk",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "717b8f220cc0bb7b44e40514c22b2e8bbc70d8b9"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "name": "MIT/X11"
			}
		  }
		],
		"purl": "pkg:npm/traverse@0.3.9",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/substack/js-traverse#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/substack/js-traverse/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+ssh://git@github.com/substack/js-traverse.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/bluebird@3.4.7",
		"name": "bluebird",
		"version": "3.4.7",
		"description": "Full featured Promises/A+ implementation with exceptionally good performance",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "f72d760be09b7f76d08ed8fae98b289a8d05fab3"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/bluebird@3.4.7",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/petkaantonov/bluebird"
		  },
		  {
			"type": "issue-tracker",
			"url": "http://github.com/petkaantonov/bluebird/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/petkaantonov/bluebird.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/buffer-indexof-polyfill@1.0.1",
		"name": "buffer-indexof-polyfill",
		"version": "1.0.1",
		"description": "This is a polyfill for Buffer#indexOf introduced in NodeJS 4.0.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "a9fb806ce8145d5428510ce72f278bb363a638bf"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/buffer-indexof-polyfill@1.0.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/sarosia/buffer-indexof-polyfill#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/sarosia/buffer-indexof-polyfill/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/sarosia/buffer-indexof-polyfill.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/fstream@1.0.12",
		"name": "fstream",
		"version": "1.0.12",
		"description": "Advanced file system stream things",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "4e8ba8ee2d48be4f7d0de505455548eae5932045"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "ISC"
			}
		  }
		],
		"purl": "pkg:npm/fstream@1.0.12",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/npm/fstream#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/npm/fstream/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/npm/fstream.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/listenercount@1.0.1",
		"name": "listenercount",
		"version": "1.0.1",
		"description": "backwards compatible version of builtin events.listenercount",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "84c8a72ab59c4725321480c975e6508342e70937"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "ISC"
			}
		  }
		],
		"purl": "pkg:npm/listenercount@1.0.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/jden/node-listenercount#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/jden/node-listenercount/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+ssh://git@github.com/jden/node-listenercount.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/setimmediate@1.0.5",
		"name": "setimmediate",
		"version": "1.0.5",
		"description": "A shim for the setImmediate efficient script yielding API",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "290cbb232e306942d7d7ea9b83732ab7856f8285"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/setimmediate@1.0.5",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/YuzuJS/setImmediate#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/YuzuJS/setImmediate/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/YuzuJS/setImmediate.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/winston@3.3.3",
		"name": "winston",
		"version": "3.3.3",
		"description": "A logger for just about everything.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "ae6172042cafb29786afa3d09c8ff833ab7c9170"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/winston@3.3.3",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/winstonjs/winston#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/winstonjs/winston/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/winstonjs/winston.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/%40dabh/diagnostics@2.0.2",
		"group": "@dabh",
		"name": "diagnostics",
		"version": "2.0.2",
		"description": "Tools for debugging your node.js modules and event loop",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "290d08f7b381b8f94607dc8f471a12c675f9db31"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/%40dabh/diagnostics@2.0.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/3rd-Eden/diagnostics"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/3rd-Eden/diagnostics/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/3rd-Eden/diagnostics.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/colorspace@1.1.2",
		"name": "colorspace",
		"version": "1.1.2",
		"description": "Generate HEX colors for a given namespace.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "e0128950d082b86a2168580796a0aa5d6c68d8c5"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/colorspace@1.1.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/3rd-Eden/colorspace"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/3rd-Eden/colorspace/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/3rd-Eden/colorspace.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/color@3.0.0",
		"name": "color",
		"version": "3.0.0",
		"description": "Color conversion and manipulation with CSS string support",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "d920b4328d534a3ac8295d68f7bd4ba6c427be9a"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/color@3.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/Qix-/color#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/Qix-/color/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/Qix-/color.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/color-string@1.5.3",
		"name": "color-string",
		"version": "1.5.3",
		"description": "Parser and generator for CSS color strings",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "c9bbc5f01b58b5492f3d6857459cb6590ce204cc"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/color-string@1.5.3",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/Qix-/color-string#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/Qix-/color-string/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/Qix-/color-string.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/simple-swizzle@0.2.2",
		"name": "simple-swizzle",
		"version": "0.2.2",
		"description": "Simply swizzle your arguments",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "a4da6b635ffcccca33f70d17cb92592de95e557a"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/simple-swizzle@0.2.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/qix-/node-simple-swizzle#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/qix-/node-simple-swizzle/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/qix-/node-simple-swizzle.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/is-arrayish@0.3.2",
		"name": "is-arrayish",
		"version": "0.3.2",
		"description": "Determines if an object can be used as an array",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "4574a2ae56f7ab206896fb431eaeed066fdf8f03"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/is-arrayish@0.3.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/qix-/node-is-arrayish#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/qix-/node-is-arrayish/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/qix-/node-is-arrayish.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/text-hex@1.0.0",
		"name": "text-hex",
		"version": "1.0.0",
		"description": "Generate a hex color from the given text",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "69dc9c1b17446ee79a92bf5b884bb4b9127506f5"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/text-hex@1.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/3rd-Eden/text-hex"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/3rd-Eden/text-hex/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/3rd-Eden/text-hex.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/enabled@2.0.0",
		"name": "enabled",
		"version": "2.0.0",
		"description": "Check if a certain debug flag is enabled.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "f9dd92ec2d6f4bbc0d5d1e64e21d61cd4665e7c2"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/enabled@2.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/3rd-Eden/enabled#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/3rd-Eden/enabled/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/3rd-Eden/enabled.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/kuler@2.0.0",
		"name": "kuler",
		"version": "2.0.0",
		"description": "Color your terminal using CSS/hex color codes",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "e2c570a3800388fb44407e851531c1d670b061b3"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/kuler@2.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/3rd-Eden/kuler"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/3rd-Eden/kuler/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/3rd-Eden/kuler.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/async@3.2.0",
		"name": "async",
		"version": "3.2.0",
		"description": "Higher-order functions and common patterns for asynchronous code",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "b3a2685c5ebb641d3de02d161002c60fc9f85720"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/async@3.2.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://caolan.github.io/async/"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/caolan/async/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/caolan/async.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/is-stream@2.0.0",
		"name": "is-stream",
		"version": "2.0.0",
		"description": "Check if something is a Node.js stream",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "bde9c32680d6fae04129d6ac9d921ce7815f78e3"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/is-stream@2.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/sindresorhus/is-stream#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/sindresorhus/is-stream/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/sindresorhus/is-stream.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/logform@2.2.0",
		"name": "logform",
		"version": "2.2.0",
		"description": "An mutable object-based log format designed for chaining & objectMode streams.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "40f036d19161fc76b68ab50fdc7fe495544492f2"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/logform@2.2.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/winstonjs/logform#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/winstonjs/logform/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/winstonjs/logform.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/fast-safe-stringify@2.0.7",
		"name": "fast-safe-stringify",
		"version": "2.0.7",
		"description": "Safely and quickly serialize JavaScript objects",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "124aa885899261f68aedb42a7c080de9da608743"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/fast-safe-stringify@2.0.7",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/davidmarkclements/fast-safe-stringify#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/davidmarkclements/fast-safe-stringify/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/davidmarkclements/fast-safe-stringify.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/fecha@4.2.0",
		"name": "fecha",
		"version": "4.2.0",
		"description": "Date formatting and parsing",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "3ffb6395453e3f3efff850404f0a59b6747f5f41"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/fecha@4.2.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/taylorhakes/fecha"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/taylorhakes/fecha/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://taylorhakes@github.com/taylorhakes/fecha.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/triple-beam@1.3.0",
		"name": "triple-beam",
		"version": "1.3.0",
		"description": "Definitions of levels for logging purposes & shareable Symbol constants.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "a595214c7298db8339eeeee083e4d10bd8cb8dd9"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/triple-beam@1.3.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/winstonjs/triple-beam#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/winstonjs/triple-beam/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/winstonjs/triple-beam.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/one-time@1.0.0",
		"name": "one-time",
		"version": "1.0.0",
		"description": "Run the supplied function exactly one time (once)",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "e06bc174aed214ed58edede573b433bbf827cb45"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/one-time@1.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/3rd-Eden/one-time#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/3rd-Eden/one-time/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/3rd-Eden/one-time.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/fn.name@1.1.0",
		"name": "fn.name",
		"version": "1.1.0",
		"description": "Extract names from functions",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "26cad8017967aea8731bc42961d04a3d5988accc"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/fn.name@1.1.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/3rd-Eden/fn.name"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/3rd-Eden/fn.name/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/3rd-Eden/fn.name.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/readable-stream@3.6.0",
		"name": "readable-stream",
		"version": "3.6.0",
		"description": "Streams3, a user-land copy of the stream library from Node.js",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "337bbda3adc0706bd3e024426a286d4b4b2c9198"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/readable-stream@3.6.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/nodejs/readable-stream#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/nodejs/readable-stream/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/nodejs/readable-stream.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/stack-trace@0.0.10",
		"name": "stack-trace",
		"version": "0.0.10",
		"description": "Get v8 stack traces as an array of CallSite objects.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "547c70b347e8d32b4e108ea1a2a159e5fdde19c0"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/stack-trace@0.0.10",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/felixge/node-stack-trace"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/felixge/node-stack-trace/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/felixge/node-stack-trace.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/winston-transport@4.4.0",
		"name": "winston-transport",
		"version": "4.4.0",
		"description": "Base stream implementations for winston@3 and up.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "17af518daa690d5b2ecccaa7acf7b20ca7925e59"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/winston-transport@4.4.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/winstonjs/winston-transport#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/winstonjs/winston-transport/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+ssh://git@github.com/winstonjs/winston-transport.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/yaml-schema-validator@1.2.2",
		"name": "yaml-schema-validator",
		"version": "1.2.2",
		"description": "Schema validator for yaml files",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "48d85ccda92fed3acc51cdf706530c2927e09807"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/yaml-schema-validator@1.2.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/ketanTechracers/schema-validator#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/ketanTechracers/schema-validator/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/ketanTechracers/schema-validator.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/commander@2.20.3",
		"name": "commander",
		"version": "2.20.3",
		"description": "the complete solution for node.js command-line programs",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "fd485e84c03eb4881c20722ba48035e8531aeb33"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/commander@2.20.3",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/tj/commander.js#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/tj/commander.js/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/tj/commander.js.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/validate@4.5.1",
		"name": "validate",
		"version": "4.5.1",
		"description": "Validate object properties in javascript.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "ba36b8450b4bad4ccf52d666ba80abb2c01cace2"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/validate@4.5.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/eivindfjeldstad/validate#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/eivindfjeldstad/validate/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/eivindfjeldstad/validate.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/component-type@1.2.1",
		"name": "component-type",
		"version": "1.2.1",
		"description": "Cross-browser type assertions (less broken typeof)",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "8a47901700238e4fc32269771230226f24b415a9"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/component-type@1.2.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/component/type#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/component/type/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/component/type.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/eivindfjeldstad-dot@0.0.1",
		"name": "eivindfjeldstad-dot",
		"version": "0.0.1",
		"description": "Get and set object properties with dot notation",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "22fc976bfaf306e0839a31db8e8213480fafb893"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/eivindfjeldstad-dot@0.0.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/eivindfjeldstad/dot"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/eivindfjeldstad/dot/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/eivindfjeldstad/dot.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/typecast@0.0.1",
		"name": "typecast",
		"version": "0.0.1",
		"description": "Simple typecasting",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "fffb75dcb6bdf1def8e293b6b6e893d6c1ed19de"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/typecast@0.0.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/eivindfjeldstad/typecast#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/eivindfjeldstad/typecast/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/eivindfjeldstad/typecast.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/z85@0.0.2",
		"name": "z85",
		"version": "0.0.2",
		"description": "ZeroMQ Base-85 Encoding",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "45d353b13e4ee3d376c3fbd37dcda85feed8b0d3"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/z85@0.0.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/msealand/z85.node"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/msealand/z85.node/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/msealand/z85.node.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/expand-template@2.0.3",
		"name": "expand-template",
		"version": "2.0.3",
		"description": "Expand placeholders in a template string",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "6e14b3fcee0f3a6340ecb57d2e8918692052a47c"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "name": "(MIT OR WTFPL)"
			}
		  }
		],
		"purl": "pkg:npm/expand-template@2.0.3",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/ralphtheninja/expand-template"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/ralphtheninja/expand-template/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/ralphtheninja/expand-template.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/github-from-package@0.0.0",
		"name": "github-from-package",
		"version": "0.0.0",
		"description": "return the github url from a package.json file",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "97fb5d96bfde8973313f20e8288ef9a167fa64ce"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/github-from-package@0.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/substack/github-from-package"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/substack/github-from-package/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/substack/github-from-package.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/iltorb@2.4.5",
		"name": "iltorb",
		"version": "2.4.5",
		"description": "Brotli compression/decompression with native bindings",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "d64434b527099125c6839ed48b666247a172ef87"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/iltorb@2.4.5",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/nstepien/iltorb"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/nstepien/iltorb/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/nstepien/iltorb.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/prebuild-install@5.3.5",
		"name": "prebuild-install",
		"version": "5.3.5",
		"description": "A command line tool to easily install prebuilt binaries for multiple version of node/iojs on a specific platform",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "e7e71e425298785ea9d22d4f958dbaccf8bb0e1b"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/prebuild-install@5.3.5",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/prebuild/prebuild-install"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/prebuild/prebuild-install/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/prebuild/prebuild-install.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/napi-build-utils@1.0.2",
		"name": "napi-build-utils",
		"version": "1.0.2",
		"description": "A set of utilities to assist developers of tools that build N-API native add-ons",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "b1fddc0b2c46e380a0b7a76f984dd47c41a13806"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/napi-build-utils@1.0.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/inspiredware/napi-build-utils#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/inspiredware/napi-build-utils/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/inspiredware/napi-build-utils.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/node-abi@2.18.0",
		"name": "node-abi",
		"version": "2.18.0",
		"description": "Get the Node ABI for a given target and runtime, and vice versa.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "1f5486cfd7d38bd4f5392fa44a4ad4d9a0dffbf4"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/node-abi@2.18.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/lgeiger/node-abi#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/lgeiger/node-abi/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/lgeiger/node-abi.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/noop-logger@0.1.1",
		"name": "noop-logger",
		"version": "0.1.1",
		"description": "A logger that does exactly nothing.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "94a2b1633c4f1317553007d8966fd0e841b6a4c2"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/noop-logger@0.1.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/segmentio/noop-logger#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/segmentio/noop-logger/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/segmentio/noop-logger.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/pump@3.0.0",
		"name": "pump",
		"version": "3.0.0",
		"description": "pipe streams together and close all of them if one of them closes",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "b4a2116815bde2f4e1ea602354e8c75565107a64"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/pump@3.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/mafintosh/pump#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/mafintosh/pump/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/mafintosh/pump.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/simple-get@3.1.0",
		"name": "simple-get",
		"version": "3.1.0",
		"description": "Simplest way to make http get requests. Supports HTTPS, redirects, gzip/deflate, streams in < 100 lines.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "b45be062435e50d159540b576202ceec40b9c6b3"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/simple-get@3.1.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/feross/simple-get"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/feross/simple-get/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/feross/simple-get.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/decompress-response@4.2.1",
		"name": "decompress-response",
		"version": "4.2.1",
		"description": "Decompress a HTTP response if needed",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "414023cc7a302da25ce2ec82d0d5238ccafd8986"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/decompress-response@4.2.1",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/sindresorhus/decompress-response#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/sindresorhus/decompress-response/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/sindresorhus/decompress-response.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/mimic-response@2.1.0",
		"name": "mimic-response",
		"version": "2.1.0",
		"description": "Mimic a Node.js HTTP response stream",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "d13763d35f613d09ec37ebb30bac0469c0ee8f43"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/mimic-response@2.1.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/sindresorhus/mimic-response#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/sindresorhus/mimic-response/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/sindresorhus/mimic-response.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/simple-concat@1.0.0",
		"name": "simple-concat",
		"version": "1.0.0",
		"description": "Super-minimalist version of concat-stream. Less than 15 lines!",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "7344cbb8b6e26fb27d66b2fc86f9f6d5997521c6"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/simple-concat@1.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/feross/simple-concat"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/feross/simple-concat/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/feross/simple-concat.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/tar-fs@2.1.0",
		"name": "tar-fs",
		"version": "2.1.0",
		"description": "filesystem bindings for tar-stream",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "d1cdd121ab465ee0eb9ccde2d35049d3f3daf0d5"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/tar-fs@2.1.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/mafintosh/tar-fs"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/mafintosh/tar-fs/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/mafintosh/tar-fs.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/mkdirp-classic@0.5.3",
		"name": "mkdirp-classic",
		"version": "0.5.3",
		"description": "Mirror of mkdirp 0.5.2",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "fa10c9115cc6d8865be221ba47ee9bed78601113"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/mkdirp-classic@0.5.3",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/mafintosh/mkdirp-classic"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/mafintosh/mkdirp-classic/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/mafintosh/mkdirp-classic.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/tar-stream@2.1.3",
		"name": "tar-stream",
		"version": "2.1.3",
		"description": "tar-stream is a streaming tar parser and generator and nothing else. It is streams2 and operates purely using streams which means you can easily extract/parse tarballs without ever hitting the file system.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "1e2022559221b7866161660f118255e20fa79e41"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/tar-stream@2.1.3",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/mafintosh/tar-stream"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/mafintosh/tar-stream/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/mafintosh/tar-stream.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/bl@4.0.2",
		"name": "bl",
		"version": "4.0.2",
		"description": "Buffer List: collect buffers and access with a standard readable Buffer interface, streamable too!",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "52b71e9088515d0606d9dd9cc7aa48dc1f98e73a"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/bl@4.0.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/rvagg/bl"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/rvagg/bl/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/rvagg/bl.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/inherits@2.0.4",
		"name": "inherits",
		"version": "2.0.4",
		"description": "Browser-friendly inheritance fully compatible with standard node.js inherits()",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "0fa2c64f932917c3433a0ded55363aae37416b7c"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "ISC"
			}
		  }
		],
		"purl": "pkg:npm/inherits@2.0.4",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/isaacs/inherits#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/isaacs/inherits/issues"
		  },
		  {
			"type": "vcs",
			"url": "git://github.com/isaacs/inherits.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/which-pm-runs@1.0.0",
		"name": "which-pm-runs",
		"version": "1.0.0",
		"description": "Detects what package manager executes the process",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "670b3afbc552e0b55df6b7780ca74615f23ad1cb"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/which-pm-runs@1.0.0",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/zkochan/which-pm-runs#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/zkochan/which-pm-runs/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/zkochan/which-pm-runs.git"
		  }
		]
	  },
	  {
		"type": "library",
		"bom-ref": "pkg:npm/uglify-to-browserify@1.0.2",
		"name": "uglify-to-browserify",
		"version": "1.0.2",
		"description": "A transform to make UglifyJS work in browserify.",
		"hashes": [
		  {
			"alg": "SHA-1",
			"content": "6e0924d6bda6b5afe349e39a6d632850a0f882b7"
		  }
		],
		"licenses": [
		  {
			"license": {
			  "id": "MIT"
			}
		  }
		],
		"purl": "pkg:npm/uglify-to-browserify@1.0.2",
		"externalReferences": [
		  {
			"type": "website",
			"url": "https://github.com/ForbesLindesay/uglify-to-browserify#readme"
		  },
		  {
			"type": "issue-tracker",
			"url": "https://github.com/ForbesLindesay/uglify-to-browserify/issues"
		  },
		  {
			"type": "vcs",
			"url": "git+https://github.com/ForbesLindesay/uglify-to-browserify.git"
		  }
		]
	  }
	]
  }`
