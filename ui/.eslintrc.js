module.exports = {
  root: true,
  env: {
    node: true,
  },
  extends: [
    'plugin:vue/essential',
    '@vue/airbnb',
    '@vue/typescript/recommended',
  ],
  parserOptions: {
    ecmaVersion: 2020,
  },
  rules: {
    'no-console': process.env.NODE_ENV === 'production' ? 'warn' : 'off',
    'no-debugger': process.env.NODE_ENV === 'production' ? 'warn' : 'off',
    'no-plusplus': 'off',
    'import/prefer-default-export': 'off',
    'computed-property-spacing': 'off',
    'array-bracket-spacing': 'off',
    'template-curly-spacing': 'off',
    'indent': ['error', 2, {
      'FunctionDeclaration': { 'parameters': 'first' },
      'FunctionExpression': { 'parameters': 'first' }
    }],
    'object-curly-newline': 'off',
    'arrow-parens': 'off',
    'lines-between-class-members': 'off',
    'class-methods-use-this': 'off',
  },
};
