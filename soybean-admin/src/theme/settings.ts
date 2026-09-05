/** Default theme settings */
export const themeSettings: App.Theme.ThemeSetting = {
  themeScheme: 'dark',
  grayscale: false,
  colourWeakness: false,
  recommendColor: false,
  themeColor: '#8b5cf6',
  otherColor: {
    info: '#38bdf8',
    success: '#34d399',
    warning: '#fbbf24',
    error: '#fb7185'
  },
  isInfoFollowPrimary: true,
  layout: {
    mode: 'vertical',
    scrollMode: 'content',
    reverseHorizontalMix: false
  },
  page: {
    animate: true,
    animateMode: 'fade-slide'
  },
  header: {
    height: 68,
    breadcrumb: {
      visible: true,
      showIcon: true
    }
  },
  tab: {
    visible: true,
    cache: true,
    height: 44,
    mode: 'button'
  },
  fixedHeaderAndTab: true,
  sider: {
    inverted: false,
    width: 252,
    collapsedWidth: 72,
    mixWidth: 96,
    mixCollapsedWidth: 72,
    mixChildMenuWidth: 220
  },
  footer: {
    visible: false,
    fixed: false,
    height: 48,
    right: true
  },
  watermark: {
    visible: false,
    text: 'SoybeanAdmin'
  },
  tokens: {
    light: {
      colors: {
        container: 'rgb(255, 255, 255)',
        layout: 'rgb(244, 247, 252)',
        inverted: 'rgb(10, 12, 24)',
        'base-text': 'rgb(18, 24, 38)'
      },
      boxShadow: {
        header: '0 12px 40px rgb(15, 23, 42, 0.08)',
        sider: '12px 0 40px rgb(15, 23, 42, 0.08)',
        tab: '0 8px 24px rgb(15, 23, 42, 0.06)'
      }
    },
    dark: {
      colors: {
        container: 'rgb(11, 15, 28)',
        layout: 'rgb(4, 7, 17)',
        inverted: 'rgb(4, 7, 17)',
        'base-text': 'rgb(226, 232, 240)'
      }
    }
  }
};

/**
 * Override theme settings
 *
 * If publish new version, use `overrideThemeSettings` to override certain theme settings
 */
export const overrideThemeSettings: Partial<App.Theme.ThemeSetting> = {
  themeScheme: 'dark',
  themeColor: '#8b5cf6',
  layout: themeSettings.layout,
  header: themeSettings.header,
  tab: themeSettings.tab,
  sider: themeSettings.sider,
  footer: themeSettings.footer,
  tokens: themeSettings.tokens
};
