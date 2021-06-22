const cacheName = "app-" + "916fdea053322b4da9a40d66c08153f7dccba8c9";

self.addEventListener("install", event => {
  console.log("installing app worker 916fdea053322b4da9a40d66c08153f7dccba8c9");
  self.skipWaiting();

  event.waitUntil(
    caches.open(cacheName).then(cache => {
      return cache.addAll([
        "",
        "/gomoku",
        "/gomoku/app.css",
        "/gomoku/app.js",
        "/gomoku/manifest.webmanifest",
        "/gomoku/wasm_exec.js",
        "/gomoku/web/app.wasm",
        "static/img/Gomoku.png",
        
      ]);
    })
  );
});

self.addEventListener("activate", event => {
  event.waitUntil(
    caches.keys().then(keyList => {
      return Promise.all(
        keyList.map(key => {
          if (key !== cacheName) {
            return caches.delete(key);
          }
        })
      );
    })
  );
  console.log("app worker 916fdea053322b4da9a40d66c08153f7dccba8c9 is activated");
});

self.addEventListener("fetch", event => {
  event.respondWith(
    caches.match(event.request).then(response => {
      return response || fetch(event.request);
    })
  );
});
