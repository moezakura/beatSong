module.exports = {
  root: true,
  env: {
    browser: true,
    node: true,
  },
  extends: [
    'eslint:recommended',
    'prettier',
    'prettier/vue',
    'plugin:vue/essential',
    '@vue/typescript',
  ],
  plugins: ['prettier'],
  // add your custom rules here
  rules: {
    'no-extra-semi': 'warn',
    quotes: ['warn', 'single'],
  },
};
