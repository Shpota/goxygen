<h1 align="center">
    <a href="https://github.com/Shpota/goxygen/tree/master/.github/README.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/gb.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/master/.github/README_zh.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/cn.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/master/.github/README_ua.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/ua.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/master/.github/README_ru.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/ru.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/master/.github/README_ko.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/kr.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/master/.github/README_pt-br.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/br.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/master/.github/README_by.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/by.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/master/.github/README_fr.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/fr.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/master/.github/README_es.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/es.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/master/.github/README_jp.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/jp.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/master/.github/README_id.md">
        <img height="20px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/id.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/tree/master/.github/README_he.md">
        <img height="25px" src="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/3.4.6/flags/4x3/il.svg">
    </a>
    <br>
    <div dir="rtl">
    Goxygen
    <a href="https://github.com/Shpota/goxygen/actions?query=workflow%3Abuild">
        <img src="https://github.com/Shpota/goxygen/workflows/build/badge.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/releases">
        <img src="https://img.shields.io/badge/version-v0.3.0-green">
    </a>
    <a href="https://gitter.im/goxygen/community">
        <img src="https://badges.gitter.im/goxygen/community.svg">
    </a>
    <a href="https://github.com/Shpota/goxygen/pulls">
        <img src="https://img.shields.io/badge/PRs-welcome-brightgreen.svg">
    </a>
    </div>
</h1>


<img src="../templates/react.webapp/src/logo.svg" align="left" width="230px" alt="goxygen logo">


<div dir="rtl">

**יצירה אוטומטית של פרויקט Web באמצעות Go ו-Angular, React או Vue.**

מטרת Goxygen היא לחסוך זמן במהלך הקמת פרויקט חדש.
הוא מייצר שלד עבור יישום עם כל הקונפיגורציה מוכנה עבורכם.
אתם יכולים ישר להתחיל לממש את ההיגיון העסקי שלכם.
Goxygen מג'נרט back-end בשפת Go, מחבר אותו עם הרכיבים של ה-front-end, מספק Dockerfile עבור היישום ומייצר קבצי docker-compose להפעלה נוחה בסביבות הפיתוח והייצור (production).
</div>

<div dir="rtl">
<table>
    <thead>
    <tr align="center">
        <td colspan=4><b>טכנולוגיות נתמכות</b></td>
    </tr>
    </thead>
    <tbody>
    <tr align="center">
        <td align="center">Front End</td>
        <td>Angular</td>
        <td>React</td>
        <td>Vue</td>
    </tr>
    <tr align="center">
        <td>Back End</td>
        <td colspan=3>Go</td>
    </tr>
    <tr align="center">
        <td>בסיס נתונים</td>
        <td>MongoDB</td>
        <td>MySQL</td>
        <td>PostgreSQL</td>
    </tr>
    </tbody>
</table>
</div>

<div dir="rtl">

## איך להשתמש
עליכם להתקין Go 1.11 ומעלה במחשב שלכם.
</div>

```go
go get -u github.com/shpota/goxygen
go run github.com/shpota/goxygen init my-app
```
<div dir="rtl">

פקודה זו תייצר פרויקט בתיקיית <span dir="ltr">`my-app`</span>.  
</div>

<div dir="rtl">

כברירת מחדל, הכלי ישתמש ב-React ו-MongoDB.  
אתם יכולים לבחור front-end framework ובסיס נתונים אחרים באמצעות הדגלים <span dir="ltr">`--frontend`</span> ו-<span dir="ltr">`--db`</span> בהתאמה.  
למשל פקודה זו תייצר פרויקט מבוסס Vue ו-PostgreSQL:
</div>

```go
go run github.com/shpota/goxygen init --frontend vue --db postgres my-app
```

<div dir="rtl">

דגל <span dir="ltr">`--frontend`</span> תומך באפשרויות <span dir="ltr">`angular`</span>, <span dir="ltr">`react`</span> ו-<span dir="ltr">`vue`</span>.
דגל <span dir="ltr">`--db`</span> תומך באפשרויות <span dir="ltr">`mongo`</span>, <span dir="ltr">`mysql`</span> ו-<span dir="ltr">`postgres`</span>.
</div>

