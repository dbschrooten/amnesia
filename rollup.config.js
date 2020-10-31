import resolve from 'rollup-plugin-node-resolve';
import commonjs from 'rollup-plugin-commonjs';
import postcss from 'rollup-plugin-postcss';

export default {
  input: 'src/dashboard/webcomponents/main.js',
  output: [
    {
      dir: 'src/dashboard/static'
    }
  ],
  plugins: [
    resolve(),
    commonjs(),
    postcss({
      extract: {
        path: 'src/dashboard/scss/main.scss'
      },
      plugins: [
        require('tailwindcss'),
        require('autoprefixer')
      ]
    })
  ]
};
