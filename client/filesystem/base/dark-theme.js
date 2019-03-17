"use strict";

const DarkThemeClassListener = "dark-theme-tag";
const DarkThemeClassTag = "dark-theme";
const DarkThemeStorageKey = "dark_theme"

function DarkThemeLoad() {
    if (DarkThemeIsOn()) {
        DarkThemeOn();
    } else {
        DarkThemeOff();
    }
}

function DarkThemeToggle() {
    if (DarkThemeIsOn()) {
        DarkThemeOff();
    } else {
        DarkThemeOn();
    }
}

function DarkThemeIsOn() {
    return localStorage.getItem(DarkThemeStorageKey) !== null;
}

function DarkThemeListeners() {
    return document.getElementsByClassName(DarkThemeClassListener);
}

function DarkThemeOn() {
    let listeners = DarkThemeListeners();
    let len = listeners.length;
    for (let i = 0; i < len; i++) {
        listeners[i].classList.add(DarkThemeClassTag);
    }
    localStorage.setItem(DarkThemeStorageKey, '');
}

function DarkThemeOff() {
    let listeners = DarkThemeListeners();
    let len = listeners.length;
    for (let i = 0; i < len; i++) {
        listeners[i].classList.remove(DarkThemeClassTag);
    }
    localStorage.removeItem(DarkThemeStorageKey);
}