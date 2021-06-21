<template>
  <div class='home'>
    <div v-for='s in songList' :key='s.dirPath'>
      <img :src='s.imageBase64'>
      <div>{{ s.name }}</div>
    </div>
  </div>
</template>

<script lang='ts'>
import { computed, defineComponent, reactive, ref } from 'vue';
import { SongInfo, ViewSongInfo } from '@/@types/global';

export default defineComponent({
  name: 'Home',
  components: {},
  setup() {
    let songs = ref<SongInfo[]>([]);
    const imageCache = reactive<Map<string, string>>(new Map<string, string>());

    const songList = computed<ViewSongInfo[]>(() => {
      return songs.value.filter(() => {
        return true;
      }).map(s => {
        return {
          ...s,
          imageBase64: imageCache.get(s.imagePath) ?? ''
        };
      });
    });

    const getSong = async () => {
      const res = await window.core.getSongList();
      console.log(res);
      songs.value = res.payload;

      for (const s of songs.value) {
        const image = await window.core.getImage(s.imagePath);
        if (!image) {
          console.log(image);
        }
        imageCache.set(s.imagePath, image);
      }
    };

    getSong();

    return {
      songList,

      getSong
    };
  }
});
</script>
