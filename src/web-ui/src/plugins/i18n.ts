// localization

import { createI18n } from 'vue-i18n'

import en from '../locales/en.json'
import ja from '../locales/ja.json'
import ru from '../locales/ru.json'

// Type-define 'en-US' as the master schema for the resource
type MessageSchema = typeof en

export default createI18n<[MessageSchema], 'en' | 'ja' | 'ru'>({
  legacy: false,
  globalInjection: true,
  locale: import.meta.env.VITE_LOCALE,
  messages: {
      'en': en,
      'ja': ja,
      'ru': ru
  }
})
