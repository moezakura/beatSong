import { Titlebar, Color } from 'custom-electron-titlebar';
import ffi from 'ffi-napi';
import { contextBridge } from 'electron';

window.addEventListener('DOMContentLoaded', () => {
  // eslint-disable-next-line no-new
  new Titlebar({
    backgroundColor: Color.fromHex('#ECECEC'),
  });

  const replaceText = (selector: string, text: string) => {
    const element = document.getElementById(selector);
    if (element) element.innerText = text;
  };

  for (const type of ['chrome', 'node', 'electron']) {
    replaceText(`${type}-version`, process.versions[type] ?? 'dev');
  }
});

contextBridge.exposeInMainWorld('core', {
  getSongList: () => {
    const dllSamplePath = 'golibs.dll';
    const dllSample = ffi.Library(dllSamplePath, {
      songList: ['string', ['string']],
    });

    const result = dllSample.songList(
      'C:/Program Files (x86)/Steam/steamapps/common/Beat Saber/Beat Saber_Data/CustomLevels/'
    );
    return JSON.parse(result);
  },
});
