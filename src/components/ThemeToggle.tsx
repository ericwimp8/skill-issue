import { useState } from 'react';

type Theme = 'light' | 'dark';

function currentTheme(): Theme {
  return document.documentElement.dataset.theme === 'dark' ? 'dark' : 'light';
}

export function ThemeToggle() {
  const [theme, setTheme] = useState<Theme>(currentTheme);
  const nextTheme = theme === 'dark' ? 'light' : 'dark';

  function toggleTheme() {
    document.documentElement.dataset.theme = nextTheme;
    localStorage.setItem('skill-issue-theme', nextTheme);
    setTheme(nextTheme);
  }

  return (
    <button
      className="icon-button"
      type="button"
      aria-label={`Switch to ${nextTheme} theme`}
      onClick={toggleTheme}
    >
      <span aria-hidden="true">{theme === 'dark' ? '☼' : '◐'}</span>
    </button>
  );
}
