import { Titlebar, Color } from 'custom-electron-titlebar';

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
