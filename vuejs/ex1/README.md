# VueJS Exercise 1: add a navbar

`index.html` in this directory is a bare-bones VueJS app: it just mounts a
single component and renders a greeting message. No build tooling, no
router, no components split into files, just Vue loaded from a CDN.

Open `index.html` in your browser (just double click it, or serve it with
anything like `npx serve .`) to confirm it works before changing anything.

## What you need to do

Add a navbar to the app with two links/buttons: **Menu** and **Home**.

That's the entire requirement on paper, but the point of this exercise is to
have fun with it, not just slap two `<a>` tags at the top and call it done.
Some ideas to get you going (you don't have to use any of these, and you're
encouraged to come up with your own):

- Style it so it actually looks like a real navbar (colors, spacing, a logo,
  a hover effect, whatever you like).
- Make it sticky/fixed to the top of the page.
- Add a mobile "hamburger" menu that toggles open/closed.
- Animate it (slide in, fade in, a burger icon that morphs into an X, etc.).
- Highlight the currently "active" link.
- Make "Menu" open a dropdown or a side drawer with more (fake) links.

Use plain Vue (`data`, methods, `v-if`/`v-show`, `@click`, etc.) and CSS. You
do not need a router or extra libraries, this is just meant to make you
comfortable with Vue's reactivity and template syntax while building
something visual.

## Submission

Once you're happy with it (or just proud of it), record a screenshot or a
short screen recording showing it off and post it in the
**Frontend/results-division-intro** channel on Discord.
