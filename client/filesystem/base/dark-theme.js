"use strict";

const DarkThemeClassListener = "dark-theme-tag";
const DarkThemeClassTag = "dark-theme";
const DarkThemeStorageKey = "dark_theme";

function DarkThemeLoad() {
    if (localStorage.getItem(DarkThemeStorageKey) !== null) {
        DarkThemeOn();
    } else {
        DarkThemeOff();
    }
}

function DarkThemeToggle() {
    if (localStorage.getItem(DarkThemeStorageKey) !== null) {
        DarkThemeOff();
    } else {
        DarkThemeOn();
    }
}

function DarkThemeOn() {
    let listeners = document.getElementsByClassName(DarkThemeClassListener);
    let len = listeners.length;
    for (let i = 0; i < len; i++) {
        listeners[i].classList.add(DarkThemeClassTag);
    }
    localStorage.setItem(DarkThemeStorageKey, '');
}

function DarkThemeOff() {
    let listeners = document.getElementsByClassName(DarkThemeClassListener);
    let len = listeners.length;
    for (let i = 0; i < len; i++) {
        listeners[i].classList.remove(DarkThemeClassTag);
    }
    localStorage.removeItem(DarkThemeStorageKey);
}