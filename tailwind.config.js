/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./ui/html/*.{tmpl,html,js}"],
  theme: {
    extend: {
      fontFamily: {
        rubik: ["Rubik", "sans-serif"],
      },
    },
  },
  plugins: [],
};
