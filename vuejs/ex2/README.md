# VueJS Exercise 2: talk to a public API

Take the app from exercise 1 (or start fresh from the `index.html` provided
here, which is the same bare-bones app minus the navbar) and make it fetch
data from a public API when the page loads, then render the text you got
back somewhere on the page.

## What you need to do

1. Pick a public, no-auth-required testing API that returns some text.
   A few free ones that work well for this:
   - `https://catfact.ninja/fact` -> `{ "fact": "...", "length": 42 }`
   - `https://api.adviceslip.com/advice` -> `{ "slip": { "id": 1, "advice": "..." } }`
   - `https://official-joke-api.appspot.com/random_joke` -> `{ "setup": "...", "punchline": "..." }`
   - `https://jsonplaceholder.typicode.com/posts/1` -> `{ "title": "...", "body": "..." }`

   Any similar API is fine, it just needs to be public and return some text
   you can display.

2. In the `mounted()` hook, use `fetch()` (or `axios` if you'd rather pull it
   in from a CDN) to call the API, parse the JSON response, and store the
   relevant text in your component's `data`.

3. Render that text somewhere in the template, so when the page loads you
   see real content coming from the API instead of the placeholder
   "Loading..." message.

## Bonus (optional)

- Add a button that re-fetches a new fact/joke/advice on click, instead of
  only fetching once on load.
- Handle the loading and error states explicitly (e.g. show "Loading..."
  while the request is in flight, and a friendly error message if it
  fails).
- Keep the navbar from exercise 1 around and put the fetched text in the
  body below it.

## Submission

Once it's working, post a screenshot or short recording of it in action in
the **Frontend/results-division-intro** channel on Discord.
