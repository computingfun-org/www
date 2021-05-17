"use strict";

const DarkThemeClass = "dark-theme";
const DarkThemeStorageKey = "dark_theme";

function DarkThemeLoad() {
    if (localStorage.getItem(DarkThemeStorageKey) !== null) {
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
    return document.body.classList.contains(DarkThemeClass);
}

function DarkThemeOn() {
    document.body.classList.add(DarkThemeClass);
    localStorage.setItem(DarkThemeStorageKey, '');
}

function DarkThemeOff() {
    document.body.classList.remove(DarkThemeClass);
    localStorage.removeItem(DarkThemeStorageKey);
}