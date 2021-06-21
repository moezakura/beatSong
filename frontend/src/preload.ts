import { Color, Titlebar } from 'custom-electron-titlebar';
import { contextBridge } from 'electron';
import ffi from 'ffi-napi';
import { GetSongListResponse } from '@/@types/global';

window.addEventListener('DOMContentLoaded', () => {
  // eslint-disable-next-line no-new
  new Titlebar({
    backgroundColor: Color.fromHex('#ECECEC')
  });
});

const goLibPath = 'golibs.dll';
const goLib = ffi.Library(goLibPath, {
  songList: ['string', ['string']],
  getImage: ['string', ['string']]
});

contextBridge.exposeInMainWorld('core', {
  getSongList: (): Promise<GetSongListResponse> => {
    return new Promise((resolve, reject) => {
      goLib.songList.async(
        'C:/Program Files (x86)/Steam/steamapps/common/Beat Saber/Beat Saber_Data/CustomLevels/',
        (err: any, res: string) => {
          if (err != null) {
            reject(err);
            return;
          }
          resolve(JSON.parse(res));
        }
      );
    });
  },
  getImage: (image: string): Promise<string> => {
    return new Promise<string>((resolve, reject) => {
      goLib.getImage.async(image, (err: any, res: string) => {
        if (err != null) {
          reject(err);
          return;
        }
        resolve(res);
      });
    });
  }
});
