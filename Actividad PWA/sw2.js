self.addEventListener("fetch", function(event) {
  if (event.request.url.includes("segato")) {
    event.respondWith(
      fetch("mia.jpg")
    );
  } else {
    event.respondWith(
      fetch(event.request.url)
      .then(data => {
        if(data.status === 404) {
          return fetch("2.jpg");
        }
        return data;
      })
   );
  }
});