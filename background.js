browser = chrome || browser; // This extension can run in Mozilla and Chrome.

const SERVER = {
    address: 'http://127.0.0.1:2727',
};

const RUN_IN_TERMINAL = {
    id: 'shellrun.terminal',
    title: 'Run in Terminal',
};

browser.contextMenus.create(
    {
        id: RUN_IN_TERMINAL.id,
        title: RUN_IN_TERMINAL.title,
        contexts: ['selection'],
    },
    () => console.log('Shellrun is ready.')
);

browser.contextMenus.onClicked.addListener((info) => {
    if (info.menuItemId == RUN_IN_TERMINAL.id) {
        fetch(`${SERVER.address}/run?command=${info.selectionText}`).then((response) => response.json());
    }
});
