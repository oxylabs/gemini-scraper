# Gemini Scraper

[![Oxylabs promo code](https://github.com/oxylabs/gemini-scraper/blob/main/Gemini-scraper-github%201090x275.png)](https://oxylabs.io/products/scraper-api/serp/gemini/web?utm_content=oxylabs-web-scraper-api&groupid=877)

[![](https://dcbadge.limes.pink/api/server/Pds3gBmKMH?style=for-the-badge&theme=discord)](https://discord.gg/Pds3gBmKMH) [![YouTube](https://img.shields.io/badge/YouTube-Oxylabs-red?style=for-the-badge&logo=youtube&logoColor=white)](https://www.youtube.com/@oxylabs)


This repository is a technical, developer-focused reference for the [**Google Gemini**](https://oxylabs.io/products/scraper-api/serp/gemini) API target available through the Oxylabs [Web Scraper API](https://oxylabs.io/products/scraper-api/webs). It explains what the target does, how requests and responses are shaped, and how the Google AI Gemini target differs from other LLM and AI scraper targets.

The `gemini` source lets you submit a prompt to **Gemini Google** and receive a parsed, structured response – including plain text and Markdown answers, plus source citations. It is intended as a single reference point you can link from other repos, docs, and discussions when working with Gemini data.

## What is Google Gemini?

Google Gemini is Google's conversational AI assistant. Like other large language models, it takes a natural-language prompt and returns a generated answer, and it can reference external sources in its responses.

The Oxylabs Gemini target wraps this interaction behind a single API call. Instead of returning raw HTML that you would have to render and parse yourself, the `gemini` source can return a structured JSON object containing the prompt, the plain-text answer, the same answer in Markdown, and any source citations Gemini included. JavaScript rendering, proxy rotation, and server-side defense handling are managed by the API, so you provide a prompt and credentials rather than maintaining browsers and proxies. Check out our [Web Scraper API documentation](https://developers.oxylabs.io/products/web-scraper-api/features) for in-depth knowledge about each feature.

## How it works

Submit a prompt together with valid Web Scraper API credentials, then make a POST request to the `/v1/queries` endpoint. The example below uses the synchronous Realtime integration method.

> [!TIP]
> Get a free trial of Oxylabs **Web Scraper API** by registering on the [Dashboard](https://dashboard.oxylabs.io/en/registration).

### Request sample (Python)

```python
import requests
from pprint import pprint

# Structure payload.
payload = {
    'source': 'gemini',
    'prompt': 'best supplements for better sleep',
    'parse': True,
    'geo_location': 'United States'
}

# Get response.
response = requests.request(
    'POST',
    'https://realtime.oxylabs.io/v1/queries',
    auth=('USERNAME', 'PASSWORD'),
    json=payload,
)

# Print prettified response to stdout.
pprint(response.json())
```

The documentation also provides equivalent examples for cURL, Node.js, PHP, Go, C#, Java, and a plain URL/JSON payload. Find them in the [Gemini documentation](https://developers.oxylabs.io/api-targets/llms-and-ai/gemini).

The examples use the [Realtime](https://developers.oxylabs.io/products/web-scraper-api/integration-methods/realtime) integration method. You can also use the [Proxy Endpoint](https://developers.oxylabs.io/products/web-scraper-api/integration-methods/proxy-endpoint) or asynchronous [Push-Pull](https://developers.oxylabs.io/products/web-scraper-api/integration-methods/push-pull) methods.

## Request parameters

Basic setup and customization options for the `gemini` source:

| Parameter            | Description                                                                | Default value |
| -------------------- | -------------------------------------------------------------------------- | ------------- |
| `source` (mandatory) | Sets the scraper target. Use `gemini`.                                     | –             |
| `prompt` (mandatory) | The prompt or question to submit to Gemini. Must be under 8000 characters. | –             |
| `parse`              | Set to `true` to return structured JSON data.                              | `false`       |
| `geo_location`       | Specify a country to route the request from.                               | –             |
| `callback_url`       | URL to your callback endpoint (used with Push-Pull).                       | –             |

## Key Gemini features

The Gemini target exposes three capabilities worth calling out:

- **Prompts.** You pass your question or instruction through the `prompt` parameter. A single prompt must be under 8000 characters. This is the only required input beyond `source`.
- **`geo_location` targeting.** You can specify a country to route the request from, so responses reflect the location you care about rather than a default region. Note that Gemini's response composition may vary depending on whether the query was made from a desktop or mobile device.
- **Citations.** When Gemini references external sources, the parsed output includes a `citations` array. Each citation object contains a `title`, `url`, `text`, and `description`, so you can trace which sources informed the answer.

## Google Gemini vs ChatGPT: how the targets differ

Gemini and ChatGPT are two of several LLM and AI targets available through the Web Scraper API. Both are prompt-in, structured-data-out sources built on the same request pattern: the same `/v1/queries` endpoint, the same authentication, the same top-level `job` and `results[]` envelope, and the same `parse_status_code` convention (`12000` for a successful parse). Each target then defines its own `content` schema and its own set of parameters.

The table below is a side-by-side specification reference, not a ranking. Which target fits your use case depends on whose responses you want to observe.

| Aspect                     | Google Gemini (`gemini`)                                     | ChatGPT (`chatgpt`)                                              |
| -------------------------- | ------------------------------------------------------------ | ---------------------------------------------------------------- |
| `source` value             | `gemini`                                                     | `chatgpt`                                                        |
| Prompt parameter           | `prompt`, within 8000 characters                             | `prompt`, within 4000 characters                                 |
| Web search trigger         | –                                                            | `search` – triggers a web search for the prompt (default `true`) |
| `geo_location` targeting   | Supported – specify a country to route the request from      | Supported – specify a country to send the prompt from            |
| JavaScript rendering       | Enabled by default; no rendering parameter needed            | Enforced by default; no rendering parameter needed               |
| `parse`                    | Set to `true` for structured JSON                            | Set to `true` for structured JSON                                |
| `callback_url`             | Supported (Push-Pull)                                        | Supported (Push-Pull)                                            |
| Citations                  | `citations` array with `title`, `url`, `text`, `description` | `citations` array with URL and text                              |
| Batch requests             | Up to 5000 requests                                          | Up to 5000 requests                                              |
| Markdown output            | `markdown_text`                                              | `markdown_text`, plus `markdown_json`                            |
| Plain-text output          | `response_text`                                              | `response_text`                                                  |
| Model identifier in output | –                                                            | `llm_model`                                                      |
| Integration methods        | Realtime, Proxy Endpoint, Push-Pull                          | Realtime, Proxy Endpoint, Push-Pull                              |

Compared with other LLM and AI targets, the practical distinctions are in the per-target `content` schema (field names and the extra data each returns) and in target-specific limits and parameters. For more information on LLM scraping targets, refer to our [documentation](https://developers.oxylabs.io/api-targets/llms-and-ai).

## Output samples

With `parse` set to `true`, the API returns a structured JSON object containing the Gemini output. A trimmed snippet looks like this:

```json
{
  "results": [
    {
      "job_id": "7470031540181837825",
      "status_code": 200,
      "url": "https://gemini.google.com/app/ca6b069d5af5c809",
      "content": {
        "prompt": "best supplements for better sleep",
        "response_text": "If you are tossing and turning, you are definitely not alone. The supplement aisle is packed with options...",
        "markdown_text": "If you are tossing and turning, you are definitely not alone. The supplement aisle is packed with options...",
        "citations": [
          {
            "title": "Hartford HealthCare",
            "url": "https://hartfordhealthcare.org/about-us/news-press/news-detail?articleId=66180",
            "text": "Can These 3 Supplements Really Improve Your Sleep? | Hartford HealthCare | CT",
            "description": "- Valerian root. This herbal sleep option likely works by acting on GABA receptors in the brain..."
          }
        ],
        "parse_status_code": 12000
      }
    }
  ]
}
```

You can also work with the answer in Markdown (`markdown_text`) for easier integration into AI and data workflows.

## Output data dictionary

All LLM targets share the same top-level `job` and `results[]` envelope. The fields below are the Gemini-specific `results[].content` fields:

| Field               | Description                                                                                                                                    | Type    |
| ------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ------- |
| `prompt`            | The submitted prompt.                                                                                                                           | string  |
| `response_text`     | Plain-text response from Gemini.                                                                                                                | string  |
| `markdown_text`     | Complete response in Markdown.                                                                                                                  | string  |
| `citations`         | Source citations referenced in the response. Each object has `title`, `url`, `text`, `description`. Returned only when present in the response. | array   |
| `parse_status_code` | `12000` indicates success. Any other value means the parser failed to extract some or all structured fields.                                    | integer |

## Notes and limitations

- **JavaScript rendering** is enabled by default for all LLM sources, so you do not need to include rendering parameters in the payload.
- **Batch requests are not supported** for the `gemini` source.
- **Desktop vs mobile** the composition of the response may vary depending on whether the query was made from a desktop or mobile device.
- **The `citations` field is conditional** and is returned only when citations are present in Gemini's response.

## Related Oxylabs LLM and AI scrapers

Gemini is part of a family of LLM and AI targets. If you are researching AI scrapers more broadly, these related repositories use the same Web Scraper API patterns:

- [ChatGPT Scraper](https://github.com/oxylabs/chatgpt-scraper) – send prompts to ChatGPT and collect responses plus structured metadata.
- [Perplexity Scraper](https://github.com/oxylabs/perplexity-scraper) – send prompts to Perplexity and collect AI-generated answers and structured metadata.

## Learn more

For the full parameter list, all language examples, and additional details, see the official [Gemini documentation](https://developers.oxylabs.io/api-targets/llms-and-ai/gemini).

## Contact us

If you have questions or need support, reach out at <support@oxylabs.io>, or via live chat in the [Oxylabs Dashboard](https://dashboard.oxylabs.io/en/). For enterprise inquiries, contact your dedicated account manager.
