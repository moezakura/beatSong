import { Titlebar, Color } from 'custom-electron-titlebar';
import { contextBridge, ipcRenderer } from 'electron';
import ffi from 'ffi-napi';

window.addEventListener('DOMContentLoaded', () => {
  // eslint-disable-next-line no-new
  new Titlebar({
    backgroundColor: Color.fromHex('#ECECEC'),
  });
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
