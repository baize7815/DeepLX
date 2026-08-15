<!--
 * @Author: Vincent Young
 * @Date: 2022-10-18 07:32:29
 * @LastEditors: Vincent Yang
 * @LastEditTime: 2026-07-08 00:00:00
 * @FilePath: /DLX/README.md
 * @Telegram: https://t.me/missuo
 * 
 * Copyright © 2022 by Vincent, All Rights Reserved. 
-->

# DLX

[![GitHub Workflow][1]](https://github.com/OwO-Network/DLX/actions)
[![Go Version][2]](https://github.com/OwO-Network/DLX/blob/main/go.mod)
[![Go Report][3]](https://goreportcard.com/badge/github.com/OwO-Network/DLX)
[![GitHub License][4]](https://github.com/OwO-Network/DLX/blob/main/LICENSE)
[![Docker Pulls][5]](https://hub.docker.com/r/missuo/deeplx)
[![Releases][6]](https://github.com/OwO-Network/DLX/releases)

[1]: https://img.shields.io/github/actions/workflow/status/OwO-Network/DLX/release.yaml?logo=github
[2]: https://img.shields.io/github/go-mod/go-version/OwO-Network/DLX?logo=go
[3]: https://goreportcard.com/badge/github.com/OwO-Network/DLX
[4]: https://img.shields.io/github/license/OwO-Network/DLX
[5]: https://img.shields.io/docker/pulls/missuo/deeplx?logo=docker
[6]: https://img.shields.io/github/v/release/OwO-Network/DLX?logo=smartthings

> [!IMPORTANT]
> **Disclaimer:** DLX is an independent, open-source project. It is **not** an official DeepL product, and it is **not** affiliated with, endorsed by, or sponsored by DeepL SE in any way. "DeepL" is a registered trademark of DeepL SE. Any reference to DeepL in this repository is made solely to describe interoperability with the DeepL translation service.

## Why was this project renamed?

In July 2026, we received a trademark notice forwarded by GitHub Trust & Safety, submitted on behalf of DeepL SE. The notice stated that this project's former name, "DeepLX", contained the registered trademark "DeepL" and might cause confusion about whether the project is authorized or endorsed by DeepL SE.

It never was, and it never claimed to be. To resolve the matter and remove any possible confusion, we renamed the repository to **DLX** and removed DeepL branding from the project. To state it plainly one more time: **this project is not an official DeepL project and has no relationship with DeepL SE whatsoever.**

DLX is a self-hosted translation API server written in Go. It exposes a simple HTTP API on port `1188`.

## Usage

### Docker

Docker Hub and GHCR image names remain `deeplx` (only the GitHub repository was renamed to DLX):

```bash
docker run -d -p 1188:1188 ghcr.io/owo-network/deeplx:latest
# or: docker run -d -p 1188:1188 missuo/deeplx:latest
```

Or use the provided [`compose.yaml`](compose.yaml):

```bash
docker compose up -d
```

### Render (Free)

This repository includes a [`render.yaml`](render.yaml) Blueprint and a
lightweight `GET /healthz` endpoint. The health endpoint only reports whether
DLX is running; it does not send a request to the translation upstream.

1. Fork this repository or push it to a repository in your own GitHub account.
2. In the Render Dashboard, choose **New > Blueprint** and connect the repository.
3. Enter a strong random value for the requested `TOKEN` environment variable.
4. Deploy the `dlx-translation-api` Free web service.
5. Copy the public URL shown by Render.

Render provides the `PORT` environment variable automatically, and DLX already
binds to it. To keep a Free service from becoming idle, an external uptime
monitor can send a `GET` request every 5 minutes to:

```text
https://<your-render-service>.onrender.com/healthz
```

The expected response is HTTP 200 with `{"status":"ok"}`. Keeping a service
continuously active consumes Free instance hours, so check the current limits
in the [Render Free services documentation](https://render.com/docs/free).

### Binary

Download the binary for your platform from [Releases](https://github.com/OwO-Network/DLX/releases) and run it (artifact names remain `deeplx_*`):

```bash
./deeplx
```

### Translate

```bash
curl -X POST http://localhost:1188/translate \
  -H "Content-Type: application/json" \
  -d '{"text": "Hello, world!", "source_lang": "EN", "target_lang": "ZH"}'
```

When `TOKEN` is configured, pass it in the authorization header:

```bash
curl -X POST https://<your-render-service>.onrender.com/translate \
  -H "Authorization: Bearer <your-token>" \
  -H "Content-Type: application/json" \
  -d '{"text":"Hello, world!","source_lang":"EN","target_lang":"ZH"}'
```

### Website integration

Do not put the DLX `TOKEN` in browser JavaScript: every visitor could read and
reuse it. The recommended design is for the website frontend to call its own
backend endpoint, and for that backend to call DLX with the secret token.

Browser code:

```js
async function translate(text, targetLang = "ZH") {
  const response = await fetch("/api/translate", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ text, target_lang: targetLang }),
  });

  const result = await response.json();
  if (!response.ok) throw new Error(result.message || "Translation failed");
  return result.data;
}
```

Example Node.js backend proxy (Node.js 18 or newer):

```js
app.post("/api/translate", async (req, res) => {
  const upstream = await fetch(`${process.env.DLX_URL}/translate`, {
    method: "POST",
    headers: {
      "Authorization": `Bearer ${process.env.DLX_TOKEN}`,
      "Content-Type": "application/json",
    },
    body: JSON.stringify(req.body),
  });

  res.status(upstream.status).send(await upstream.text());
});
```

Configure `DLX_URL` and `DLX_TOKEN` only on the website server. Validate input
and add rate limiting on `/api/translate`. A single request is limited to 1500
characters; for article or page translation, split content into paragraphs and
translate them sequentially instead of sending the full HTML document.

## Discussion Group
[Telegram Group](https://t.me/+8KDGHKJCxEVkNzll)

## Acknowledgements

### Contributors

<a href="https://github.com/OwO-Network/DLX/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=OwO-Network/DLX&anon=0" />
</a>

## Activity
![Alt](https://repobeats.axiom.co/api/embed/5f473f85db27cb30028a2f3db7a560f3577a4860.svg "Repobeats analytics image")

## License
[MIT](LICENSE)
