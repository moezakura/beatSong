declare global {
  interface Window {
    core: CoreAPI;
  }
}

export interface CoreAPI {
  getSongList: () => Promise<GetSongListResponse>;
  getImage: (imagePath: string) => Promise<string>;
}

export interface GetSongListResponse {
  error: object | null;
  payload: SongInfo[];
  status: boolean;
}

export interface SongInfo {
  id: string;
  isValidID: boolean;
  difficulty: string[];
  dirPath: string;
  imagePath: string;
  name: string;
}

export interface ViewSongInfo extends SongInfo {
  imageBase64: string;
}