<div dir="rtl">

הפרויקט שנוצר מוכן להפעלה עם <span dir="ltr">`docker-compose`</span>:  
</div>

```sh
cd my-app
docker-compose up
```

<div dir="rtl">

אחרי שה-build הסתיים, היישום נגיש ב-http://localhost:8080.
</div>

<div dir="rtl">

אתם יכולים למצוא פרטים נוספים על איך לעבוד עם הפרויקט שנוצר בקובץ ה-README שלו.  
</div>

![Showcase](showcase.gif)

<div dir="rtl">

## מבנה של פרויקט שג'ונרט (דוגמת React/MongoDB)
</div>


    my-app
    ├── server                   # Go קבצי פרויקט
    │   ├── db                   # MongoDB תקשורת
    │   ├── model                # domain אובייקטי
    │   ├── web                  # REST APIs, web שרת
    │   ├── server.go            # נקודת ההתחלה של השרת
    │   └── go.mod               # של השרת dependencies
    ├── webapp                    
    │   ├── public               # index.html-אייקונים, קבצים סטטיים ו
    │   ├── src                       
    │   │   ├── App.js           # React הרכיב המרכזי של
    │   │   ├── App.css          # App סגנונות ספציפיים עבור רכיב
    │   │   ├── index.js         # נקודת הכניסה של היישום          
    │   │   └── index.css        # סגנונות גלובליים
    │   ├── package.json         # front end-של ה dependencies
    │   ├── .env.development     # dev-לסביבת ה API-מחזיק בנקודת הקצה של ה
    │   └── .env.production      # prod-לסביבת ה API-מחזיק בנקודת הקצה של ה
    ├── Dockerfile               # ביחד front end-וה back end-של ה build
    ├── docker-compose.yml       # prod תיאור פריסת סביבת
    ├── docker-compose-dev.yml   # מקומי לצרכי פיתוח MongoDB מפעיל
    ├── init-db.js               # עם נתוני בדיקה MongoDB-יוצר אסופה ב
    ├── .dockerignore            # Docker builds מציין קבצים שיש להתעלם מהם בעת
    ├── .gitignore
    └── README.md                # שנוצר repo-מדריך כיצד להשתמש ב

<div dir="rtl">

קבצים כמו קבצי unit tests או רכיבי בסיס אינם כלולים כאן לצורך פשטות.
</div>

<div dir="rtl">

## תלויות - Dependencies
</div>

<div dir="rtl">

Goxygen מג'נרט שלד בסיסי של פרויקט ואינו מכריח אתכם להשתמש בקבוצת כלים מסוימת. לכן, הוא לא מוסיף תלויות בלתי נחוצות לפרויקט שלכם.
הוא משתמש רק ב-driver עבור בסיס נתונים ב-back end וב-[axios](https://github.com/axios/axios) בפרויקטים מבוססי React או Vue. בפרויקטים מבוססי Angular הוא משתמש רק בספריות ספציפיות של Angular.
</div>

<div dir="rtl">

## איך לתרום
</div>

<div dir="rtl">

אם מצאתם באג או שיש לכם רעיון כיצד לשפר את הפרויקט [open an issue](https://github.com/Shpota/goxygen/issues) ואנחנו נתקן אותו בהקדם האפשרי. אתם יכולים גם להציע שינויים באמצעות Pull Request. תעשו Fork ל-repository, בצעו שינויים, שלחו לנו בקשת pull ואנחנו נבדוק אותה בקרוב. יש לנו גם [Gitter chat](https://gitter.im/goxygen/community) בו אנו דנים בכל השינויים.
</div>

<div dir="rtl">

### קרדיטים
</div>

<div dir="rtl">

הלוגו של Goxygen נוצר על ידי [Egon Elbre](https://twitter.com/egonelbre).
</div>
