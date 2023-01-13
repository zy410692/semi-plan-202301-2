---
title: "{{ replace .Name "-" " " | title }}"
subtitle: "{{ replace .TranslationBaseName "-" " " | title }}"
categories:
  - only-one
tags:
  - tag01
  - tag02
date: "{{ .Date }}"
lastmod: "{{ .Date }}"
toc: true
draft: false
hiddenFromHomePage: false
pinTop: false
---



# {{ replace .Name "-" " " | title }}
