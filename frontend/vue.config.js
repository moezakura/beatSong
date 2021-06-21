module.exports = {
  pluginOptions: {
    electronBuilder: {
      preload: 'src/preload.ts',
      nodeintegration: true,
      chainWebpackMainProcess: (config) => {
        config.module
          .rule('ts')
          .test(/\.ts$/)
          .use('ts-loader')
          .loader('ts-loader')
          .end();
      },
      externals: ['ffi-napi'],
      extraResources: [
        'src/preload.ts',
        'src/preload.js',
        ],
    },
  },
};
