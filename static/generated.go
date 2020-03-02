// THIS CODE IS GENERATED, DO NOT EDIT

// Package 'static' contains static assets such as
// source code, text files or images generated form
// the 'templates' folder in the root of the repository.
// If a change is made in templates regenerate this file
// by running 'transform/transformer.go'.
package static

func Sources() map[string]string {
	return map[string]string{
		".dockerignore": `webapp/node_modules/
webapp/build/
`,
		".gitignore": `*.exe
*.exe~
*.dll
*.so
*.dylib
*.test
*.out

webapp/node_modules
webapp/build
webapp/npm-debug.log*
`,
		"Dockerfile": `FROM node:12.16 AS JS_BUILD
COPY webapp /webapp
WORKDIR webapp
RUN npm install && npm run build

FROM golang:1.14.0-alpine AS GO_BUILD
RUN apk update && apk add build-base
COPY server /server
WORKDIR /server
RUN go build -o /go/bin/server

FROM alpine:3.11
COPY --from=JS_BUILD /webapp/build* ./webapp/
COPY --from=GO_BUILD /go/bin/server ./
CMD ./server
`,
		"README.md": `# project-name

## Environment setup

You need to have [Go](https://golang.org/),
[Node.js](https://nodejs.org/),
[Docker](https://www.docker.com/), and
[Docker Compose](https://docs.docker.com/compose/)
(comes pre-installed with Docker on Mac and Windows)
installed on your computer.

Verify the tools by running the following commands:

` + "`" + `` + "`" + `` + "`" + `sh
go version
npm --version
docker --version
docker-compose --version
` + "`" + `` + "`" + `` + "`" + `

If you are using Windows you will also need
[gcc](https://gcc.gnu.org/). It comes installed
on Mac and almost all Linux distributions.

## Start in development mode

In the project directory run the command (you might
need to prepend it with ` + "`" + `sudo` + "`" + ` depending on your setup):
` + "`" + `` + "`" + `` + "`" + `sh
docker-compose -f docker-compose-dev.yml up
` + "`" + `` + "`" + `` + "`" + `

This starts a local MongoDB on ` + "`" + `localhost:27017` + "`" + `.
The DB will be populated with test records from the
[init-db.js](init-db.js) file.

Navigate to the ` + "`" + `server` + "`" + ` folder and start the back end:

` + "`" + `` + "`" + `` + "`" + `sh
cd server
go run server.go
` + "`" + `` + "`" + `` + "`" + `
The back end will serve on http://localhost:8080.

Navigate to the ` + "`" + `webapp` + "`" + ` folder, install dependencies,
and start the front end development server by running:

` + "`" + `` + "`" + `` + "`" + `sh
cd webapp
npm install
npm start
` + "`" + `` + "`" + `` + "`" + `
The application will be available on http://localhost:3000.
 
## Start in production mode

Perform:
` + "`" + `` + "`" + `` + "`" + `sh
docker-compose up
` + "`" + `` + "`" + `` + "`" + `
This will build the application and start it together with
its database. Access the application on http://localhost:8080.
`,
		"angular.webapp/.editorconfig": `# Editor configuration, see https://editorconfig.org
root = true

[*]
charset = utf-8
indent_style = space
indent_size = 2
insert_final_newline = true
trim_trailing_whitespace = true

[*.md]
max_line_length = off
trim_trailing_whitespace = false
`,
		"angular.webapp/.gitignore": `# See http://help.github.com/ignore-files/ for more about ignoring files.

# compiled output
/build
/tmp
/out-tsc
# Only exists if Bazel was run
/bazel-out

# dependencies
/node_modules

# profiling files
chrome-profiler-events*.json
speed-measure-plugin*.json

# IDEs and editors
/.idea
.project
.classpath
.c9/
*.launch
.settings/
*.sublime-workspace

# IDE - VSCode
.vscode/*
!.vscode/settings.json
!.vscode/tasks.json
!.vscode/launch.json
!.vscode/extensions.json
.history/*

# misc
/.sass-cache
/connect.lock
/coverage
/libpeerconnection.log
npm-debug.log
yarn-error.log
testem.log
/typings

# System Files
.DS_Store
Thumbs.db
`,
		"angular.webapp/angular.json": `{
  "$schema": "./node_modules/@angular/cli/lib/config/schema.json",
  "version": 1,
  "newProjectRoot": "projects",
  "projects": {
    "project-name": {
      "projectType": "application",
      "schematics": {
        "@schematics/angular:component": {
          "style": "scss"
        }
      },
      "root": "",
      "sourceRoot": "src",
      "prefix": "app",
      "architect": {
        "build": {
          "builder": "@angular-devkit/build-angular:browser",
          "options": {
            "outputPath": "build",
            "index": "src/index.html",
            "main": "src/main.ts",
            "polyfills": "src/polyfills.ts",
            "tsConfig": "tsconfig.app.json",
            "aot": false,
            "assets": [
              "src/favicon.ico",
              "src/robot.txt",
              "src/assets"
            ],
            "styles": [
              "src/styles.scss"
            ],
            "scripts": []
          },
          "configurations": {
            "production": {
              "fileReplacements": [
                {
                  "replace": "src/environments/environment.ts",
                  "with": "src/environments/environment.prod.ts"
                }
              ],
              "optimization": true,
              "outputHashing": "all",
              "sourceMap": false,
              "extractCss": true,
              "namedChunks": false,
              "aot": true,
              "extractLicenses": true,
              "vendorChunk": false,
              "buildOptimizer": true,
              "budgets": [
                {
                  "type": "initial",
                  "maximumWarning": "2mb",
                  "maximumError": "5mb"
                },
                {
                  "type": "anyComponentStyle",
                  "maximumWarning": "6kb",
                  "maximumError": "10kb"
                }
              ]
            }
          }
        },
        "serve": {
          "builder": "@angular-devkit/build-angular:dev-server",
          "options": {
            "browserTarget": "project-name:build"
          },
          "configurations": {
            "production": {
              "browserTarget": "project-name:build:production"
            }
          }
        },
        "extract-i18n": {
          "builder": "@angular-devkit/build-angular:extract-i18n",
          "options": {
            "browserTarget": "project-name:build"
          }
        },
        "test": {
          "builder": "@angular-devkit/build-angular:karma",
          "options": {
            "main": "src/test.ts",
            "polyfills": "src/polyfills.ts",
            "tsConfig": "tsconfig.spec.json",
            "karmaConfig": "karma.conf.js",
            "assets": [
              "src/favicon.ico",
              "src/assets"
            ],
            "styles": [
              "src/styles.scss"
            ],
            "scripts": []
          }
        },
        "lint": {
          "builder": "@angular-devkit/build-angular:tslint",
          "options": {
            "tsConfig": [
              "tsconfig.app.json",
              "tsconfig.spec.json",
              "e2e/tsconfig.json"
            ],
            "exclude": [
              "**/node_modules/**"
            ]
          }
        },
        "e2e": {
          "builder": "@angular-devkit/build-angular:protractor",
          "options": {
            "protractorConfig": "e2e/protractor.conf.js",
            "devServerTarget": "project-name:serve"
          },
          "configurations": {
            "production": {
              "devServerTarget": "project-name:serve:production"
            }
          }
        }
      }
    }},
  "defaultProject": "project-name"
}
`,
		"angular.webapp/browserslist": `# This file is used by the build system to adjust CSS and JS output to support the specified browsers below.
# For additional information regarding the format and rule options, please see:
# https://github.com/browserslist/browserslist#queries

# You can see what browsers were selected by your queries by running:
#   npx browserslist

> 0.5%
last 2 versions
Firefox ESR
not dead
not IE 9-11 # For IE 9-11 support, remove 'not'.`,
		"angular.webapp/e2e/protractor.conf.js": `// @ts-check
// Protractor configuration file, see link for more information
// https://github.com/angular/protractor/blob/master/lib/config.ts

const { SpecReporter } = require('jasmine-spec-reporter');

/**
 * @type { import("protractor").Config }
 */
exports.config = {
  allScriptsTimeout: 11000,
  specs: [
    './src/**/*.e2e-spec.ts'
  ],
  capabilities: {
    browserName: 'chrome'
  },
  directConnect: true,
  baseUrl: 'http://localhost:4200/',
  framework: 'jasmine',
  jasmineNodeOpts: {
    showColors: true,
    defaultTimeoutInterval: 30000,
    print: function() {}
  },
  onPrepare() {
    require('ts-node').register({
      project: require('path').join(__dirname, './tsconfig.json')
    });
    jasmine.getEnv().addReporter(new SpecReporter({ spec: { displayStacktrace: true } }));
  }
};`,
		"angular.webapp/e2e/src/app.e2e-spec.ts": `import { AppPage } from './app.po';
import { browser, logging } from 'protractor';

describe('workspace-project App', () => {
  let page: AppPage;

  beforeEach(() => {
    page = new AppPage();
  });

  it('should display project name', () => {
    page.navigateTo();
    expect(page.getTitleText()).toEqual('project-name');
  });

  afterEach(async () => {
    // Assert that there are no errors emitted from the browser
    const logs = await browser.manage().logs().get(logging.Type.BROWSER);
    expect(logs).not.toContain(jasmine.objectContaining({
      level: logging.Level.SEVERE,
    } as logging.Entry));
  });
});
`,
		"angular.webapp/e2e/src/app.po.ts": `import { browser, by, element } from 'protractor';

export class AppPage {
  navigateTo() {
    return browser.get(browser.baseUrl) as Promise<any>;
  }

  getTitleText() {
    return element(by.css('app-root .title')).getText() as Promise<string>;
  }
}
`,
		"angular.webapp/e2e/tsconfig.json": `{
  "extends": "../tsconfig.json",
  "compilerOptions": {
    "outDir": "../out-tsc/e2e",
    "module": "commonjs",
    "target": "es5",
    "types": [
      "jasmine",
      "jasminewd2",
      "node"
    ]
  }
}
`,
		"angular.webapp/karma.conf.js": `// Karma configuration file, see link for more information
// https://karma-runner.github.io/1.0/config/configuration-file.html

module.exports = function (config) {
  config.set({
    basePath: '',
    frameworks: ['jasmine', '@angular-devkit/build-angular'],
    plugins: [
      require('karma-jasmine'),
      require('karma-chrome-launcher'),
      require('karma-jasmine-html-reporter'),
      require('karma-coverage-istanbul-reporter'),
      require('@angular-devkit/build-angular/plugins/karma')
    ],
    client: {
      clearContext: false // leave Jasmine Spec Runner output visible in browser
    },
    coverageIstanbulReporter: {
      dir: require('path').join(__dirname, './coverage/project-name'),
      reports: ['html', 'lcovonly', 'text-summary'],
      fixWebpackSourcePaths: true
    },
    reporters: ['progress', 'kjhtml'],
    port: 9876,
    colors: true,
    logLevel: config.LOG_INFO,
    autoWatch: true,
    browsers: ['Chrome'],
    singleRun: false,
    restartOnFileChange: true
  });
};
`,
		"angular.webapp/package.json": `{
  "name": "project-name",
  "version": "0.1.0",
  "scripts": {
    "ng": "ng",
    "start": "ng serve --port 3000",
    "build": "ng build --prod",
    "test": "ng test",
    "lint": "ng lint",
    "e2e": "ng e2e"
  },
  "private": true,
  "dependencies": {
    "@angular/animations": "~8.2.14",
    "@angular/common": "~8.2.14",
    "@angular/compiler": "~8.2.14",
    "@angular/core": "~8.2.14",
    "@angular/forms": "~8.2.14",
    "@angular/platform-browser": "~8.2.14",
    "@angular/platform-browser-dynamic": "~8.2.14",
    "@angular/router": "~8.2.14",
    "rxjs": "~6.4.0",
    "tslib": "~1.10.0",
    "zone.js": "~0.9.1"
  },
  "devDependencies": {
    "@angular-devkit/build-angular": "~0.803.23",
    "@angular/cli": "~8.3.23",
    "@angular/compiler-cli": "~8.2.14",
    "@angular/language-service": "~8.2.14",
    "@types/node": "~8.9.4",
    "@types/jasmine": "~3.3.8",
    "@types/jasminewd2": "~2.0.3",
    "codelyzer": "~5.0.0",
    "jasmine-core": "~3.4.0",
    "jasmine-spec-reporter": "~4.2.1",
    "karma": "~4.1.0",
    "karma-chrome-launcher": "~2.2.0",
    "karma-coverage-istanbul-reporter": "~2.0.1",
    "karma-jasmine": "~2.0.1",
    "karma-jasmine-html-reporter": "~1.4.0",
    "protractor": "~5.4.0",
    "ts-node": "~7.0.0",
    "tslint": "~5.15.0",
    "typescript": "~3.5.3"
  }
}
`,
		"angular.webapp/src/app/app.component.html": `<div class="app">
  <h2 class="title">{{title}}</h2>
  <div class="logo"><img src="assets/logo.svg" height="150px" alt="logo" /></div>
  <div>
    This project is generated with <b><a href="https://github.com/shpota/goxygen">goxygen</a></b>.
    <p>
      The following list of technologies comes from
      a REST API call to the Go-based back end. Find
      and change the corresponding code in
      <code>webapp/src/app/tech/tech.component.ts</code>
      and <code>server/web/app.go</code>.
    </p>
    <app-tech></app-tech>
  </div>
</div>
`,
		"angular.webapp/src/app/app.component.scss": `code {
    font-family: source-code-pro, Menlo, Monaco, Consolas, 'Courier New',
    monospace;
    background-color: #b3e6ff;
}

.title {
    text-align: center;
}

.logo {
    text-align: center;
}
`,
		"angular.webapp/src/app/app.component.spec.ts": `import { TestBed, async } from '@angular/core/testing';
import { AppComponent } from './app.component';
import { TechComponent } from './tech/tech.component';
import { HttpClientTestingModule } from '@angular/common/http/testing';

describe('AppComponent', () => {
  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [
        AppComponent,
        TechComponent
      ],
      imports: [
        HttpClientTestingModule
      ]
    }).compileComponents();
  }));

  it('should create the app', () => {
    const fixture = TestBed.createComponent(AppComponent);
    const app = fixture.debugElement.componentInstance;
    expect(app).toBeTruthy();
  });

  it(` + "`" + `should have as title 'project-name'` + "`" + `, () => {
    const fixture = TestBed.createComponent(AppComponent);
    const app = fixture.debugElement.componentInstance;
    expect(app.title).toEqual('project-name');
  });

  it('should render title', () => {
    const fixture = TestBed.createComponent(AppComponent);
    fixture.detectChanges();
    const compiled = fixture.debugElement.nativeElement;
    expect(compiled.querySelector('.title').textContent).toContain('project-name');
  });

  it('should render goxygen link', () => {
    const fixture = TestBed.createComponent(AppComponent);
    fixture.detectChanges();
    const compiled = fixture.debugElement.nativeElement;
    expect(compiled.querySelector('a').textContent).toContain('goxygen');
  });
});
`,
		"angular.webapp/src/app/app.component.ts": `import { Component } from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent {
  title = 'project-name';
}
`,
		"angular.webapp/src/app/app.module.ts": `import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppComponent } from './app.component';
import { TechComponent } from './tech/tech.component';
import { HttpClientModule } from '@angular/common/http';

@NgModule({
  declarations: [
    AppComponent,
    TechComponent
  ],
  imports: [
    BrowserModule,
    HttpClientModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
`,
		"angular.webapp/src/app/tech/tech.component.html": `<ul>
  <li *ngFor="let tech of technologies">
    <b>{{tech.name}}</b>: {{tech.details}}
  </li>
</ul>
`,
		"angular.webapp/src/app/tech/tech.component.ts": `import { Component, OnInit } from '@angular/core';
import { TechService } from './tech.service';
import { Technology } from './tech.model';

@Component({
  selector: 'app-tech',
  templateUrl: './tech.component.html'
})
export class TechComponent implements OnInit {

  technologies: Technology[] = [];

  constructor(private readonly techService: TechService) { }

  ngOnInit() {
    this.techService.getTechnologies().subscribe(value => {
      this.technologies = value;
    });
  }
}
`,
		"angular.webapp/src/app/tech/tech.model.ts": `export interface Technology {
  name: string;
  details: string;
}
`,
		"angular.webapp/src/app/tech/tech.service.ts": `import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { Technology } from './tech.model';
import { environment } from '../../environments/environment';

@Injectable({
  providedIn: 'root'
})
export class TechService {

  constructor(private readonly httpClient: HttpClient) { }

  getTechnologies(): Observable<Technology[]> {
    return this.httpClient.get<Technology[]>(` + "`" + `${environment.apiUrl}/api/technologies` + "`" + `);
  }
}
`,
		"angular.webapp/src/assets/logo.svg": `<?xml version="1.0" encoding="UTF-8" standalone="no"?>

<svg
        xmlns:dc="http://purl.org/dc/elements/1.1/"
        xmlns:cc="http://creativecommons.org/ns#"
        xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#"
        xmlns:svg="http://www.w3.org/2000/svg"
        xmlns="http://www.w3.org/2000/svg"
        xmlns:xlink="http://www.w3.org/1999/xlink"
        xmlns:sodipodi="http://sodipodi.sourceforge.net/DTD/sodipodi-0.dtd"
        xmlns:inkscape="http://www.inkscape.org/namespaces/inkscape"
        width="131.16136mm"
        height="95.620743mm"
        viewBox="0 0 464.74498 338.81365"
        id="svg2"
        version="1.1"
        inkscape:version="0.91 r13725"
        sodipodi:docname="surfing-js.svg"
        inkscape:export-filename="F:\dev\graphics\gophers\svg-rasterized\surfing-js.png"
        inkscape:export-xdpi="134.17"
        inkscape:export-ydpi="134.17">
    <defs
            id="defs4">
        <linearGradient
                inkscape:collect="always"
                id="linearGradient4289">
            <stop
                    style="stop-color:#00dce2;stop-opacity:1;"
                    offset="0"
                    id="stop4291" />
            <stop
                    style="stop-color:#97f0ff;stop-opacity:1"
                    offset="1"
                    id="stop4293" />
        </linearGradient>
        <linearGradient
                inkscape:collect="always"
                xlink:href="#linearGradient4289"
                id="linearGradient4295"
                x1="431.3338"
                y1="531.39313"
                x2="400.66592"
                y2="386.12817"
                gradientUnits="userSpaceOnUse"
                gradientTransform="translate(129.45303,193.73997)" />
    </defs>
    <sodipodi:namedview
            id="base"
            pagecolor="#ffffff"
            bordercolor="#666666"
            borderopacity="1.0"
            inkscape:pageopacity="0.0"
            inkscape:pageshadow="2"
            inkscape:zoom="0.50000001"
            inkscape:cx="-305.29686"
            inkscape:cy="443.22335"
            inkscape:document-units="px"
            inkscape:current-layer="g4764"
            showgrid="false"
            inkscape:window-width="1920"
            inkscape:window-height="1017"
            inkscape:window-x="1912"
            inkscape:window-y="-8"
            inkscape:window-maximized="1"
            showguides="false"
            fit-margin-top="5"
            fit-margin-left="5"
            fit-margin-right="5"
            fit-margin-bottom="5" />
    <metadata
            id="metadata7">
        <rdf:RDF>
            <cc:Work
                    rdf:about="">
                <dc:format>image/svg+xml</dc:format>
                <dc:type
                        rdf:resource="http://purl.org/dc/dcmitype/StillImage" />
                <dc:title></dc:title>
            </cc:Work>
        </rdf:RDF>
    </metadata>
    <g
            inkscape:groupmode="layer"
            id="layer6"
            inkscape:label="water"
            transform="translate(-111.7365,-176.02344)"
            style="display:inline">
        <path
                style="fill:url(#linearGradient4295);fill-opacity:1;fill-rule:evenodd;stroke:none;stroke-width:1px;stroke-linecap:butt;stroke-linejoin:miter;stroke-opacity:1"
                d="m 132.04163,433.46271 c 23.83971,-32.31562 103.54234,-15.62414 130.05635,0.6066 89.59503,44.01946 169.11278,-19.05037 186.40202,-33.70711 42.5191,-37.92107 60.76259,-37.33059 75.65481,-30.1241 4.45055,2.15366 10.79479,11.46502 14.40769,12.4991 98.10479,111.17073 -186.00633,131.62933 -317.79746,102.34326 -60.45211,-13.43341 -102.37596,-33.11121 -88.72341,-51.61775 z"
                id="path4287"
                inkscape:connector-curvature="0"
                sodipodi:nodetypes="sccscss" />
    </g>
    <g
            style="display:none"
            transform="translate(-111.7365,-176.02344)"
            inkscape:label="water yellow"
            id="g4764"
            inkscape:groupmode="layer">
        <path
                sodipodi:nodetypes="sccscss"
                inkscape:connector-curvature="0"
                id="path4766"
                d="m 132.04163,433.46271 c 23.83971,-32.31562 103.54234,-15.62414 130.05635,0.6066 89.59503,44.01946 169.11278,-19.05037 186.40202,-33.70711 42.5191,-37.92107 60.76259,-37.33059 75.65481,-30.1241 4.45055,2.15366 10.79479,11.46502 14.40769,12.4991 98.10479,111.17073 -186.00633,131.62933 -317.79746,102.34326 -60.45211,-13.43341 -102.37596,-33.11121 -88.72341,-51.61775 z"
                style="fill:#f3df49;fill-opacity:1;fill-rule:evenodd;stroke:none;stroke-width:1px;stroke-linecap:butt;stroke-linejoin:miter;stroke-opacity:1"
                inkscape:export-filename="F:\BB.png"
                inkscape:export-xdpi="67.080002"
                inkscape:export-ydpi="67.080002" />
    </g>
    <g
            inkscape:groupmode="layer"
            id="layer1"
            inkscape:label="JS"
            transform="translate(17.716534,17.716534)">
        <path
                style="display:inline;fill:#2e2e2c;fill-opacity:1"
                d="m 397.82357,203.03086 c -2.15833,0.003 -4.04434,0.32029 -5.59908,0.97122 -1.55476,0.65095 -2.77858,1.62984 -3.64506,2.85687 -0.86648,1.22704 -1.37588,2.69856 -1.52935,4.30587 -0.15349,1.60734 0.0486,3.34871 0.58179,5.12605 0.57889,1.92971 1.30687,3.54962 2.19183,4.91079 0.88495,1.36117 1.92582,2.45369 3.10744,3.34777 1.1816,0.89409 2.50309,1.58986 3.93146,2.16479 1.42836,0.57497 2.96253,1.02966 4.56608,1.44355 l 0.5531,0.14648 0.54835,0.14938 0.5438,0.15226 0.53943,0.15513 c 0.98305,0.29659 1.84327,0.58763 2.58417,0.91306 0.74091,0.32542 1.36474,0.67712 1.87085,1.08622 0.50609,0.40908 0.89446,0.87518 1.15261,1.45254 0.25814,0.57738 0.3863,1.2648 0.35803,2.15226 -0.0236,0.73958 -0.18511,1.45783 -0.4697,2.13139 -0.28461,0.67354 -0.69231,1.30244 -1.2083,1.86375 -0.51597,0.56132 -1.14021,1.05509 -1.85805,1.45869 -0.71783,0.40365 -1.52914,0.71711 -2.42057,0.91847 -1.06033,0.23937 -2.00599,0.29259 -2.85348,0.18258 -0.8475,-0.11037 -1.59686,-0.384 -2.26928,-0.78971 -0.67241,-0.40574 -1.26797,-0.9433 -1.81208,-1.5773 -0.54411,-0.63399 -1.0369,-1.3642 -1.50671,-2.15437 l -1.98495,1.8073 -2.03798,1.86352 -2.09039,1.91483 -2.14196,1.95887 c 0.65321,1.39463 1.47304,2.66169 2.46907,3.74115 0.99602,1.07948 2.16837,1.97097 3.52109,2.61786 1.35269,0.64687 2.8856,1.04859 4.59719,1.15901 1.71161,0.11038 3.60137,-0.0712 5.66759,-0.57092 2.1138,-0.51178 4.08903,-1.29641 5.86582,-2.33188 1.7768,-1.0355 3.35427,-2.32145 4.66312,-3.83355 1.30887,-1.51209 2.3483,-3.25065 3.05413,-5.18713 0.70586,-1.93648 1.09715,-4.13213 1.04675,-6.31823 -0.0468,-2.02974 -0.39254,-3.57535 -0.95727,-4.92971 -0.56472,-1.35433 -1.35013,-2.50497 -2.32537,-3.5646 -0.97522,-1.05963 -2.1226,-1.878 -3.4771,-2.66038 -1.35451,-0.78239 -2.90711,-1.48633 -4.75517,-2.0661 l -0.54674,-0.16732 -0.55374,-0.16013 -0.5599,-0.15682 -0.57027,-0.14008 c -0.98308,-0.24486 -1.85406,-0.48079 -2.61385,-0.73567 -0.75976,-0.25485 -1.4082,-0.52844 -1.95199,-0.84849 -0.54379,-0.32004 -0.98289,-0.6866 -1.33095,-1.12796 -0.34807,-0.44137 -0.60508,-0.95729 -0.79206,-1.57778 -0.15132,-0.5022 -0.19217,-0.98554 -0.12441,-1.42792 0.0677,-0.44234 0.24435,-0.84353 0.52553,-1.18313 0.28118,-0.33959 0.66708,-0.61742 1.15061,-0.81709 0.48353,-0.19967 1.06476,-0.32107 1.73278,-0.35417 0.6517,-0.032 1.26098,-0.004 1.84305,0.0895 0.58209,0.094 1.13698,0.25443 1.67692,0.48587 0.53993,0.23142 1.04863,0.5269 1.55151,0.89696 0.50288,0.37001 1.00095,0.81468 1.51335,1.346 l 1.16331,-1.30657 1.08059,-1.39481 1.12596,-1.44869 1.30228,-1.46491 c -0.99177,-1.03735 -1.863,-1.84161 -2.76967,-2.49747 -0.90665,-0.65588 -1.86899,-1.17004 -3.04093,-1.61412 -1.17194,-0.44408 -2.5291,-0.80759 -3.91164,-1.03264 -1.38253,-0.22502 -2.43022,-0.27847 -4.37161,-0.33004 -1e-4,0 1.89422,-0.003 0,-3e-4 z m -35.90052,4.17824 2.95712,8.9242 2.48436,8.91613 1.61482,9.05055 0.64779,9.27226 c 0.0956,1.37067 -0.12944,2.54398 -0.43868,3.521 -0.30923,0.97704 -0.76799,1.76798 -1.35219,2.40818 -0.58422,0.64021 -1.29441,1.1298 -2.11506,1.50615 -0.82064,0.37633 -1.75212,0.63954 -2.78055,0.82727 -1.07289,0.19601 -1.99147,0.19576 -2.80086,0.0437 -0.80936,-0.15194 -1.50963,-0.45558 -2.14514,-0.8682 -0.63548,-0.41261 -1.20617,-0.93423 -1.75535,-1.52351 -0.54917,-0.58926 -1.07689,-1.24617 -1.62546,-1.9304 l -2.48262,1.7634 -2.42689,1.73727 -2.35737,1.70957 -2.2764,1.6812 c -1.21926,0.90046 1.79287,2.46599 2.93076,3.51442 1.13789,1.04841 2.4418,1.95339 3.9335,2.65 1.4917,0.6966 3.17113,1.18495 5.06833,1.39489 1.8972,0.20996 4.01297,0.14132 6.38083,-0.28115 2.62207,-0.46777 5.04063,-1.25295 7.2081,-2.3517 2.16746,-1.09874 4.08565,-2.5115 5.70493,-4.23195 1.61926,-1.72043 2.94249,-3.74916 3.88298,-6.06916 0.94049,-2.32001 1.68304,-4.9681 1.51137,-7.80022 l -0.4731,-9.29203 -1.58977,-8.97986 -2.62278,-8.76892 -3.25911,-8.72796 -2.8918,0.47986 -2.94,0.47991 -2.98043,0.47627 z"
                id="j"
                inkscape:connector-curvature="0"
                sodipodi:nodetypes="cssssssscccccssssssssssscccccssssssssssscccccssssssssssscccccssssscccsssssssscccccsssssssscccccccccc"
                inkscape:export-filename="F:\BB.png"
                inkscape:export-xdpi="67.080002"
                inkscape:export-ydpi="67.080002" />
    </g>
    <g
            inkscape:groupmode="layer"
            id="layer5"
            inkscape:label="board back"
            transform="translate(-111.7365,-176.02344)">
        <path
                style="display:inline;fill:#388e3c;fill-opacity:1;fill-rule:evenodd;stroke:#000000;stroke-width:2;stroke-linecap:butt;stroke-linejoin:round;stroke-miterlimit:4;stroke-dasharray:none;stroke-opacity:1"
                d="m 198.7904,380.42903 c 0.25809,-2.63099 4.09727,-13.36941 31.37399,-0.9521 87.58706,39.87267 172.37344,28.76155 238.08524,1.51305 -35.44555,31.73992 -105.09012,70.96031 -151.49963,82.62222 -33.17285,-0.25002 -69.22045,-8.33225 -91.92216,-33.7906 -10.3767,-11.63672 -27.08255,-38.73851 -26.03744,-49.39257 z"
                id="path4159"
                inkscape:connector-curvature="0"
                sodipodi:nodetypes="ssccss"
                inkscape:export-filename="F:\BB.png"
                inkscape:export-xdpi="67.080002"
                inkscape:export-ydpi="67.080002" />
        <ellipse
                style="opacity:1;fill:#1b5e2f;fill-opacity:0.94117647;stroke:none;stroke-width:2;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:4;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1"
                id="path4284"
                cx="280.87637"
                cy="474.38382"
                rx="83.288368"
                ry="22.619684"
                transform="matrix(0.98520272,-0.17139312,0.13103475,0.99137778,0,0)"
                inkscape:export-filename="F:\BB.png"
                inkscape:export-xdpi="67.080002"
                inkscape:export-ydpi="67.080002" />
    </g>
    <g
            inkscape:groupmode="layer"
            id="layer3"
            inkscape:label="gopher"
            transform="translate(-111.7365,-176.02344)">
        <path
                style="display:inline;opacity:1;fill:#8ed4fe;fill-opacity:1;fill-rule:evenodd;stroke:#000000;stroke-width:2.36734128;stroke-linecap:round;stroke-linejoin:miter;stroke-miterlimit:4;stroke-dasharray:none;stroke-opacity:1"
                d="m 369.5092,223.97146 c 8.70159,-7.65557 16.28692,-10.65856 23.07967,-7.25757 4.97464,2.49072 3.8793,9.20637 0.18568,14.34944 -4.40361,6.13167 -8.48975,9.32099 -8.7434,9.31074"
                id="path4191"
                inkscape:connector-curvature="0"
                sodipodi:nodetypes="cssc"
                inkscape:export-filename="F:\BB.png"
                inkscape:export-xdpi="67.080002"
                inkscape:export-ydpi="67.080002" />
        <path
                style="display:inline;opacity:1;fill:#000000;fill-opacity:1;fill-rule:evenodd;stroke:#000000;stroke-width:0.39455688px;stroke-linecap:butt;stroke-linejoin:miter;stroke-opacity:1"
                d="m 377.17602,235.83381 c 0.0393,0.82372 4.24935,3.42187 5.45713,3.096 1.76689,-0.47671 8.64143,-9.05634 8.60055,-10.83044 -0.11459,-4.97101 -14.08317,7.20081 -14.05768,7.73444 z"
                id="path4201"
                inkscape:connector-curvature="0"
                sodipodi:nodetypes="ssss"
                inkscape:export-filename="F:\BB.png"
                inkscape:export-xdpi="67.080002"
                inkscape:export-ydpi="67.080002" />
        <path
                sodipodi:nodetypes="ccccc"
                inkscape:connector-curvature="0"
                id="path4251"
                d="m 388.09226,315.2492 c 1.49571,-1.40459 4.86627,-1.9416 7.21199,-3.03825 4.64492,-2.17612 7.61104,-5.53269 6.6251,-7.49725 -0.98527,-1.96513 -5.55031,-1.79401 -10.19585,0.38219 -2.24303,1.05203 -4.38565,1.92955 -6.40835,2.1379"
                style="display:inline;opacity:1;fill:#8ed4fd;fill-opacity:1;stroke:#000000;stroke-width:2;stroke-linecap:round;stroke-linejoin:bevel;stroke-miterlimit:4;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1"
                inkscape:export-filename="F:\BB.png"
                inkscape:export-xdpi="67.080002"
                inkscape:export-ydpi="67.080002" />
        <path
                style="fill:#8ed4fd;fill-opacity:1;fill-rule:evenodd;stroke:#000000;stroke-width:4;stroke-linecap:butt;stroke-linejoin:miter;stroke-miterlimit:4;stroke-dasharray:none;stroke-opacity:1"
                d="m 298.5481,433.97311 c -13.21919,-2.14324 -19.67913,-12.58085 -26.60671,-23.44515 -39.79944,-62.41617 -31.2814,-150.25148 -11.38378,-173.74419 28.95072,-34.18152 95.75894,-33.50993 125.69239,5.32843 23.50633,26.83283 -1.18688,59.97242 -0.60355,88.63909 C 427.86396,348.35484 422,405.1122 422,405.1122 c 0,0 -29.16421,-3.51472 -46.41421,0.48528 -13.46876,3.12319 -29.73182,10.11303 -48.81044,19.56333 -9.17029,4.54236 -16.53921,10.70729 -28.22725,8.8123 z"
                id="path4235"
                inkscape:connector-curvature="0"
                sodipodi:nodetypes="ssscccsss"
                inkscape:export-filename="F:\BB.png"
                inkscape:export-xdpi="67.080002"
                inkscape:export-ydpi="67.080002" />
        <path
                style="display:inline;opacity:1;fill:#8ed4fd;fill-opacity:1;stroke:#000000;stroke-width:4;stroke-linecap:round;stroke-linejoin:bevel;stroke-miterlimit:4;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1"
                d="m 250.54937,329.78787 c -2.01484,1.9022 -6.56468,2.63951 -9.72851,4.1289 -6.26491,2.95542 -10.25844,7.49896 -8.91992,10.14844 1.33762,2.65027 7.50185,2.40279 13.76758,-0.55273 3.02532,-1.42878 5.91567,-2.62152 8.64648,-2.91016"
                id="path4183"
                inkscape:connector-curvature="0"
                sodipodi:nodetypes="ccccc"
                inkscape:export-filename="F:\BB.png"
                inkscape:export-xdpi="67.080002"
                inkscape:export-ydpi="67.080002" />
        <path
                style="fill:none;fill-rule:evenodd;stroke:#000000;stroke-width:1px;stroke-linecap:butt;stroke-linejoin:miter;stroke-opacity:1"
                d="m 391.3836,334.29527 c -9.82295,-5.14566 -22.27386,-2.47488 -22.27386,-2.47488"
                id="path4253"
                inkscape:connector-curvature="0"
                sodipodi:nodetypes="cc"
                inkscape:export-filename="F:\BB.png"
                inkscape:export-xdpi="67.080002"
                inkscape:export-ydpi="67.080002" />
        <path
                sodipodi:nodetypes="csc"
                inkscape:connector-curvature="0"
                id="path4255"
                d="m 286.83033,422.58661 c 6.99565,10.91962 1.84084,23.66997 8.24051,25.19347 5.23572,1.24641 10.82844,-11.01538 16.59576,-20.23475"
                style="display:inline;opacity:1;fill:#8ed4fd;fill-opacity:1;stroke:#000000;stroke-width:4;stroke-linecap:round;stroke-linejoin:bevel;stroke-miterlimit:4;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1"
                inkscape:export-filename="F:\BB.png"
                inkscape:export-xdpi="67.080002"
                inkscape:export-ydpi="67.080002" />
        <circle
                style="opacity:1;fill:#ffffff;fill-opacity:1;stroke:#000000;stroke-width:3;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:4;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1"
                id="path4212"
                cx="312.05698"
                cy="259.30524"
                r="33.4146"
                inkscape:export-filename="F:\BB.png"
                inkscape:export-xdpi="67.080002"
                inkscape:export-ydpi="67.080002" />
        <ellipse
                cy="275.77502"
                cx="383.2565"
                id="circle4214"
                style="opacity:1;fill:#ffffff;fill-opacity:1;stroke:#000000;stroke-width:2.80603194;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:4;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1"
                rx="23.076626"
                ry="24.170307"
                inkscape:export-filename="F:\BB.png"
                inkscape:export-xdpi="67.080002"
                inkscape:export-ydpi="67.080002" />
        <g
                id="g4257"
                transform="translate(4.3284271,0.4571068)"
                inkscape:export-filename="F:\BB.png"
                inkscape:export-xdpi="67.080002"
                inkscape:export-ydpi="67.080002">
            <path
                    inkscape:export-ydpi="210.08769"
                    inkscape:export-xdpi="210.08769"
                    inkscape:export-filename="F:\C-Gopher.png"
                    sodipodi:nodetypes="cccscc"
                    inkscape:connector-curvature="0"
                    id="path4224"
                    d="m 335.72791,298.82062 c -1.10745,3.63148 -0.69833,12.2121 -0.69833,12.2121 l 10.47516,3.15242 c 0,0 4.85288,-5.71445 5.44452,-7.78161 0.59168,-2.06713 0.3377,-3.2066 0.3377,-3.2066 z"
                    style="display:inline;opacity:1;fill:#ffffff;fill-opacity:1;fill-rule:evenodd;stroke:#000000;stroke-width:2.56399131;stroke-linecap:butt;stroke-linejoin:miter;stroke-miterlimit:4;stroke-dasharray:none;stroke-opacity:1" />
            <path
                    inkscape:export-ydpi="210.08769"
                    inkscape:export-xdpi="210.08769"
                    inkscape:export-filename="F:\C-Gopher.png"
                    sodipodi:nodetypes="ssss"
                    inkscape:connector-curvature="0"
                    id="path4222"
                    d="m 327.71579,292.84858 c -7.61561,7.89636 4.59118,8.08453 14.16297,10.19598 7.57686,1.67137 16.83011,10.15944 19.09861,2.50827 3.92037,-13.22268 -23.49485,-22.83102 -33.26158,-12.70425 z"
                    style="display:inline;opacity:1;fill:#d3b78d;fill-opacity:1;fill-rule:evenodd;stroke:#000000;stroke-width:2.56399131;stroke-linecap:butt;stroke-linejoin:miter;stroke-miterlimit:4;stroke-dasharray:none;stroke-opacity:1" />
            <ellipse
                    inkscape:export-ydpi="210.08769"
                    inkscape:export-xdpi="210.08769"
                    inkscape:export-filename="F:\C-Gopher.png"
                    transform="matrix(0.9497031,0.31315174,0.31315174,-0.9497031,0,0)"
                    ry="6.4980431"
                    rx="11.350795"
                    cy="-167.76379"
                    cx="420.1701"
                    id="path4220"
                    style="display:inline;opacity:1;fill:#000000;fill-opacity:1;stroke:#000000;stroke-width:1.92299342;stroke-linecap:round;stroke-linejoin:bevel;stroke-miterlimit:4;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1" />
        </g>
        <circle
                style="opacity:1;fill:#000000;fill-opacity:1;stroke:#000000;stroke-width:2.28475904;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:4;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1"
                id="path4231"
                cx="324.33755"
                cy="267.0163"
                r="8.8534412"
                inkscape:export-filename="F:\BB.png"
                inkscape:export-xdpi="67.080002"
                inkscape:export-ydpi="67.080002" />
        <circle
                r="6.7808099"
                cy="283.75888"
                cx="392.33401"
                id="circle4233"
                style="opacity:1;fill:#000000;fill-opacity:1;stroke:#000000;stroke-width:1.74988639;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:4;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1"
                inkscape:export-filename="F:\BB.png"
                inkscape:export-xdpi="67.080002"
                inkscape:export-ydpi="67.080002" />
        <path
                style="display:inline;opacity:1;fill:#8ed4fe;fill-opacity:1;fill-rule:evenodd;stroke:#000000;stroke-width:2.91332936;stroke-linecap:round;stroke-linejoin:miter;stroke-miterlimit:4;stroke-dasharray:none;stroke-opacity:1"
                d="m 278.19769,228.79463 c -3.90156,-10.60895 -2.25892,-20.22855 2.5085,-28.16278 3.34565,-5.56805 7.80324,-6.18771 11.59575,-4.82156 6.3323,2.28107 9.1159,9.46446 12.68017,20.66394"
                id="path4189"
                inkscape:connector-curvature="0"
                sodipodi:nodetypes="cssc"
                inkscape:export-filename="F:\BB.png"
                inkscape:export-xdpi="67.080002"
                inkscape:export-ydpi="67.080002" />
        <path
                sodipodi:nodetypes="ssss"
                inkscape:connector-curvature="0"
                id="path4212-4"
                d="m 288.66496,223.71799 c 0.78826,0.63922 5.95007,-1.75896 7.31549,-2.47008 2.72013,-1.41667 0.13056,-13.61893 -2.11014,-15.23773 -4.9601,-3.58345 -5.71599,17.2937 -5.20535,17.70781 z"
                style="display:inline;opacity:1;fill:#000000;fill-opacity:1;fill-rule:evenodd;stroke:#000000;stroke-width:0.4855549px;stroke-linecap:butt;stroke-linejoin:miter;stroke-opacity:1"
                inkscape:export-filename="F:\BB.png"
                inkscape:export-xdpi="67.080002"
                inkscape:export-ydpi="67.080002" />
    </g>
    <g
            inkscape:groupmode="layer"
            id="layer2"
            inkscape:label="board"
            style="display:inline"
            transform="translate(-111.7365,-176.02344)">
        <path
                style="fill:#4caf50;fill-opacity:1;fill-rule:evenodd;stroke:#000000;stroke-width:2;stroke-linecap:butt;stroke-linejoin:round;stroke-miterlimit:4;stroke-dasharray:none;stroke-opacity:1"
                d="m 199.50538,378.71575 c -3.00355,5.85609 -4.3397,9.43062 -3.46266,28.28214 21.66777,47.76999 122.45022,97.48537 235.38585,8.5786 16.78571,-13.21429 43.01355,-41.15656 70.71429,-46.42857 20.43968,-3.89009 35.55187,12.63729 35.55187,12.63729 l -9.16997,-19.11611 c -10.61241,-21.21673 -55.7985,-16.39216 -82.11055,2.94329 -35.08701,25.78374 -70.29979,86.10442 -142.84278,83.53553 -79.71358,-2.82281 -108.8202,-63.39239 -104.06605,-70.43217 z"
                id="path4149"
                inkscape:connector-curvature="0"
                sodipodi:nodetypes="ccssccssc"
                inkscape:export-filename="F:\BB.png"
                inkscape:export-xdpi="67.080002"
                inkscape:export-ydpi="67.080002" />
    </g>
</svg>`,
		"angular.webapp/src/environments/environment.prod.ts": `export const environment = {
  production: true,
  apiUrl: ''
};
`,
		"angular.webapp/src/environments/environment.ts": `// This file can be replaced during build by using the ` + "`" + `fileReplacements` + "`" + ` array.
// ` + "`" + `ng build --prod` + "`" + ` replaces ` + "`" + `environment.ts` + "`" + ` with ` + "`" + `environment.prod.ts` + "`" + `.
// The list of file replacements can be found in ` + "`" + `angular.json` + "`" + `.

export const environment = {
  production: false,
  apiUrl: 'http://localhost:8080'
};

/*
 * For easier debugging in development mode, you can import the following file
 * to ignore zone related error stack frames such as ` + "`" + `zone.run` + "`" + `, ` + "`" + `zoneDelegate.invokeTask` + "`" + `.
 *
 * This import should be commented out in production mode because it will have a negative impact
 * on performance if an error is thrown.
 */
// import 'zone.js/dist/zone-error';  // Included with Angular CLI.
`,
		"angular.webapp/src/index.html": `<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <title>project-name</title>
  <base href="/">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <link rel="icon" type="image/x-icon" href="favicon.ico">
</head>
<body>
  <app-root></app-root>
</body>
</html>
`,
		"angular.webapp/src/main.ts": `import { enableProdMode } from '@angular/core';
import { platformBrowserDynamic } from '@angular/platform-browser-dynamic';

import { AppModule } from './app/app.module';
import { environment } from './environments/environment';

if (environment.production) {
  enableProdMode();
}

platformBrowserDynamic().bootstrapModule(AppModule)
  .catch(err => console.error(err));
`,
		"angular.webapp/src/polyfills.ts": `/**
 * This file includes polyfills needed by Angular and is loaded before the app.
 * You can add your own extra polyfills to this file.
 *
 * This file is divided into 2 sections:
 *   1. Browser polyfills. These are applied before loading ZoneJS and are sorted by browsers.
 *   2. Application imports. Files imported after ZoneJS that should be loaded before your main
 *      file.
 *
 * The current setup is for so-called "evergreen" browsers; the last versions of browsers that
 * automatically update themselves. This includes Safari >= 10, Chrome >= 55 (including Opera),
 * Edge >= 13 on the desktop, and iOS 10 and Chrome on mobile.
 *
 * Learn more in https://angular.io/guide/browser-support
 */

/***************************************************************************************************
 * BROWSER POLYFILLS
 */

/** IE10 and IE11 requires the following for NgClass support on SVG elements */
// import 'classlist.js';  // Run ` + "`" + `npm install --save classlist.js` + "`" + `.

/**
 * Web Animations ` + "`" + `@angular/platform-browser/animations` + "`" + `
 * Only required if AnimationBuilder is used within the application and using IE/Edge or Safari.
 * Standard animation support in Angular DOES NOT require any polyfills (as of Angular 6.0).
 */
// import 'web-animations-js';  // Run ` + "`" + `npm install --save web-animations-js` + "`" + `.

/**
 * By default, zone.js will patch all possible macroTask and DomEvents
 * user can disable parts of macroTask/DomEvents patch by setting following flags
 * because those flags need to be set before ` + "`" + `zone.js` + "`" + ` being loaded, and webpack
 * will put import in the top of bundle, so user need to create a separate file
 * in this directory (for example: zone-flags.ts), and put the following flags
 * into that file, and then add the following code before importing zone.js.
 * import './zone-flags.ts';
 *
 * The flags allowed in zone-flags.ts are listed here.
 *
 * The following flags will work for all browsers.
 *
 * (window as any).__Zone_disable_requestAnimationFrame = true; // disable patch requestAnimationFrame
 * (window as any).__Zone_disable_on_property = true; // disable patch onProperty such as onclick
 * (window as any).__zone_symbol__UNPATCHED_EVENTS = ['scroll', 'mousemove']; // disable patch specified eventNames
 *
 *  in IE/Edge developer tools, the addEventListener will also be wrapped by zone.js
 *  with the following flag, it will bypass ` + "`" + `zone.js` + "`" + ` patch for IE/Edge
 *
 *  (window as any).__Zone_enable_cross_context_check = true;
 *
 */

/***************************************************************************************************
 * Zone JS is required by default for Angular itself.
 */
import 'zone.js/dist/zone';  // Included with Angular CLI.


/***************************************************************************************************
 * APPLICATION IMPORTS
 */
`,
		"angular.webapp/src/robots.txt": `# https://www.robotstxt.org/robotstxt.html
User-agent: *
Disallow:
`,
		"angular.webapp/src/styles.scss": `body {
    margin: 0;
    font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', 'Roboto', 'Oxygen',
    'Ubuntu', 'Cantarell', 'Fira Sans', 'Droid Sans', 'Helvetica Neue',
    sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
}
body {
    margin-top: 5%;
    padding-right: 5%;
    padding-left: 5%;
    font-size: larger;
}

@media screen and (min-width: 800px) {
    body {
        padding-right: 15%;
        padding-left: 15%;
    }
}

@media screen and (min-width: 1600px) {
    body {
        padding-right: 30%;
        padding-left: 30%;
    }
}
`,
		"angular.webapp/src/test.ts": `// This file is required by karma.conf.js and loads recursively all the .spec and framework files

import 'zone.js/dist/zone-testing';
import { getTestBed } from '@angular/core/testing';
import {
  BrowserDynamicTestingModule,
  platformBrowserDynamicTesting
} from '@angular/platform-browser-dynamic/testing';

declare const require: any;

// First, initialize the Angular testing environment.
getTestBed().initTestEnvironment(
  BrowserDynamicTestingModule,
  platformBrowserDynamicTesting()
);
// Then we find all the tests.
const context = require.context('./', true, /\.spec\.ts$/);
// And load the modules.
context.keys().map(context);
`,
		"angular.webapp/tsconfig.app.json": `{
  "extends": "./tsconfig.json",
  "compilerOptions": {
    "outDir": "./out-tsc/app",
    "types": []
  },
  "files": [
    "src/main.ts",
    "src/polyfills.ts"
  ],
  "include": [
    "src/**/*.ts"
  ],
  "exclude": [
    "src/test.ts",
    "src/**/*.spec.ts"
  ]
}
`,
		"angular.webapp/tsconfig.json": `{
  "compileOnSave": false,
  "compilerOptions": {
    "baseUrl": "./",
    "outDir": "./dist/out-tsc",
    "sourceMap": true,
    "declaration": false,
    "downlevelIteration": true,
    "experimentalDecorators": true,
    "module": "esnext",
    "moduleResolution": "node",
    "importHelpers": true,
    "target": "es2015",
    "typeRoots": [
      "node_modules/@types"
    ],
    "lib": [
      "es2018",
      "dom"
    ]
  },
  "angularCompilerOptions": {
    "fullTemplateTypeCheck": true,
    "strictInjectionParameters": true
  }
}
`,
		"angular.webapp/tsconfig.spec.json": `{
  "extends": "./tsconfig.json",
  "compilerOptions": {
    "outDir": "./out-tsc/spec",
    "types": [
      "jasmine",
      "node"
    ]
  },
  "files": [
    "src/test.ts",
    "src/polyfills.ts"
  ],
  "include": [
    "src/**/*.spec.ts",
    "src/**/*.d.ts"
  ]
}
`,
		"angular.webapp/tslint.json": `{
  "extends": "tslint:recommended",
  "rules": {
    "array-type": false,
    "arrow-parens": false,
    "deprecation": {
      "severity": "warning"
    },
    "component-class-suffix": true,
    "contextual-lifecycle": true,
    "directive-class-suffix": true,
    "directive-selector": [
      true,
      "attribute",
      "app",
      "camelCase"
    ],
    "component-selector": [
      true,
      "element",
      "app",
      "kebab-case"
    ],
    "import-blacklist": [
      true,
      "rxjs/Rx"
    ],
    "interface-name": false,
    "max-classes-per-file": false,
    "max-line-length": [
      true,
      140
    ],
    "member-access": false,
    "member-ordering": [
      true,
      {
        "order": [
          "static-field",
          "instance-field",
          "static-method",
          "instance-method"
        ]
      }
    ],
    "no-consecutive-blank-lines": false,
    "no-console": [
      true,
      "debug",
      "info",
      "time",
      "timeEnd",
      "trace"
    ],
    "no-empty": false,
    "no-inferrable-types": [
      true,
      "ignore-params"
    ],
    "no-non-null-assertion": true,
    "no-redundant-jsdoc": true,
    "no-switch-case-fall-through": true,
    "no-var-requires": false,
    "object-literal-key-quotes": [
      true,
      "as-needed"
    ],
    "object-literal-sort-keys": false,
    "ordered-imports": false,
    "quotemark": [
      true,
      "single"
    ],
    "trailing-comma": false,
    "no-conflicting-lifecycle": true,
    "no-host-metadata-property": true,
    "no-input-rename": true,
    "no-inputs-metadata-property": true,
    "no-output-native": true,
    "no-output-on-prefix": true,
    "no-output-rename": true,
    "no-outputs-metadata-property": true,
    "template-banana-in-box": true,
    "template-no-negated-async": true,
    "use-lifecycle-interface": true,
    "use-pipe-transform-interface": true
  },
  "rulesDirectory": [
    "codelyzer"
  ]
}`,
		"docker-compose-dev.yml": `version: "3.7"
services:
  dev_db:
    image: mongo:4.2.3
    environment:
      MONGO_INITDB_DATABASE: tech
    ports:
      - 27017:27017
    volumes:
      - ./init-db.js:/docker-entrypoint-initdb.d/init.js
`,
		"docker-compose.yml": `version: "3.7"
services:
  app:
    build: .
    container_name: app
    ports:
      - 8080:8080
    depends_on:
      - db
    environment:
      profile: prod
  db:
    image: mongo:4.2.3
    container_name: db
    environment:
      MONGO_INITDB_DATABASE: tech
    volumes:
      - ./init-db.js:/docker-entrypoint-initdb.d/init.js
`,
		"init-db.js": `db.tech.insert({
    "name": "Go",
    "details": "An open source programming language that makes it easy to build simple and efficient software."
});
db.tech.insert({
    "name": "JavaScript",
    "details": "A lightweight, interpreted, or just-in-time compiled programming language with first-class functions."
});
db.tech.insert({
    "name": "MongoDB",
    "details": "A general purpose, document-based, distributed database."
});
`,
		"react.webapp/.env.development": `REACT_APP_API_URL=http://localhost:8080`,
		"react.webapp/.env.production":  `REACT_APP_API_URL=`,
		"react.webapp/package.json": `{
  "name": "project-name",
  "version": "0.1.0",
  "private": true,
  "dependencies": {
    "axios": "~0.19.2",
    "react": "~16.12.0",
    "react-dom": "~16.12.0",
    "react-scripts": "3.4.0"
  },
  "devDependencies": {
    "@testing-library/jest-dom": "~4.2.4",
    "@testing-library/react": "~9.4.0",
    "@testing-library/user-event": "~7.2.1"
  },
  "scripts": {
    "start": "react-scripts start",
    "build": "react-scripts build",
    "test": "react-scripts test",
    "eject": "react-scripts eject"
  },
  "eslintConfig": {
    "extends": "react-app"
  },
  "browserslist": {
    "production": [
      ">0.2%",
      "not dead",
      "not op_mini all"
    ],
    "development": [
      "last 1 chrome version",
      "last 1 firefox version",
      "last 1 safari version"
    ]
  }
}
`,
		"react.webapp/public/index.html": `<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8" />
    <link rel="icon" href="%PUBLIC_URL%/favicon.ico" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />
    <meta name="theme-color" content="#000000" />
    <meta
      name="description"
      content="Web site created using goxygen"
    />
    <link rel="apple-touch-icon" href="%PUBLIC_URL%/logo192.png" />
    <!--
      manifest.json provides metadata used when your web app is installed on a
      user's mobile device or desktop. See https://developers.google.com/web/fundamentals/web-app-manifest/
    -->
    <link rel="manifest" href="%PUBLIC_URL%/manifest.json" />
    <!--
      Notice the use of %PUBLIC_URL% in the tags above.
      It will be replaced with the URL of the ` + "`" + `public` + "`" + ` folder during the build.
      Only files inside the ` + "`" + `public` + "`" + ` folder can be referenced from the HTML.

      Unlike "/favicon.ico" or "favicon.ico", "%PUBLIC_URL%/favicon.ico" will
      work correctly both with client-side routing and a non-root public URL.
      Learn how to configure a non-root public URL by running ` + "`" + `npm run build` + "`" + `.
    -->
    <title>project-name</title>
  </head>
  <body>
    <noscript>You need to enable JavaScript to run this app.</noscript>
    <div id="root"></div>
    <!--
      This HTML file is a template.
      If you open it directly in the browser, you will see an empty page.

      You can add webfonts, meta tags, or analytics to this file.
      The build step will place the bundled scripts into the <body> tag.

      To begin the development, run ` + "`" + `npm start` + "`" + `.
      To create a production bundle, use ` + "`" + `npm run build` + "`" + `.
    -->
  </body>
</html>
`,
		"react.webapp/public/manifest.json": `{
  "short_name": "project-name",
  "name": "project-name",
  "icons": [
    {
      "src": "favicon.ico",
      "sizes": "64x64 32x32 24x24 16x16",
      "type": "image/x-icon"
    },
    {
      "src": "logo192.png",
      "type": "image/png",
      "sizes": "192x192"
    },
    {
      "src": "logo512.png",
      "type": "image/png",
      "sizes": "512x512"
    }
  ],
  "start_url": ".",
  "display": "standalone",
  "theme_color": "#000000",
  "background_color": "#ffffff"
}
`,
		"react.webapp/public/robots.txt": `# https://www.robotstxt.org/robotstxt.html
User-agent: *
Disallow:
`,
		"react.webapp/src/App.css": `body {
    margin-top: 5%;
    padding-right: 5%;
    padding-left: 5%;
    font-size: larger;
}

@media screen and (min-width: 800px) {
    body {
        padding-right: 15%;
        padding-left: 15%;
    }
}

@media screen and (min-width: 1600px) {
    body {
        padding-right: 30%;
        padding-left: 30%;
    }
}

code {
    font-family: source-code-pro, Menlo, Monaco, Consolas, 'Courier New',
    monospace;
    background-color: #b3e6ff;
}

.title {
    text-align: center;
}

.logo {
    text-align: center;
}
`,
		"react.webapp/src/App.js": `import React from 'react';
import './App.css';
import logo from './logo.svg';
import {Tech} from "./tech/Tech";

export function App() {
    return (
        <div className="app">
            <h2 className="title">project-name</h2>
            <div className="logo"><img src={logo} height="150px" alt="logo"/></div>
            <div>
                This project is generated with <b><a
                href="https://github.com/shpota/goxygen">goxygen</a></b>.
                <p/>The following list of technologies comes from
                a REST API call to the Go-based back end. Find
                and change the corresponding code
                in <code>webapp/src/tech/Tech.js
                </code> and <code>server/web/app.go</code>.
                <Tech/>
            </div>
        </div>
    );
}
`,
		"react.webapp/src/App.test.js": `import React from 'react';
import {render} from '@testing-library/react';
import {App} from './App';

test('renders learn react link', () => {
    const {getByText} = render(<App/>);
    const linkElement = getByText(/goxygen/i);
    expect(linkElement).toBeInTheDocument();
});
`,
		"react.webapp/src/index.css": `body {
    margin: 0;
    font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', 'Roboto', 'Oxygen',
    'Ubuntu', 'Cantarell', 'Fira Sans', 'Droid Sans', 'Helvetica Neue',
    sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
}
`,
		"react.webapp/src/index.js": `import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import {App} from './App';

ReactDOM.render(<App/>, document.getElementById('root'));
`,
		"react.webapp/src/logo.svg": `<?xml version="1.0" encoding="UTF-8" standalone="no"?>

<svg
        xmlns:dc="http://purl.org/dc/elements/1.1/"
        xmlns:cc="http://creativecommons.org/ns#"
        xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#"
        xmlns:svg="http://www.w3.org/2000/svg"
        xmlns="http://www.w3.org/2000/svg"
        xmlns:xlink="http://www.w3.org/1999/xlink"
        xmlns:sodipodi="http://sodipodi.sourceforge.net/DTD/sodipodi-0.dtd"
        xmlns:inkscape="http://www.inkscape.org/namespaces/inkscape"
        width="131.16136mm"
        height="95.620743mm"
        viewBox="0 0 464.74498 338.81365"
        id="svg2"
        version="1.1"
        inkscape:version="0.91 r13725"
        sodipodi:docname="surfing-js.svg"
        inkscape:export-filename="F:\dev\graphics\gophers\svg-rasterized\surfing-js.png"
        inkscape:export-xdpi="134.17"
        inkscape:export-ydpi="134.17">
    <defs
            id="defs4">
        <linearGradient
                inkscape:collect="always"
                id="linearGradient4289">
            <stop
                    style="stop-color:#00dce2;stop-opacity:1;"
                    offset="0"
                    id="stop4291" />
            <stop
                    style="stop-color:#97f0ff;stop-opacity:1"
                    offset="1"
                    id="stop4293" />
        </linearGradient>
        <linearGradient
                inkscape:collect="always"
                xlink:href="#linearGradient4289"
                id="linearGradient4295"
                x1="431.3338"
                y1="531.39313"
                x2="400.66592"
                y2="386.12817"
                gradientUnits="userSpaceOnUse"
                gradientTransform="translate(129.45303,193.73997)" />
    </defs>
    <sodipodi:namedview
            id="base"
            pagecolor="#ffffff"
            bordercolor="#666666"
            borderopacity="1.0"
            inkscape:pageopacity="0.0"
            inkscape:pageshadow="2"
            inkscape:zoom="0.50000001"
            inkscape:cx="-305.29686"
            inkscape:cy="443.22335"
            inkscape:document-units="px"
            inkscape:current-layer="g4764"
            showgrid="false"
            inkscape:window-width="1920"
            inkscape:window-height="1017"
            inkscape:window-x="1912"
            inkscape:window-y="-8"
            inkscape:window-maximized="1"
            showguides="false"
            fit-margin-top="5"
            fit-margin-left="5"
            fit-margin-right="5"
            fit-margin-bottom="5" />
    <metadata
            id="metadata7">
        <rdf:RDF>
            <cc:Work
                    rdf:about="">
                <dc:format>image/svg+xml</dc:format>
                <dc:type
                        rdf:resource="http://purl.org/dc/dcmitype/StillImage" />
                <dc:title></dc:title>
            </cc:Work>
        </rdf:RDF>
    </metadata>
    <g
            inkscape:groupmode="layer"
            id="layer6"
            inkscape:label="water"
            transform="translate(-111.7365,-176.02344)"
            style="display:inline">
        <path
                style="fill:url(#linearGradient4295);fill-opacity:1;fill-rule:evenodd;stroke:none;stroke-width:1px;stroke-linecap:butt;stroke-linejoin:miter;stroke-opacity:1"
                d="m 132.04163,433.46271 c 23.83971,-32.31562 103.54234,-15.62414 130.05635,0.6066 89.59503,44.01946 169.11278,-19.05037 186.40202,-33.70711 42.5191,-37.92107 60.76259,-37.33059 75.65481,-30.1241 4.45055,2.15366 10.79479,11.46502 14.40769,12.4991 98.10479,111.17073 -186.00633,131.62933 -317.79746,102.34326 -60.45211,-13.43341 -102.37596,-33.11121 -88.72341,-51.61775 z"
                id="path4287"
                inkscape:connector-curvature="0"
                sodipodi:nodetypes="sccscss" />
    </g>
    <g
            style="display:none"
            transform="translate(-111.7365,-176.02344)"
            inkscape:label="water yellow"
            id="g4764"
            inkscape:groupmode="layer">
        <path
                sodipodi:nodetypes="sccscss"
                inkscape:connector-curvature="0"
                id="path4766"
                d="m 132.04163,433.46271 c 23.83971,-32.31562 103.54234,-15.62414 130.05635,0.6066 89.59503,44.01946 169.11278,-19.05037 186.40202,-33.70711 42.5191,-37.92107 60.76259,-37.33059 75.65481,-30.1241 4.45055,2.15366 10.79479,11.46502 14.40769,12.4991 98.10479,111.17073 -186.00633,131.62933 -317.79746,102.34326 -60.45211,-13.43341 -102.37596,-33.11121 -88.72341,-51.61775 z"
                style="fill:#f3df49;fill-opacity:1;fill-rule:evenodd;stroke:none;stroke-width:1px;stroke-linecap:butt;stroke-linejoin:miter;stroke-opacity:1"
                inkscape:export-filename="F:\BB.png"
                inkscape:export-xdpi="67.080002"
                inkscape:export-ydpi="67.080002" />
    </g>
    <g
            inkscape:groupmode="layer"
            id="layer1"
            inkscape:label="JS"
            transform="translate(17.716534,17.716534)">
        <path
                style="display:inline;fill:#2e2e2c;fill-opacity:1"
                d="m 397.82357,203.03086 c -2.15833,0.003 -4.04434,0.32029 -5.59908,0.97122 -1.55476,0.65095 -2.77858,1.62984 -3.64506,2.85687 -0.86648,1.22704 -1.37588,2.69856 -1.52935,4.30587 -0.15349,1.60734 0.0486,3.34871 0.58179,5.12605 0.57889,1.92971 1.30687,3.54962 2.19183,4.91079 0.88495,1.36117 1.92582,2.45369 3.10744,3.34777 1.1816,0.89409 2.50309,1.58986 3.93146,2.16479 1.42836,0.57497 2.96253,1.02966 4.56608,1.44355 l 0.5531,0.14648 0.54835,0.14938 0.5438,0.15226 0.53943,0.15513 c 0.98305,0.29659 1.84327,0.58763 2.58417,0.91306 0.74091,0.32542 1.36474,0.67712 1.87085,1.08622 0.50609,0.40908 0.89446,0.87518 1.15261,1.45254 0.25814,0.57738 0.3863,1.2648 0.35803,2.15226 -0.0236,0.73958 -0.18511,1.45783 -0.4697,2.13139 -0.28461,0.67354 -0.69231,1.30244 -1.2083,1.86375 -0.51597,0.56132 -1.14021,1.05509 -1.85805,1.45869 -0.71783,0.40365 -1.52914,0.71711 -2.42057,0.91847 -1.06033,0.23937 -2.00599,0.29259 -2.85348,0.18258 -0.8475,-0.11037 -1.59686,-0.384 -2.26928,-0.78971 -0.67241,-0.40574 -1.26797,-0.9433 -1.81208,-1.5773 -0.54411,-0.63399 -1.0369,-1.3642 -1.50671,-2.15437 l -1.98495,1.8073 -2.03798,1.86352 -2.09039,1.91483 -2.14196,1.95887 c 0.65321,1.39463 1.47304,2.66169 2.46907,3.74115 0.99602,1.07948 2.16837,1.97097 3.52109,2.61786 1.35269,0.64687 2.8856,1.04859 4.59719,1.15901 1.71161,0.11038 3.60137,-0.0712 5.66759,-0.57092 2.1138,-0.51178 4.08903,-1.29641 5.86582,-2.33188 1.7768,-1.0355 3.35427,-2.32145 4.66312,-3.83355 1.30887,-1.51209 2.3483,-3.25065 3.05413,-5.18713 0.70586,-1.93648 1.09715,-4.13213 1.04675,-6.31823 -0.0468,-2.02974 -0.39254,-3.57535 -0.95727,-4.92971 -0.56472,-1.35433 -1.35013,-2.50497 -2.32537,-3.5646 -0.97522,-1.05963 -2.1226,-1.878 -3.4771,-2.66038 -1.35451,-0.78239 -2.90711,-1.48633 -4.75517,-2.0661 l -0.54674,-0.16732 -0.55374,-0.16013 -0.5599,-0.15682 -0.57027,-0.14008 c -0.98308,-0.24486 -1.85406,-0.48079 -2.61385,-0.73567 -0.75976,-0.25485 -1.4082,-0.52844 -1.95199,-0.84849 -0.54379,-0.32004 -0.98289,-0.6866 -1.33095,-1.12796 -0.34807,-0.44137 -0.60508,-0.95729 -0.79206,-1.57778 -0.15132,-0.5022 -0.19217,-0.98554 -0.12441,-1.42792 0.0677,-0.44234 0.24435,-0.84353 0.52553,-1.18313 0.28118,-0.33959 0.66708,-0.61742 1.15061,-0.81709 0.48353,-0.19967 1.06476,-0.32107 1.73278,-0.35417 0.6517,-0.032 1.26098,-0.004 1.84305,0.0895 0.58209,0.094 1.13698,0.25443 1.67692,0.48587 0.53993,0.23142 1.04863,0.5269 1.55151,0.89696 0.50288,0.37001 1.00095,0.81468 1.51335,1.346 l 1.16331,-1.30657 1.08059,-1.39481 1.12596,-1.44869 1.30228,-1.46491 c -0.99177,-1.03735 -1.863,-1.84161 -2.76967,-2.49747 -0.90665,-0.65588 -1.86899,-1.17004 -3.04093,-1.61412 -1.17194,-0.44408 -2.5291,-0.80759 -3.91164,-1.03264 -1.38253,-0.22502 -2.43022,-0.27847 -4.37161,-0.33004 -1e-4,0 1.89422,-0.003 0,-3e-4 z m -35.90052,4.17824 2.95712,8.9242 2.48436,8.91613 1.61482,9.05055 0.64779,9.27226 c 0.0956,1.37067 -0.12944,2.54398 -0.43868,3.521 -0.30923,0.97704 -0.76799,1.76798 -1.35219,2.40818 -0.58422,0.64021 -1.29441,1.1298 -2.11506,1.50615 -0.82064,0.37633 -1.75212,0.63954 -2.78055,0.82727 -1.07289,0.19601 -1.99147,0.19576 -2.80086,0.0437 -0.80936,-0.15194 -1.50963,-0.45558 -2.14514,-0.8682 -0.63548,-0.41261 -1.20617,-0.93423 -1.75535,-1.52351 -0.54917,-0.58926 -1.07689,-1.24617 -1.62546,-1.9304 l -2.48262,1.7634 -2.42689,1.73727 -2.35737,1.70957 -2.2764,1.6812 c -1.21926,0.90046 1.79287,2.46599 2.93076,3.51442 1.13789,1.04841 2.4418,1.95339 3.9335,2.65 1.4917,0.6966 3.17113,1.18495 5.06833,1.39489 1.8972,0.20996 4.01297,0.14132 6.38083,-0.28115 2.62207,-0.46777 5.04063,-1.25295 7.2081,-2.3517 2.16746,-1.09874 4.08565,-2.5115 5.70493,-4.23195 1.61926,-1.72043 2.94249,-3.74916 3.88298,-6.06916 0.94049,-2.32001 1.68304,-4.9681 1.51137,-7.80022 l -0.4731,-9.29203 -1.58977,-8.97986 -2.62278,-8.76892 -3.25911,-8.72796 -2.8918,0.47986 -2.94,0.47991 -2.98043,0.47627 z"
                id="j"
                inkscape:connector-curvature="0"
                sodipodi:nodetypes="cssssssscccccssssssssssscccccssssssssssscccccssssssssssscccccssssscccsssssssscccccsssssssscccccccccc"
                inkscape:export-filename="F:\BB.png"
                inkscape:export-xdpi="67.080002"
                inkscape:export-ydpi="67.080002" />
    </g>
    <g
            inkscape:groupmode="layer"
            id="layer5"
            inkscape:label="board back"
            transform="translate(-111.7365,-176.02344)">
        <path
                style="display:inline;fill:#388e3c;fill-opacity:1;fill-rule:evenodd;stroke:#000000;stroke-width:2;stroke-linecap:butt;stroke-linejoin:round;stroke-miterlimit:4;stroke-dasharray:none;stroke-opacity:1"
                d="m 198.7904,380.42903 c 0.25809,-2.63099 4.09727,-13.36941 31.37399,-0.9521 87.58706,39.87267 172.37344,28.76155 238.08524,1.51305 -35.44555,31.73992 -105.09012,70.96031 -151.49963,82.62222 -33.17285,-0.25002 -69.22045,-8.33225 -91.92216,-33.7906 -10.3767,-11.63672 -27.08255,-38.73851 -26.03744,-49.39257 z"
                id="path4159"
                inkscape:connector-curvature="0"
                sodipodi:nodetypes="ssccss"
                inkscape:export-filename="F:\BB.png"
                inkscape:export-xdpi="67.080002"
                inkscape:export-ydpi="67.080002" />
        <ellipse
                style="opacity:1;fill:#1b5e2f;fill-opacity:0.94117647;stroke:none;stroke-width:2;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:4;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1"
                id="path4284"
                cx="280.87637"
                cy="474.38382"
                rx="83.288368"
                ry="22.619684"
                transform="matrix(0.98520272,-0.17139312,0.13103475,0.99137778,0,0)"
                inkscape:export-filename="F:\BB.png"
                inkscape:export-xdpi="67.080002"
                inkscape:export-ydpi="67.080002" />
    </g>
    <g
            inkscape:groupmode="layer"
            id="layer3"
            inkscape:label="gopher"
            transform="translate(-111.7365,-176.02344)">
        <path
                style="display:inline;opacity:1;fill:#8ed4fe;fill-opacity:1;fill-rule:evenodd;stroke:#000000;stroke-width:2.36734128;stroke-linecap:round;stroke-linejoin:miter;stroke-miterlimit:4;stroke-dasharray:none;stroke-opacity:1"
                d="m 369.5092,223.97146 c 8.70159,-7.65557 16.28692,-10.65856 23.07967,-7.25757 4.97464,2.49072 3.8793,9.20637 0.18568,14.34944 -4.40361,6.13167 -8.48975,9.32099 -8.7434,9.31074"
                id="path4191"
                inkscape:connector-curvature="0"
                sodipodi:nodetypes="cssc"
                inkscape:export-filename="F:\BB.png"
                inkscape:export-xdpi="67.080002"
                inkscape:export-ydpi="67.080002" />
        <path
                style="display:inline;opacity:1;fill:#000000;fill-opacity:1;fill-rule:evenodd;stroke:#000000;stroke-width:0.39455688px;stroke-linecap:butt;stroke-linejoin:miter;stroke-opacity:1"
                d="m 377.17602,235.83381 c 0.0393,0.82372 4.24935,3.42187 5.45713,3.096 1.76689,-0.47671 8.64143,-9.05634 8.60055,-10.83044 -0.11459,-4.97101 -14.08317,7.20081 -14.05768,7.73444 z"
                id="path4201"
                inkscape:connector-curvature="0"
                sodipodi:nodetypes="ssss"
                inkscape:export-filename="F:\BB.png"
                inkscape:export-xdpi="67.080002"
                inkscape:export-ydpi="67.080002" />
        <path
                sodipodi:nodetypes="ccccc"
                inkscape:connector-curvature="0"
                id="path4251"
                d="m 388.09226,315.2492 c 1.49571,-1.40459 4.86627,-1.9416 7.21199,-3.03825 4.64492,-2.17612 7.61104,-5.53269 6.6251,-7.49725 -0.98527,-1.96513 -5.55031,-1.79401 -10.19585,0.38219 -2.24303,1.05203 -4.38565,1.92955 -6.40835,2.1379"
                style="display:inline;opacity:1;fill:#8ed4fd;fill-opacity:1;stroke:#000000;stroke-width:2;stroke-linecap:round;stroke-linejoin:bevel;stroke-miterlimit:4;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1"
                inkscape:export-filename="F:\BB.png"
                inkscape:export-xdpi="67.080002"
                inkscape:export-ydpi="67.080002" />
        <path
                style="fill:#8ed4fd;fill-opacity:1;fill-rule:evenodd;stroke:#000000;stroke-width:4;stroke-linecap:butt;stroke-linejoin:miter;stroke-miterlimit:4;stroke-dasharray:none;stroke-opacity:1"
                d="m 298.5481,433.97311 c -13.21919,-2.14324 -19.67913,-12.58085 -26.60671,-23.44515 -39.79944,-62.41617 -31.2814,-150.25148 -11.38378,-173.74419 28.95072,-34.18152 95.75894,-33.50993 125.69239,5.32843 23.50633,26.83283 -1.18688,59.97242 -0.60355,88.63909 C 427.86396,348.35484 422,405.1122 422,405.1122 c 0,0 -29.16421,-3.51472 -46.41421,0.48528 -13.46876,3.12319 -29.73182,10.11303 -48.81044,19.56333 -9.17029,4.54236 -16.53921,10.70729 -28.22725,8.8123 z"
                id="path4235"
                inkscape:connector-curvature="0"
                sodipodi:nodetypes="ssscccsss"
                inkscape:export-filename="F:\BB.png"
                inkscape:export-xdpi="67.080002"
                inkscape:export-ydpi="67.080002" />
        <path
                style="display:inline;opacity:1;fill:#8ed4fd;fill-opacity:1;stroke:#000000;stroke-width:4;stroke-linecap:round;stroke-linejoin:bevel;stroke-miterlimit:4;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1"
                d="m 250.54937,329.78787 c -2.01484,1.9022 -6.56468,2.63951 -9.72851,4.1289 -6.26491,2.95542 -10.25844,7.49896 -8.91992,10.14844 1.33762,2.65027 7.50185,2.40279 13.76758,-0.55273 3.02532,-1.42878 5.91567,-2.62152 8.64648,-2.91016"
                id="path4183"
                inkscape:connector-curvature="0"
                sodipodi:nodetypes="ccccc"
                inkscape:export-filename="F:\BB.png"
                inkscape:export-xdpi="67.080002"
                inkscape:export-ydpi="67.080002" />
        <path
                style="fill:none;fill-rule:evenodd;stroke:#000000;stroke-width:1px;stroke-linecap:butt;stroke-linejoin:miter;stroke-opacity:1"
                d="m 391.3836,334.29527 c -9.82295,-5.14566 -22.27386,-2.47488 -22.27386,-2.47488"
                id="path4253"
                inkscape:connector-curvature="0"
                sodipodi:nodetypes="cc"
                inkscape:export-filename="F:\BB.png"
                inkscape:export-xdpi="67.080002"
                inkscape:export-ydpi="67.080002" />
        <path
                sodipodi:nodetypes="csc"
                inkscape:connector-curvature="0"
                id="path4255"
                d="m 286.83033,422.58661 c 6.99565,10.91962 1.84084,23.66997 8.24051,25.19347 5.23572,1.24641 10.82844,-11.01538 16.59576,-20.23475"
                style="display:inline;opacity:1;fill:#8ed4fd;fill-opacity:1;stroke:#000000;stroke-width:4;stroke-linecap:round;stroke-linejoin:bevel;stroke-miterlimit:4;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1"
                inkscape:export-filename="F:\BB.png"
                inkscape:export-xdpi="67.080002"
                inkscape:export-ydpi="67.080002" />
        <circle
                style="opacity:1;fill:#ffffff;fill-opacity:1;stroke:#000000;stroke-width:3;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:4;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1"
                id="path4212"
                cx="312.05698"
                cy="259.30524"
                r="33.4146"
                inkscape:export-filename="F:\BB.png"
                inkscape:export-xdpi="67.080002"
                inkscape:export-ydpi="67.080002" />
        <ellipse
                cy="275.77502"
                cx="383.2565"
                id="circle4214"
                style="opacity:1;fill:#ffffff;fill-opacity:1;stroke:#000000;stroke-width:2.80603194;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:4;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1"
                rx="23.076626"
                ry="24.170307"
                inkscape:export-filename="F:\BB.png"
                inkscape:export-xdpi="67.080002"
                inkscape:export-ydpi="67.080002" />
        <g
                id="g4257"
                transform="translate(4.3284271,0.4571068)"
                inkscape:export-filename="F:\BB.png"
                inkscape:export-xdpi="67.080002"
                inkscape:export-ydpi="67.080002">
            <path
                    inkscape:export-ydpi="210.08769"
                    inkscape:export-xdpi="210.08769"
                    inkscape:export-filename="F:\C-Gopher.png"
                    sodipodi:nodetypes="cccscc"
                    inkscape:connector-curvature="0"
                    id="path4224"
                    d="m 335.72791,298.82062 c -1.10745,3.63148 -0.69833,12.2121 -0.69833,12.2121 l 10.47516,3.15242 c 0,0 4.85288,-5.71445 5.44452,-7.78161 0.59168,-2.06713 0.3377,-3.2066 0.3377,-3.2066 z"
                    style="display:inline;opacity:1;fill:#ffffff;fill-opacity:1;fill-rule:evenodd;stroke:#000000;stroke-width:2.56399131;stroke-linecap:butt;stroke-linejoin:miter;stroke-miterlimit:4;stroke-dasharray:none;stroke-opacity:1" />
            <path
                    inkscape:export-ydpi="210.08769"
                    inkscape:export-xdpi="210.08769"
                    inkscape:export-filename="F:\C-Gopher.png"
                    sodipodi:nodetypes="ssss"
                    inkscape:connector-curvature="0"
                    id="path4222"
                    d="m 327.71579,292.84858 c -7.61561,7.89636 4.59118,8.08453 14.16297,10.19598 7.57686,1.67137 16.83011,10.15944 19.09861,2.50827 3.92037,-13.22268 -23.49485,-22.83102 -33.26158,-12.70425 z"
                    style="display:inline;opacity:1;fill:#d3b78d;fill-opacity:1;fill-rule:evenodd;stroke:#000000;stroke-width:2.56399131;stroke-linecap:butt;stroke-linejoin:miter;stroke-miterlimit:4;stroke-dasharray:none;stroke-opacity:1" />
            <ellipse
                    inkscape:export-ydpi="210.08769"
                    inkscape:export-xdpi="210.08769"
                    inkscape:export-filename="F:\C-Gopher.png"
                    transform="matrix(0.9497031,0.31315174,0.31315174,-0.9497031,0,0)"
                    ry="6.4980431"
                    rx="11.350795"
                    cy="-167.76379"
                    cx="420.1701"
                    id="path4220"
                    style="display:inline;opacity:1;fill:#000000;fill-opacity:1;stroke:#000000;stroke-width:1.92299342;stroke-linecap:round;stroke-linejoin:bevel;stroke-miterlimit:4;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1" />
        </g>
        <circle
                style="opacity:1;fill:#000000;fill-opacity:1;stroke:#000000;stroke-width:2.28475904;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:4;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1"
                id="path4231"
                cx="324.33755"
                cy="267.0163"
                r="8.8534412"
                inkscape:export-filename="F:\BB.png"
                inkscape:export-xdpi="67.080002"
                inkscape:export-ydpi="67.080002" />
        <circle
                r="6.7808099"
                cy="283.75888"
                cx="392.33401"
                id="circle4233"
                style="opacity:1;fill:#000000;fill-opacity:1;stroke:#000000;stroke-width:1.74988639;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:4;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1"
                inkscape:export-filename="F:\BB.png"
                inkscape:export-xdpi="67.080002"
                inkscape:export-ydpi="67.080002" />
        <path
                style="display:inline;opacity:1;fill:#8ed4fe;fill-opacity:1;fill-rule:evenodd;stroke:#000000;stroke-width:2.91332936;stroke-linecap:round;stroke-linejoin:miter;stroke-miterlimit:4;stroke-dasharray:none;stroke-opacity:1"
                d="m 278.19769,228.79463 c -3.90156,-10.60895 -2.25892,-20.22855 2.5085,-28.16278 3.34565,-5.56805 7.80324,-6.18771 11.59575,-4.82156 6.3323,2.28107 9.1159,9.46446 12.68017,20.66394"
                id="path4189"
                inkscape:connector-curvature="0"
                sodipodi:nodetypes="cssc"
                inkscape:export-filename="F:\BB.png"
                inkscape:export-xdpi="67.080002"
                inkscape:export-ydpi="67.080002" />
        <path
                sodipodi:nodetypes="ssss"
                inkscape:connector-curvature="0"
                id="path4212-4"
                d="m 288.66496,223.71799 c 0.78826,0.63922 5.95007,-1.75896 7.31549,-2.47008 2.72013,-1.41667 0.13056,-13.61893 -2.11014,-15.23773 -4.9601,-3.58345 -5.71599,17.2937 -5.20535,17.70781 z"
                style="display:inline;opacity:1;fill:#000000;fill-opacity:1;fill-rule:evenodd;stroke:#000000;stroke-width:0.4855549px;stroke-linecap:butt;stroke-linejoin:miter;stroke-opacity:1"
                inkscape:export-filename="F:\BB.png"
                inkscape:export-xdpi="67.080002"
                inkscape:export-ydpi="67.080002" />
    </g>
    <g
            inkscape:groupmode="layer"
            id="layer2"
            inkscape:label="board"
            style="display:inline"
            transform="translate(-111.7365,-176.02344)">
        <path
                style="fill:#4caf50;fill-opacity:1;fill-rule:evenodd;stroke:#000000;stroke-width:2;stroke-linecap:butt;stroke-linejoin:round;stroke-miterlimit:4;stroke-dasharray:none;stroke-opacity:1"
                d="m 199.50538,378.71575 c -3.00355,5.85609 -4.3397,9.43062 -3.46266,28.28214 21.66777,47.76999 122.45022,97.48537 235.38585,8.5786 16.78571,-13.21429 43.01355,-41.15656 70.71429,-46.42857 20.43968,-3.89009 35.55187,12.63729 35.55187,12.63729 l -9.16997,-19.11611 c -10.61241,-21.21673 -55.7985,-16.39216 -82.11055,2.94329 -35.08701,25.78374 -70.29979,86.10442 -142.84278,83.53553 -79.71358,-2.82281 -108.8202,-63.39239 -104.06605,-70.43217 z"
                id="path4149"
                inkscape:connector-curvature="0"
                sodipodi:nodetypes="ccssccssc"
                inkscape:export-filename="F:\BB.png"
                inkscape:export-xdpi="67.080002"
                inkscape:export-ydpi="67.080002" />
    </g>
</svg>`,
		"react.webapp/src/setupTests.js": `// jest-dom adds custom jest matchers for asserting on DOM nodes.
// allows you to do things like:
// expect(element).toHaveTextContent(/react/i)
// learn more: https://github.com/testing-library/jest-dom
import '@testing-library/jest-dom/extend-expect';`,
		"react.webapp/src/tech/Tech.css": `.technologies {
    margin-top: 5px;
}
`,
		"react.webapp/src/tech/Tech.js": `import React, {Component} from "react";
import axios from "axios";
import "./Tech.css"

export class Tech extends Component {
    state = {
        technologies: []
    };

    componentDidMount() {
        axios.get(` + "`" + `${process.env.REACT_APP_API_URL}/api/technologies` + "`" + `)
            .then(resp => this.setState({
                technologies: resp.data
            }));
    }

    render() {
        const technologies = this.state.technologies.map((technology, i) =>
            <li key={i}>
                <b>{technology.name}</b>: {technology.details}
            </li>
        );
        return (
            <ul className="technologies">
                {technologies}
            </ul>
        );
    }
}
`,
		"server/db/db.go": `package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"project-name/model"
)

type DB interface {
	GetTechnologies() ([]*model.Technology, error)
}

type MongoDB struct {
	collection *mongo.Collection
}

func NewMongo(client *mongo.Client) DB {
	tech := client.Database("tech").Collection("tech")
	return MongoDB{collection: tech}
}

func (m MongoDB) GetTechnologies() ([]*model.Technology, error) {
	res, err := m.collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Println("Error while fetching technologies:", err.Error())
		return nil, err
	}
	var tech []*model.Technology
	err = res.All(context.TODO(), &tech)
	if err != nil {
		log.Println("Error while decoding technologies:", err.Error())
		return nil, err
	}
	return tech, nil
}
`,
		"server/go.mod": `module project-name

go 1.14

require go.mongodb.org/mongo-driver v1.3.0
`,
		"server/model/technology.go": `package model

type Technology struct {
	Name    string ` + "`" + `json:"name"` + "`" + `
	Details string ` + "`" + `json:"details"` + "`" + `
}
`,
		"server/server.go": `package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"project-name/db"
	"project-name/web"
)

func main() {
	client, err := mongo.Connect(context.TODO(), clientOptions())
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())
	mongoDB := db.NewMongo(client)
	// CORS is enabled only in prod profile
	cors := os.Getenv("profile") == "prod"
	app := web.NewApp(mongoDB, cors)
	err = app.Serve()
	log.Println("Error", err)
}

func clientOptions() *options.ClientOptions {
	host := "db"
	if os.Getenv("profile") != "prod" {
		host = "localhost"
	}
	return options.Client().ApplyURI(
		"mongodb://" + host + ":27017",
	)
}
`,
		"server/web/app.go": `package web

import (
	"encoding/json"
	"log"
	"net/http"
	"project-name/db"
)

type App struct {
	d        db.DB
	handlers map[string]http.HandlerFunc
}

func NewApp(d db.DB, cors bool) App {
	app := App{
		d:        d,
		handlers: make(map[string]http.HandlerFunc),
	}
	techHandler := app.GetTechnologies
	if !cors {
		techHandler = disableCors(techHandler)
	}
	app.handlers["/api/technologies"] = techHandler
	app.handlers["/"] = http.FileServer(http.Dir("/webapp")).ServeHTTP
	return app
}

func (a *App) Serve() error {
	for path, handler := range a.handlers {
		http.Handle(path, handler)
	}
	log.Println("Web server is available on port 8080")
	return http.ListenAndServe(":8080", nil)
}

func (a *App) GetTechnologies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	technologies, err := a.d.GetTechnologies()
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	err = json.NewEncoder(w).Encode(technologies)
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func sendErr(w http.ResponseWriter, code int, message string) {
	resp, _ := json.Marshal(map[string]string{"error": message})
	http.Error(w, string(resp), code)
}

// Needed in order to disable CORS for local development
func disableCors(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		h(w, r)
	}
}
`,
		"vue.webapp/.env.development": `NODE_ENV=development
VUE_APP_API_URL=http://localhost:8080`,
		"vue.webapp/.env.production": `NODE_ENV=production
VUE_APP_API_URL=`,
		"vue.webapp/babel.config.js": `module.exports = {
  presets: [
    '@vue/cli-plugin-babel/preset'
  ]
}
`,
		"vue.webapp/package.json": `{
  "name": "project-name",
  "version": "0.1.0",
  "private": true,
  "scripts": {
    "start": "vue-cli-service serve --port 3000",
    "build": "vue-cli-service build --dest build",
    "lint": "vue-cli-service lint"
  },
  "dependencies": {
    "axios": "~0.19.2",
    "core-js": "~3.6.4",
    "vue": "~2.6.11"
  },
  "devDependencies": {
    "@vue/cli-plugin-babel": "~4.2.0",
    "@vue/cli-plugin-eslint": "~4.2.0",
    "@vue/cli-service": "~4.2.0",
    "babel-eslint": "~10.0.3",
    "eslint": "~6.7.2",
    "eslint-plugin-vue": "~6.1.2",
    "vue-template-compiler": "~2.6.11"
  },
  "eslintConfig": {
    "root": true,
    "env": {
      "node": true
    },
    "extends": [
      "plugin:vue/essential",
      "eslint:recommended"
    ],
    "parserOptions": {
      "parser": "babel-eslint"
    },
    "rules": {}
  },
  "browserslist": [
    "> 1%",
    "last 2 versions"
  ]
}
`,
		"vue.webapp/public/index.html": `<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width,initial-scale=1.0">
    <link rel="icon" href="<%= BASE_URL %>favicon.ico">
    <title><%= htmlWebpackPlugin.options.title %></title>
  </head>
  <body>
    <noscript>
      <strong>We're sorry but <%= htmlWebpackPlugin.options.title %> doesn't work properly without JavaScript enabled. Please enable it to continue.</strong>
    </noscript>
    <div id="app"></div>
  </body>
</html>
`,
		"vue.webapp/src/App.vue": `<template>
  <div id="app">
    <h2 class="title">project-name</h2>
    <div class="logo">
      <img :src="logoSVG" height="150px" alt="logo" />
    </div>
    <div>
      This project is generated with
      <b>
        <a href="https://github.com/shpota/goxygen">goxygen</a>
      </b>.
      <p />The following list of technologies comes from
      a REST API call to the Go-based back end. Find
      and change the corresponding code in
      <code>webapp/src/components/Tech.vue</code>
      and <code>server/web/app.go</code>.
      <Tech />
    </div>
  </div>
</template>

<script>
import Tech from './components/Tech.vue'

export default {
  name: 'App',
  components: {
    Tech
  },
  data() {
    return {
      logoSVG: require('./assets/logo.svg')
    }
  }
};
</script>

<style>
body {
  margin-top: 5%;
  padding-right: 5%;
  padding-left: 5%;
  font-size: larger;
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', 'Roboto', 'Oxygen',
    'Ubuntu', 'Cantarell', 'Fira Sans', 'Droid Sans', 'Helvetica Neue',
    sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}

@media screen and (min-width: 800px) {
  body {
    padding-right: 15%;
    padding-left: 15%;
  }
}

@media screen and (min-width: 1600px) {
  body {
    padding-right: 30%;
    padding-left: 30%;
  }
}

code {
  font-family: source-code-pro, Menlo, Monaco, Consolas, "Courier New",
    monospace;
  background-color: #b3e6ff;
}

.title {
  text-align: center;
}

.logo {
  text-align: center;
}
</style>
`,
		"vue.webapp/src/assets/logo.svg": `<?xml version="1.0" encoding="UTF-8" standalone="no"?>

<svg
        xmlns:dc="http://purl.org/dc/elements/1.1/"
        xmlns:cc="http://creativecommons.org/ns#"
        xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#"
        xmlns:svg="http://www.w3.org/2000/svg"
        xmlns="http://www.w3.org/2000/svg"
        xmlns:xlink="http://www.w3.org/1999/xlink"
        xmlns:sodipodi="http://sodipodi.sourceforge.net/DTD/sodipodi-0.dtd"
        xmlns:inkscape="http://www.inkscape.org/namespaces/inkscape"
        width="131.16136mm"
        height="95.620743mm"
        viewBox="0 0 464.74498 338.81365"
        id="svg2"
        version="1.1"
        inkscape:version="0.91 r13725"
        sodipodi:docname="surfing-js.svg"
        inkscape:export-filename="F:\dev\graphics\gophers\svg-rasterized\surfing-js.png"
        inkscape:export-xdpi="134.17"
        inkscape:export-ydpi="134.17">
    <defs
            id="defs4">
        <linearGradient
                inkscape:collect="always"
                id="linearGradient4289">
            <stop
                    style="stop-color:#00dce2;stop-opacity:1;"
                    offset="0"
                    id="stop4291" />
            <stop
                    style="stop-color:#97f0ff;stop-opacity:1"
                    offset="1"
                    id="stop4293" />
        </linearGradient>
        <linearGradient
                inkscape:collect="always"
                xlink:href="#linearGradient4289"
                id="linearGradient4295"
                x1="431.3338"
                y1="531.39313"
                x2="400.66592"
                y2="386.12817"
                gradientUnits="userSpaceOnUse"
                gradientTransform="translate(129.45303,193.73997)" />
    </defs>
    <sodipodi:namedview
            id="base"
            pagecolor="#ffffff"
            bordercolor="#666666"
            borderopacity="1.0"
            inkscape:pageopacity="0.0"
            inkscape:pageshadow="2"
            inkscape:zoom="0.50000001"
            inkscape:cx="-305.29686"
            inkscape:cy="443.22335"
            inkscape:document-units="px"
            inkscape:current-layer="g4764"
            showgrid="false"
            inkscape:window-width="1920"
            inkscape:window-height="1017"
            inkscape:window-x="1912"
            inkscape:window-y="-8"
            inkscape:window-maximized="1"
            showguides="false"
            fit-margin-top="5"
            fit-margin-left="5"
            fit-margin-right="5"
            fit-margin-bottom="5" />
    <metadata
            id="metadata7">
        <rdf:RDF>
            <cc:Work
                    rdf:about="">
                <dc:format>image/svg+xml</dc:format>
                <dc:type
                        rdf:resource="http://purl.org/dc/dcmitype/StillImage" />
                <dc:title></dc:title>
            </cc:Work>
        </rdf:RDF>
    </metadata>
    <g
            inkscape:groupmode="layer"
            id="layer6"
            inkscape:label="water"
            transform="translate(-111.7365,-176.02344)"
            style="display:inline">
        <path
                style="fill:url(#linearGradient4295);fill-opacity:1;fill-rule:evenodd;stroke:none;stroke-width:1px;stroke-linecap:butt;stroke-linejoin:miter;stroke-opacity:1"
                d="m 132.04163,433.46271 c 23.83971,-32.31562 103.54234,-15.62414 130.05635,0.6066 89.59503,44.01946 169.11278,-19.05037 186.40202,-33.70711 42.5191,-37.92107 60.76259,-37.33059 75.65481,-30.1241 4.45055,2.15366 10.79479,11.46502 14.40769,12.4991 98.10479,111.17073 -186.00633,131.62933 -317.79746,102.34326 -60.45211,-13.43341 -102.37596,-33.11121 -88.72341,-51.61775 z"
                id="path4287"
                inkscape:connector-curvature="0"
                sodipodi:nodetypes="sccscss" />
    </g>
    <g
            style="display:none"
            transform="translate(-111.7365,-176.02344)"
            inkscape:label="water yellow"
            id="g4764"
            inkscape:groupmode="layer">
        <path
                sodipodi:nodetypes="sccscss"
                inkscape:connector-curvature="0"
                id="path4766"
                d="m 132.04163,433.46271 c 23.83971,-32.31562 103.54234,-15.62414 130.05635,0.6066 89.59503,44.01946 169.11278,-19.05037 186.40202,-33.70711 42.5191,-37.92107 60.76259,-37.33059 75.65481,-30.1241 4.45055,2.15366 10.79479,11.46502 14.40769,12.4991 98.10479,111.17073 -186.00633,131.62933 -317.79746,102.34326 -60.45211,-13.43341 -102.37596,-33.11121 -88.72341,-51.61775 z"
                style="fill:#f3df49;fill-opacity:1;fill-rule:evenodd;stroke:none;stroke-width:1px;stroke-linecap:butt;stroke-linejoin:miter;stroke-opacity:1"
                inkscape:export-filename="F:\BB.png"
                inkscape:export-xdpi="67.080002"
                inkscape:export-ydpi="67.080002" />
    </g>
    <g
            inkscape:groupmode="layer"
            id="layer1"
            inkscape:label="JS"
            transform="translate(17.716534,17.716534)">
        <path
                style="display:inline;fill:#2e2e2c;fill-opacity:1"
                d="m 397.82357,203.03086 c -2.15833,0.003 -4.04434,0.32029 -5.59908,0.97122 -1.55476,0.65095 -2.77858,1.62984 -3.64506,2.85687 -0.86648,1.22704 -1.37588,2.69856 -1.52935,4.30587 -0.15349,1.60734 0.0486,3.34871 0.58179,5.12605 0.57889,1.92971 1.30687,3.54962 2.19183,4.91079 0.88495,1.36117 1.92582,2.45369 3.10744,3.34777 1.1816,0.89409 2.50309,1.58986 3.93146,2.16479 1.42836,0.57497 2.96253,1.02966 4.56608,1.44355 l 0.5531,0.14648 0.54835,0.14938 0.5438,0.15226 0.53943,0.15513 c 0.98305,0.29659 1.84327,0.58763 2.58417,0.91306 0.74091,0.32542 1.36474,0.67712 1.87085,1.08622 0.50609,0.40908 0.89446,0.87518 1.15261,1.45254 0.25814,0.57738 0.3863,1.2648 0.35803,2.15226 -0.0236,0.73958 -0.18511,1.45783 -0.4697,2.13139 -0.28461,0.67354 -0.69231,1.30244 -1.2083,1.86375 -0.51597,0.56132 -1.14021,1.05509 -1.85805,1.45869 -0.71783,0.40365 -1.52914,0.71711 -2.42057,0.91847 -1.06033,0.23937 -2.00599,0.29259 -2.85348,0.18258 -0.8475,-0.11037 -1.59686,-0.384 -2.26928,-0.78971 -0.67241,-0.40574 -1.26797,-0.9433 -1.81208,-1.5773 -0.54411,-0.63399 -1.0369,-1.3642 -1.50671,-2.15437 l -1.98495,1.8073 -2.03798,1.86352 -2.09039,1.91483 -2.14196,1.95887 c 0.65321,1.39463 1.47304,2.66169 2.46907,3.74115 0.99602,1.07948 2.16837,1.97097 3.52109,2.61786 1.35269,0.64687 2.8856,1.04859 4.59719,1.15901 1.71161,0.11038 3.60137,-0.0712 5.66759,-0.57092 2.1138,-0.51178 4.08903,-1.29641 5.86582,-2.33188 1.7768,-1.0355 3.35427,-2.32145 4.66312,-3.83355 1.30887,-1.51209 2.3483,-3.25065 3.05413,-5.18713 0.70586,-1.93648 1.09715,-4.13213 1.04675,-6.31823 -0.0468,-2.02974 -0.39254,-3.57535 -0.95727,-4.92971 -0.56472,-1.35433 -1.35013,-2.50497 -2.32537,-3.5646 -0.97522,-1.05963 -2.1226,-1.878 -3.4771,-2.66038 -1.35451,-0.78239 -2.90711,-1.48633 -4.75517,-2.0661 l -0.54674,-0.16732 -0.55374,-0.16013 -0.5599,-0.15682 -0.57027,-0.14008 c -0.98308,-0.24486 -1.85406,-0.48079 -2.61385,-0.73567 -0.75976,-0.25485 -1.4082,-0.52844 -1.95199,-0.84849 -0.54379,-0.32004 -0.98289,-0.6866 -1.33095,-1.12796 -0.34807,-0.44137 -0.60508,-0.95729 -0.79206,-1.57778 -0.15132,-0.5022 -0.19217,-0.98554 -0.12441,-1.42792 0.0677,-0.44234 0.24435,-0.84353 0.52553,-1.18313 0.28118,-0.33959 0.66708,-0.61742 1.15061,-0.81709 0.48353,-0.19967 1.06476,-0.32107 1.73278,-0.35417 0.6517,-0.032 1.26098,-0.004 1.84305,0.0895 0.58209,0.094 1.13698,0.25443 1.67692,0.48587 0.53993,0.23142 1.04863,0.5269 1.55151,0.89696 0.50288,0.37001 1.00095,0.81468 1.51335,1.346 l 1.16331,-1.30657 1.08059,-1.39481 1.12596,-1.44869 1.30228,-1.46491 c -0.99177,-1.03735 -1.863,-1.84161 -2.76967,-2.49747 -0.90665,-0.65588 -1.86899,-1.17004 -3.04093,-1.61412 -1.17194,-0.44408 -2.5291,-0.80759 -3.91164,-1.03264 -1.38253,-0.22502 -2.43022,-0.27847 -4.37161,-0.33004 -1e-4,0 1.89422,-0.003 0,-3e-4 z m -35.90052,4.17824 2.95712,8.9242 2.48436,8.91613 1.61482,9.05055 0.64779,9.27226 c 0.0956,1.37067 -0.12944,2.54398 -0.43868,3.521 -0.30923,0.97704 -0.76799,1.76798 -1.35219,2.40818 -0.58422,0.64021 -1.29441,1.1298 -2.11506,1.50615 -0.82064,0.37633 -1.75212,0.63954 -2.78055,0.82727 -1.07289,0.19601 -1.99147,0.19576 -2.80086,0.0437 -0.80936,-0.15194 -1.50963,-0.45558 -2.14514,-0.8682 -0.63548,-0.41261 -1.20617,-0.93423 -1.75535,-1.52351 -0.54917,-0.58926 -1.07689,-1.24617 -1.62546,-1.9304 l -2.48262,1.7634 -2.42689,1.73727 -2.35737,1.70957 -2.2764,1.6812 c -1.21926,0.90046 1.79287,2.46599 2.93076,3.51442 1.13789,1.04841 2.4418,1.95339 3.9335,2.65 1.4917,0.6966 3.17113,1.18495 5.06833,1.39489 1.8972,0.20996 4.01297,0.14132 6.38083,-0.28115 2.62207,-0.46777 5.04063,-1.25295 7.2081,-2.3517 2.16746,-1.09874 4.08565,-2.5115 5.70493,-4.23195 1.61926,-1.72043 2.94249,-3.74916 3.88298,-6.06916 0.94049,-2.32001 1.68304,-4.9681 1.51137,-7.80022 l -0.4731,-9.29203 -1.58977,-8.97986 -2.62278,-8.76892 -3.25911,-8.72796 -2.8918,0.47986 -2.94,0.47991 -2.98043,0.47627 z"
                id="j"
                inkscape:connector-curvature="0"
                sodipodi:nodetypes="cssssssscccccssssssssssscccccssssssssssscccccssssssssssscccccssssscccsssssssscccccsssssssscccccccccc"
                inkscape:export-filename="F:\BB.png"
                inkscape:export-xdpi="67.080002"
                inkscape:export-ydpi="67.080002" />
    </g>
    <g
            inkscape:groupmode="layer"
            id="layer5"
            inkscape:label="board back"
            transform="translate(-111.7365,-176.02344)">
        <path
                style="display:inline;fill:#388e3c;fill-opacity:1;fill-rule:evenodd;stroke:#000000;stroke-width:2;stroke-linecap:butt;stroke-linejoin:round;stroke-miterlimit:4;stroke-dasharray:none;stroke-opacity:1"
                d="m 198.7904,380.42903 c 0.25809,-2.63099 4.09727,-13.36941 31.37399,-0.9521 87.58706,39.87267 172.37344,28.76155 238.08524,1.51305 -35.44555,31.73992 -105.09012,70.96031 -151.49963,82.62222 -33.17285,-0.25002 -69.22045,-8.33225 -91.92216,-33.7906 -10.3767,-11.63672 -27.08255,-38.73851 -26.03744,-49.39257 z"
                id="path4159"
                inkscape:connector-curvature="0"
                sodipodi:nodetypes="ssccss"
                inkscape:export-filename="F:\BB.png"
                inkscape:export-xdpi="67.080002"
                inkscape:export-ydpi="67.080002" />
        <ellipse
                style="opacity:1;fill:#1b5e2f;fill-opacity:0.94117647;stroke:none;stroke-width:2;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:4;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1"
                id="path4284"
                cx="280.87637"
                cy="474.38382"
                rx="83.288368"
                ry="22.619684"
                transform="matrix(0.98520272,-0.17139312,0.13103475,0.99137778,0,0)"
                inkscape:export-filename="F:\BB.png"
                inkscape:export-xdpi="67.080002"
                inkscape:export-ydpi="67.080002" />
    </g>
    <g
            inkscape:groupmode="layer"
            id="layer3"
            inkscape:label="gopher"
            transform="translate(-111.7365,-176.02344)">
        <path
                style="display:inline;opacity:1;fill:#8ed4fe;fill-opacity:1;fill-rule:evenodd;stroke:#000000;stroke-width:2.36734128;stroke-linecap:round;stroke-linejoin:miter;stroke-miterlimit:4;stroke-dasharray:none;stroke-opacity:1"
                d="m 369.5092,223.97146 c 8.70159,-7.65557 16.28692,-10.65856 23.07967,-7.25757 4.97464,2.49072 3.8793,9.20637 0.18568,14.34944 -4.40361,6.13167 -8.48975,9.32099 -8.7434,9.31074"
                id="path4191"
                inkscape:connector-curvature="0"
                sodipodi:nodetypes="cssc"
                inkscape:export-filename="F:\BB.png"
                inkscape:export-xdpi="67.080002"
                inkscape:export-ydpi="67.080002" />
        <path
                style="display:inline;opacity:1;fill:#000000;fill-opacity:1;fill-rule:evenodd;stroke:#000000;stroke-width:0.39455688px;stroke-linecap:butt;stroke-linejoin:miter;stroke-opacity:1"
                d="m 377.17602,235.83381 c 0.0393,0.82372 4.24935,3.42187 5.45713,3.096 1.76689,-0.47671 8.64143,-9.05634 8.60055,-10.83044 -0.11459,-4.97101 -14.08317,7.20081 -14.05768,7.73444 z"
                id="path4201"
                inkscape:connector-curvature="0"
                sodipodi:nodetypes="ssss"
                inkscape:export-filename="F:\BB.png"
                inkscape:export-xdpi="67.080002"
                inkscape:export-ydpi="67.080002" />
        <path
                sodipodi:nodetypes="ccccc"
                inkscape:connector-curvature="0"
                id="path4251"
                d="m 388.09226,315.2492 c 1.49571,-1.40459 4.86627,-1.9416 7.21199,-3.03825 4.64492,-2.17612 7.61104,-5.53269 6.6251,-7.49725 -0.98527,-1.96513 -5.55031,-1.79401 -10.19585,0.38219 -2.24303,1.05203 -4.38565,1.92955 -6.40835,2.1379"
                style="display:inline;opacity:1;fill:#8ed4fd;fill-opacity:1;stroke:#000000;stroke-width:2;stroke-linecap:round;stroke-linejoin:bevel;stroke-miterlimit:4;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1"
                inkscape:export-filename="F:\BB.png"
                inkscape:export-xdpi="67.080002"
                inkscape:export-ydpi="67.080002" />
        <path
                style="fill:#8ed4fd;fill-opacity:1;fill-rule:evenodd;stroke:#000000;stroke-width:4;stroke-linecap:butt;stroke-linejoin:miter;stroke-miterlimit:4;stroke-dasharray:none;stroke-opacity:1"
                d="m 298.5481,433.97311 c -13.21919,-2.14324 -19.67913,-12.58085 -26.60671,-23.44515 -39.79944,-62.41617 -31.2814,-150.25148 -11.38378,-173.74419 28.95072,-34.18152 95.75894,-33.50993 125.69239,5.32843 23.50633,26.83283 -1.18688,59.97242 -0.60355,88.63909 C 427.86396,348.35484 422,405.1122 422,405.1122 c 0,0 -29.16421,-3.51472 -46.41421,0.48528 -13.46876,3.12319 -29.73182,10.11303 -48.81044,19.56333 -9.17029,4.54236 -16.53921,10.70729 -28.22725,8.8123 z"
                id="path4235"
                inkscape:connector-curvature="0"
                sodipodi:nodetypes="ssscccsss"
                inkscape:export-filename="F:\BB.png"
                inkscape:export-xdpi="67.080002"
                inkscape:export-ydpi="67.080002" />
        <path
                style="display:inline;opacity:1;fill:#8ed4fd;fill-opacity:1;stroke:#000000;stroke-width:4;stroke-linecap:round;stroke-linejoin:bevel;stroke-miterlimit:4;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1"
                d="m 250.54937,329.78787 c -2.01484,1.9022 -6.56468,2.63951 -9.72851,4.1289 -6.26491,2.95542 -10.25844,7.49896 -8.91992,10.14844 1.33762,2.65027 7.50185,2.40279 13.76758,-0.55273 3.02532,-1.42878 5.91567,-2.62152 8.64648,-2.91016"
                id="path4183"
                inkscape:connector-curvature="0"
                sodipodi:nodetypes="ccccc"
                inkscape:export-filename="F:\BB.png"
                inkscape:export-xdpi="67.080002"
                inkscape:export-ydpi="67.080002" />
        <path
                style="fill:none;fill-rule:evenodd;stroke:#000000;stroke-width:1px;stroke-linecap:butt;stroke-linejoin:miter;stroke-opacity:1"
                d="m 391.3836,334.29527 c -9.82295,-5.14566 -22.27386,-2.47488 -22.27386,-2.47488"
                id="path4253"
                inkscape:connector-curvature="0"
                sodipodi:nodetypes="cc"
                inkscape:export-filename="F:\BB.png"
                inkscape:export-xdpi="67.080002"
                inkscape:export-ydpi="67.080002" />
        <path
                sodipodi:nodetypes="csc"
                inkscape:connector-curvature="0"
                id="path4255"
                d="m 286.83033,422.58661 c 6.99565,10.91962 1.84084,23.66997 8.24051,25.19347 5.23572,1.24641 10.82844,-11.01538 16.59576,-20.23475"
                style="display:inline;opacity:1;fill:#8ed4fd;fill-opacity:1;stroke:#000000;stroke-width:4;stroke-linecap:round;stroke-linejoin:bevel;stroke-miterlimit:4;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1"
                inkscape:export-filename="F:\BB.png"
                inkscape:export-xdpi="67.080002"
                inkscape:export-ydpi="67.080002" />
        <circle
                style="opacity:1;fill:#ffffff;fill-opacity:1;stroke:#000000;stroke-width:3;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:4;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1"
                id="path4212"
                cx="312.05698"
                cy="259.30524"
                r="33.4146"
                inkscape:export-filename="F:\BB.png"
                inkscape:export-xdpi="67.080002"
                inkscape:export-ydpi="67.080002" />
        <ellipse
                cy="275.77502"
                cx="383.2565"
                id="circle4214"
                style="opacity:1;fill:#ffffff;fill-opacity:1;stroke:#000000;stroke-width:2.80603194;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:4;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1"
                rx="23.076626"
                ry="24.170307"
                inkscape:export-filename="F:\BB.png"
                inkscape:export-xdpi="67.080002"
                inkscape:export-ydpi="67.080002" />
        <g
                id="g4257"
                transform="translate(4.3284271,0.4571068)"
                inkscape:export-filename="F:\BB.png"
                inkscape:export-xdpi="67.080002"
                inkscape:export-ydpi="67.080002">
            <path
                    inkscape:export-ydpi="210.08769"
                    inkscape:export-xdpi="210.08769"
                    inkscape:export-filename="F:\C-Gopher.png"
                    sodipodi:nodetypes="cccscc"
                    inkscape:connector-curvature="0"
                    id="path4224"
                    d="m 335.72791,298.82062 c -1.10745,3.63148 -0.69833,12.2121 -0.69833,12.2121 l 10.47516,3.15242 c 0,0 4.85288,-5.71445 5.44452,-7.78161 0.59168,-2.06713 0.3377,-3.2066 0.3377,-3.2066 z"
                    style="display:inline;opacity:1;fill:#ffffff;fill-opacity:1;fill-rule:evenodd;stroke:#000000;stroke-width:2.56399131;stroke-linecap:butt;stroke-linejoin:miter;stroke-miterlimit:4;stroke-dasharray:none;stroke-opacity:1" />
            <path
                    inkscape:export-ydpi="210.08769"
                    inkscape:export-xdpi="210.08769"
                    inkscape:export-filename="F:\C-Gopher.png"
                    sodipodi:nodetypes="ssss"
                    inkscape:connector-curvature="0"
                    id="path4222"
                    d="m 327.71579,292.84858 c -7.61561,7.89636 4.59118,8.08453 14.16297,10.19598 7.57686,1.67137 16.83011,10.15944 19.09861,2.50827 3.92037,-13.22268 -23.49485,-22.83102 -33.26158,-12.70425 z"
                    style="display:inline;opacity:1;fill:#d3b78d;fill-opacity:1;fill-rule:evenodd;stroke:#000000;stroke-width:2.56399131;stroke-linecap:butt;stroke-linejoin:miter;stroke-miterlimit:4;stroke-dasharray:none;stroke-opacity:1" />
            <ellipse
                    inkscape:export-ydpi="210.08769"
                    inkscape:export-xdpi="210.08769"
                    inkscape:export-filename="F:\C-Gopher.png"
                    transform="matrix(0.9497031,0.31315174,0.31315174,-0.9497031,0,0)"
                    ry="6.4980431"
                    rx="11.350795"
                    cy="-167.76379"
                    cx="420.1701"
                    id="path4220"
                    style="display:inline;opacity:1;fill:#000000;fill-opacity:1;stroke:#000000;stroke-width:1.92299342;stroke-linecap:round;stroke-linejoin:bevel;stroke-miterlimit:4;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1" />
        </g>
        <circle
                style="opacity:1;fill:#000000;fill-opacity:1;stroke:#000000;stroke-width:2.28475904;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:4;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1"
                id="path4231"
                cx="324.33755"
                cy="267.0163"
                r="8.8534412"
                inkscape:export-filename="F:\BB.png"
                inkscape:export-xdpi="67.080002"
                inkscape:export-ydpi="67.080002" />
        <circle
                r="6.7808099"
                cy="283.75888"
                cx="392.33401"
                id="circle4233"
                style="opacity:1;fill:#000000;fill-opacity:1;stroke:#000000;stroke-width:1.74988639;stroke-linecap:round;stroke-linejoin:round;stroke-miterlimit:4;stroke-dasharray:none;stroke-dashoffset:0;stroke-opacity:1"
                inkscape:export-filename="F:\BB.png"
                inkscape:export-xdpi="67.080002"
                inkscape:export-ydpi="67.080002" />
        <path
                style="display:inline;opacity:1;fill:#8ed4fe;fill-opacity:1;fill-rule:evenodd;stroke:#000000;stroke-width:2.91332936;stroke-linecap:round;stroke-linejoin:miter;stroke-miterlimit:4;stroke-dasharray:none;stroke-opacity:1"
                d="m 278.19769,228.79463 c -3.90156,-10.60895 -2.25892,-20.22855 2.5085,-28.16278 3.34565,-5.56805 7.80324,-6.18771 11.59575,-4.82156 6.3323,2.28107 9.1159,9.46446 12.68017,20.66394"
                id="path4189"
                inkscape:connector-curvature="0"
                sodipodi:nodetypes="cssc"
                inkscape:export-filename="F:\BB.png"
                inkscape:export-xdpi="67.080002"
                inkscape:export-ydpi="67.080002" />
        <path
                sodipodi:nodetypes="ssss"
                inkscape:connector-curvature="0"
                id="path4212-4"
                d="m 288.66496,223.71799 c 0.78826,0.63922 5.95007,-1.75896 7.31549,-2.47008 2.72013,-1.41667 0.13056,-13.61893 -2.11014,-15.23773 -4.9601,-3.58345 -5.71599,17.2937 -5.20535,17.70781 z"
                style="display:inline;opacity:1;fill:#000000;fill-opacity:1;fill-rule:evenodd;stroke:#000000;stroke-width:0.4855549px;stroke-linecap:butt;stroke-linejoin:miter;stroke-opacity:1"
                inkscape:export-filename="F:\BB.png"
                inkscape:export-xdpi="67.080002"
                inkscape:export-ydpi="67.080002" />
    </g>
    <g
            inkscape:groupmode="layer"
            id="layer2"
            inkscape:label="board"
            style="display:inline"
            transform="translate(-111.7365,-176.02344)">
        <path
                style="fill:#4caf50;fill-opacity:1;fill-rule:evenodd;stroke:#000000;stroke-width:2;stroke-linecap:butt;stroke-linejoin:round;stroke-miterlimit:4;stroke-dasharray:none;stroke-opacity:1"
                d="m 199.50538,378.71575 c -3.00355,5.85609 -4.3397,9.43062 -3.46266,28.28214 21.66777,47.76999 122.45022,97.48537 235.38585,8.5786 16.78571,-13.21429 43.01355,-41.15656 70.71429,-46.42857 20.43968,-3.89009 35.55187,12.63729 35.55187,12.63729 l -9.16997,-19.11611 c -10.61241,-21.21673 -55.7985,-16.39216 -82.11055,2.94329 -35.08701,25.78374 -70.29979,86.10442 -142.84278,83.53553 -79.71358,-2.82281 -108.8202,-63.39239 -104.06605,-70.43217 z"
                id="path4149"
                inkscape:connector-curvature="0"
                sodipodi:nodetypes="ccssccssc"
                inkscape:export-filename="F:\BB.png"
                inkscape:export-xdpi="67.080002"
                inkscape:export-ydpi="67.080002" />
    </g>
</svg>`,
		"vue.webapp/src/components/Tech.vue": `<template>
  <ul class="technologies">
    <li v-for="technology in technologies" v-bind:key="technology.name">
      <b>{{technology.name}}</b>: {{technology.details}}
    </li>
  </ul>
</template>

<script>
import axios from 'axios'

export default {
  name: 'Tech',
  data() {
    return {
      technologies: []
    };
  },
  mounted() {
    axios
      .get(` + "`" + `${process.env.VUE_APP_API_URL}/api/technologies` + "`" + `)
      .then(response => (this.technologies = response.data));
  }
};
</script>

<style scoped>
.technologies {
  margin-top: 5px;
}
</style>
`,
		"vue.webapp/src/main.js": `import Vue from 'vue'
import App from './App.vue'

Vue.config.productionTip = false

new Vue({
  render: h => h(App),
}).$mount('#app')
`,
	}
}

func Images() map[string][]byte {
	return map[string][]byte{
		"angular.webapp/src/favicon.ico":  {0, 0, 1, 0, 1, 0, 16, 16, 0, 0, 1, 0, 32, 0, 86, 4, 0, 0, 22, 0, 0, 0, 40, 0, 0, 0, 16, 0, 0, 0, 32, 0, 0, 0, 1, 0, 32, 0, 0, 0, 0, 0, 0, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 19, 18, 11, 17, 65, 61, 38, 60, 111, 103, 65, 102, 147, 134, 86, 133, 167, 153, 98, 153, 176, 161, 103, 164, 173, 160, 102, 165, 162, 153, 96, 157, 144, 135, 85, 139, 115, 109, 68, 110, 75, 71, 44, 69, 30, 28, 17, 23, 0, 0, 0, 0, 0, 0, 0, 0, 59, 56, 35, 55, 171, 161, 101, 164, 237, 222, 140, 230, 255, 248, 157, 255, 242, 237, 146, 255, 193, 214, 124, 255, 168, 198, 111, 255, 187, 210, 121, 255, 230, 232, 141, 255, 255, 250, 158, 255, 255, 255, 161, 255, 255, 253, 159, 255, 255, 244, 154, 255, 216, 204, 130, 236, 171, 161, 101, 162, 37, 35, 22, 30, 225, 212, 133, 217, 255, 255, 168, 255, 255, 248, 157, 255, 159, 189, 106, 255, 66, 140, 59, 255, 47, 100, 36, 255, 95, 112, 59, 255, 38, 94, 29, 255, 48, 118, 42, 255, 111, 170, 84, 255, 210, 216, 130, 255, 255, 244, 155, 255, 253, 238, 150, 255, 171, 163, 110, 255, 181, 172, 116, 255, 164, 156, 102, 210, 37, 35, 22, 34, 100, 90, 58, 86, 104, 130, 71, 185, 58, 146, 58, 255, 48, 135, 49, 255, 146, 145, 85, 255, 255, 210, 143, 255, 205, 175, 114, 255, 116, 120, 64, 255, 70, 108, 47, 255, 81, 158, 71, 255, 170, 201, 112, 255, 255, 247, 157, 255, 205, 192, 126, 255, 149, 142, 99, 255, 179, 170, 111, 228, 0, 0, 0, 0, 0, 0, 0, 0, 14, 44, 16, 114, 42, 108, 41, 206, 55, 80, 39, 172, 245, 204, 137, 254, 255, 216, 145, 255, 255, 217, 146, 255, 255, 216, 146, 255, 255, 209, 142, 255, 182, 168, 106, 243, 53, 137, 56, 231, 144, 196, 103, 255, 232, 227, 140, 255, 230, 222, 139, 255, 82, 78, 49, 96, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 5, 107, 86, 59, 149, 255, 222, 149, 255, 253, 212, 142, 255, 253, 212, 142, 255, 254, 213, 142, 255, 255, 225, 151, 255, 129, 104, 71, 164, 0, 5, 0, 27, 38, 97, 40, 162, 74, 141, 64, 217, 60, 100, 48, 168, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 18, 159, 133, 89, 221, 255, 220, 148, 255, 255, 213, 142, 255, 255, 214, 143, 255, 255, 221, 147, 255, 165, 137, 90, 206, 0, 0, 0, 19, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 166, 139, 93, 199, 255, 221, 146, 255, 245, 206, 140, 255, 194, 172, 130, 255, 147, 139, 118, 255, 148, 130, 100, 221, 0, 0, 0, 19, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 151, 126, 82, 187, 246, 214, 159, 255, 218, 216, 213, 255, 138, 134, 124, 255, 174, 161, 138, 255, 211, 213, 217, 255, 39, 39, 39, 82, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 83, 68, 42, 122, 235, 211, 168, 255, 246, 250, 255, 255, 203, 200, 194, 255, 251, 215, 155, 255, 166, 155, 136, 223, 8, 9, 11, 20, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 6, 136, 114, 78, 200, 193, 175, 145, 255, 194, 171, 133, 235, 152, 126, 82, 190, 88, 72, 44, 159, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 42, 35, 23, 73, 40, 31, 17, 81, 1, 0, 0, 16, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		"react.webapp/public/favicon.ico": {0, 0, 1, 0, 1, 0, 16, 16, 0, 0, 1, 0, 32, 0, 86, 4, 0, 0, 22, 0, 0, 0, 40, 0, 0, 0, 16, 0, 0, 0, 32, 0, 0, 0, 1, 0, 32, 0, 0, 0, 0, 0, 0, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 19, 18, 11, 17, 65, 61, 38, 60, 111, 103, 65, 102, 147, 134, 86, 133, 167, 153, 98, 153, 176, 161, 103, 164, 173, 160, 102, 165, 162, 153, 96, 157, 144, 135, 85, 139, 115, 109, 68, 110, 75, 71, 44, 69, 30, 28, 17, 23, 0, 0, 0, 0, 0, 0, 0, 0, 59, 56, 35, 55, 171, 161, 101, 164, 237, 222, 140, 230, 255, 248, 157, 255, 242, 237, 146, 255, 193, 214, 124, 255, 168, 198, 111, 255, 187, 210, 121, 255, 230, 232, 141, 255, 255, 250, 158, 255, 255, 255, 161, 255, 255, 253, 159, 255, 255, 244, 154, 255, 216, 204, 130, 236, 171, 161, 101, 162, 37, 35, 22, 30, 225, 212, 133, 217, 255, 255, 168, 255, 255, 248, 157, 255, 159, 189, 106, 255, 66, 140, 59, 255, 47, 100, 36, 255, 95, 112, 59, 255, 38, 94, 29, 255, 48, 118, 42, 255, 111, 170, 84, 255, 210, 216, 130, 255, 255, 244, 155, 255, 253, 238, 150, 255, 171, 163, 110, 255, 181, 172, 116, 255, 164, 156, 102, 210, 37, 35, 22, 34, 100, 90, 58, 86, 104, 130, 71, 185, 58, 146, 58, 255, 48, 135, 49, 255, 146, 145, 85, 255, 255, 210, 143, 255, 205, 175, 114, 255, 116, 120, 64, 255, 70, 108, 47, 255, 81, 158, 71, 255, 170, 201, 112, 255, 255, 247, 157, 255, 205, 192, 126, 255, 149, 142, 99, 255, 179, 170, 111, 228, 0, 0, 0, 0, 0, 0, 0, 0, 14, 44, 16, 114, 42, 108, 41, 206, 55, 80, 39, 172, 245, 204, 137, 254, 255, 216, 145, 255, 255, 217, 146, 255, 255, 216, 146, 255, 255, 209, 142, 255, 182, 168, 106, 243, 53, 137, 56, 231, 144, 196, 103, 255, 232, 227, 140, 255, 230, 222, 139, 255, 82, 78, 49, 96, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 5, 107, 86, 59, 149, 255, 222, 149, 255, 253, 212, 142, 255, 253, 212, 142, 255, 254, 213, 142, 255, 255, 225, 151, 255, 129, 104, 71, 164, 0, 5, 0, 27, 38, 97, 40, 162, 74, 141, 64, 217, 60, 100, 48, 168, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 18, 159, 133, 89, 221, 255, 220, 148, 255, 255, 213, 142, 255, 255, 214, 143, 255, 255, 221, 147, 255, 165, 137, 90, 206, 0, 0, 0, 19, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 166, 139, 93, 199, 255, 221, 146, 255, 245, 206, 140, 255, 194, 172, 130, 255, 147, 139, 118, 255, 148, 130, 100, 221, 0, 0, 0, 19, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 151, 126, 82, 187, 246, 214, 159, 255, 218, 216, 213, 255, 138, 134, 124, 255, 174, 161, 138, 255, 211, 213, 217, 255, 39, 39, 39, 82, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 83, 68, 42, 122, 235, 211, 168, 255, 246, 250, 255, 255, 203, 200, 194, 255, 251, 215, 155, 255, 166, 155, 136, 223, 8, 9, 11, 20, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 6, 136, 114, 78, 200, 193, 175, 145, 255, 194, 171, 133, 235, 152, 126, 82, 190, 88, 72, 44, 159, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 42, 35, 23, 73, 40, 31, 17, 81, 1, 0, 0, 16, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		"react.webapp/public/logo192.png": {137, 80, 78, 71, 13, 10, 26, 10, 0, 0, 0, 13, 73, 72, 68, 82, 0, 0, 0, 192, 0, 0, 0, 192, 8, 6, 0, 0, 0, 82, 220, 108, 7, 0, 0, 0, 4, 115, 66, 73, 84, 8, 8, 8, 8, 124, 8, 100, 136, 0, 0, 0, 9, 112, 72, 89, 115, 0, 0, 6, 49, 0, 0, 6, 49, 1, 132, 67, 140, 150, 0, 0, 0, 25, 116, 69, 88, 116, 83, 111, 102, 116, 119, 97, 114, 101, 0, 119, 119, 119, 46, 105, 110, 107, 115, 99, 97, 112, 101, 46, 111, 114, 103, 155, 238, 60, 26, 0, 0, 32, 0, 73, 68, 65, 84, 120, 156, 237, 157, 103, 116, 149, 197, 214, 128, 159, 57, 189, 164, 247, 132, 36, 132, 132, 94, 133, 132, 34, 29, 65, 69, 138, 88, 81, 64, 68, 197, 11, 216, 174, 189, 93, 133, 107, 185, 214, 79, 5, 11, 10, 246, 130, 93, 81, 170, 34, 85, 170, 244, 22, 122, 11, 33, 144, 222, 203, 233, 103, 190, 31, 65, 32, 38, 164, 135, 132, 195, 251, 172, 149, 181, 78, 230, 157, 178, 223, 115, 102, 79, 221, 179, 7, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 154, 6, 26, 224, 30, 181, 70, 183, 72, 171, 55, 237, 86, 169, 84, 95, 3, 215, 2, 162, 145, 229, 82, 80, 104, 112, 204, 26, 173, 126, 173, 86, 111, 112, 244, 184, 238, 14, 121, 213, 61, 211, 100, 187, 1, 195, 157, 66, 168, 220, 106, 141, 110, 9, 224, 219, 216, 2, 42, 212, 47, 74, 171, 118, 14, 42, 181, 246, 107, 147, 143, 223, 205, 247, 127, 181, 90, 27, 18, 211, 230, 76, 248, 201, 253, 59, 248, 232, 158, 17, 142, 226, 220, 204, 245, 46, 167, 99, 48, 224, 106, 60, 41, 21, 234, 19, 85, 99, 11, 208, 132, 232, 235, 118, 57, 198, 222, 246, 250, 156, 50, 149, 31, 160, 89, 219, 203, 184, 231, 147, 165, 90, 161, 82, 245, 5, 30, 108, 28, 241, 20, 26, 2, 117, 99, 11, 208, 84, 208, 104, 245, 95, 183, 233, 115, 85, 196, 85, 83, 166, 86, 216, 40, 120, 249, 7, 35, 84, 42, 213, 209, 173, 107, 122, 75, 183, 251, 29, 192, 113, 129, 69, 84, 104, 0, 148, 30, 160, 148, 118, 78, 135, 173, 239, 224, 187, 159, 210, 252, 243, 193, 182, 197, 223, 178, 234, 139, 183, 0, 232, 59, 230, 62, 84, 26, 141, 17, 184, 227, 2, 203, 167, 208, 64, 40, 10, 80, 202, 205, 126, 225, 81, 246, 216, 248, 126, 101, 2, 139, 115, 51, 249, 245, 213, 135, 88, 252, 246, 127, 56, 178, 249, 79, 12, 94, 190, 92, 126, 243, 36, 141, 90, 167, 127, 28, 101, 254, 228, 17, 148, 107, 241, 46, 69, 52, 58, 195, 176, 14, 3, 71, 234, 254, 25, 190, 253, 247, 31, 40, 202, 201, 4, 96, 207, 159, 11, 136, 235, 62, 128, 222, 183, 76, 97, 245, 87, 111, 55, 7, 250, 3, 167, 128, 72, 32, 10, 8, 3, 180, 128, 247, 57, 89, 20, 0, 25, 64, 22, 144, 13, 236, 7, 50, 27, 242, 93, 20, 106, 134, 162, 0, 128, 116, 187, 58, 68, 182, 239, 86, 46, 60, 253, 200, 222, 51, 159, 83, 246, 108, 3, 32, 36, 166, 13, 97, 113, 237, 101, 218, 145, 189, 43, 169, 93, 47, 144, 6, 236, 62, 253, 247, 39, 176, 2, 40, 170, 69, 62, 10, 245, 128, 162, 0, 224, 229, 114, 58, 188, 2, 155, 181, 40, 247, 32, 63, 227, 212, 153, 207, 89, 39, 14, 159, 249, 28, 63, 242, 54, 177, 116, 214, 255, 232, 56, 104, 36, 190, 161, 145, 248, 133, 71, 225, 29, 24, 134, 193, 219, 23, 149, 234, 236, 186, 130, 173, 164, 144, 226, 220, 44, 138, 114, 51, 41, 202, 206, 32, 237, 200, 94, 82, 15, 237, 14, 43, 204, 74, 15, 3, 174, 4, 30, 65, 8, 7, 82, 174, 5, 150, 0, 223, 0, 39, 106, 40, 127, 47, 149, 74, 51, 92, 103, 54, 119, 83, 169, 212, 58, 107, 113, 209, 118, 183, 211, 254, 61, 176, 181, 134, 249, 92, 146, 40, 10, 112, 122, 200, 162, 51, 121, 149, 123, 80, 156, 155, 117, 230, 179, 181, 168, 16, 187, 181, 4, 157, 193, 196, 101, 215, 140, 102, 221, 183, 239, 115, 253, 51, 239, 98, 246, 11, 172, 113, 129, 197, 185, 89, 156, 216, 187, 141, 131, 27, 150, 178, 127, 221, 31, 218, 180, 67, 137, 131, 128, 65, 192, 75, 192, 239, 192, 199, 192, 66, 192, 89, 73, 54, 87, 233, 205, 222, 239, 170, 52, 218, 168, 174, 87, 223, 172, 137, 108, 223, 77, 171, 214, 234, 72, 217, 187, 109, 224, 230, 95, 191, 120, 0, 220, 63, 91, 139, 139, 198, 3, 178, 198, 2, 94, 66, 40, 19, 57, 240, 1, 242, 239, 251, 124, 37, 113, 9, 3, 202, 60, 120, 237, 218, 14, 164, 31, 221, 7, 128, 201, 199, 159, 105, 43, 78, 160, 51, 152, 168, 239, 58, 149, 159, 113, 146, 221, 203, 126, 229, 175, 159, 63, 230, 212, 129, 93, 165, 129, 66, 156, 68, 202, 23, 129, 79, 41, 187, 228, 106, 210, 25, 205, 159, 104, 244, 134, 81, 163, 30, 127, 195, 24, 63, 98, 28, 42, 117, 217, 118, 204, 82, 152, 199, 140, 91, 122, 22, 103, 38, 31, 154, 2, 204, 169, 87, 97, 61, 12, 101, 21, 8, 138, 85, 42, 181, 243, 220, 225, 206, 223, 56, 157, 165, 245, 206, 39, 56, 156, 190, 227, 30, 64, 103, 48, 210, 16, 13, 170, 111, 72, 51, 250, 142, 189, 143, 199, 126, 222, 206, 67, 223, 110, 160, 215, 141, 19, 209, 104, 117, 17, 192, 44, 16, 251, 129, 113, 148, 254, 86, 173, 117, 70, 175, 93, 45, 186, 245, 25, 245, 204, 226, 131, 198, 238, 163, 38, 148, 171, 252, 0, 70, 111, 63, 250, 142, 187, 223, 108, 244, 242, 189, 161, 222, 133, 245, 48, 46, 245, 33, 80, 87, 224, 67, 149, 70, 163, 73, 61, 184, 27, 134, 141, 41, 243, 176, 125, 191, 107, 240, 9, 10, 167, 215, 205, 19, 49, 251, 5, 149, 75, 44, 165, 196, 90, 152, 71, 73, 65, 46, 58, 163, 25, 163, 183, 31, 26, 157, 190, 78, 2, 69, 119, 234, 65, 116, 167, 30, 92, 117, 207, 84, 177, 228, 253, 23, 216, 60, 239, 139, 22, 110, 151, 107, 14, 240, 144, 90, 163, 109, 127, 229, 164, 167, 245, 87, 220, 253, 148, 90, 136, 202, 59, 111, 141, 86, 7, 66, 104, 235, 36, 204, 37, 192, 165, 60, 4, 186, 91, 8, 241, 158, 148, 82, 31, 20, 221, 18, 163, 183, 31, 15, 127, 191, 233, 156, 199, 21, 183, 244, 39, 247, 109, 39, 113, 229, 124, 14, 172, 249, 141, 83, 135, 18, 177, 91, 45, 103, 158, 169, 212, 106, 130, 154, 197, 16, 215, 243, 10, 58, 14, 26, 69, 171, 158, 131, 208, 232, 13, 117, 18, 50, 51, 233, 32, 243, 223, 124, 130, 61, 43, 23, 32, 132, 202, 54, 252, 193, 151, 242, 6, 77, 124, 34, 244, 124, 10, 224, 114, 216, 81, 107, 117, 124, 245, 248, 88, 235, 206, 37, 63, 188, 224, 118, 187, 95, 169, 147, 0, 30, 206, 165, 168, 0, 6, 224, 61, 96, 162, 222, 236, 45, 199, 190, 252, 185, 48, 152, 125, 152, 245, 175, 171, 120, 254, 207, 84, 188, 2, 130, 203, 37, 144, 82, 178, 123, 249, 175, 44, 121, 119, 42, 169, 71, 246, 17, 209, 44, 146, 81, 215, 142, 36, 62, 62, 158, 232, 232, 104, 124, 125, 125, 41, 42, 42, 34, 43, 43, 139, 61, 123, 246, 176, 112, 209, 98, 182, 111, 219, 138, 193, 236, 205, 128, 9, 143, 48, 240, 142, 71, 208, 25, 205, 117, 18, 122, 203, 252, 175, 152, 251, 242, 131, 210, 90, 148, 47, 226, 186, 15, 176, 222, 245, 246, 92, 131, 209, 199, 255, 204, 243, 227, 59, 255, 226, 143, 89, 255, 227, 248, 174, 191, 152, 60, 123, 9, 239, 222, 222, 215, 230, 180, 219, 90, 3, 201, 117, 42, 216, 195, 185, 212, 20, 192, 4, 44, 2, 6, 134, 181, 236, 32, 239, 156, 241, 147, 8, 142, 105, 141, 219, 237, 98, 211, 220, 207, 232, 126, 221, 4, 212, 154, 178, 163, 134, 180, 195, 123, 248, 241, 191, 147, 72, 218, 181, 145, 209, 183, 220, 194, 163, 143, 60, 66, 66, 66, 2, 85, 13, 65, 82, 82, 82, 248, 252, 243, 207, 121, 237, 245, 255, 67, 99, 48, 49, 226, 177, 55, 232, 54, 124, 76, 165, 105, 170, 34, 47, 61, 133, 111, 159, 185, 147, 67, 127, 173, 32, 164, 69, 91, 247, 228, 217, 191, 169, 252, 194, 163, 89, 246, 225, 203, 172, 248, 244, 117, 108, 197, 133, 116, 30, 114, 3, 110, 183, 203, 122, 96, 221, 31, 31, 57, 108, 150, 127, 215, 169, 192, 75, 128, 75, 73, 1, 12, 192, 124, 224, 202, 246, 253, 135, 113, 251, 155, 223, 85, 217, 42, 239, 93, 189, 152, 175, 159, 24, 71, 167, 14, 237, 153, 249, 222, 187, 36, 36, 36, 212, 184, 208, 172, 172, 44, 254, 251, 223, 255, 242, 193, 7, 31, 48, 96, 194, 195, 140, 120, 228, 213, 50, 123, 5, 53, 197, 237, 114, 242, 243, 75, 15, 176, 225, 135, 15, 241, 14, 12, 149, 81, 29, 18, 196, 222, 213, 139, 128, 210, 201, 122, 175, 27, 39, 186, 87, 125, 49, 61, 221, 110, 41, 110, 15, 228, 213, 186, 160, 75, 132, 75, 197, 26, 84, 11, 204, 5, 134, 182, 190, 124, 8, 119, 189, 51, 23, 173, 193, 84, 105, 130, 191, 126, 250, 152, 175, 159, 26, 207, 184, 177, 99, 248, 249, 167, 159, 136, 142, 142, 174, 85, 193, 38, 147, 137, 225, 195, 135, 211, 170, 85, 43, 102, 190, 60, 149, 148, 196, 45, 116, 185, 250, 38, 132, 170, 118, 11, 112, 66, 165, 162, 195, 128, 17, 104, 180, 122, 18, 87, 45, 16, 89, 201, 135, 0, 208, 25, 76, 244, 29, 123, 31, 43, 62, 253, 63, 171, 195, 102, 25, 128, 50, 244, 169, 22, 151, 138, 2, 188, 0, 76, 108, 217, 125, 0, 19, 103, 206, 71, 107, 48, 86, 26, 249, 224, 134, 101, 204, 121, 114, 28, 207, 62, 243, 12, 51, 102, 204, 64, 163, 169, 251, 98, 89, 167, 78, 157, 24, 60, 120, 48, 51, 94, 121, 158, 194, 156, 76, 218, 245, 187, 166, 78, 249, 197, 198, 247, 197, 47, 44, 146, 196, 85, 11, 16, 42, 21, 221, 71, 77, 96, 253, 15, 179, 173, 78, 155, 245, 102, 96, 109, 157, 5, 190, 68, 184, 20, 20, 160, 15, 240, 137, 111, 72, 132, 184, 239, 243, 149, 194, 224, 229, 83, 105, 228, 172, 228, 195, 124, 56, 121, 40, 55, 92, 127, 29, 239, 190, 243, 78, 149, 99, 253, 154, 16, 21, 21, 69, 235, 214, 173, 121, 235, 191, 79, 224, 23, 22, 73, 100, 187, 174, 117, 202, 47, 178, 93, 87, 84, 106, 53, 135, 54, 174, 224, 212, 129, 157, 46, 151, 195, 126, 7, 165, 61, 157, 66, 53, 241, 116, 5, 208, 128, 88, 42, 132, 42, 232, 206, 119, 230, 138, 176, 184, 246, 85, 38, 248, 254, 217, 187, 240, 81, 187, 88, 180, 112, 1, 90, 109, 253, 47, 163, 119, 232, 208, 129, 226, 226, 98, 190, 126, 247, 85, 122, 222, 120, 247, 233, 157, 229, 218, 19, 151, 208, 159, 130, 204, 84, 78, 36, 110, 81, 81, 122, 102, 249, 107, 20, 243, 135, 106, 227, 233, 59, 193, 227, 64, 182, 74, 24, 53, 158, 86, 61, 7, 85, 25, 249, 232, 214, 53, 236, 94, 185, 128, 183, 222, 124, 3, 163, 177, 242, 97, 82, 93, 152, 54, 109, 26, 102, 163, 129, 101, 31, 190, 92, 47, 249, 221, 248, 236, 123, 180, 190, 124, 8, 192, 16, 96, 90, 189, 100, 122, 137, 224, 201, 171, 64, 26, 96, 159, 74, 173, 142, 123, 122, 225, 62, 17, 24, 21, 87, 101, 130, 217, 119, 95, 73, 184, 17, 86, 174, 88, 222, 224, 194, 205, 156, 57, 147, 135, 31, 121, 148, 169, 203, 146, 240, 10, 8, 169, 115, 126, 69, 57, 25, 188, 121, 83, 188, 204, 207, 56, 37, 129, 193, 192, 170, 58, 103, 122, 9, 224, 201, 61, 192, 141, 64, 203, 132, 107, 199, 87, 171, 242, 23, 231, 101, 115, 120, 203, 106, 238, 187, 247, 158, 134, 151, 12, 184, 243, 206, 59, 209, 168, 213, 236, 89, 181, 176, 94, 242, 243, 10, 8, 225, 182, 215, 231, 8, 132, 80, 1, 211, 241, 236, 223, 182, 222, 240, 100, 91, 160, 59, 1, 250, 223, 86, 189, 189, 160, 61, 43, 23, 160, 86, 171, 185, 250, 234, 171, 27, 84, 168, 191, 49, 153, 76, 12, 30, 60, 152, 61, 43, 231, 209, 243, 134, 187, 106, 156, 62, 47, 61, 133, 195, 27, 87, 146, 115, 234, 56, 37, 121, 217, 152, 252, 2, 9, 136, 104, 78, 251, 126, 215, 176, 119, 245, 226, 203, 128, 177, 40, 150, 160, 85, 226, 169, 10, 208, 12, 24, 210, 172, 109, 23, 34, 218, 116, 169, 86, 130, 131, 27, 150, 50, 112, 192, 64, 188, 189, 189, 171, 142, 92, 79, 140, 28, 57, 130, 7, 30, 124, 8, 183, 219, 85, 237, 205, 177, 253, 235, 254, 96, 249, 172, 23, 57, 186, 99, 3, 65, 33, 161, 196, 198, 197, 17, 18, 20, 68, 198, 209, 173, 108, 253, 241, 8, 89, 25, 233, 232, 245, 122, 108, 54, 219, 243, 40, 10, 80, 37, 158, 218, 77, 222, 6, 168, 187, 143, 154, 80, 237, 4, 89, 199, 246, 211, 165, 75, 231, 134, 147, 168, 2, 58, 119, 238, 140, 221, 106, 33, 247, 100, 82, 149, 113, 173, 69, 5, 124, 254, 224, 245, 124, 114, 223, 72, 174, 232, 222, 145, 141, 27, 55, 146, 158, 122, 138, 191, 214, 173, 101, 254, 188, 95, 249, 107, 221, 90, 210, 83, 79, 177, 113, 227, 70, 198, 143, 31, 143, 74, 165, 138, 213, 235, 245, 171, 81, 188, 217, 85, 138, 167, 246, 0, 163, 132, 80, 209, 117, 216, 173, 213, 78, 80, 144, 157, 78, 68, 68, 68, 189, 20, 190, 99, 199, 14, 150, 44, 89, 130, 197, 98, 161, 107, 215, 174, 140, 24, 49, 2, 181, 186, 124, 11, 255, 119, 121, 133, 217, 233, 84, 54, 79, 41, 200, 56, 197, 199, 83, 134, 97, 86, 187, 216, 185, 99, 7, 29, 58, 116, 168, 48, 158, 16, 130, 238, 221, 187, 211, 189, 123, 119, 30, 124, 240, 65, 174, 191, 254, 250, 222, 41, 41, 41, 155, 173, 86, 235, 32, 224, 100, 189, 188, 156, 135, 225, 137, 10, 224, 15, 244, 104, 214, 238, 50, 188, 3, 67, 171, 157, 200, 97, 181, 96, 48, 212, 205, 116, 217, 225, 112, 112, 255, 253, 247, 51, 103, 206, 28, 28, 14, 7, 46, 151, 11, 179, 217, 76, 108, 108, 44, 243, 230, 205, 163, 121, 243, 230, 101, 226, 155, 205, 165, 182, 72, 118, 139, 165, 162, 236, 0, 112, 218, 172, 124, 254, 224, 13, 52, 11, 244, 230, 183, 69, 11, 241, 247, 247, 63, 111, 220, 115, 233, 216, 177, 35, 155, 55, 111, 86, 15, 29, 58, 52, 102, 199, 142, 29, 139, 109, 54, 91, 79, 192, 90, 235, 151, 243, 80, 60, 113, 8, 52, 4, 80, 183, 233, 115, 85, 141, 18, 233, 140, 38, 44, 149, 84, 196, 234, 240, 252, 243, 207, 51, 103, 206, 28, 74, 74, 74, 112, 56, 28, 184, 221, 110, 10, 11, 11, 73, 76, 76, 100, 216, 176, 97, 56, 157, 101, 143, 248, 22, 23, 23, 159, 46, 251, 252, 123, 14, 11, 223, 122, 18, 91, 78, 42, 243, 126, 153, 91, 237, 202, 255, 55, 126, 126, 126, 204, 159, 63, 95, 235, 239, 239, 223, 86, 171, 213, 190, 85, 243, 55, 242, 124, 60, 81, 1, 174, 2, 104, 211, 251, 202, 26, 37, 242, 9, 10, 227, 228, 201, 218, 143, 18, 178, 179, 179, 121, 227, 141, 55, 40, 41, 41, 41, 247, 204, 229, 114, 145, 156, 156, 204, 119, 223, 125, 87, 38, 252, 239, 242, 124, 130, 195, 43, 204, 51, 43, 249, 48, 235, 127, 152, 205, 199, 31, 206, 38, 52, 180, 250, 189, 217, 185, 132, 132, 132, 240, 217, 103, 159, 233, 92, 46, 215, 100, 160, 85, 173, 50, 241, 96, 60, 81, 1, 122, 107, 116, 122, 217, 226, 178, 222, 53, 74, 20, 212, 162, 45, 219, 119, 236, 172, 117, 161, 219, 182, 109, 171, 116, 8, 85, 84, 84, 196, 170, 85, 171, 202, 132, 237, 220, 185, 19, 189, 201, 140, 127, 120, 243, 10, 211, 172, 250, 236, 13, 18, 186, 247, 96, 216, 176, 97, 181, 150, 11, 96, 232, 208, 161, 244, 232, 209, 195, 169, 213, 106, 159, 174, 83, 70, 30, 136, 167, 41, 128, 23, 208, 38, 162, 77, 23, 161, 214, 150, 115, 244, 86, 41, 109, 251, 14, 101, 245, 234, 63, 201, 203, 171, 157, 9, 125, 85, 118, 67, 66, 8, 116, 186, 178, 50, 205, 155, 55, 159, 54, 189, 175, 172, 208, 52, 90, 186, 221, 236, 93, 181, 128, 127, 77, 172, 249, 30, 65, 69, 76, 154, 52, 73, 39, 132, 184, 1, 207, 251, 205, 235, 132, 167, 125, 25, 221, 0, 117, 116, 199, 154, 31, 92, 105, 63, 96, 56, 72, 248, 253, 247, 223, 107, 85, 112, 66, 66, 2, 14, 199, 249, 29, 70, 155, 205, 102, 174, 185, 230, 172, 9, 116, 81, 81, 17, 43, 87, 173, 164, 195, 160, 81, 21, 198, 63, 117, 96, 39, 121, 153, 105, 101, 210, 212, 133, 107, 174, 185, 6, 135, 195, 225, 75, 169, 35, 0, 133, 211, 120, 154, 2, 36, 0, 68, 182, 143, 175, 113, 66, 163, 183, 31, 173, 123, 93, 193, 123, 51, 223, 175, 85, 193, 94, 94, 94, 188, 249, 230, 155, 103, 86, 118, 206, 69, 175, 215, 147, 144, 144, 192, 200, 145, 35, 207, 132, 125, 244, 209, 71, 72, 4, 237, 251, 87, 60, 188, 201, 62, 153, 132, 143, 175, 31, 225, 225, 21, 207, 15, 106, 74, 88, 88, 24, 102, 179, 217, 1, 196, 212, 75, 134, 30, 130, 167, 41, 64, 103, 128, 102, 237, 46, 171, 85, 226, 171, 238, 123, 142, 245, 235, 214, 50, 127, 254, 252, 90, 165, 159, 60, 121, 50, 255, 251, 223, 255, 48, 153, 76, 248, 248, 248, 224, 237, 237, 141, 201, 100, 98, 228, 200, 145, 252, 242, 203, 47, 103, 226, 229, 229, 229, 241, 226, 255, 94, 162, 223, 109, 255, 198, 236, 95, 222, 221, 10, 148, 26, 183, 5, 133, 212, 221, 72, 238, 92, 66, 66, 66, 156, 148, 58, 241, 85, 56, 141, 167, 41, 64, 75, 132, 32, 184, 121, 237, 22, 59, 162, 59, 245, 160, 203, 149, 55, 240, 248, 19, 79, 158, 89, 162, 172, 9, 66, 8, 30, 122, 232, 33, 146, 146, 146, 248, 244, 211, 79, 153, 57, 115, 38, 155, 55, 111, 230, 199, 31, 127, 196, 207, 207, 239, 76, 188, 169, 83, 167, 226, 148, 112, 197, 196, 39, 206, 155, 151, 70, 171, 199, 110, 179, 213, 168, 252, 21, 43, 86, 208, 190, 125, 123, 194, 195, 195, 105, 217, 178, 37, 223, 124, 243, 77, 153, 231, 86, 171, 85, 160, 236, 5, 148, 193, 211, 54, 194, 98, 189, 3, 67, 234, 228, 130, 100, 228, 227, 111, 240, 246, 45, 61, 24, 63, 254, 118, 126, 254, 249, 167, 90, 157, 8, 11, 14, 14, 230, 198, 27, 111, 172, 240, 217, 151, 95, 126, 201, 204, 153, 51, 25, 251, 202, 23, 24, 189, 253, 42, 140, 3, 224, 29, 20, 70, 102, 70, 58, 82, 202, 106, 201, 176, 113, 227, 70, 70, 143, 30, 77, 118, 118, 246, 153, 176, 251, 238, 187, 15, 33, 4, 99, 198, 148, 122, 163, 200, 206, 206, 214, 2, 233, 53, 123, 27, 207, 198, 147, 122, 0, 19, 16, 22, 20, 213, 178, 78, 153, 248, 135, 71, 115, 251, 244, 31, 89, 176, 112, 33, 79, 62, 249, 36, 82, 214, 223, 225, 170, 85, 171, 86, 49, 105, 242, 100, 6, 222, 241, 40, 241, 35, 198, 85, 26, 55, 172, 101, 123, 108, 86, 43, 187, 118, 237, 170, 86, 222, 143, 63, 254, 120, 153, 202, 15, 165, 67, 173, 105, 211, 74, 207, 199, 236, 222, 189, 27, 155, 205, 166, 6, 18, 107, 37, 188, 135, 226, 73, 61, 64, 44, 32, 2, 35, 203, 187, 57, 175, 113, 70, 241, 253, 24, 253, 194, 71, 188, 53, 117, 34, 201, 39, 78, 240, 217, 167, 159, 214, 249, 132, 216, 135, 31, 126, 200, 253, 15, 60, 64, 167, 193, 215, 49, 252, 225, 151, 177, 21, 23, 178, 113, 238, 167, 236, 248, 253, 7, 172, 69, 5, 152, 3, 130, 105, 221, 243, 10, 46, 27, 58, 154, 224, 152, 214, 248, 135, 71, 19, 217, 166, 19, 11, 22, 44, 160, 75, 151, 170, 45, 90, 83, 82, 82, 42, 12, 255, 123, 40, 183, 96, 193, 2, 140, 70, 227, 17, 139, 197, 146, 84, 167, 23, 241, 48, 60, 169, 7, 136, 4, 240, 143, 168, 120, 83, 169, 166, 36, 140, 188, 141, 73, 179, 127, 103, 209, 239, 127, 208, 235, 242, 222, 172, 94, 189, 186, 86, 249, 164, 164, 164, 48, 254, 246, 219, 153, 50, 101, 10, 131, 238, 122, 130, 113, 175, 127, 131, 219, 233, 228, 141, 27, 187, 177, 232, 237, 103, 72, 218, 249, 23, 105, 71, 246, 114, 100, 243, 159, 252, 254, 254, 243, 188, 61, 174, 55, 239, 223, 53, 132, 188, 244, 20, 186, 92, 51, 134, 153, 31, 204, 170, 150, 137, 198, 249, 12, 228, 130, 130, 130, 176, 90, 173, 188, 243, 206, 59, 118, 139, 197, 242, 73, 173, 94, 194, 131, 241, 36, 5, 8, 2, 42, 116, 109, 88, 91, 90, 245, 28, 196, 131, 223, 254, 133, 205, 20, 200, 128, 1, 3, 24, 49, 242, 90, 86, 173, 90, 133, 203, 85, 245, 53, 193, 7, 15, 30, 228, 169, 167, 158, 162, 85, 235, 214, 252, 190, 98, 53, 119, 204, 248, 153, 161, 247, 63, 79, 234, 193, 93, 252, 240, 223, 73, 228, 165, 157, 192, 97, 45, 91, 177, 165, 219, 77, 73, 126, 46, 71, 54, 175, 226, 173, 155, 187, 19, 219, 173, 47, 22, 155, 131, 233, 211, 167, 87, 89, 222, 59, 239, 188, 67, 88, 88, 217, 5, 158, 160, 160, 32, 102, 205, 154, 197, 244, 233, 211, 201, 205, 205, 45, 1, 222, 173, 209, 23, 112, 9, 224, 73, 103, 130, 31, 2, 166, 143, 123, 245, 43, 226, 71, 140, 173, 247, 204, 15, 172, 95, 202, 239, 111, 63, 195, 241, 61, 91, 241, 15, 8, 100, 228, 136, 225, 36, 36, 36, 16, 29, 29, 141, 217, 108, 198, 110, 183, 147, 150, 150, 198, 222, 189, 123, 89, 176, 104, 49, 7, 247, 239, 195, 39, 32, 152, 129, 119, 61, 65, 223, 177, 247, 97, 41, 204, 227, 167, 105, 19, 57, 176, 97, 57, 157, 219, 183, 32, 113, 223, 81, 172, 182, 202, 111, 90, 245, 15, 143, 102, 232, 253, 207, 243, 227, 115, 147, 88, 190, 108, 25, 253, 251, 247, 175, 52, 126, 82, 82, 18, 15, 61, 244, 16, 137, 137, 137, 196, 196, 196, 240, 198, 27, 111, 80, 84, 84, 196, 192, 129, 3, 221, 46, 151, 107, 2, 202, 1, 153, 114, 120, 146, 2, 188, 8, 60, 59, 105, 214, 98, 218, 246, 109, 184, 99, 141, 153, 73, 7, 79, 123, 135, 94, 76, 234, 161, 68, 10, 115, 207, 78, 60, 181, 122, 3, 97, 45, 218, 16, 219, 99, 16, 29, 175, 184, 150, 22, 221, 250, 162, 82, 169, 41, 206, 205, 226, 237, 91, 187, 211, 57, 38, 128, 103, 239, 29, 69, 112, 128, 15, 15, 190, 240, 25, 235, 182, 30, 192, 102, 63, 191, 18, 232, 140, 102, 238, 152, 241, 19, 251, 254, 92, 200, 238, 223, 190, 229, 215, 95, 230, 50, 96, 192, 128, 243, 198, 255, 39, 171, 87, 175, 102, 196, 136, 17, 142, 146, 146, 146, 79, 92, 46, 215, 133, 57, 236, 124, 145, 225, 73, 10, 240, 62, 112, 207, 67, 223, 253, 69, 116, 199, 238, 23, 172, 80, 135, 205, 130, 211, 102, 69, 168, 84, 24, 188, 42, 62, 124, 245, 249, 3, 215, 226, 103, 57, 193, 187, 211, 110, 71, 117, 122, 73, 211, 102, 119, 242, 248, 43, 95, 178, 101, 247, 17, 10, 139, 43, 94, 154, 23, 66, 208, 103, 204, 61, 92, 247, 212, 12, 230, 191, 254, 40, 127, 253, 248, 33, 255, 157, 54, 141, 71, 30, 121, 164, 82, 195, 59, 171, 213, 202, 244, 233, 211, 153, 54, 109, 154, 75, 74, 249, 129, 203, 229, 122, 8, 168, 122, 220, 118, 9, 226, 73, 142, 177, 198, 0, 157, 7, 221, 241, 104, 173, 238, 237, 170, 45, 106, 141, 22, 173, 222, 136, 70, 87, 113, 133, 148, 110, 55, 95, 63, 61, 129, 215, 159, 28, 75, 72, 192, 89, 175, 116, 26, 181, 138, 107, 6, 118, 37, 54, 42, 148, 3, 199, 82, 145, 18, 52, 167, 79, 141, 233, 117, 26, 76, 70, 61, 87, 92, 222, 137, 245, 203, 150, 18, 209, 182, 43, 253, 110, 251, 55, 190, 97, 81, 124, 62, 253, 37, 102, 205, 250, 128, 236, 172, 44, 188, 188, 188, 208, 235, 245, 152, 205, 102, 50, 51, 51, 73, 76, 76, 228, 131, 15, 62, 96, 194, 132, 9, 252, 254, 251, 239, 5, 14, 135, 227, 30, 41, 229, 43, 40, 142, 178, 206, 139, 39, 245, 0, 223, 3, 163, 159, 93, 114, 132, 128, 102, 49, 141, 45, 75, 25, 126, 154, 54, 145, 67, 107, 22, 48, 102, 88, 79, 250, 38, 180, 161, 85, 76, 56, 26, 77, 217, 182, 39, 51, 167, 128, 61, 135, 78, 144, 155, 95, 140, 151, 201, 64, 66, 231, 56, 252, 125, 204, 124, 61, 111, 13, 95, 46, 221, 199, 99, 243, 74, 175, 108, 181, 22, 21, 176, 225, 199, 15, 73, 252, 227, 39, 146, 18, 183, 148, 217, 167, 16, 66, 96, 52, 153, 41, 41, 41, 182, 34, 101, 4, 144, 123, 33, 223, 243, 98, 196, 147, 20, 224, 103, 224, 134, 105, 203, 143, 227, 23, 26, 217, 216, 178, 148, 65, 186, 221, 108, 91, 252, 29, 219, 230, 125, 198, 145, 109, 235, 144, 110, 23, 177, 81, 161, 116, 108, 29, 69, 187, 150, 205, 232, 214, 161, 5, 173, 98, 42, 54, 122, 43, 182, 216, 184, 252, 198, 103, 120, 118, 201, 145, 114, 75, 188, 150, 194, 60, 242, 210, 82, 40, 202, 206, 192, 43, 48, 4, 135, 205, 194, 140, 91, 123, 65, 169, 27, 248, 155, 128, 56, 32, 0, 216, 135, 162, 12, 21, 226, 73, 27, 97, 26, 0, 117, 5, 151, 198, 53, 54, 66, 165, 34, 126, 196, 88, 226, 71, 140, 197, 237, 114, 242, 201, 148, 97, 116, 140, 12, 160, 89, 179, 8, 254, 220, 178, 153, 215, 63, 92, 64, 68, 104, 0, 183, 14, 239, 197, 117, 87, 245, 192, 108, 60, 123, 207, 152, 201, 160, 67, 171, 213, 98, 45, 46, 40, 151, 175, 209, 219, 175, 212, 156, 226, 180, 233, 211, 186, 239, 62, 0, 64, 107, 48, 134, 107, 180, 250, 140, 160, 232, 150, 46, 179, 95, 0, 201, 187, 55, 171, 74, 10, 114, 251, 2, 123, 203, 101, 114, 137, 211, 244, 106, 75, 237, 209, 0, 168, 52, 77, 251, 94, 56, 149, 90, 131, 70, 171, 161, 111, 223, 190, 252, 251, 223, 165, 78, 187, 138, 138, 138, 248, 254, 251, 239, 153, 49, 253, 45, 222, 253, 242, 119, 110, 25, 222, 155, 49, 35, 251, 16, 22, 236, 199, 246, 61, 199, 80, 169, 53, 149, 122, 141, 248, 155, 67, 27, 87, 0, 112, 229, 228, 103, 18, 250, 143, 127, 80, 252, 237, 120, 247, 147, 251, 175, 205, 223, 179, 106, 81, 44, 138, 2, 148, 195, 243, 20, 160, 2, 247, 35, 77, 143, 178, 35, 79, 47, 47, 47, 38, 78, 156, 200, 192, 129, 3, 233, 208, 161, 35, 43, 246, 229, 241, 217, 207, 47, 17, 30, 22, 66, 110, 94, 1, 87, 221, 247, 223, 106, 121, 145, 62, 178, 101, 53, 222, 129, 161, 12, 190, 251, 41, 241, 183, 1, 221, 202, 207, 222, 180, 29, 219, 190, 225, 56, 176, 184, 33, 222, 228, 98, 199, 147, 20, 192, 14, 165, 183, 36, 54, 117, 84, 26, 45, 243, 23, 44, 160, 101, 203, 150, 92, 113, 197, 21, 24, 12, 6, 236, 118, 59, 183, 142, 187, 141, 174, 195, 110, 101, 244, 139, 159, 80, 152, 149, 70, 234, 161, 61, 120, 5, 6, 19, 209, 186, 106, 135, 93, 217, 41, 71, 41, 206, 205, 162, 211, 224, 81, 8, 33, 72, 217, 187, 149, 185, 175, 60, 152, 159, 121, 236, 224, 246, 146, 252, 156, 107, 1, 119, 195, 191, 217, 197, 135, 39, 41, 64, 9, 128, 221, 82, 222, 43, 67, 83, 163, 239, 237, 143, 176, 109, 225, 28, 198, 221, 49, 17, 91, 113, 33, 131, 135, 92, 137, 10, 73, 90, 110, 17, 247, 189, 251, 54, 80, 106, 14, 237, 29, 84, 253, 179, 43, 39, 18, 183, 0, 96, 179, 148, 184, 94, 188, 178, 69, 161, 195, 106, 73, 177, 20, 228, 60, 235, 114, 185, 230, 53, 200, 75, 120, 8, 30, 167, 0, 14, 107, 211, 87, 128, 184, 132, 254, 196, 37, 244, 71, 186, 221, 156, 216, 179, 133, 189, 171, 22, 114, 116, 203, 159, 76, 120, 123, 46, 122, 147, 87, 173, 242, 76, 78, 220, 12, 192, 209, 173, 107, 231, 57, 109, 150, 71, 128, 227, 245, 40, 178, 199, 226, 121, 10, 96, 171, 155, 115, 171, 11, 137, 80, 169, 206, 220, 12, 95, 87, 146, 119, 111, 6, 144, 78, 155, 101, 18, 144, 93, 69, 116, 133, 211, 120, 146, 53, 232, 69, 51, 4, 106, 8, 78, 238, 223, 33, 129, 163, 40, 149, 191, 70, 120, 146, 2, 20, 66, 233, 78, 233, 165, 70, 73, 126, 14, 182, 226, 66, 1, 28, 105, 108, 89, 46, 54, 60, 73, 1, 210, 0, 10, 178, 210, 26, 91, 142, 11, 78, 238, 169, 51, 195, 125, 101, 220, 95, 67, 60, 73, 1, 82, 1, 10, 50, 83, 27, 91, 142, 11, 142, 165, 48, 255, 239, 143, 89, 141, 41, 199, 197, 136, 199, 41, 64, 214, 241, 195, 124, 63, 245, 110, 102, 253, 235, 170, 122, 61, 208, 222, 148, 113, 57, 207, 156, 41, 112, 86, 22, 79, 161, 60, 158, 178, 10, 100, 4, 122, 1, 108, 93, 56, 7, 41, 101, 189, 172, 172, 52, 53, 118, 47, 255, 149, 156, 148, 99, 64, 233, 10, 146, 211, 110, 67, 165, 214, 144, 151, 126, 250, 64, 188, 16, 79, 33, 101, 48, 240, 50, 112, 162, 209, 4, 189, 136, 184, 216, 173, 65, 53, 192, 100, 33, 196, 11, 82, 202, 0, 128, 240, 86, 29, 185, 124, 244, 100, 46, 191, 249, 95, 168, 155, 184, 93, 80, 77, 177, 22, 21, 80, 209, 77, 247, 133, 153, 105, 252, 254, 254, 243, 108, 255, 237, 59, 105, 45, 42, 16, 66, 8, 187, 148, 114, 58, 165, 119, 6, 55, 253, 173, 241, 70, 228, 98, 86, 128, 8, 74, 205, 126, 227, 117, 6, 147, 188, 252, 150, 201, 162, 251, 181, 227, 171, 125, 41, 158, 39, 226, 180, 89, 89, 251, 221, 7, 252, 249, 197, 155, 50, 63, 35, 85, 0, 219, 128, 113, 192, 254, 70, 22, 173, 201, 114, 177, 42, 64, 115, 132, 88, 131, 148, 81, 151, 93, 125, 51, 163, 158, 124, 19, 223, 144, 102, 141, 45, 83, 147, 193, 90, 148, 207, 79, 47, 220, 203, 182, 197, 223, 33, 132, 176, 72, 41, 239, 2, 190, 171, 50, 225, 37, 200, 197, 96, 58, 89, 17, 115, 128, 238, 87, 77, 121, 150, 27, 167, 206, 196, 96, 62, 59, 44, 40, 202, 205, 100, 237, 55, 51, 89, 252, 206, 84, 126, 126, 241, 62, 10, 179, 82, 105, 119, 30, 15, 204, 158, 138, 70, 103, 160, 243, 149, 55, 18, 24, 21, 199, 129, 117, 127, 104, 92, 78, 231, 13, 64, 10, 176, 189, 177, 101, 107, 106, 92, 140, 147, 224, 24, 96, 68, 108, 124, 63, 134, 222, 255, 60, 80, 122, 226, 234, 208, 198, 149, 108, 158, 247, 5, 59, 151, 252, 136, 211, 97, 71, 8, 65, 72, 108, 59, 98, 106, 120, 83, 140, 39, 145, 48, 242, 54, 130, 162, 227, 196, 135, 147, 135, 99, 45, 202, 255, 24, 208, 3, 31, 52, 182, 92, 77, 137, 139, 113, 8, 228, 143, 16, 233, 58, 131, 73, 19, 213, 33, 65, 184, 93, 78, 210, 142, 236, 145, 150, 130, 188, 211, 239, 34, 136, 77, 232, 203, 132, 55, 191, 175, 209, 45, 145, 158, 204, 201, 125, 219, 153, 245, 175, 171, 101, 113, 94, 182, 4, 174, 5, 22, 53, 182, 76, 77, 133, 139, 113, 31, 32, 23, 41, 111, 179, 91, 75, 142, 31, 217, 242, 167, 60, 182, 125, 157, 219, 82, 144, 119, 12, 248, 2, 184, 5, 36, 122, 163, 89, 169, 252, 231, 208, 172, 93, 87, 38, 190, 247, 171, 80, 107, 180, 66, 8, 49, 135, 210, 179, 194, 10, 92, 156, 61, 192, 185, 232, 41, 125, 135, 115, 28, 235, 136, 116, 179, 95, 64, 240, 11, 107, 210, 69, 109, 92, 155, 123, 50, 235, 191, 159, 197, 79, 47, 222, 7, 176, 3, 232, 1, 84, 238, 154, 174, 97, 208, 0, 29, 129, 142, 90, 47, 109, 119, 181, 78, 221, 65, 186, 100, 132, 219, 229, 14, 148, 110, 169, 151, 46, 169, 17, 106, 225, 66, 96, 87, 107, 212, 5, 66, 37, 146, 157, 86, 231, 86, 167, 197, 185, 14, 88, 14, 212, 171, 177, 151, 39, 214, 144, 95, 128, 235, 158, 94, 180, 191, 214, 23, 101, 120, 50, 115, 158, 24, 199, 182, 197, 223, 1, 60, 14, 188, 113, 129, 138, 237, 164, 49, 104, 110, 84, 25, 52, 35, 220, 54, 103, 107, 239, 104, 111, 187, 95, 43, 127, 131, 111, 172, 159, 217, 20, 110, 198, 224, 111, 64, 239, 111, 64, 173, 59, 187, 38, 227, 118, 184, 112, 20, 57, 40, 73, 47, 161, 240, 68, 1, 217, 123, 178, 243, 51, 182, 167, 11, 33, 56, 96, 207, 183, 191, 72, 233, 48, 174, 206, 167, 220, 60, 81, 1, 158, 0, 94, 27, 243, 210, 103, 116, 31, 117, 123, 99, 203, 210, 228, 40, 202, 201, 224, 149, 225, 237, 164, 165, 40, 223, 138, 148, 237, 104, 56, 3, 186, 214, 90, 179, 246, 30, 9, 99, 76, 193, 70, 109, 100, 255, 104, 159, 192, 46, 193, 26, 159, 24, 31, 132, 170, 150, 213, 78, 66, 206, 254, 108, 14, 253, 112, 32, 47, 239, 80, 94, 158, 163, 216, 62, 26, 216, 92, 23, 33, 47, 214, 101, 208, 202, 80, 1, 119, 122, 7, 133, 210, 97, 192, 136, 198, 150, 165, 201, 161, 51, 154, 209, 155, 189, 197, 190, 213, 139, 181, 128, 63, 240, 107, 61, 102, 175, 86, 171, 213, 163, 181, 62, 218, 239, 12, 126, 134, 71, 227, 110, 104, 221, 175, 243, 61, 151, 121, 199, 141, 106, 101, 244, 111, 23, 160, 50, 248, 27, 106, 117, 227, 206, 25, 4, 24, 131, 77, 68, 14, 140, 50, 4, 118, 8, 244, 203, 222, 153, 57, 90, 170, 165, 202, 109, 119, 175, 169, 67, 150, 30, 135, 17, 200, 143, 104, 221, 89, 251, 216, 92, 101, 217, 187, 34, 220, 110, 23, 175, 12, 111, 75, 118, 202, 49, 39, 82, 182, 164, 238, 189, 128, 70, 173, 85, 223, 37, 116, 234, 231, 2, 59, 4, 152, 90, 143, 110, 235, 235, 27, 119, 254, 235, 159, 234, 11, 151, 213, 197, 134, 105, 107, 11, 75, 82, 139, 223, 178, 23, 217, 159, 171, 77, 30, 158, 216, 3, 56, 129, 17, 69, 185, 89, 205, 250, 223, 246, 0, 90, 253, 249, 157, 200, 94, 170, 8, 161, 66, 171, 51, 176, 247, 207, 133, 42, 74, 123, 204, 218, 93, 142, 12, 160, 102, 152, 214, 164, 91, 26, 210, 61, 236, 218, 248, 199, 18, 2, 99, 134, 198, 26, 12, 1, 23, 230, 59, 87, 105, 84, 52, 235, 31, 165, 79, 89, 145, 220, 213, 81, 228, 88, 77, 45, 12, 0, 47, 198, 101, 208, 234, 176, 76, 186, 93, 28, 88, 247, 71, 99, 203, 209, 100, 233, 62, 234, 118, 188, 3, 67, 37, 66, 76, 4, 106, 115, 171, 96, 168, 214, 91, 187, 194, 47, 206, 255, 155, 203, 255, 215, 39, 170, 219, 35, 9, 222, 166, 208, 218, 95, 78, 88, 91, 212, 122, 53, 93, 31, 237, 238, 171, 53, 235, 190, 165, 22, 13, 186, 39, 246, 0, 80, 186, 44, 58, 81, 171, 55, 210, 105, 200, 245, 141, 45, 75, 147, 68, 165, 209, 80, 148, 157, 46, 146, 118, 172, 215, 1, 123, 128, 221, 213, 77, 171, 214, 169, 111, 209, 24, 53, 139, 219, 79, 232, 216, 186, 211, 228, 46, 38, 189, 95, 227, 246, 178, 134, 0, 3, 121, 135, 114, 133, 37, 181, 228, 184, 148, 178, 218, 239, 1, 158, 219, 3, 108, 68, 136, 236, 125, 107, 126, 147, 110, 183, 226, 22, 255, 124, 116, 59, 123, 147, 206, 109, 213, 76, 162, 209, 154, 180, 239, 155, 195, 205, 179, 251, 191, 117, 69, 64, 212, 144, 230, 13, 110, 74, 227, 40, 118, 224, 40, 178, 227, 44, 169, 252, 172, 79, 155, 49, 237, 124, 212, 6, 245, 75, 212, 112, 94, 123, 49, 218, 2, 85, 7, 23, 82, 46, 46, 206, 203, 30, 159, 188, 107, 227, 37, 109, 15, 84, 25, 205, 218, 94, 70, 104, 108, 59, 210, 143, 237, 31, 130, 148, 94, 64, 81, 37, 209, 67, 180, 94, 186, 223, 194, 123, 133, 183, 233, 56, 185, 139, 89, 165, 169, 255, 182, 211, 89, 226, 36, 115, 71, 6, 249, 7, 138, 200, 59, 144, 143, 53, 167, 4, 163, 191, 31, 26, 157, 14, 135, 197, 138, 37, 191, 0, 99, 144, 25, 255, 54, 190, 4, 247, 8, 192, 191, 93, 192, 153, 85, 37, 239, 104, 111, 124, 91, 249, 249, 101, 238, 206, 24, 133, 171, 250, 43, 91, 158, 170, 0, 80, 186, 81, 50, 126, 239, 234, 223, 20, 5, 168, 132, 182, 125, 175, 38, 253, 232, 62, 45, 48, 128, 243, 219, 8, 181, 211, 152, 117, 43, 58, 220, 213, 49, 40, 114, 80, 116, 189, 214, 25, 233, 150, 164, 111, 74, 35, 109, 77, 14, 5, 73, 249, 116, 24, 50, 148, 1, 163, 174, 161, 101, 239, 126, 248, 133, 151, 55, 113, 207, 61, 149, 194, 254, 149, 75, 217, 244, 203, 87, 236, 126, 111, 45, 205, 6, 69, 200, 152, 225, 209, 66, 99, 210, 210, 102, 92, 123, 223, 252, 23, 242, 94, 179, 23, 218, 171, 173, 0, 77, 97, 25, 52, 82, 99, 212, 60, 44, 180, 170, 190, 184, 100, 40, 42, 10, 65, 149, 229, 178, 185, 86, 187, 237, 206, 37, 148, 110, 116, 212, 102, 203, 222, 15, 68, 102, 104, 108, 27, 205, 147, 243, 247, 212, 179, 200, 158, 195, 254, 181, 75, 248, 112, 202, 48, 128, 183, 41, 189, 104, 240, 159, 244, 211, 122, 107, 231, 117, 127, 170, 151, 95, 64, 251, 192, 122, 171, 47, 78, 171, 147, 228, 37, 201, 238, 147, 203, 82, 85, 45, 226, 47, 103, 208, 221, 15, 209, 178, 119, 255, 26, 57, 55, 46, 206, 205, 97, 217, 251, 255, 199, 250, 175, 63, 33, 114, 72, 168, 163, 197, 181, 113, 218, 245, 83, 215, 228, 229, 238, 203, 25, 13, 44, 173, 78, 30, 141, 170, 0, 106, 181, 122, 180, 202, 168, 154, 29, 49, 36, 202, 203, 39, 206, 71, 163, 247, 55, 224, 44, 113, 226, 40, 114, 80, 116, 188, 192, 149, 187, 39, 39, 191, 56, 165, 72, 43, 180, 234, 117, 206, 66, 251, 27, 192, 10, 106, 118, 221, 207, 175, 192, 168, 7, 191, 89, 79, 243, 206, 61, 27, 230, 37, 46, 114, 108, 37, 69, 252, 167, 87, 0, 210, 237, 218, 4, 148, 249, 146, 212, 106, 245, 205, 26, 111, 205, 135, 189, 158, 239, 235, 231, 29, 237, 93, 47, 229, 185, 157, 110, 146, 22, 37, 185, 83, 254, 56, 165, 234, 126, 227, 109, 92, 117, 255, 147, 248, 132, 86, 223, 7, 106, 69, 148, 228, 229, 242, 211, 212, 135, 220, 123, 87, 47, 180, 68, 15, 141, 50, 31, 248, 118, 223, 14, 123, 129, 189, 107, 117, 210, 54, 166, 2, 180, 86, 235, 213, 219, 47, 123, 58, 222, 164, 245, 209, 157, 55, 146, 116, 75, 242, 246, 229, 146, 186, 234, 100, 110, 241, 169, 98, 155, 116, 241, 190, 219, 238, 252, 16, 72, 175, 70, 25, 215, 3, 115, 123, 143, 158, 204, 77, 211, 222, 175, 47, 185, 61, 142, 215, 70, 117, 34, 253, 232, 62, 27, 82, 122, 115, 186, 183, 85, 27, 212, 119, 234, 125, 12, 51, 122, 191, 220, 207, 199, 16, 88, 63, 171, 60, 89, 59, 51, 57, 240, 197, 97, 217, 105, 240, 181, 98, 228, 83, 47, 225, 21, 84, 127, 119, 58, 3, 28, 248, 115, 185, 115, 214, 157, 215, 89, 236, 249, 37, 102, 183, 203, 221, 31, 88, 87, 85, 154, 70, 83, 0, 181, 70, 253, 179, 33, 220, 120, 93, 167, 135, 47, 171, 246, 108, 202, 158, 111, 39, 99, 67, 154, 45, 109, 125, 170, 21, 41, 214, 56, 139, 237, 47, 3, 27, 42, 73, 162, 69, 136, 83, 6, 47, 159, 192, 231, 87, 157, 20, 90, 189, 177, 238, 130, 123, 32, 95, 63, 117, 59, 91, 23, 126, 13, 208, 25, 216, 173, 210, 107, 239, 53, 6, 232, 95, 233, 253, 74, 63, 31, 189, 175, 190, 138, 212, 85, 99, 205, 182, 178, 255, 139, 3, 46, 85, 137, 175, 122, 252, 140, 207, 136, 236, 116, 89, 157, 243, 60, 31, 121, 167, 78, 57, 95, 26, 212, 57, 211, 86, 92, 112, 216, 81, 228, 168, 252, 98, 101, 26, 111, 25, 84, 37, 52, 98, 168, 90, 171, 174, 145, 2, 234, 124, 117, 68, 14, 141, 214, 199, 63, 215, 195, 55, 246, 166, 216, 225, 94, 81, 94, 139, 180, 38, 237, 17, 181, 90, 125, 55, 165, 38, 16, 255, 196, 129, 148, 115, 172, 133, 249, 34, 113, 197, 252, 250, 145, 220, 3, 9, 141, 109, 251, 247, 199, 86, 26, 179, 230, 81, 115, 136, 225, 149, 190, 175, 247, 175, 151, 202, 127, 226, 143, 227, 174, 205, 255, 221, 38, 7, 142, 126, 70, 253, 212, 178, 109, 13, 90, 249, 1, 252, 34, 34, 52, 47, 110, 57, 26, 218, 182, 239, 80, 159, 207, 164, 172, 178, 235, 106, 44, 5, 232, 168, 243, 213, 9, 91, 78, 197, 247, 227, 86, 133, 80, 9, 2, 186, 4, 137, 142, 15, 95, 230, 223, 254, 129, 78, 177, 129, 9, 65, 211, 213, 70, 205, 41, 181, 81, 61, 147, 242, 135, 61, 62, 7, 216, 244, 203, 103, 117, 147, 216, 131, 249, 251, 250, 37, 181, 94, 61, 197, 20, 98, 158, 214, 231, 213, 254, 62, 90, 175, 243, 15, 75, 171, 131, 163, 200, 206, 182, 215, 182, 59, 28, 135, 3, 212, 207, 174, 222, 35, 250, 140, 159, 88, 55, 67, 184, 26, 96, 240, 246, 81, 221, 251, 221, 252, 46, 182, 124, 190, 148, 82, 86, 90, 104, 163, 40, 128, 74, 171, 186, 34, 114, 96, 180, 90, 99, 212, 98, 205, 170, 157, 18, 252, 141, 49, 212, 68, 236, 45, 173, 188, 186, 77, 237, 238, 23, 57, 180, 249, 36, 125, 128, 126, 139, 214, 75, 187, 87, 165, 85, 253, 27, 8, 6, 118, 2, 91, 15, 108, 88, 70, 218, 97, 101, 53, 168, 34, 252, 35, 162, 81, 235, 212, 120, 69, 121, 15, 232, 253, 114, 63, 31, 141, 169, 110, 254, 148, 178, 247, 100, 177, 225, 169, 141, 238, 190, 215, 63, 172, 189, 255, 135, 165, 120, 5, 6, 213, 147, 164, 53, 67, 194, 205, 179, 11, 152, 90, 89, 156, 70, 153, 3, 232, 252, 12, 95, 118, 158, 210, 101, 124, 73, 122, 177, 204, 218, 151, 37, 98, 71, 183, 172, 215, 252, 75, 82, 139, 201, 220, 146, 97, 205, 218, 146, 225, 64, 136, 125, 142, 34, 251, 22, 220, 220, 219, 125, 212, 237, 140, 121, 73, 233, 9, 254, 201, 247, 207, 79, 228, 224, 174, 249, 244, 122, 161, 79, 153, 67, 41, 53, 70, 194, 193, 239, 15, 56, 242, 118, 216, 53, 147, 62, 251, 69, 52, 107, 223, 169, 254, 132, 172, 61, 82, 72, 174, 159, 226, 47, 42, 188, 41, 167, 113, 122, 0, 21, 33, 58, 111, 29, 209, 87, 198, 136, 188, 189, 57, 210, 101, 171, 95, 115, 5, 83, 184, 153, 230, 35, 91, 24, 226, 159, 235, 233, 221, 122, 66, 219, 30, 193, 9, 161, 183, 11, 149, 112, 111, 93, 244, 205, 185, 158, 148, 47, 121, 164, 148, 204, 125, 229, 62, 78, 28, 95, 197, 229, 47, 246, 173, 83, 229, 119, 90, 156, 108, 121, 121, 139, 205, 108, 107, 163, 121, 102, 213, 174, 166, 82, 249, 1, 132, 20, 124, 246, 65, 174, 140, 169, 232, 97, 163, 40, 128, 116, 203, 96, 173, 151, 14, 141, 81, 67, 228, 192, 104, 113, 106, 229, 201, 134, 41, 72, 128, 119, 11, 31, 226, 110, 109, 229, 21, 115, 93, 172, 202, 237, 116, 242, 231, 151, 51, 26, 166, 172, 139, 12, 41, 37, 63, 190, 56, 137, 227, 199, 151, 209, 245, 137, 206, 168, 180, 181, 175, 10, 197, 167, 138, 88, 255, 244, 6, 71, 175, 145, 247, 234, 39, 125, 254, 139, 208, 25, 171, 190, 209, 242, 2, 227, 143, 224, 139, 231, 164, 44, 247, 146, 141, 51, 9, 150, 120, 105, 76, 165, 59, 234, 173, 111, 105, 75, 214, 230, 12, 28, 5, 13, 235, 194, 50, 184, 87, 40, 58, 31, 29, 127, 253, 252, 49, 197, 185, 151, 182, 23, 113, 233, 118, 243, 205, 212, 241, 156, 202, 88, 67, 231, 135, 58, 32, 212, 181, 175, 6, 25, 219, 210, 229, 150, 23, 183, 57, 39, 204, 248, 70, 123, 245, 131, 79, 215, 163, 148, 213, 71, 186, 171, 117, 52, 184, 127, 104, 1, 247, 255, 51, 176, 113, 108, 129, 4, 14, 233, 42, 21, 90, 99, 210, 208, 234, 166, 86, 50, 121, 65, 146, 136, 27, 215, 186, 193, 138, 84, 105, 84, 132, 245, 139, 32, 121, 81, 18, 171, 190, 156, 206, 240, 7, 95, 106, 176, 178, 154, 50, 110, 183, 139, 57, 79, 143, 33, 207, 158, 72, 167, 7, 218, 215, 254, 124, 46, 112, 252, 183, 99, 142, 212, 229, 185, 226, 137, 223, 55, 107, 130, 98, 98, 235, 81, 202, 170, 41, 202, 76, 103, 235, 231, 179, 72, 254, 107, 13, 69, 153, 233, 8, 149, 10, 189, 143, 47, 65, 173, 218, 17, 221, 179, 15, 237, 71, 222, 132, 246, 31, 61, 145, 148, 188, 60, 59, 71, 254, 58, 57, 64, 36, 255, 29, 214, 40, 147, 96, 189, 159, 126, 199, 229, 255, 235, 215, 197, 171, 153, 215, 105, 193, 36, 235, 159, 94, 67, 240, 229, 33, 4, 116, 110, 184, 21, 3, 151, 213, 197, 142, 87, 136, 6, 242, 27, 0, 0, 22, 48, 73, 68, 65, 84, 182, 34, 109, 130, 167, 22, 236, 33, 160, 89, 139, 6, 43, 171, 41, 226, 118, 187, 152, 243, 212, 173, 20, 184, 247, 210, 110, 82, 155, 218, 47, 75, 74, 56, 240, 237, 1, 155, 43, 197, 71, 115, 255, 119, 191, 171, 141, 190, 117, 59, 254, 232, 114, 56, 56, 181, 99, 11, 121, 201, 73, 228, 30, 63, 66, 110, 210, 81, 114, 146, 14, 211, 239, 193, 167, 105, 57, 248, 154, 114, 241, 139, 210, 79, 241, 243, 148, 113, 148, 100, 159, 191, 39, 55, 250, 7, 48, 100, 218, 235, 68, 198, 151, 51, 129, 249, 238, 30, 63, 49, 230, 239, 127, 26, 69, 1, 12, 1, 134, 141, 61, 166, 246, 238, 225, 19, 115, 214, 167, 103, 73, 122, 49, 27, 166, 174, 165, 253, 3, 157, 209, 122, 53, 156, 91, 243, 140, 141, 233, 28, 253, 254, 16, 157, 175, 188, 129, 59, 166, 255, 216, 96, 229, 52, 52, 150, 130, 92, 82, 15, 37, 146, 126, 100, 31, 217, 41, 71, 41, 202, 205, 164, 56, 55, 139, 226, 188, 28, 156, 118, 27, 80, 122, 99, 166, 16, 2, 141, 206, 128, 148, 146, 130, 156, 100, 2, 187, 250, 210, 113, 82, 167, 90, 255, 242, 46, 171, 139, 29, 111, 110, 183, 69, 181, 236, 167, 187, 253, 221, 47, 132, 90, 91, 247, 223, 74, 186, 221, 124, 121, 227, 224, 114, 21, 58, 186, 103, 95, 134, 255, 95, 121, 79, 142, 43, 95, 157, 198, 254, 197, 191, 84, 153, 175, 193, 215, 159, 113, 223, 255, 134, 206, 84, 230, 164, 154, 148, 130, 222, 247, 250, 138, 191, 160, 177, 134, 64, 18, 187, 219, 89, 118, 220, 102, 10, 53, 211, 250, 150, 118, 28, 254, 234, 0, 109, 39, 119, 168, 83, 215, 92, 25, 33, 61, 66, 73, 95, 151, 202, 174, 165, 115, 57, 188, 121, 21, 45, 187, 15, 108, 144, 114, 234, 27, 135, 205, 194, 145, 205, 171, 217, 183, 230, 55, 246, 173, 253, 141, 172, 227, 135, 171, 76, 163, 210, 151, 254, 188, 110, 155, 19, 141, 81, 67, 212, 224, 230, 116, 152, 88, 251, 213, 25, 107, 182, 149, 173, 175, 110, 115, 244, 189, 245, 62, 253, 176, 199, 166, 213, 58, 159, 127, 34, 84, 42, 98, 251, 15, 33, 241, 151, 178, 14, 172, 79, 108, 94, 79, 73, 118, 6, 166, 192, 144, 50, 225, 169, 59, 183, 148, 203, 195, 63, 38, 142, 128, 22, 113, 100, 31, 57, 72, 94, 114, 18, 66, 8, 6, 61, 253, 194, 63, 43, 127, 105, 113, 110, 158, 5, 70, 64, 35, 41, 128, 148, 50, 167, 162, 93, 224, 168, 193, 209, 228, 29, 202, 33, 121, 97, 18, 205, 175, 109, 160, 225, 137, 128, 152, 235, 98, 217, 51, 115, 23, 191, 190, 250, 48, 143, 252, 184, 5, 149, 170, 105, 158, 12, 181, 21, 23, 178, 243, 143, 159, 217, 249, 199, 143, 28, 222, 180, 10, 135, 173, 244, 59, 19, 26, 21, 250, 48, 47, 180, 254, 6, 180, 126, 70, 52, 62, 122, 84, 6, 13, 42, 131, 26, 149, 94, 83, 166, 241, 112, 91, 157, 228, 174, 74, 34, 102, 104, 115, 226, 70, 213, 222, 81, 88, 254, 209, 124, 118, 190, 181, 219, 125, 203, 75, 51, 181, 221, 174, 27, 93, 231, 119, 251, 39, 113, 131, 174, 42, 167, 0, 210, 237, 230, 192, 146, 133, 116, 29, 123, 87, 153, 112, 151, 163, 172, 117, 124, 104, 251, 206, 92, 255, 254, 87, 8, 149, 10, 41, 37, 251, 23, 205, 165, 36, 59, 139, 152, 222, 3, 43, 44, 75, 10, 134, 205, 206, 145, 157, 38, 7, 136, 221, 141, 162, 0, 182, 60, 219, 235, 7, 190, 221, 215, 63, 180, 71, 88, 185, 193, 99, 135, 187, 59, 179, 254, 217, 53, 100, 108, 76, 39, 164, 103, 195, 248, 247, 244, 142, 245, 33, 176, 75, 16, 167, 118, 236, 226, 175, 159, 62, 166, 247, 232, 201, 13, 82, 78, 109, 144, 110, 55, 135, 55, 255, 201, 230, 121, 95, 176, 107, 233, 207, 103, 238, 61, 214, 248, 232, 241, 138, 13, 198, 16, 233, 141, 62, 204, 171, 90, 43, 55, 206, 34, 59, 185, 171, 146, 104, 55, 182, 45, 145, 131, 162, 107, 45, 83, 250, 166, 52, 121, 232, 171, 99, 220, 243, 229, 34, 85, 76, 124, 195, 92, 61, 21, 222, 57, 30, 83, 64, 32, 37, 57, 101, 175, 57, 62, 178, 106, 105, 57, 5, 240, 143, 137, 163, 40, 227, 236, 109, 160, 46, 187, 13, 161, 42, 253, 62, 132, 16, 180, 27, 113, 99, 85, 197, 9, 183, 224, 110, 224, 193, 198, 106, 250, 78, 72, 151, 188, 222, 43, 210, 187, 153, 87, 51, 175, 50, 99, 29, 161, 18, 132, 245, 136, 96, 255, 23, 123, 209, 152, 53, 24, 67, 27, 102, 77, 217, 43, 218, 155, 204, 191, 210, 57, 178, 105, 53, 241, 35, 111, 171, 240, 234, 161, 11, 73, 81, 78, 6, 171, 190, 152, 206, 183, 207, 220, 193, 154, 175, 223, 229, 212, 129, 93, 160, 23, 120, 181, 11, 194, 191, 119, 20, 190, 93, 195, 49, 68, 250, 160, 241, 209, 87, 107, 120, 104, 207, 42, 33, 111, 245, 113, 58, 223, 211, 133, 136, 222, 181, 191, 60, 228, 216, 252, 99, 174, 180, 101, 133, 170, 135, 231, 175, 17, 225, 109, 219, 215, 58, 159, 170, 16, 66, 144, 115, 244, 48, 217, 135, 15, 148, 9, 183, 100, 103, 210, 254, 218, 155, 208, 158, 51, 148, 177, 23, 23, 145, 252, 215, 89, 95, 88, 37, 57, 217, 148, 100, 101, 18, 210, 182, 125, 153, 120, 149, 23, 72, 92, 247, 87, 158, 155, 222, 104, 125, 191, 203, 238, 90, 159, 189, 59, 115, 124, 204, 53, 177, 250, 127, 158, 47, 85, 235, 213, 132, 245, 140, 96, 239, 71, 187, 209, 7, 24, 48, 4, 215, 191, 25, 179, 198, 168, 65, 173, 83, 147, 181, 43, 141, 140, 164, 3, 196, 15, 31, 91, 117, 162, 6, 224, 228, 254, 29, 44, 153, 249, 28, 223, 61, 123, 23, 7, 55, 44, 195, 86, 82, 136, 177, 185, 47, 126, 241, 17, 248, 247, 108, 134, 33, 194, 27, 181, 161, 102, 29, 181, 245, 120, 30, 133, 91, 79, 209, 253, 233, 158, 4, 117, 172, 221, 170, 154, 116, 73, 246, 125, 182, 223, 45, 211, 252, 213, 15, 253, 178, 18, 159, 224, 134, 247, 182, 237, 180, 89, 57, 250, 231, 178, 114, 225, 254, 205, 91, 16, 220, 186, 253, 57, 255, 199, 176, 119, 193, 92, 156, 182, 179, 195, 232, 204, 131, 123, 217, 245, 227, 28, 146, 214, 175, 166, 56, 51, 29, 159, 136, 40, 244, 94, 149, 30, 226, 49, 153, 45, 44, 171, 245, 76, 243, 57, 41, 53, 97, 69, 180, 118, 187, 185, 76, 64, 28, 110, 66, 16, 132, 33, 8, 69, 98, 162, 212, 215, 204, 25, 147, 66, 1, 5, 82, 80, 136, 164, 72, 66, 129, 16, 164, 189, 123, 211, 149, 9, 118, 243, 137, 238, 109, 39, 180, 171, 208, 244, 176, 36, 189, 152, 191, 158, 91, 71, 212, 176, 230, 248, 119, 12, 172, 173, 168, 231, 71, 194, 158, 247, 118, 81, 120, 172, 128, 113, 175, 126, 73, 252, 136, 113, 245, 95, 70, 69, 197, 186, 221, 36, 174, 156, 207, 234, 175, 222, 225, 200, 150, 63, 1, 208, 120, 233, 48, 183, 13, 194, 220, 58, 16, 85, 29, 76, 18, 138, 246, 100, 224, 202, 40, 162, 231, 212, 222, 212, 246, 32, 139, 53, 219, 202, 174, 183, 19, 105, 219, 107, 24, 183, 188, 250, 126, 141, 142, 41, 214, 5, 107, 126, 46, 159, 143, 26, 88, 110, 99, 171, 69, 223, 65, 12, 125, 249, 157, 50, 97, 71, 86, 46, 97, 233, 115, 143, 159, 247, 42, 92, 161, 86, 19, 219, 127, 48, 189, 38, 63, 140, 79, 68, 100, 133, 113, 164, 228, 133, 106, 43, 192, 115, 82, 170, 130, 11, 233, 165, 114, 51, 10, 24, 68, 169, 139, 235, 58, 53, 205, 14, 171, 149, 151, 6, 118, 34, 118, 92, 24, 129, 231, 105, 169, 44, 89, 22, 54, 190, 176, 158, 144, 203, 195, 8, 233, 85, 255, 173, 80, 201, 169, 18, 18, 103, 236, 64, 111, 242, 225, 177, 159, 182, 225, 31, 209, 188, 222, 203, 248, 27, 151, 211, 193, 182, 133, 223, 176, 236, 227, 87, 201, 76, 58, 8, 128, 62, 212, 140, 87, 187, 96, 140, 205, 125, 161, 14, 230, 194, 210, 225, 38, 127, 83, 10, 70, 31, 29, 9, 143, 117, 71, 109, 168, 93, 165, 205, 220, 158, 193, 254, 79, 14, 115, 227, 11, 211, 233, 126, 227, 133, 239, 21, 127, 154, 52, 134, 204, 253, 137, 101, 194, 76, 129, 65, 76, 248, 101, 101, 185, 184, 7, 150, 204, 103, 237, 219, 175, 98, 47, 42, 60, 111, 126, 90, 147, 23, 195, 94, 123, 143, 136, 46, 241, 229, 158, 73, 193, 130, 42, 191, 241, 247, 11, 100, 27, 225, 230, 97, 74, 143, 23, 134, 84, 21, 191, 166, 156, 220, 179, 139, 119, 111, 189, 146, 30, 47, 116, 227, 124, 14, 150, 28, 69, 118, 54, 190, 184, 1, 239, 56, 31, 154, 13, 137, 170, 151, 221, 11, 183, 75, 82, 152, 97, 167, 40, 219, 78, 225, 158, 44, 242, 54, 158, 164, 121, 231, 158, 220, 255, 69, 233, 106, 203, 209, 173, 107, 104, 215, 127, 88, 189, 216, 176, 59, 109, 86, 254, 154, 251, 41, 43, 63, 251, 63, 114, 79, 37, 35, 132, 192, 216, 194, 15, 175, 14, 193, 232, 2, 235, 62, 199, 113, 228, 89, 201, 95, 151, 76, 243, 33, 49, 180, 186, 185, 117, 173, 190, 31, 233, 146, 28, 254, 254, 40, 5, 123, 29, 76, 249, 114, 62, 33, 113, 141, 227, 90, 126, 205, 140, 151, 73, 156, 251, 109, 185, 240, 59, 231, 175, 198, 224, 231, 95, 46, 220, 90, 144, 207, 206, 239, 62, 231, 240, 242, 223, 40, 72, 173, 216, 166, 204, 224, 235, 207, 152, 57, 243, 48, 248, 150, 75, 127, 228, 188, 205, 196, 251, 121, 50, 126, 196, 83, 207, 189, 43, 36, 239, 1, 221, 169, 157, 251, 188, 42, 241, 9, 9, 197, 43, 32, 148, 181, 51, 190, 35, 188, 95, 24, 162, 130, 67, 98, 106, 157, 154, 102, 253, 163, 56, 181, 58, 133, 140, 77, 105, 248, 182, 245, 175, 147, 241, 86, 73, 174, 147, 172, 99, 86, 108, 69, 46, 144, 2, 93, 176, 25, 71, 142, 133, 172, 131, 71, 200, 57, 121, 140, 130, 204, 84, 190, 125, 230, 46, 2, 155, 197, 208, 172, 109, 237, 79, 48, 217, 74, 138, 88, 51, 231, 29, 190, 124, 236, 86, 118, 46, 249, 17, 91, 73, 33, 230, 86, 1, 4, 12, 108, 142, 185, 117, 32, 234, 58, 218, 221, 3, 148, 28, 205, 165, 104, 203, 41, 186, 62, 20, 79, 228, 160, 218, 53, 14, 150, 204, 18, 182, 191, 190, 155, 200, 232, 62, 220, 243, 213, 2, 124, 66, 26, 239, 118, 157, 226, 140, 116, 142, 175, 255, 179, 92, 120, 116, 207, 62, 248, 132, 151, 31, 202, 104, 244, 6, 34, 227, 123, 209, 249, 230, 219, 104, 209, 111, 48, 222, 161, 225, 88, 114, 178, 177, 228, 229, 156, 137, 227, 180, 89, 209, 26, 77, 68, 92, 150, 240, 207, 228, 229, 7, 119, 211, 115, 165, 223, 168, 167, 159, 123, 79, 148, 94, 166, 214, 158, 11, 176, 91, 28, 217, 177, 11, 57, 73, 39, 216, 183, 112, 29, 33, 61, 131, 43, 108, 117, 85, 26, 21, 17, 125, 154, 225, 178, 184, 56, 248, 213, 62, 204, 81, 222, 232, 124, 107, 118, 106, 73, 186, 33, 55, 197, 74, 65, 186, 157, 127, 14, 29, 141, 145, 190, 216, 78, 21, 146, 178, 115, 27, 102, 255, 32, 50, 142, 29, 224, 228, 129, 93, 244, 185, 117, 74, 149, 251, 4, 5, 153, 169, 232, 205, 103, 39, 92, 46, 135, 157, 245, 223, 207, 226, 243, 135, 111, 38, 113, 249, 60, 28, 118, 43, 94, 109, 131, 8, 28, 16, 131, 41, 206, 255, 204, 6, 85, 93, 112, 219, 156, 228, 111, 58, 137, 202, 98, 231, 242, 231, 123, 227, 29, 93, 187, 85, 172, 148, 21, 41, 236, 157, 125, 144, 27, 159, 157, 193, 208, 135, 255, 131, 90, 211, 184, 174, 162, 92, 14, 7, 251, 23, 205, 45, 23, 30, 220, 186, 61, 161, 237, 59, 87, 154, 214, 20, 16, 68, 120, 231, 110, 116, 184, 238, 22, 116, 94, 94, 156, 216, 180, 254, 236, 67, 9, 109, 175, 25, 245, 207, 36, 101, 23, 147, 103, 229, 202, 46, 6, 193, 86, 1, 19, 185, 192, 102, 18, 55, 188, 240, 22, 161, 97, 9, 236, 251, 248, 64, 165, 142, 79, 98, 134, 197, 210, 237, 209, 30, 28, 253, 254, 16, 167, 86, 164, 32, 93, 213, 243, 146, 226, 118, 73, 50, 143, 148, 80, 146, 91, 177, 139, 61, 161, 85, 17, 56, 164, 5, 90, 63, 3, 219, 22, 125, 139, 79, 112, 24, 217, 39, 142, 112, 112, 125, 249, 85, 137, 115, 217, 240, 195, 135, 60, 55, 40, 146, 125, 107, 126, 199, 237, 118, 177, 101, 254, 87, 188, 58, 178, 61, 115, 95, 254, 55, 197, 249, 89, 120, 119, 12, 33, 252, 166, 118, 248, 245, 104, 134, 218, 92, 63, 38, 30, 214, 19, 5, 100, 47, 57, 66, 116, 223, 8, 122, 191, 212, 247, 188, 67, 199, 202, 176, 229, 89, 217, 254, 218, 46, 236, 251, 125, 249, 207, 242, 29, 116, 187, 238, 230, 122, 145, 173, 174, 4, 196, 182, 172, 176, 1, 44, 201, 169, 190, 5, 175, 16, 130, 46, 163, 111, 199, 20, 120, 214, 235, 132, 37, 47, 183, 162, 168, 37, 103, 20, 224, 253, 28, 217, 79, 10, 86, 3, 23, 214, 172, 239, 52, 66, 8, 238, 120, 127, 14, 94, 34, 142, 125, 159, 86, 174, 4, 126, 45, 253, 232, 255, 198, 64, 132, 83, 144, 248, 246, 78, 10, 147, 170, 184, 54, 74, 66, 214, 49, 11, 118, 75, 229, 102, 179, 106, 163, 150, 224, 107, 90, 162, 11, 49, 147, 155, 90, 234, 105, 123, 239, 234, 243, 95, 168, 88, 148, 155, 201, 194, 25, 255, 65, 103, 52, 147, 153, 116, 128, 87, 71, 182, 231, 155, 255, 220, 65, 206, 169, 36, 204, 173, 3, 9, 187, 161, 29, 190, 9, 17, 168, 106, 184, 140, 121, 62, 220, 54, 39, 249, 27, 78, 224, 76, 206, 165, 207, 75, 125, 137, 29, 25, 87, 171, 57, 74, 234, 250, 83, 108, 122, 118, 27, 131, 198, 61, 205, 3, 63, 46, 171, 179, 95, 158, 250, 68, 103, 50, 163, 53, 123, 149, 11, 119, 88, 138, 43, 140, 127, 62, 83, 104, 233, 118, 227, 58, 109, 19, 5, 160, 247, 42, 63, 130, 151, 144, 174, 2, 152, 157, 35, 59, 9, 21, 139, 128, 70, 221, 13, 82, 169, 213, 76, 254, 226, 87, 252, 244, 29, 72, 124, 127, 111, 165, 173, 187, 198, 164, 165, 211, 228, 46, 116, 125, 32, 158, 228, 121, 73, 36, 253, 124, 4, 103, 113, 197, 173, 123, 97, 134, 29, 123, 73, 245, 174, 147, 82, 233, 53, 4, 15, 109, 137, 79, 151, 80, 16, 130, 99, 219, 206, 239, 90, 102, 241, 140, 103, 176, 20, 228, 130, 148, 252, 250, 218, 35, 100, 159, 56, 130, 41, 46, 128, 208, 235, 218, 226, 223, 59, 170, 94, 198, 248, 80, 234, 27, 169, 120, 95, 38, 217, 75, 14, 19, 213, 39, 130, 62, 175, 245, 199, 20, 86, 243, 41, 89, 73, 122, 49, 219, 94, 221, 73, 225, 38, 61, 79, 252, 190, 137, 62, 227, 239, 174, 23, 249, 234, 27, 131, 143, 111, 185, 48, 123, 113, 73, 185, 176, 162, 204, 116, 126, 250, 215, 173, 100, 30, 220, 87, 238, 89, 226, 47, 223, 99, 43, 60, 219, 48, 6, 198, 181, 41, 23, 71, 8, 246, 169, 222, 58, 33, 141, 110, 21, 191, 0, 245, 227, 250, 171, 142, 168, 52, 26, 38, 126, 244, 61, 81, 205, 251, 179, 237, 213, 157, 56, 75, 42, 247, 138, 232, 215, 218, 159, 254, 111, 14, 194, 191, 85, 0, 137, 111, 239, 36, 229, 183, 228, 114, 158, 132, 139, 178, 107, 118, 216, 70, 168, 4, 62, 93, 195, 209, 120, 235, 200, 72, 58, 64, 69, 30, 166, 143, 239, 218, 200, 198, 185, 159, 1, 2, 167, 180, 99, 110, 19, 72, 232, 117, 109, 9, 232, 23, 141, 198, 167, 238, 238, 68, 254, 198, 114, 34, 159, 172, 197, 7, 241, 246, 213, 50, 96, 250, 21, 181, 106, 245, 221, 78, 55, 71, 231, 30, 99, 219, 75, 137, 92, 51, 249, 21, 30, 158, 183, 154, 128, 168, 134, 91, 238, 173, 43, 21, 41, 128, 163, 164, 108, 15, 144, 151, 156, 196, 47, 247, 222, 78, 214, 161, 125, 44, 124, 116, 18, 7, 151, 44, 160, 48, 237, 20, 153, 251, 19, 89, 51, 253, 37, 214, 189, 243, 74, 153, 248, 45, 250, 93, 81, 46, 79, 33, 89, 171, 49, 122, 243, 32, 77, 236, 222, 88, 161, 82, 49, 246, 205, 143, 88, 49, 107, 58, 203, 166, 190, 206, 101, 143, 118, 196, 28, 81, 190, 91, 60, 19, 95, 45, 136, 27, 213, 138, 230, 87, 181, 224, 216, 194, 35, 36, 78, 223, 73, 80, 124, 48, 97, 253, 35, 80, 27, 53, 184, 43, 247, 172, 125, 94, 180, 126, 6, 44, 201, 249, 20, 102, 165, 225, 27, 114, 214, 156, 192, 237, 118, 241, 211, 139, 247, 34, 165, 27, 223, 132, 8, 188, 218, 5, 87, 184, 122, 85, 23, 172, 39, 10, 40, 222, 159, 137, 41, 192, 64, 175, 105, 189, 241, 142, 170, 93, 251, 148, 185, 35, 131, 131, 95, 30, 165, 211, 144, 81, 252, 107, 253, 235, 24, 188, 27, 215, 228, 163, 58, 232, 188, 203, 43, 128, 234, 156, 201, 121, 230, 254, 68, 22, 62, 126, 47, 214, 252, 210, 113, 189, 53, 63, 143, 229, 47, 253, 231, 188, 249, 133, 118, 232, 76, 100, 247, 10, 28, 36, 171, 89, 172, 1, 110, 173, 179, 196, 13, 196, 21, 83, 30, 38, 172, 77, 7, 190, 124, 224, 118, 226, 110, 105, 78, 68, 223, 136, 74, 227, 107, 140, 26, 90, 221, 220, 134, 22, 195, 227, 56, 186, 224, 48, 137, 211, 119, 226, 19, 231, 131, 182, 153, 47, 194, 183, 230, 235, 237, 42, 99, 233, 16, 230, 159, 10, 176, 246, 155, 153, 156, 220, 183, 3, 99, 140, 31, 222, 29, 235, 111, 107, 68, 186, 37, 150, 163, 185, 148, 28, 200, 194, 59, 218, 135, 110, 255, 238, 138, 127, 235, 128, 90, 229, 85, 144, 84, 192, 225, 111, 142, 161, 87, 7, 113, 207, 23, 139, 137, 234, 92, 45, 87, 153, 77, 2, 85, 5, 134, 126, 70, 255, 179, 223, 195, 218, 119, 95, 63, 83, 249, 171, 194, 224, 227, 203, 224, 103, 94, 174, 168, 215, 252, 107, 138, 183, 216, 171, 1, 154, 244, 101, 186, 237, 7, 93, 197, 211, 203, 182, 49, 251, 142, 235, 200, 217, 189, 143, 182, 19, 90, 82, 149, 223, 26, 141, 73, 67, 235, 91, 218, 210, 242, 198, 214, 164, 111, 74, 37, 105, 73, 18, 197, 169, 41, 24, 98, 252, 48, 182, 240, 175, 246, 216, 92, 109, 44, 109, 117, 10, 50, 207, 90, 30, 230, 158, 58, 206, 111, 239, 78, 67, 165, 211, 224, 215, 163, 246, 70, 102, 231, 226, 200, 177, 96, 77, 206, 199, 122, 60, 143, 224, 203, 66, 232, 242, 92, 239, 74, 123, 188, 202, 176, 102, 91, 57, 242, 67, 18, 69, 199, 108, 220, 244, 226, 12, 58, 15, 189, 182, 94, 100, 172, 111, 236, 37, 197, 21, 217, 234, 35, 93, 46, 114, 142, 150, 63, 235, 16, 126, 206, 78, 110, 159, 251, 31, 231, 215, 7, 238, 192, 101, 175, 124, 104, 235, 29, 22, 193, 208, 151, 223, 193, 55, 178, 252, 112, 79, 192, 155, 80, 122, 30, 224, 36, 77, 92, 9, 124, 195, 34, 120, 108, 209, 6, 150, 190, 243, 26, 43, 159, 122, 155, 182, 19, 226, 8, 142, 175, 122, 179, 70, 165, 81, 17, 222, 187, 25, 225, 189, 155, 97, 201, 180, 144, 188, 44, 137, 147, 171, 143, 227, 150, 18, 125, 184, 23, 186, 48, 111, 116, 33, 230, 243, 90, 87, 170, 79, 123, 71, 203, 74, 46, 253, 65, 156, 54, 43, 95, 62, 54, 6, 91, 113, 33, 126, 189, 34, 235, 52, 201, 117, 22, 216, 176, 30, 207, 195, 154, 156, 143, 222, 79, 79, 212, 160, 40, 34, 30, 238, 134, 174, 18, 71, 193, 149, 97, 205, 177, 146, 180, 32, 153, 204, 205, 217, 92, 243, 200, 84, 250, 77, 152, 92, 102, 216, 208, 84, 144, 82, 178, 251, 231, 175, 217, 252, 201, 76, 174, 121, 181, 188, 137, 194, 206, 31, 190, 44, 99, 234, 12, 160, 243, 242, 38, 186, 103, 223, 51, 255, 135, 180, 235, 196, 176, 87, 103, 178, 252, 165, 167, 43, 60, 22, 169, 53, 121, 209, 241, 250, 91, 136, 31, 127, 55, 90, 83, 133, 13, 201, 230, 201, 190, 252, 60, 5, 16, 179, 242, 229, 116, 41, 43, 244, 11, 223, 36, 73, 63, 116, 128, 207, 238, 29, 139, 219, 80, 64, 171, 113, 177, 152, 106, 97, 46, 109, 201, 180, 144, 177, 45, 157, 180, 77, 169, 228, 31, 206, 67, 23, 108, 66, 237, 103, 64, 227, 107, 64, 27, 96, 68, 115, 186, 226, 59, 243, 109, 164, 253, 178, 143, 230, 157, 123, 50, 250, 249, 217, 204, 123, 253, 49, 14, 110, 88, 134, 49, 198, 143, 192, 129, 49, 53, 42, 211, 153, 111, 197, 153, 93, 130, 51, 167, 4, 203, 169, 34, 244, 126, 122, 154, 245, 141, 36, 162, 111, 36, 198, 58, 88, 187, 90, 50, 75, 72, 154, 159, 66, 246, 142, 28, 134, 220, 251, 56, 253, 238, 156, 66, 19, 116, 75, 2, 148, 154, 45, 175, 124, 229, 89, 146, 55, 174, 5, 64, 163, 215, 211, 229, 214, 59, 136, 232, 146, 64, 73, 118, 38, 71, 86, 45, 37, 105, 93, 121, 155, 159, 30, 19, 239, 39, 126, 66, 249, 51, 27, 246, 146, 98, 142, 173, 89, 78, 230, 129, 189, 184, 236, 54, 76, 1, 65, 132, 180, 237, 72, 179, 248, 94, 104, 244, 231, 93, 136, 112, 184, 37, 61, 238, 243, 23, 59, 0, 196, 71, 217, 50, 210, 169, 102, 31, 80, 187, 62, 183, 17, 144, 82, 178, 233, 199, 57, 252, 250, 226, 147, 132, 94, 30, 72, 204, 200, 230, 181, 62, 71, 44, 93, 110, 114, 15, 230, 82, 144, 148, 79, 222, 225, 60, 242, 143, 229, 99, 205, 178, 160, 245, 213, 163, 241, 53, 96, 77, 43, 194, 158, 109, 57, 19, 95, 31, 98, 38, 96, 64, 115, 132, 246, 236, 238, 176, 16, 160, 194, 13, 72, 164, 205, 133, 171, 196, 129, 219, 226, 192, 85, 104, 199, 158, 103, 195, 154, 101, 193, 43, 220, 139, 192, 142, 65, 4, 180, 15, 196, 191, 109, 0, 154, 58, 238, 13, 228, 31, 201, 227, 196, 146, 84, 10, 14, 22, 113, 213, 191, 255, 67, 159, 219, 38, 86, 246, 163, 55, 58, 41, 91, 254, 98, 249, 255, 158, 42, 119, 224, 165, 42, 66, 218, 117, 226, 186, 119, 63, 71, 173, 171, 155, 175, 210, 115, 120, 242, 30, 63, 241, 250, 223, 255, 8, 128, 15, 242, 228, 104, 224, 59, 46, 240, 238, 111, 93, 177, 21, 23, 241, 199, 59, 175, 177, 246, 171, 217, 52, 27, 20, 78, 243, 97, 209, 252, 237, 111, 168, 46, 72, 41, 41, 73, 43, 161, 232, 68, 33, 150, 172, 18, 44, 89, 22, 74, 210, 138, 113, 148, 56, 113, 217, 93, 56, 45, 78, 156, 197, 103, 151, 103, 213, 6, 53, 106, 157, 10, 141, 81, 139, 206, 87, 143, 193, 223, 128, 49, 216, 136, 57, 204, 11, 175, 72, 47, 76, 161, 166, 58, 249, 222, 57, 35, 151, 203, 77, 234, 134, 84, 82, 150, 164, 99, 208, 249, 115, 229, 125, 79, 210, 245, 218, 155, 168, 143, 131, 233, 13, 205, 214, 47, 63, 100, 211, 199, 239, 214, 40, 77, 96, 92, 107, 70, 188, 49, 171, 204, 142, 110, 29, 249, 110, 138, 47, 99, 133, 16, 103, 54, 152, 206, 84, 248, 15, 114, 229, 195, 8, 222, 228, 34, 83, 2, 0, 75, 126, 30, 75, 222, 121, 149, 13, 223, 124, 74, 232, 229, 33, 68, 15, 109, 134, 49, 184, 105, 14, 3, 106, 67, 225, 241, 2, 82, 215, 164, 147, 182, 33, 157, 54, 253, 6, 51, 228, 222, 39, 136, 238, 210, 173, 177, 197, 170, 49, 71, 86, 46, 97, 229, 171, 211, 112, 88, 202, 111, 106, 157, 139, 80, 169, 104, 55, 226, 6, 122, 223, 251, 232, 249, 198, 240, 53, 70, 192, 143, 194, 151, 113, 147, 133, 112, 252, 35, 252, 44, 31, 228, 203, 241, 72, 62, 2, 154, 110, 95, 90, 9, 182, 226, 34, 214, 207, 249, 132, 229, 179, 222, 196, 20, 161, 39, 108, 64, 16, 33, 221, 66, 235, 125, 141, 254, 66, 96, 205, 182, 144, 182, 62, 141, 212, 53, 153, 152, 188, 131, 232, 51, 118, 18, 9, 55, 142, 193, 236, 95, 187, 101, 209, 166, 66, 254, 201, 19, 108, 254, 244, 61, 142, 173, 94, 142, 211, 102, 43, 243, 204, 59, 44, 130, 152, 62, 3, 233, 48, 106, 52, 254, 49, 245, 186, 53, 245, 77, 186, 47, 19, 158, 19, 162, 220, 142, 80, 185, 154, 49, 43, 87, 118, 145, 130, 111, 129, 118, 245, 41, 193, 133, 68, 186, 221, 236, 255, 115, 57, 171, 191, 152, 201, 209, 77, 235, 8, 233, 30, 74, 72, 247, 211, 215, 106, 54, 144, 187, 149, 250, 160, 32, 169, 128, 204, 45, 153, 100, 109, 205, 5, 167, 134, 248, 235, 199, 208, 123, 204, 93, 4, 199, 214, 175, 247, 236, 166, 128, 203, 110, 35, 251, 200, 65, 236, 197, 197, 168, 245, 58, 124, 194, 35, 49, 7, 213, 251, 113, 19, 164, 228, 221, 32, 63, 30, 30, 45, 68, 133, 30, 152, 43, 172, 13, 179, 79, 73, 147, 52, 49, 77, 194, 195, 156, 115, 172, 241, 98, 196, 146, 159, 199, 142, 133, 191, 176, 249, 215, 57, 164, 236, 222, 73, 80, 151, 16, 252, 218, 155, 9, 234, 28, 84, 43, 43, 202, 250, 164, 36, 173, 152, 156, 61, 217, 228, 237, 183, 144, 179, 47, 147, 128, 200, 230, 116, 27, 121, 11, 151, 13, 187, 222, 35, 43, 253, 5, 38, 95, 192, 221, 83, 252, 196, 79, 149, 69, 170, 180, 57, 60, 125, 26, 108, 58, 80, 222, 63, 221, 69, 136, 173, 184, 136, 253, 171, 150, 177, 103, 229, 98, 246, 255, 185, 20, 151, 211, 138, 95, 107, 127, 204, 45, 244, 248, 198, 250, 226, 19, 227, 131, 74, 219, 48, 231, 95, 45, 153, 37, 228, 31, 203, 167, 40, 169, 132, 226, 100, 43, 121, 71, 115, 240, 13, 13, 167, 77, 223, 33, 180, 237, 55, 132, 150, 189, 251, 93, 20, 102, 10, 23, 9, 155, 37, 220, 122, 175, 159, 56, 90, 85, 196, 106, 141, 7, 102, 21, 200, 203, 165, 139, 255, 32, 74, 189, 105, 121, 10, 5, 25, 233, 36, 109, 219, 68, 210, 182, 141, 28, 221, 178, 134, 180, 3, 7, 144, 110, 39, 94, 17, 190, 24, 66, 244, 104, 253, 4, 122, 63, 29, 58, 63, 61, 122, 95, 61, 42, 141, 10, 181, 65, 205, 63, 189, 88, 56, 138, 29, 184, 29, 110, 236, 5, 118, 236, 133, 118, 236, 249, 54, 44, 153, 54, 108, 89, 78, 172, 89, 86, 236, 133, 54, 124, 195, 195, 136, 238, 18, 79, 139, 174, 151, 19, 213, 185, 27, 17, 29, 58, 161, 175, 174, 11, 15, 133, 106, 33, 32, 207, 13, 175, 6, 249, 50, 125, 180, 16, 213, 178, 128, 172, 209, 128, 120, 118, 142, 236, 235, 86, 241, 8, 112, 45, 208, 52, 221, 169, 213, 17, 167, 205, 70, 230, 177, 195, 100, 30, 59, 66, 222, 169, 147, 228, 103, 156, 34, 63, 227, 36, 133, 89, 233, 216, 45, 37, 56, 44, 22, 28, 86, 11, 182, 146, 18, 116, 102, 19, 66, 10, 164, 148, 104, 244, 58, 76, 190, 1, 4, 53, 143, 35, 164, 69, 27, 124, 67, 195, 9, 138, 137, 37, 40, 38, 246, 162, 159, 184, 94, 4, 184, 37, 124, 38, 212, 252, 231, 30, 111, 145, 81, 147, 132, 181, 154, 17, 190, 159, 39, 99, 85, 130, 7, 164, 228, 118, 64, 249, 117, 21, 26, 11, 41, 5, 11, 85, 146, 23, 166, 248, 137, 242, 14, 67, 171, 65, 157, 150, 68, 126, 144, 82, 151, 147, 203, 213, 168, 24, 47, 97, 20, 23, 249, 132, 89, 225, 162, 193, 141, 100, 177, 16, 60, 95, 219, 138, 255, 55, 245, 182, 38, 248, 94, 129, 12, 84, 75, 198, 32, 25, 15, 52, 140, 3, 73, 133, 75, 26, 1, 121, 72, 190, 70, 197, 219, 83, 124, 197, 161, 122, 202, 179, 254, 153, 93, 32, 219, 186, 93, 140, 21, 130, 145, 18, 26, 246, 102, 100, 5, 79, 71, 2, 171, 17, 124, 98, 41, 224, 167, 71, 162, 132, 165, 202, 20, 53, 160, 193, 119, 133, 102, 102, 203, 40, 181, 154, 225, 72, 70, 72, 193, 21, 212, 209, 155, 156, 194, 37, 195, 62, 33, 249, 89, 165, 230, 171, 73, 62, 226, 96, 67, 21, 114, 65, 183, 69, 223, 58, 33, 141, 6, 111, 250, 8, 201, 72, 4, 215, 3, 81, 23, 178, 124, 133, 38, 142, 100, 47, 240, 163, 20, 44, 184, 215, 79, 108, 189, 16, 69, 54, 170, 93, 192, 123, 217, 178, 131, 70, 195, 0, 183, 164, 191, 144, 244, 71, 16, 222, 152, 242, 40, 92, 112, 242, 129, 21, 192, 82, 9, 75, 170, 179, 113, 85, 223, 52, 41, 195, 152, 15, 11, 100, 107, 167, 155, 254, 42, 201, 0, 41, 24, 128, 210, 67, 120, 26, 22, 96, 179, 16, 44, 71, 176, 52, 205, 155, 205, 21, 25, 168, 93, 72, 154, 148, 2, 252, 147, 143, 242, 100, 11, 71, 105, 207, 208, 75, 64, 60, 208, 5, 101, 169, 245, 98, 226, 132, 132, 245, 42, 201, 6, 169, 98, 131, 202, 135, 237, 255, 52, 71, 110, 108, 154, 180, 2, 252, 147, 217, 82, 106, 201, 167, 147, 11, 18, 84, 16, 47, 161, 19, 165, 254, 75, 203, 251, 209, 80, 184, 144, 184, 37, 28, 86, 193, 78, 55, 236, 16, 42, 118, 97, 103, 251, 61, 65, 162, 98, 119, 205, 77, 136, 139, 74, 1, 206, 199, 236, 28, 25, 45, 213, 180, 71, 210, 17, 104, 47, 37, 173, 17, 196, 1, 77, 199, 231, 159, 103, 96, 7, 142, 10, 216, 47, 225, 160, 128, 131, 66, 144, 104, 180, 144, 120, 123, 152, 168, 216, 119, 97, 19, 199, 35, 20, 224, 124, 204, 204, 144, 94, 106, 45, 113, 82, 208, 82, 66, 75, 1, 205, 145, 68, 33, 136, 6, 34, 81, 204, 56, 202, 35, 73, 149, 130, 100, 1, 199, 37, 36, 11, 193, 113, 9, 199, 36, 28, 200, 244, 33, 169, 177, 199, 236, 245, 141, 71, 43, 64, 85, 204, 62, 37, 77, 120, 17, 237, 114, 19, 137, 32, 76, 184, 9, 22, 130, 80, 41, 8, 21, 110, 130, 165, 32, 4, 8, 5, 252, 105, 34, 174, 35, 107, 73, 17, 144, 11, 228, 0, 39, 17, 100, 72, 73, 170, 74, 146, 142, 32, 13, 73, 170, 74, 77, 154, 214, 155, 228, 59, 133, 40, 127, 127, 173, 7, 115, 73, 43, 64, 77, 120, 78, 74, 77, 64, 17, 254, 6, 137, 191, 75, 226, 7, 248, 35, 241, 149, 224, 167, 146, 24, 221, 2, 131, 0, 63, 33, 48, 32, 49, 73, 240, 17, 160, 150, 2, 29, 242, 204, 229, 34, 26, 106, 160, 72, 82, 96, 17, 146, 211, 151, 3, 83, 34, 36, 54, 0, 9, 121, 82, 98, 85, 9, 74, 36, 228, 9, 137, 197, 45, 176, 0, 185, 42, 65, 49, 130, 92, 55, 228, 8, 65, 174, 202, 139, 220, 166, 54, 241, 84, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 80, 248, 127, 174, 49, 51, 149, 202, 238, 0, 111, 0, 0, 0, 0, 73, 69, 78, 68, 174, 66, 96, 130},
		"react.webapp/public/logo512.png": {137, 80, 78, 71, 13, 10, 26, 10, 0, 0, 0, 13, 73, 72, 68, 82, 0, 0, 2, 0, 0, 0, 2, 0, 8, 6, 0, 0, 0, 244, 120, 212, 250, 0, 0, 0, 4, 115, 66, 73, 84, 8, 8, 8, 8, 124, 8, 100, 136, 0, 0, 0, 9, 112, 72, 89, 115, 0, 0, 16, 130, 0, 0, 16, 130, 1, 203, 80, 102, 0, 0, 0, 0, 25, 116, 69, 88, 116, 83, 111, 102, 116, 119, 97, 114, 101, 0, 119, 119, 119, 46, 105, 110, 107, 115, 99, 97, 112, 101, 46, 111, 114, 103, 155, 238, 60, 26, 0, 0, 32, 0, 73, 68, 65, 84, 120, 156, 236, 221, 119, 120, 20, 213, 215, 192, 241, 239, 110, 122, 39, 9, 9, 132, 14, 1, 66, 11, 157, 208, 148, 142, 116, 41, 210, 53, 136, 136, 138, 96, 69, 121, 197, 142, 13, 21, 123, 65, 127, 22, 170, 20, 1, 65, 186, 5, 164, 137, 16, 170, 16, 8, 6, 66, 11, 157, 4, 72, 239, 237, 253, 99, 201, 102, 107, 234, 214, 228, 124, 158, 103, 31, 118, 238, 220, 153, 57, 27, 2, 115, 246, 206, 45, 32, 132, 16, 66, 8, 33, 132, 16, 66, 8, 33, 132, 16, 66, 8, 33, 132, 16, 66, 8, 33, 132, 16, 66, 8, 33, 132, 16, 66, 8, 33, 132, 16, 66, 8, 33, 132, 16, 66, 8, 33, 132, 16, 66, 8, 33, 132, 16, 66, 8, 33, 132, 16, 66, 8, 33, 132, 16, 66, 8, 33, 132, 16, 66, 8, 33, 132, 16, 66, 8, 33, 132, 16, 66, 8, 33, 132, 16, 66, 8, 33, 132, 16, 66, 8, 33, 132, 16, 66, 8, 33, 132, 16, 66, 8, 33, 132, 16, 66, 8, 33, 132, 16, 66, 8, 33, 132, 16, 66, 8, 33, 132, 16, 66, 8, 33, 132, 16, 66, 8, 33, 132, 16, 66, 8, 33, 132, 16, 66, 8, 33, 132, 16, 66, 8, 33, 132, 16, 66, 8, 33, 132, 16, 66, 8, 33, 132, 16, 66, 8, 33, 132, 16, 66, 8, 33, 132, 16, 66, 8, 33, 132, 16, 66, 8, 33, 132, 16, 66, 8, 33, 132, 16, 66, 8, 33, 132, 16, 66, 8, 33, 132, 16, 66, 8, 33, 132, 16, 66, 8, 33, 132, 16, 66, 8, 33, 132, 16, 66, 8, 33, 132, 16, 66, 8, 33, 132, 16, 66, 8, 33, 132, 16, 66, 8, 33, 132, 16, 66, 8, 33, 132, 16, 66, 8, 33, 132, 16, 66, 8, 33, 132, 16, 66, 8, 33, 132, 16, 66, 8, 33, 132, 16, 66, 8, 33, 132, 16, 66, 8, 33, 132, 16, 66, 8, 33, 132, 16, 66, 8, 33, 132, 16, 66, 8, 33, 132, 16, 66, 8, 33, 132, 16, 66, 8, 33, 132, 16, 66, 8, 33, 132, 16, 66, 8, 33, 132, 16, 66, 8, 33, 132, 16, 66, 8, 33, 132, 16, 66, 8, 33, 132, 16, 66, 8, 33, 132, 16, 66, 8, 33, 132, 16, 66, 8, 33, 132, 16, 66, 8, 33, 132, 16, 66, 8, 33, 132, 16, 66, 8, 33, 132, 16, 66, 8, 33, 132, 16, 66, 8, 33, 132, 16, 66, 8, 33, 132, 16, 66, 8, 33, 132, 16, 66, 8, 33, 132, 16, 66, 8, 33, 132, 16, 66, 8, 33, 132, 16, 66, 8, 33, 132, 176, 3, 10, 107, 7, 32, 132, 5, 40, 128, 126, 64, 127, 160, 51, 80, 11, 112, 0, 146, 128, 243, 192, 63, 192, 22, 224, 180, 181, 2, 20, 66, 8, 33, 132, 105, 61, 4, 68, 1, 5, 165, 120, 237, 7, 30, 64, 18, 99, 33, 132, 16, 194, 110, 249, 3, 191, 82, 186, 27, 191, 238, 107, 47, 16, 98, 249, 144, 133, 16, 194, 114, 228, 155, 142, 168, 140, 106, 3, 187, 129, 96, 205, 66, 7, 71, 39, 234, 182, 234, 68, 205, 224, 22, 40, 148, 74, 82, 239, 196, 115, 53, 250, 24, 119, 174, 94, 48, 116, 142, 52, 96, 10, 176, 218, 252, 225, 10, 33, 132, 229, 73, 2, 32, 42, 155, 26, 168, 110, 254, 234, 111, 240, 78, 46, 110, 244, 158, 50, 139, 123, 38, 206, 192, 211, 55, 64, 239, 128, 27, 103, 163, 216, 191, 230, 7, 34, 126, 249, 129, 156, 172, 12, 205, 93, 249, 192, 12, 224, 127, 102, 142, 89, 8, 33, 44, 78, 18, 0, 81, 153, 40, 129, 157, 64, 143, 194, 130, 192, 134, 205, 120, 248, 211, 213, 4, 53, 105, 85, 226, 193, 9, 215, 47, 177, 250, 205, 199, 57, 189, 239, 79, 205, 226, 60, 96, 36, 176, 201, 196, 177, 10, 33, 132, 85, 57, 88, 59, 0, 33, 76, 104, 22, 48, 181, 112, 35, 160, 65, 83, 102, 44, 218, 65, 245, 186, 193, 197, 28, 82, 196, 205, 203, 135, 14, 67, 31, 36, 47, 39, 155, 11, 71, 247, 22, 22, 43, 129, 251, 129, 63, 128, 235, 38, 142, 87, 8, 33, 172, 70, 18, 0, 81, 89, 52, 3, 86, 1, 142, 0, 174, 158, 62, 60, 243, 211, 94, 170, 213, 172, 91, 166, 147, 40, 20, 10, 154, 118, 233, 75, 102, 90, 10, 177, 199, 247, 23, 22, 59, 163, 26, 70, 248, 35, 144, 99, 186, 144, 133, 16, 194, 122, 36, 1, 16, 149, 197, 15, 128, 186, 157, 127, 236, 156, 239, 8, 238, 212, 179, 220, 39, 11, 233, 218, 143, 27, 49, 39, 185, 121, 254, 191, 194, 34, 95, 192, 5, 216, 86, 145, 32, 133, 16, 194, 86, 72, 31, 0, 81, 25, 116, 67, 53, 153, 15, 0, 205, 186, 15, 224, 241, 239, 126, 171, 240, 73, 179, 210, 83, 153, 55, 34, 148, 132, 107, 177, 133, 69, 185, 168, 38, 18, 58, 90, 225, 147, 11, 33, 132, 149, 41, 173, 29, 128, 16, 38, 240, 65, 225, 27, 133, 82, 201, 144, 231, 223, 55, 201, 73, 93, 220, 61, 121, 224, 181, 175, 53, 139, 28, 129, 79, 77, 114, 114, 33, 132, 176, 50, 73, 0, 132, 189, 235, 11, 220, 91, 184, 209, 126, 240, 68, 106, 55, 107, 107, 178, 147, 183, 232, 49, 132, 182, 3, 198, 106, 22, 245, 188, 251, 18, 66, 8, 187, 38, 9, 128, 176, 119, 79, 22, 190, 113, 112, 116, 98, 208, 211, 111, 153, 252, 2, 131, 158, 121, 7, 165, 82, 171, 187, 204, 27, 38, 191, 136, 16, 66, 88, 152, 36, 0, 194, 158, 213, 2, 134, 23, 110, 180, 234, 59, 2, 191, 218, 13, 203, 125, 178, 130, 130, 2, 131, 229, 1, 245, 155, 208, 110, 240, 4, 205, 162, 62, 64, 215, 114, 95, 72, 8, 33, 108, 128, 36, 0, 194, 158, 61, 198, 221, 97, 127, 0, 221, 198, 62, 81, 230, 19, 20, 20, 20, 176, 246, 221, 167, 152, 59, 184, 41, 159, 60, 208, 142, 200, 237, 191, 26, 172, 215, 127, 218, 171, 186, 173, 0, 51, 202, 124, 49, 33, 132, 176, 33, 146, 0, 8, 123, 54, 169, 240, 77, 141, 70, 205, 105, 28, 214, 187, 204, 39, 216, 189, 244, 51, 14, 109, 88, 194, 173, 75, 103, 185, 118, 38, 146, 95, 231, 62, 67, 220, 133, 104, 189, 122, 129, 13, 66, 104, 222, 115, 136, 102, 209, 104, 64, 127, 94, 97, 33, 132, 176, 19, 146, 0, 8, 123, 213, 28, 104, 84, 184, 17, 54, 114, 50, 10, 69, 217, 70, 181, 230, 102, 103, 177, 127, 245, 247, 100, 103, 164, 169, 203, 146, 226, 174, 178, 97, 222, 11, 6, 235, 119, 27, 243, 184, 230, 166, 11, 240, 72, 153, 46, 40, 132, 16, 54, 68, 18, 0, 97, 175, 6, 107, 110, 180, 232, 49, 196, 88, 61, 163, 14, 109, 88, 194, 157, 107, 23, 245, 202, 175, 157, 57, 65, 226, 205, 43, 122, 229, 205, 238, 25, 168, 219, 199, 224, 49, 100, 46, 13, 33, 132, 157, 146, 4, 64, 216, 171, 65, 133, 111, 252, 106, 55, 164, 70, 112, 139, 50, 159, 224, 223, 223, 126, 38, 47, 39, 91, 175, 60, 233, 230, 21, 118, 45, 214, 31, 238, 175, 80, 42, 233, 58, 122, 170, 102, 81, 99, 160, 99, 153, 47, 44, 132, 16, 54, 192, 177, 228, 42, 66, 216, 28, 23, 52, 198, 254, 183, 232, 49, 184, 152, 170, 134, 21, 228, 231, 147, 120, 67, 255, 91, 126, 161, 11, 255, 238, 53, 88, 222, 126, 200, 68, 182, 126, 249, 154, 230, 136, 129, 113, 192, 161, 18, 46, 231, 9, 212, 4, 2, 1, 47, 192, 27, 112, 3, 92, 1, 31, 140, 39, 226, 9, 64, 42, 144, 162, 241, 186, 13, 220, 4, 50, 140, 28, 35, 132, 16, 165, 34, 9, 128, 176, 71, 33, 168, 22, 232, 1, 160, 81, 135, 123, 139, 169, 106, 88, 220, 197, 211, 100, 36, 39, 24, 221, 159, 28, 127, 157, 148, 91, 55, 240, 170, 94, 83, 171, 220, 183, 86, 125, 234, 181, 238, 76, 236, 241, 136, 194, 162, 49, 192, 39, 64, 240, 221, 87, 227, 187, 127, 54, 0, 106, 163, 234, 40, 232, 86, 230, 0, 75, 118, 7, 213, 234, 132, 87, 239, 254, 121, 14, 56, 13, 156, 185, 251, 74, 55, 195, 53, 133, 16, 149, 136, 36, 0, 194, 30, 181, 212, 220, 8, 106, 210, 202, 88, 61, 163, 206, 29, 218, 77, 122, 210, 29, 163, 251, 147, 227, 174, 17, 253, 207, 159, 116, 26, 62, 73, 111, 95, 187, 129, 99, 53, 19, 128, 122, 192, 181, 50, 7, 80, 113, 126, 119, 95, 45, 13, 236, 43, 0, 46, 3, 255, 1, 71, 52, 94, 177, 6, 234, 10, 33, 170, 40, 233, 3, 32, 236, 145, 250, 166, 231, 224, 232, 68, 245, 122, 141, 203, 124, 130, 51, 17, 127, 25, 157, 248, 7, 84, 243, 3, 68, 255, 109, 120, 65, 161, 182, 3, 199, 161, 80, 218, 244, 63, 29, 5, 170, 196, 100, 0, 240, 10, 176, 22, 184, 8, 196, 1, 91, 239, 150, 117, 67, 190, 0, 8, 81, 165, 217, 244, 255, 98, 66, 24, 209, 188, 240, 77, 245, 250, 77, 112, 112, 114, 46, 174, 174, 65, 201, 241, 37, 127, 105, 55, 52, 66, 0, 192, 59, 32, 136, 134, 237, 186, 151, 249, 154, 54, 32, 0, 85, 231, 201, 247, 80, 173, 158, 152, 0, 252, 6, 252, 31, 208, 206, 138, 113, 9, 33, 172, 64, 190, 1, 8, 123, 164, 158, 128, 199, 175, 86, 253, 114, 157, 32, 51, 53, 165, 196, 58, 25, 41, 137, 70, 247, 181, 29, 56, 150, 243, 71, 254, 54, 186, 95, 161, 80, 224, 233, 23, 136, 135, 95, 0, 94, 254, 129, 120, 87, 175, 137, 135, 111, 0, 158, 126, 1, 56, 58, 185, 224, 226, 225, 133, 147, 171, 27, 78, 206, 174, 184, 122, 249, 160, 80, 40, 201, 206, 72, 211, 26, 149, 144, 147, 149, 65, 102, 90, 10, 89, 105, 41, 100, 164, 36, 146, 153, 154, 76, 122, 210, 29, 146, 226, 174, 145, 28, 127, 157, 164, 184, 171, 228, 102, 101, 150, 237, 131, 107, 243, 4, 6, 222, 125, 125, 8, 92, 2, 54, 222, 125, 237, 6, 244, 135, 72, 216, 14, 71, 160, 19, 208, 6, 213, 148, 208, 133, 89, 96, 42, 170, 214, 142, 163, 64, 52, 144, 111, 141, 224, 132, 176, 7, 146, 0, 8, 123, 228, 85, 248, 198, 197, 221, 171, 184, 122, 6, 229, 231, 229, 146, 157, 145, 90, 98, 189, 236, 140, 116, 178, 210, 82, 112, 241, 208, 191, 70, 219, 129, 99, 89, 255, 193, 243, 228, 231, 229, 2, 80, 179, 113, 75, 58, 143, 154, 66, 245, 186, 193, 84, 175, 23, 140, 127, 157, 70, 56, 186, 184, 150, 57, 182, 178, 74, 79, 186, 67, 226, 141, 203, 196, 95, 140, 33, 62, 246, 12, 113, 23, 78, 19, 119, 81, 245, 103, 102, 106, 82, 89, 79, 87, 15, 120, 234, 238, 43, 9, 216, 2, 44, 3, 182, 1, 185, 38, 13, 188, 124, 106, 162, 154, 253, 177, 47, 208, 29, 240, 40, 161, 254, 13, 96, 61, 240, 29, 112, 204, 188, 161, 9, 97, 127, 36, 1, 16, 246, 168, 40, 1, 48, 112, 115, 46, 73, 226, 141, 43, 228, 100, 150, 60, 138, 46, 59, 35, 141, 132, 235, 151, 168, 217, 88, 191, 159, 157, 167, 111, 0, 193, 29, 123, 16, 115, 96, 7, 0, 169, 183, 227, 232, 17, 254, 108, 153, 103, 35, 172, 40, 119, 31, 63, 220, 125, 252, 168, 21, 210, 70, 111, 223, 237, 203, 231, 184, 28, 117, 148, 203, 81, 135, 185, 114, 234, 8, 151, 163, 142, 150, 37, 41, 240, 1, 38, 222, 125, 221, 0, 126, 6, 126, 66, 245, 205, 218, 210, 238, 5, 166, 3, 163, 208, 24, 253, 81, 10, 53, 129, 105, 119, 95, 127, 0, 51, 129, 83, 38, 143, 78, 8, 59, 37, 9, 128, 176, 71, 21, 74, 0, 18, 174, 199, 150, 234, 70, 152, 155, 157, 69, 102, 90, 178, 209, 253, 237, 6, 143, 87, 39, 0, 25, 41, 137, 228, 229, 100, 227, 232, 236, 82, 230, 120, 204, 197, 191, 110, 48, 254, 117, 131, 105, 59, 112, 12, 160, 154, 251, 224, 122, 204, 9, 206, 30, 220, 69, 204, 193, 157, 156, 63, 252, 119, 177, 143, 57, 52, 212, 4, 158, 187, 251, 58, 1, 124, 131, 170, 101, 160, 228, 102, 148, 138, 9, 3, 62, 67, 213, 97, 81, 139, 135, 111, 117, 130, 59, 244, 160, 97, 187, 110, 248, 213, 105, 120, 183, 37, 168, 128, 148, 219, 113, 196, 199, 198, 112, 241, 216, 62, 206, 29, 218, 77, 94, 110, 78, 225, 33, 3, 80, 141, 132, 120, 4, 85, 50, 35, 68, 149, 39, 9, 128, 176, 71, 234, 101, 249, 202, 211, 27, 63, 241, 230, 85, 114, 74, 241, 236, 188, 160, 160, 0, 140, 142, 20, 40, 160, 227, 208, 7, 137, 222, 251, 59, 23, 143, 71, 208, 247, 209, 151, 108, 234, 230, 111, 136, 66, 169, 164, 86, 72, 27, 106, 133, 180, 161, 71, 248, 179, 228, 231, 231, 113, 45, 250, 24, 167, 246, 252, 70, 212, 206, 77, 92, 57, 117, 164, 216, 145, 17, 119, 133, 2, 223, 2, 31, 0, 11, 81, 37, 3, 103, 77, 28, 106, 29, 224, 125, 224, 65, 52, 166, 90, 118, 245, 244, 161, 211, 240, 135, 233, 60, 234, 17, 130, 154, 182, 46, 177, 181, 37, 35, 57, 129, 136, 117, 11, 216, 189, 244, 115, 146, 227, 174, 129, 106, 226, 165, 101, 168, 38, 84, 218, 98, 226, 152, 133, 176, 59, 50, 143, 185, 176, 71, 49, 168, 38, 220, 161, 243, 200, 41, 140, 123, 231, 199, 50, 29, 188, 107, 201, 167, 108, 252, 232, 197, 18, 235, 57, 187, 186, 51, 237, 199, 63, 105, 208, 86, 247, 11, 104, 137, 55, 73, 187, 148, 20, 119, 149, 168, 93, 155, 57, 185, 99, 3, 49, 17, 59, 52, 191, 61, 23, 39, 31, 213, 208, 194, 247, 129, 125, 21, 12, 65, 129, 170, 169, 255, 67, 52, 158, 239, 251, 213, 110, 72, 191, 199, 102, 211, 126, 200, 68, 156, 221, 74, 122, 236, 175, 47, 43, 61, 149, 53, 115, 158, 224, 232, 214, 149, 133, 69, 55, 129, 102, 64, 169, 154, 63, 132, 168, 172, 164, 5, 64, 216, 35, 245, 20, 126, 233, 201, 198, 39, 243, 49, 38, 37, 254, 70, 169, 234, 185, 120, 122, 225, 19, 88, 91, 163, 164, 114, 222, 248, 11, 249, 4, 214, 166, 219, 216, 39, 232, 54, 246, 9, 82, 19, 226, 249, 119, 235, 42, 142, 108, 94, 206, 165, 19, 7, 139, 59, 76, 9, 12, 189, 251, 250, 11, 120, 7, 213, 8, 130, 178, 170, 137, 170, 69, 65, 189, 198, 131, 171, 167, 55, 253, 30, 123, 153, 30, 225, 207, 85, 168, 117, 197, 197, 221, 147, 137, 31, 44, 37, 53, 225, 22, 103, 246, 111, 3, 168, 129, 170, 95, 192, 7, 229, 62, 169, 16, 149, 128, 204, 3, 32, 236, 145, 250, 174, 159, 158, 100, 124, 58, 95, 99, 82, 19, 227, 75, 85, 207, 221, 199, 31, 223, 90, 245, 81, 221, 248, 43, 247, 205, 95, 151, 167, 111, 0, 247, 62, 248, 20, 207, 173, 220, 207, 236, 205, 167, 232, 247, 248, 43, 120, 7, 4, 149, 116, 88, 95, 96, 23, 170, 4, 160, 79, 25, 46, 119, 63, 16, 137, 198, 205, 191, 227, 253, 147, 120, 121, 203, 105, 250, 152, 232, 209, 138, 82, 233, 192, 200, 217, 159, 105, 62, 54, 152, 80, 225, 147, 10, 97, 231, 36, 1, 16, 246, 72, 157, 0, 164, 220, 190, 89, 230, 131, 75, 51, 118, 222, 221, 199, 143, 46, 163, 166, 80, 213, 110, 252, 134, 4, 54, 8, 97, 240, 51, 239, 240, 250, 182, 11, 60, 52, 111, 57, 13, 218, 116, 45, 233, 144, 30, 168, 90, 3, 182, 0, 197, 45, 211, 168, 0, 222, 70, 53, 84, 47, 0, 192, 163, 154, 63, 147, 63, 91, 195, 196, 185, 139, 241, 242, 175, 97, 130, 232, 139, 212, 8, 110, 161, 185, 106, 100, 40, 37, 15, 35, 20, 162, 82, 147, 4, 64, 216, 35, 117, 239, 243, 91, 177, 49, 100, 103, 150, 109, 221, 155, 188, 92, 227, 67, 218, 61, 170, 249, 83, 191, 77, 23, 30, 154, 183, 140, 158, 15, 63, 95, 254, 8, 43, 33, 7, 71, 39, 218, 15, 30, 207, 51, 203, 247, 50, 115, 245, 65, 58, 222, 31, 142, 131, 163, 83, 113, 135, 12, 6, 142, 163, 234, 40, 24, 160, 179, 207, 3, 88, 3, 188, 206, 221, 190, 72, 77, 186, 244, 229, 197, 117, 199, 104, 221, 255, 1, 51, 68, 175, 226, 95, 55, 184, 240, 173, 2, 213, 90, 10, 66, 84, 89, 146, 0, 8, 123, 226, 14, 124, 4, 76, 41, 44, 200, 207, 207, 227, 198, 217, 168, 50, 157, 196, 80, 239, 113, 223, 160, 122, 180, 234, 51, 156, 25, 139, 119, 242, 236, 242, 127, 104, 214, 125, 64, 69, 99, 173, 212, 234, 180, 232, 192, 196, 185, 139, 121, 121, 75, 52, 93, 199, 62, 94, 220, 116, 204, 142, 192, 147, 168, 58, 110, 206, 68, 53, 130, 163, 30, 176, 23, 120, 0, 84, 127, 31, 3, 166, 191, 201, 180, 31, 254, 212, 233, 115, 97, 122, 110, 158, 62, 154, 155, 190, 102, 189, 152, 16, 54, 78, 58, 1, 10, 123, 209, 3, 248, 30, 213, 82, 192, 90, 174, 69, 31, 163, 94, 171, 78, 165, 62, 81, 80, 147, 86, 68, 239, 253, 3, 87, 47, 111, 124, 2, 107, 211, 172, 251, 125, 220, 51, 113, 6, 158, 126, 129, 38, 12, 183, 106, 240, 171, 221, 128, 49, 111, 124, 203, 192, 25, 115, 216, 189, 244, 115, 254, 94, 246, 21, 57, 89, 6, 39, 89, 242, 65, 181, 108, 242, 195, 64, 32, 170, 78, 127, 56, 187, 186, 51, 97, 238, 98, 218, 220, 55, 218, 34, 241, 234, 172, 0, 89, 242, 124, 208, 66, 84, 98, 50, 12, 80, 216, 58, 55, 224, 83, 224, 9, 140, 252, 190, 118, 31, 63, 157, 7, 94, 251, 186, 212, 39, 204, 205, 202, 228, 218, 153, 227, 248, 213, 105, 136, 167, 175, 110, 203, 116, 197, 228, 231, 229, 146, 116, 243, 42, 105, 137, 183, 200, 74, 79, 35, 59, 35, 157, 236, 244, 84, 20, 14, 14, 184, 122, 120, 227, 234, 233, 133, 147, 171, 59, 62, 129, 181, 112, 247, 169, 124, 45, 208, 137, 55, 46, 179, 229, 243, 87, 57, 186, 101, 69, 137, 115, 10, 248, 212, 168, 195, 212, 175, 55, 80, 187, 185, 229, 214, 33, 250, 96, 104, 115, 226, 46, 158, 6, 213, 58, 7, 238, 64, 158, 197, 46, 46, 132, 141, 145, 22, 0, 97, 203, 90, 162, 154, 181, 173, 85, 113, 149, 238, 14, 237, 42, 65, 209, 205, 200, 209, 197, 133, 122, 161, 97, 21, 139, 12, 213, 248, 242, 243, 71, 246, 114, 254, 232, 223, 196, 157, 143, 38, 238, 226, 25, 110, 93, 58, 171, 181, 160, 79, 113, 60, 124, 171, 19, 216, 176, 25, 129, 13, 154, 82, 187, 89, 91, 26, 135, 245, 50, 56, 237, 176, 61, 169, 86, 179, 46, 15, 126, 176, 148, 123, 31, 122, 154, 141, 31, 205, 42, 118, 193, 164, 220, 204, 244, 132, 156, 236, 76, 139, 53, 195, 167, 220, 190, 73, 252, 165, 152, 194, 205, 115, 200, 205, 95, 84, 113, 210, 2, 32, 108, 213, 19, 168, 166, 129, 117, 51, 180, 211, 187, 122, 16, 201, 183, 174, 171, 183, 95, 218, 24, 69, 141, 70, 205, 13, 212, 44, 160, 0, 211, 253, 162, 223, 60, 255, 31, 255, 110, 253, 153, 51, 17, 59, 184, 116, 226, 160, 122, 49, 32, 83, 241, 170, 94, 147, 38, 97, 189, 104, 222, 99, 8, 161, 253, 70, 224, 236, 234, 110, 210, 243, 91, 90, 228, 182, 117, 108, 152, 247, 2, 9, 215, 47, 25, 220, 175, 80, 42, 243, 7, 61, 245, 150, 178, 239, 212, 151, 203, 53, 171, 99, 89, 68, 252, 242, 3, 171, 231, 60, 81, 184, 249, 25, 170, 62, 9, 66, 84, 89, 146, 0, 8, 91, 227, 138, 106, 66, 24, 131, 227, 180, 29, 93, 92, 185, 111, 218, 107, 180, 236, 53, 140, 143, 70, 22, 45, 128, 51, 228, 185, 185, 244, 157, 58, 91, 163, 166, 233, 134, 239, 165, 37, 220, 226, 232, 214, 159, 57, 178, 105, 25, 151, 78, 30, 50, 217, 121, 75, 226, 226, 225, 69, 235, 254, 163, 232, 116, 255, 36, 130, 59, 245, 180, 248, 66, 67, 166, 146, 149, 150, 194, 214, 47, 95, 103, 239, 202, 249, 20, 228, 27, 94, 157, 183, 113, 88, 111, 30, 252, 96, 105, 133, 59, 1, 22, 228, 231, 115, 116, 235, 74, 206, 236, 223, 134, 127, 221, 96, 250, 63, 241, 154, 250, 231, 246, 205, 148, 190, 156, 61, 184, 179, 176, 106, 111, 84, 115, 22, 8, 81, 101, 217, 231, 255, 40, 162, 178, 10, 2, 126, 5, 58, 27, 218, 89, 171, 105, 107, 194, 63, 90, 174, 30, 203, 253, 193, 208, 22, 133, 207, 115, 105, 208, 166, 43, 207, 44, 255, 7, 83, 222, 248, 19, 111, 92, 102, 231, 162, 143, 137, 248, 101, 129, 177, 142, 109, 106, 190, 190, 190, 52, 109, 218, 148, 102, 205, 154, 17, 18, 18, 66, 245, 234, 213, 169, 86, 173, 26, 158, 158, 158, 184, 187, 187, 147, 147, 147, 67, 106, 106, 42, 105, 105, 105, 36, 37, 37, 113, 225, 194, 5, 78, 159, 62, 77, 116, 116, 52, 23, 46, 92, 32, 183, 152, 161, 137, 0, 65, 77, 66, 233, 251, 216, 108, 218, 14, 28, 131, 82, 233, 80, 108, 93, 91, 21, 123, 60, 130, 85, 115, 158, 224, 70, 204, 73, 131, 251, 61, 170, 249, 51, 225, 189, 69, 180, 232, 57, 180, 204, 231, 206, 203, 205, 225, 239, 101, 95, 114, 116, 235, 207, 92, 143, 57, 65, 94, 78, 54, 74, 71, 39, 30, 124, 127, 41, 237, 6, 141, 227, 210, 137, 131, 124, 62, 161, 75, 97, 245, 203, 64, 35, 108, 99, 137, 99, 33, 172, 70, 18, 0, 97, 43, 218, 1, 27, 128, 186, 186, 59, 20, 10, 5, 221, 199, 63, 201, 253, 47, 126, 132, 163, 139, 171, 186, 124, 211, 39, 47, 177, 115, 209, 199, 170, 58, 74, 37, 239, 254, 19, 143, 155, 87, 181, 10, 7, 146, 112, 45, 150, 63, 191, 125, 135, 195, 155, 151, 27, 125, 158, 31, 16, 16, 64, 239, 222, 189, 233, 221, 187, 55, 125, 250, 244, 161, 105, 211, 166, 229, 190, 94, 122, 122, 58, 251, 246, 237, 99, 199, 142, 29, 236, 220, 185, 147, 67, 135, 14, 145, 151, 103, 248, 241, 116, 245, 122, 141, 233, 247, 216, 203, 116, 26, 62, 201, 236, 77, 230, 230, 144, 151, 155, 195, 182, 255, 189, 203, 246, 239, 223, 39, 63, 95, 255, 51, 42, 148, 74, 6, 63, 253, 14, 125, 166, 206, 46, 117, 139, 71, 204, 193, 157, 108, 250, 120, 22, 215, 78, 31, 39, 95, 227, 231, 166, 84, 42, 25, 255, 222, 98, 58, 14, 123, 136, 5, 79, 13, 39, 106, 215, 166, 194, 93, 207, 1, 95, 84, 248, 195, 8, 97, 231, 36, 1, 16, 182, 96, 40, 170, 206, 126, 122, 51, 179, 121, 84, 243, 103, 220, 219, 63, 210, 170, 207, 253, 122, 7, 93, 142, 58, 204, 231, 227, 187, 80, 80, 80, 128, 171, 167, 15, 115, 118, 93, 169, 208, 51, 243, 188, 220, 28, 254, 249, 249, 91, 126, 251, 242, 117, 178, 210, 245, 87, 186, 117, 117, 117, 101, 216, 176, 97, 132, 135, 135, 51, 104, 208, 32, 28, 29, 205, 211, 135, 246, 246, 237, 219, 172, 93, 187, 150, 165, 75, 151, 242, 207, 63, 255, 24, 172, 83, 167, 69, 123, 30, 120, 237, 107, 234, 183, 54, 216, 88, 98, 243, 206, 29, 222, 195, 242, 217, 147, 72, 188, 113, 217, 224, 254, 246, 131, 39, 48, 238, 157, 31, 113, 114, 49, 216, 5, 4, 128, 236, 204, 116, 214, 189, 247, 52, 81, 187, 54, 145, 150, 112, 75, 111, 127, 141, 70, 205, 121, 126, 245, 33, 98, 34, 254, 98, 193, 83, 195, 11, 139, 111, 1, 13, 128, 180, 138, 126, 6, 33, 236, 157, 36, 0, 194, 218, 198, 162, 90, 162, 85, 111, 74, 185, 128, 6, 77, 153, 250, 245, 6, 2, 26, 24, 255, 118, 125, 252, 207, 181, 196, 68, 252, 69, 167, 17, 15, 87, 232, 102, 120, 238, 240, 110, 214, 188, 53, 157, 184, 11, 209, 122, 251, 234, 212, 169, 195, 172, 89, 179, 152, 60, 121, 50, 222, 222, 222, 229, 190, 70, 121, 252, 251, 239, 191, 124, 244, 209, 71, 172, 94, 189, 90, 175, 85, 64, 169, 116, 160, 219, 184, 39, 24, 252, 236, 123, 184, 122, 90, 54, 46, 83, 200, 76, 77, 98, 245, 156, 105, 28, 251, 125, 181, 193, 253, 181, 155, 181, 101, 202, 87, 235, 241, 13, 170, 167, 183, 239, 202, 169, 163, 172, 122, 125, 42, 215, 206, 28, 55, 56, 220, 208, 193, 193, 145, 1, 79, 205, 225, 222, 137, 79, 51, 111, 68, 168, 102, 39, 196, 39, 80, 205, 39, 33, 68, 149, 39, 9, 128, 176, 166, 71, 128, 31, 80, 205, 14, 167, 165, 121, 143, 65, 132, 207, 91, 97, 246, 27, 91, 126, 126, 30, 127, 126, 243, 54, 219, 190, 159, 171, 215, 65, 173, 81, 163, 70, 188, 252, 242, 203, 76, 154, 52, 9, 103, 103, 163, 51, 221, 89, 68, 76, 76, 12, 31, 124, 240, 1, 75, 151, 46, 213, 235, 47, 16, 80, 191, 9, 147, 62, 94, 105, 209, 241, 244, 166, 180, 119, 229, 55, 108, 248, 112, 166, 193, 229, 135, 125, 2, 107, 243, 196, 247, 191, 107, 13, 143, 252, 123, 217, 151, 236, 92, 252, 137, 209, 214, 3, 128, 70, 29, 238, 101, 250, 162, 29, 172, 152, 61, 73, 115, 25, 224, 189, 168, 38, 148, 146, 5, 30, 132, 64, 18, 0, 97, 61, 211, 128, 249, 24, 152, 142, 186, 235, 152, 199, 120, 224, 181, 175, 81, 58, 152, 119, 154, 138, 212, 59, 113, 44, 123, 41, 156, 51, 251, 183, 107, 149, 59, 59, 59, 243, 252, 243, 207, 51, 103, 206, 28, 92, 93, 93, 141, 28, 109, 29, 145, 145, 145, 76, 159, 62, 93, 239, 209, 128, 163, 179, 11, 67, 103, 126, 64, 143, 135, 158, 177, 82, 100, 21, 115, 225, 232, 94, 22, 61, 55, 134, 212, 59, 113, 122, 251, 220, 188, 125, 121, 236, 219, 205, 212, 105, 222, 158, 21, 175, 76, 226, 212, 238, 173, 100, 103, 24, 111, 193, 247, 175, 211, 136, 233, 11, 255, 226, 191, 189, 191, 241, 203, 219, 211, 11, 139, 211, 129, 142, 192, 127, 230, 136, 95, 8, 123, 36, 9, 128, 176, 134, 199, 129, 239, 116, 11, 21, 10, 5, 247, 207, 250, 152, 158, 147, 158, 51, 123, 0, 87, 78, 29, 229, 135, 39, 135, 234, 173, 38, 120, 223, 125, 247, 241, 205, 55, 223, 16, 28, 28, 108, 228, 72, 235, 43, 40, 40, 96, 209, 162, 69, 188, 248, 226, 139, 36, 36, 104, 47, 135, 220, 105, 196, 195, 140, 123, 235, 123, 179, 39, 79, 230, 112, 231, 234, 5, 22, 60, 53, 146, 235, 49, 39, 244, 246, 185, 184, 123, 226, 87, 187, 1, 55, 206, 157, 50, 58, 148, 16, 84, 67, 39, 135, 206, 252, 128, 186, 45, 58, 240, 245, 228, 94, 154, 43, 63, 78, 6, 150, 152, 35, 110, 33, 236, 149, 36, 0, 194, 210, 70, 162, 90, 5, 78, 171, 217, 95, 161, 84, 50, 250, 181, 249, 116, 29, 251, 184, 217, 3, 56, 123, 104, 23, 11, 159, 30, 69, 102, 106, 146, 186, 204, 209, 209, 145, 87, 95, 125, 149, 55, 222, 120, 3, 165, 157, 244, 174, 191, 124, 249, 50, 227, 199, 143, 103, 223, 190, 125, 90, 229, 45, 122, 14, 97, 210, 39, 63, 219, 229, 36, 66, 89, 105, 41, 44, 155, 29, 78, 212, 206, 77, 37, 87, 54, 160, 237, 192, 177, 140, 120, 233, 51, 62, 27, 23, 70, 82, 220, 213, 194, 226, 5, 192, 84, 83, 197, 40, 68, 101, 97, 159, 3, 138, 133, 189, 234, 129, 106, 237, 119, 173, 7, 234, 74, 165, 3, 227, 223, 254, 145, 206, 15, 60, 106, 246, 0, 34, 183, 173, 99, 209, 179, 15, 104, 53, 33, 215, 171, 87, 143, 45, 91, 182, 16, 30, 30, 110, 87, 147, 237, 248, 248, 248, 48, 105, 210, 36, 210, 210, 210, 136, 136, 136, 80, 151, 199, 199, 198, 112, 246, 224, 110, 90, 247, 27, 89, 108, 47, 122, 91, 228, 232, 236, 66, 219, 129, 99, 72, 186, 121, 149, 171, 255, 253, 91, 166, 99, 235, 133, 134, 49, 113, 238, 18, 126, 156, 49, 140, 248, 187, 243, 67, 0, 7, 129, 113, 200, 152, 127, 33, 244, 72, 2, 32, 44, 165, 53, 240, 7, 224, 169, 89, 168, 84, 58, 240, 208, 188, 101, 116, 24, 250, 160, 217, 3, 56, 181, 103, 43, 75, 102, 142, 213, 26, 219, 223, 170, 85, 43, 118, 238, 220, 73, 139, 22, 45, 204, 126, 125, 115, 112, 112, 112, 96, 192, 128, 1, 52, 108, 216, 144, 205, 155, 55, 147, 127, 183, 121, 60, 241, 198, 101, 206, 30, 220, 69, 251, 33, 19, 138, 91, 170, 215, 38, 41, 20, 74, 90, 246, 26, 70, 86, 90, 50, 177, 199, 35, 74, 62, 0, 168, 25, 220, 130, 41, 95, 173, 103, 213, 27, 83, 57, 127, 84, 189, 254, 192, 117, 160, 63, 112, 199, 248, 145, 66, 84, 93, 146, 0, 8, 75, 240, 5, 118, 160, 154, 233, 79, 203, 136, 217, 159, 209, 121, 212, 20, 179, 7, 16, 27, 121, 128, 5, 51, 134, 147, 155, 173, 126, 38, 76, 207, 158, 61, 217, 182, 109, 27, 1, 1, 166, 93, 17, 208, 26, 218, 180, 105, 67, 187, 118, 237, 88, 191, 126, 189, 122, 148, 64, 82, 220, 85, 174, 157, 62, 78, 219, 129, 99, 237, 110, 210, 32, 133, 66, 65, 179, 238, 3, 112, 114, 113, 227, 76, 196, 95, 197, 214, 245, 175, 211, 136, 135, 62, 90, 206, 230, 79, 103, 243, 223, 223, 91, 11, 139, 147, 129, 1, 192, 105, 227, 71, 10, 81, 181, 73, 2, 32, 204, 77, 1, 172, 198, 192, 244, 190, 247, 61, 249, 58, 125, 31, 125, 201, 236, 1, 220, 60, 119, 138, 111, 167, 222, 71, 102, 106, 178, 186, 108, 192, 128, 1, 108, 222, 188, 25, 79, 79, 207, 98, 142, 180, 47, 33, 33, 33, 220, 115, 207, 61, 172, 90, 181, 74, 157, 4, 196, 199, 198, 144, 112, 253, 18, 161, 125, 71, 88, 57, 186, 242, 105, 216, 190, 59, 46, 238, 94, 156, 222, 103, 120, 197, 71, 79, 223, 0, 198, 188, 249, 63, 14, 172, 253, 145, 227, 127, 254, 82, 88, 156, 129, 106, 114, 169, 210, 53, 31, 8, 81, 69, 73, 2, 32, 204, 237, 21, 224, 73, 221, 194, 174, 99, 30, 99, 248, 172, 143, 205, 126, 241, 236, 140, 52, 254, 55, 117, 128, 214, 152, 241, 176, 176, 48, 182, 108, 217, 130, 187, 187, 253, 117, 146, 43, 73, 131, 6, 13, 104, 211, 166, 13, 107, 214, 172, 81, 63, 14, 184, 118, 250, 56, 30, 213, 170, 155, 100, 9, 100, 107, 104, 208, 182, 43, 0, 231, 14, 237, 214, 219, 231, 232, 236, 66, 106, 226, 45, 205, 155, 127, 14, 48, 26, 40, 205, 26, 209, 66, 84, 105, 146, 0, 8, 115, 106, 13, 172, 64, 231, 247, 44, 184, 99, 79, 38, 125, 178, 210, 34, 205, 210, 171, 222, 120, 140, 152, 3, 59, 212, 219, 77, 154, 52, 97, 251, 246, 237, 248, 250, 90, 108, 25, 122, 139, 107, 218, 180, 41, 141, 26, 53, 98, 253, 250, 245, 234, 178, 51, 251, 183, 19, 210, 181, 31, 213, 106, 214, 177, 98, 100, 229, 215, 184, 83, 47, 178, 211, 211, 184, 120, 108, 191, 86, 121, 78, 86, 166, 102, 135, 191, 60, 32, 28, 88, 103, 225, 240, 132, 176, 75, 146, 0, 8, 115, 113, 0, 54, 161, 179, 184, 143, 151, 127, 13, 166, 253, 240, 59, 174, 94, 62, 102, 15, 224, 192, 186, 133, 108, 251, 238, 61, 245, 182, 143, 143, 15, 127, 255, 253, 55, 117, 234, 216, 231, 77, 176, 44, 90, 183, 110, 77, 90, 90, 154, 122, 136, 96, 65, 126, 62, 103, 15, 237, 38, 108, 196, 36, 28, 157, 109, 107, 114, 163, 210, 106, 218, 181, 31, 201, 241, 215, 184, 114, 234, 168, 161, 221, 121, 168, 122, 251, 175, 177, 108, 84, 66, 216, 47, 73, 0, 132, 185, 204, 68, 53, 249, 138, 154, 66, 169, 100, 234, 55, 27, 9, 106, 18, 106, 246, 139, 39, 222, 188, 194, 130, 25, 247, 107, 245, 248, 95, 180, 104, 17, 61, 122, 244, 48, 251, 181, 109, 69, 239, 222, 189, 217, 177, 99, 7, 151, 47, 171, 30, 127, 100, 36, 39, 144, 158, 116, 135, 150, 189, 202, 190, 220, 174, 45, 80, 40, 20, 52, 239, 49, 152, 43, 81, 71, 136, 143, 141, 209, 221, 173, 4, 142, 2, 134, 87, 79, 18, 66, 232, 177, 175, 174, 193, 194, 94, 212, 1, 222, 214, 45, 188, 103, 194, 116, 130, 59, 246, 180, 72, 0, 235, 63, 120, 94, 107, 69, 191, 25, 51, 102, 48, 97, 194, 4, 139, 92, 219, 86, 56, 58, 58, 178, 114, 229, 74, 170, 87, 175, 174, 46, 139, 88, 187, 160, 212, 67, 235, 108, 145, 106, 216, 232, 114, 106, 4, 27, 28, 182, 57, 23, 85, 231, 63, 33, 68, 41, 72, 2, 32, 204, 225, 109, 64, 171, 135, 157, 111, 80, 61, 6, 63, 243, 174, 69, 46, 126, 122, 223, 54, 34, 183, 21, 61, 6, 174, 95, 191, 62, 31, 126, 248, 161, 69, 174, 109, 107, 234, 214, 173, 203, 39, 159, 124, 162, 222, 46, 200, 207, 103, 237, 123, 79, 147, 159, 159, 87, 204, 81, 182, 205, 213, 211, 155, 71, 191, 250, 21, 119, 31, 63, 221, 93, 74, 96, 57, 208, 220, 242, 81, 9, 97, 127, 36, 1, 16, 166, 214, 10, 152, 164, 91, 248, 192, 235, 243, 113, 241, 240, 50, 251, 197, 243, 243, 114, 89, 55, 87, 123, 65, 156, 249, 243, 231, 227, 225, 225, 97, 246, 107, 219, 170, 240, 240, 112, 173, 71, 31, 87, 78, 29, 229, 192, 218, 133, 86, 140, 168, 226, 170, 215, 107, 204, 195, 159, 174, 54, 180, 230, 129, 55, 176, 1, 213, 220, 19, 66, 136, 98, 72, 2, 32, 76, 237, 3, 116, 250, 150, 52, 235, 126, 31, 45, 122, 12, 182, 200, 197, 255, 253, 109, 21, 241, 23, 207, 168, 183, 71, 142, 28, 201, 144, 33, 67, 44, 114, 109, 91, 165, 80, 40, 248, 230, 155, 111, 112, 114, 114, 82, 151, 109, 255, 126, 174, 86, 255, 8, 123, 212, 164, 115, 111, 70, 204, 254, 204, 224, 46, 96, 25, 178, 214, 137, 16, 197, 146, 78, 128, 194, 148, 186, 0, 90, 109, 237, 10, 165, 146, 201, 159, 173, 198, 171, 122, 77, 179, 95, 188, 32, 63, 159, 101, 47, 133, 147, 122, 39, 30, 80, 77, 147, 187, 97, 195, 6, 252, 252, 244, 154, 138, 171, 156, 192, 192, 64, 174, 95, 191, 206, 225, 195, 135, 1, 200, 76, 77, 162, 122, 221, 96, 106, 55, 107, 107, 229, 200, 42, 166, 94, 104, 39, 146, 226, 12, 142, 12, 104, 2, 196, 3, 135, 44, 31, 149, 16, 246, 65, 90, 0, 132, 41, 205, 214, 45, 232, 56, 236, 33, 106, 133, 180, 177, 200, 197, 79, 238, 220, 200, 141, 179, 81, 234, 237, 241, 227, 199, 211, 184, 113, 99, 139, 92, 219, 30, 204, 158, 61, 27, 103, 231, 162, 117, 1, 182, 125, 63, 215, 174, 251, 2, 20, 122, 224, 181, 175, 141, 117, 46, 253, 24, 104, 105, 225, 112, 132, 176, 27, 146, 0, 8, 83, 105, 6, 12, 211, 44, 80, 58, 56, 50, 96, 250, 155, 22, 11, 224, 239, 101, 95, 170, 223, 43, 20, 10, 102, 207, 214, 203, 71, 170, 180, 122, 245, 234, 49, 113, 226, 68, 245, 246, 173, 75, 103, 57, 189, 247, 79, 43, 70, 100, 26, 14, 142, 78, 76, 250, 120, 5, 158, 126, 129, 186, 187, 92, 129, 69, 72, 75, 167, 16, 6, 73, 2, 32, 76, 101, 22, 58, 191, 79, 237, 6, 141, 195, 175, 118, 3, 139, 92, 60, 225, 90, 44, 231, 142, 168, 87, 129, 99, 200, 144, 33, 180, 106, 213, 202, 34, 215, 182, 39, 255, 247, 127, 255, 167, 181, 125, 104, 227, 82, 43, 69, 98, 90, 94, 213, 107, 50, 225, 189, 69, 134, 150, 115, 238, 4, 60, 110, 133, 144, 132, 176, 121, 146, 0, 8, 83, 168, 1, 104, 173, 231, 171, 80, 40, 232, 253, 200, 11, 22, 11, 224, 208, 198, 159, 40, 184, 59, 247, 61, 192, 35, 143, 60, 98, 177, 107, 219, 147, 230, 205, 155, 211, 185, 115, 209, 186, 76, 39, 119, 110, 36, 35, 37, 209, 138, 17, 153, 78, 243, 123, 7, 210, 249, 129, 71, 13, 237, 122, 23, 168, 110, 104, 135, 16, 85, 153, 36, 0, 194, 20, 30, 1, 92, 52, 11, 66, 186, 223, 103, 177, 103, 255, 0, 71, 54, 47, 87, 191, 247, 243, 243, 171, 242, 61, 255, 139, 51, 105, 82, 209, 40, 205, 220, 172, 76, 173, 57, 19, 236, 221, 240, 89, 31, 227, 19, 88, 91, 183, 216, 15, 85, 18, 32, 132, 208, 32, 9, 128, 168, 40, 37, 240, 152, 110, 97, 175, 135, 103, 90, 44, 128, 184, 139, 167, 181, 134, 254, 141, 27, 55, 14, 23, 23, 151, 98, 142, 168, 218, 198, 141, 27, 167, 213, 25, 48, 106, 215, 38, 43, 70, 99, 90, 46, 30, 94, 12, 123, 193, 224, 164, 79, 143, 2, 193, 22, 14, 71, 8, 155, 166, 55, 139, 134, 16, 101, 212, 31, 104, 164, 89, 224, 95, 167, 17, 77, 58, 247, 177, 88, 0, 103, 15, 236, 212, 218, 30, 62, 124, 184, 197, 174, 109, 143, 252, 253, 253, 233, 222, 189, 59, 59, 119, 170, 126, 110, 231, 14, 237, 33, 63, 63, 15, 165, 210, 54, 251, 202, 101, 36, 39, 112, 243, 252, 127, 196, 95, 140, 33, 229, 246, 77, 178, 210, 83, 201, 205, 206, 66, 161, 84, 226, 230, 85, 13, 23, 119, 79, 170, 215, 107, 76, 96, 195, 16, 124, 107, 213, 167, 253, 144, 9, 68, 252, 242, 3, 103, 181, 151, 15, 118, 4, 94, 6, 166, 90, 231, 83, 8, 97, 123, 36, 1, 16, 21, 165, 215, 193, 170, 235, 216, 199, 45, 178, 212, 111, 161, 152, 131, 69, 9, 128, 147, 147, 19, 247, 220, 115, 143, 197, 174, 109, 175, 122, 247, 238, 173, 78, 0, 50, 82, 18, 185, 250, 223, 191, 212, 109, 217, 209, 202, 81, 169, 100, 166, 38, 19, 189, 247, 119, 206, 236, 223, 206, 153, 136, 191, 184, 115, 245, 98, 169, 143, 117, 113, 247, 36, 184, 99, 15, 106, 55, 111, 167, 155, 0, 128, 106, 134, 202, 247, 128, 11, 166, 139, 86, 8, 251, 37, 9, 128, 168, 8, 63, 116, 22, 95, 113, 112, 114, 38, 108, 196, 195, 22, 11, 160, 160, 160, 128, 115, 7, 139, 254, 163, 15, 11, 11, 171, 210, 211, 254, 150, 86, 223, 190, 125, 121, 227, 141, 55, 212, 219, 49, 7, 118, 90, 61, 1, 56, 127, 228, 111, 34, 214, 46, 32, 242, 207, 181, 100, 103, 166, 151, 235, 28, 89, 233, 169, 156, 218, 179, 21, 246, 108, 53, 180, 219, 9, 120, 14, 120, 182, 2, 97, 10, 81, 105, 72, 2, 32, 42, 226, 1, 192, 89, 179, 32, 180, 207, 112, 67, 227, 177, 205, 38, 241, 198, 101, 82, 19, 226, 213, 219, 61, 123, 90, 102, 181, 65, 123, 215, 169, 83, 39, 220, 221, 221, 73, 79, 87, 221, 104, 13, 204, 164, 103, 49, 49, 17, 59, 248, 125, 254, 28, 46, 252, 91, 252, 74, 190, 78, 78, 78, 4, 5, 5, 225, 233, 233, 137, 155, 155, 27, 121, 121, 121, 36, 39, 39, 115, 235, 214, 45, 146, 147, 147, 75, 123, 185, 73, 168, 38, 172, 202, 168, 96, 216, 66, 216, 61, 73, 0, 68, 69, 140, 215, 45, 232, 48, 244, 65, 67, 245, 204, 70, 179, 243, 31, 64, 104, 104, 168, 69, 175, 111, 175, 156, 156, 156, 8, 9, 9, 225, 223, 127, 255, 5, 244, 127, 142, 150, 112, 251, 202, 121, 126, 121, 123, 58, 167, 247, 109, 51, 184, 191, 105, 211, 166, 244, 235, 215, 143, 94, 189, 122, 17, 26, 26, 74, 112, 112, 176, 214, 122, 6, 154, 174, 95, 191, 206, 169, 83, 167, 216, 191, 127, 63, 219, 182, 109, 35, 34, 34, 130, 236, 108, 131, 107, 29, 84, 67, 181, 94, 133, 180, 2, 136, 42, 79, 18, 0, 81, 94, 65, 128, 214, 215, 109, 55, 111, 95, 154, 221, 51, 192, 162, 65, 196, 93, 56, 173, 181, 29, 18, 18, 98, 209, 235, 219, 51, 221, 4, 160, 32, 63, 223, 34, 125, 55, 10, 10, 10, 216, 179, 244, 115, 126, 251, 234, 13, 189, 166, 126, 63, 63, 63, 194, 195, 195, 153, 50, 101, 10, 173, 91, 183, 46, 245, 57, 131, 130, 130, 8, 10, 10, 162, 111, 223, 190, 188, 246, 218, 107, 220, 185, 115, 135, 229, 203, 151, 179, 96, 193, 2, 142, 31, 63, 174, 91, 253, 25, 160, 30, 240, 4, 16, 87, 193, 143, 35, 132, 221, 146, 97, 128, 162, 188, 198, 162, 51, 197, 106, 235, 126, 35, 113, 112, 114, 54, 82, 221, 60, 226, 46, 22, 37, 0, 10, 133, 130, 38, 77, 154, 88, 244, 250, 246, 76, 51, 89, 202, 206, 76, 39, 41, 238, 170, 217, 175, 153, 149, 150, 194, 146, 153, 99, 217, 240, 209, 139, 90, 55, 127, 127, 127, 127, 222, 124, 243, 77, 206, 159, 63, 207, 231, 159, 127, 94, 166, 155, 191, 33, 126, 126, 126, 60, 253, 244, 211, 28, 59, 118, 140, 109, 219, 182, 17, 22, 22, 166, 91, 101, 4, 112, 20, 232, 90, 161, 11, 9, 97, 199, 36, 1, 16, 229, 53, 82, 183, 160, 221, 32, 189, 39, 2, 102, 151, 28, 119, 77, 253, 62, 48, 48, 16, 79, 79, 79, 139, 199, 96, 175, 130, 131, 181, 135, 197, 39, 222, 188, 98, 214, 235, 221, 190, 114, 158, 79, 199, 118, 210, 154, 120, 72, 161, 80, 240, 196, 19, 79, 112, 238, 220, 57, 230, 204, 153, 131, 143, 143, 143, 201, 175, 219, 175, 95, 63, 14, 28, 56, 192, 146, 37, 75, 168, 94, 93, 107, 66, 192, 218, 192, 46, 192, 114, 189, 86, 133, 176, 33, 242, 8, 64, 148, 135, 47, 208, 93, 179, 192, 211, 55, 128, 198, 157, 123, 89, 60, 144, 204, 180, 20, 245, 123, 115, 220, 60, 172, 33, 49, 49, 145, 200, 200, 72, 174, 94, 189, 74, 114, 114, 50, 254, 254, 254, 212, 170, 85, 139, 246, 237, 219, 227, 234, 234, 106, 178, 235, 232, 254, 188, 178, 211, 211, 76, 118, 110, 93, 113, 23, 162, 249, 246, 209, 251, 180, 90, 25, 234, 212, 169, 195, 242, 229, 203, 233, 209, 163, 135, 217, 174, 171, 105, 210, 164, 73, 12, 26, 52, 136, 169, 83, 167, 178, 113, 227, 198, 194, 98, 103, 84, 11, 6, 121, 1, 95, 91, 36, 16, 33, 108, 132, 36, 0, 162, 60, 6, 161, 243, 187, 211, 188, 231, 96, 171, 76, 36, 147, 149, 158, 170, 126, 239, 229, 229, 101, 241, 235, 155, 74, 126, 126, 62, 63, 255, 252, 51, 139, 23, 47, 102, 215, 174, 93, 228, 228, 228, 232, 213, 241, 240, 240, 96, 224, 192, 129, 60, 249, 228, 147, 244, 237, 219, 183, 194, 215, 212, 253, 121, 105, 38, 83, 166, 20, 119, 241, 52, 95, 63, 220, 155, 212, 59, 69, 143, 219, 123, 245, 234, 197, 170, 85, 171, 8, 12, 180, 220, 136, 17, 128, 128, 128, 0, 214, 175, 95, 207, 220, 185, 115, 121, 227, 141, 55, 200, 87, 173, 31, 161, 0, 190, 66, 245, 59, 253, 185, 69, 3, 18, 194, 138, 228, 17, 128, 40, 143, 161, 186, 5, 45, 122, 88, 103, 238, 125, 205, 4, 192, 94, 155, 255, 35, 34, 34, 232, 216, 177, 35, 15, 62, 248, 32, 219, 182, 109, 51, 120, 243, 7, 72, 75, 75, 99, 237, 218, 181, 244, 235, 215, 143, 129, 3, 7, 114, 254, 252, 249, 10, 93, 87, 55, 1, 200, 50, 67, 2, 144, 150, 112, 139, 5, 51, 134, 107, 221, 252, 135, 15, 31, 206, 239, 191, 255, 110, 241, 155, 127, 33, 133, 66, 193, 171, 175, 190, 202, 138, 21, 43, 116, 71, 21, 124, 130, 106, 104, 171, 16, 85, 130, 36, 0, 162, 172, 28, 129, 129, 154, 5, 14, 78, 206, 52, 235, 126, 159, 85, 130, 201, 215, 184, 89, 106, 206, 111, 111, 47, 150, 44, 89, 66, 175, 94, 189, 212, 189, 241, 75, 235, 143, 63, 254, 32, 44, 44, 76, 61, 155, 95, 121, 232, 254, 188, 114, 179, 179, 202, 125, 46, 67, 242, 114, 178, 249, 113, 198, 253, 196, 199, 198, 168, 203, 194, 195, 195, 89, 187, 118, 173, 77, 172, 213, 48, 110, 220, 56, 214, 174, 93, 171, 249, 115, 80, 2, 63, 1, 182, 49, 37, 162, 16, 102, 38, 9, 128, 40, 171, 206, 168, 250, 0, 168, 5, 119, 236, 129, 139, 135, 117, 154, 223, 157, 220, 220, 213, 239, 211, 210, 204, 247, 12, 219, 28, 22, 46, 92, 200, 228, 201, 147, 201, 202, 42, 223, 141, 247, 246, 237, 219, 12, 28, 56, 144, 61, 123, 246, 148, 235, 248, 212, 212, 84, 173, 109, 83, 255, 29, 110, 249, 226, 85, 98, 35, 15, 168, 183, 7, 12, 24, 192, 194, 133, 11, 113, 112, 176, 157, 53, 7, 134, 13, 27, 198, 226, 197, 139, 81, 40, 20, 133, 69, 110, 192, 207, 168, 250, 4, 8, 81, 169, 73, 2, 32, 202, 170, 183, 110, 65, 139, 30, 131, 173, 17, 7, 0, 174, 26, 55, 173, 148, 20, 243, 60, 195, 54, 135, 189, 123, 247, 242, 228, 147, 79, 86, 248, 60, 217, 217, 217, 140, 30, 61, 154, 216, 216, 216, 50, 31, 171, 251, 243, 114, 53, 97, 2, 112, 102, 255, 95, 236, 94, 90, 244, 56, 189, 85, 171, 86, 172, 89, 179, 6, 71, 71, 219, 235, 118, 52, 97, 194, 4, 94, 121, 229, 21, 205, 162, 96, 224, 11, 43, 133, 35, 132, 197, 72, 2, 32, 202, 74, 111, 153, 191, 38, 93, 44, 183, 242, 159, 46, 103, 247, 162, 121, 255, 117, 191, 209, 218, 170, 188, 188, 60, 166, 77, 155, 102, 108, 166, 186, 50, 139, 143, 143, 103, 230, 204, 178, 47, 191, 172, 155, 0, 40, 16, 238, 223, 0, 0, 32, 0, 73, 68, 65, 84, 184, 120, 152, 166, 15, 69, 94, 78, 54, 107, 223, 157, 65, 129, 170, 131, 29, 174, 174, 174, 44, 95, 190, 220, 166, 59, 105, 190, 253, 246, 219, 244, 234, 213, 75, 179, 232, 17, 84, 43, 93, 10, 81, 105, 73, 2, 32, 202, 194, 21, 157, 137, 83, 60, 125, 3, 168, 217, 184, 149, 149, 194, 1, 55, 175, 106, 234, 247, 241, 241, 241, 20, 20, 20, 88, 45, 150, 210, 90, 178, 100, 9, 81, 81, 81, 38, 61, 231, 175, 191, 254, 74, 68, 68, 68, 153, 142, 137, 139, 211, 158, 4, 207, 213, 211, 52, 195, 40, 119, 45, 254, 84, 235, 185, 255, 188, 121, 243, 42, 60, 177, 143, 185, 41, 149, 74, 150, 44, 89, 162, 59, 52, 242, 99, 116, 38, 187, 18, 162, 50, 145, 4, 64, 148, 69, 87, 84, 73, 128, 90, 112, 88, 79, 205, 231, 167, 22, 87, 189, 110, 209, 100, 54, 169, 169, 169, 92, 187, 118, 173, 152, 218, 182, 97, 193, 130, 5, 38, 63, 103, 65, 65, 1, 63, 254, 248, 99, 153, 142, 57, 125, 90, 123, 26, 101, 255, 58, 141, 42, 28, 71, 106, 66, 60, 219, 127, 120, 95, 189, 221, 190, 125, 123, 102, 204, 152, 81, 225, 243, 90, 66, 189, 122, 245, 180, 86, 72, 4, 90, 3, 83, 172, 20, 142, 16, 102, 39, 9, 128, 40, 139, 94, 186, 5, 77, 194, 244, 186, 4, 88, 84, 64, 3, 237, 185, 255, 117, 111, 106, 182, 230, 230, 205, 155, 101, 254, 166, 94, 90, 155, 54, 109, 42, 28, 215, 94, 42, 154, 63, 43, 239, 192, 90, 184, 122, 122, 87, 56, 134, 127, 86, 126, 163, 30, 154, 169, 80, 40, 248, 226, 139, 47, 80, 90, 96, 125, 1, 83, 121, 250, 233, 167, 117, 215, 147, 152, 133, 252, 63, 41, 42, 41, 249, 197, 22, 101, 209, 77, 183, 160, 113, 88, 47, 43, 132, 81, 36, 176, 97, 83, 173, 237, 232, 232, 104, 43, 69, 82, 58, 7, 15, 30, 44, 211, 77, 186, 44, 226, 226, 226, 184, 112, 225, 66, 169, 235, 159, 57, 83, 180, 2, 96, 96, 131, 166, 197, 212, 44, 157, 236, 204, 116, 246, 174, 252, 70, 189, 61, 96, 192, 0, 238, 185, 231, 158, 10, 159, 215, 146, 156, 156, 156, 116, 91, 1, 154, 96, 96, 218, 107, 33, 42, 3, 73, 0, 68, 105, 41, 1, 173, 21, 85, 60, 170, 249, 235, 125, 3, 183, 180, 192, 6, 33, 90, 43, 216, 29, 58, 116, 200, 138, 209, 148, 236, 234, 85, 243, 46, 184, 83, 218, 243, 199, 197, 197, 113, 241, 226, 69, 245, 118, 96, 195, 102, 21, 190, 118, 228, 159, 107, 73, 75, 184, 165, 222, 126, 241, 197, 23, 43, 124, 78, 107, 24, 59, 118, 44, 245, 235, 215, 215, 44, 170, 248, 112, 13, 33, 108, 144, 36, 0, 162, 180, 90, 2, 90, 109, 196, 245, 91, 119, 182, 234, 243, 127, 80, 141, 93, 175, 213, 180, 168, 131, 217, 142, 29, 59, 172, 24, 77, 201, 18, 18, 18, 204, 122, 254, 219, 183, 111, 151, 170, 222, 142, 29, 59, 180, 58, 76, 54, 106, 95, 241, 111, 234, 71, 183, 172, 80, 191, 15, 13, 13, 53, 201, 116, 197, 214, 224, 232, 232, 168, 219, 111, 161, 23, 170, 229, 175, 133, 168, 84, 36, 1, 16, 165, 165, 183, 108, 106, 253, 54, 93, 172, 17, 135, 158, 198, 157, 139, 250, 33, 92, 186, 116, 169, 194, 83, 228, 154, 147, 191, 191, 191, 89, 207, 31, 16, 16, 80, 170, 122, 154, 137, 146, 66, 161, 168, 240, 163, 156, 212, 59, 113, 156, 137, 40, 58, 231, 196, 137, 19, 43, 116, 62, 107, 27, 63, 126, 188, 102, 223, 5, 7, 96, 140, 21, 195, 17, 194, 44, 36, 1, 16, 165, 213, 89, 183, 160, 126, 107, 189, 34, 171, 104, 210, 89, 123, 30, 130, 237, 219, 183, 91, 41, 146, 146, 213, 170, 85, 203, 38, 206, 255, 215, 95, 127, 169, 223, 7, 54, 106, 142, 119, 64, 197, 190, 224, 158, 222, 183, 157, 252, 188, 92, 64, 149, 80, 76, 152, 48, 161, 66, 231, 179, 182, 186, 117, 235, 234, 246, 95, 24, 100, 173, 88, 132, 48, 23, 73, 0, 68, 105, 105, 221, 237, 21, 74, 37, 245, 66, 59, 89, 43, 22, 45, 141, 218, 223, 131, 131, 83, 209, 188, 246, 43, 87, 174, 180, 98, 52, 197, 235, 218, 181, 171, 217, 102, 195, 171, 85, 171, 22, 13, 27, 54, 44, 177, 222, 129, 3, 7, 180, 90, 73, 154, 118, 169, 120, 83, 253, 185, 195, 187, 213, 239, 155, 55, 111, 174, 251, 12, 221, 46, 13, 24, 48, 64, 115, 243, 30, 100, 245, 84, 81, 201, 72, 2, 32, 74, 195, 21, 208, 234, 237, 23, 216, 32, 196, 100, 19, 199, 84, 148, 171, 167, 55, 205, 239, 45, 90, 159, 104, 247, 238, 221, 101, 234, 13, 111, 73, 254, 254, 254, 116, 237, 170, 247, 52, 197, 36, 134, 13, 27, 86, 170, 62, 25, 63, 253, 244, 147, 214, 118, 187, 193, 227, 42, 124, 237, 243, 71, 246, 170, 223, 247, 232, 209, 163, 194, 231, 179, 5, 58, 51, 3, 122, 2, 237, 173, 19, 137, 16, 230, 33, 9, 128, 40, 141, 22, 232, 124, 251, 169, 21, 98, 91, 51, 187, 117, 28, 22, 174, 126, 95, 80, 80, 192, 138, 21, 43, 138, 169, 109, 93, 143, 63, 254, 184, 201, 207, 169, 80, 40, 120, 236, 177, 199, 74, 172, 151, 147, 147, 195, 170, 85, 171, 212, 219, 254, 117, 131, 169, 223, 186, 98, 125, 57, 114, 179, 50, 137, 143, 45, 26, 82, 216, 173, 155, 222, 104, 81, 187, 212, 169, 83, 39, 221, 21, 19, 67, 173, 21, 139, 16, 230, 32, 9, 128, 40, 141, 182, 186, 5, 181, 155, 233, 21, 89, 85, 139, 158, 67, 112, 247, 241, 83, 111, 255, 240, 195, 15, 228, 104, 44, 21, 108, 75, 38, 78, 156, 72, 251, 246, 166, 253, 50, 57, 97, 194, 4, 58, 116, 232, 80, 98, 189, 53, 107, 214, 112, 235, 86, 209, 80, 189, 142, 195, 30, 170, 240, 72, 142, 91, 151, 207, 169, 231, 253, 7, 213, 35, 128, 202, 192, 201, 201, 137, 70, 141, 180, 102, 71, 108, 98, 173, 88, 132, 48, 7, 73, 0, 68, 105, 180, 209, 45, 168, 21, 162, 87, 100, 85, 142, 206, 46, 116, 24, 90, 212, 243, 60, 54, 54, 150, 229, 203, 151, 91, 49, 34, 227, 148, 74, 37, 223, 125, 247, 29, 110, 110, 110, 38, 57, 95, 173, 90, 181, 248, 232, 163, 143, 74, 172, 151, 159, 159, 207, 220, 185, 115, 53, 226, 112, 160, 211, 240, 240, 98, 142, 40, 157, 91, 151, 206, 105, 109, 55, 110, 220, 184, 194, 231, 180, 21, 77, 155, 106, 77, 144, 36, 9, 128, 168, 84, 36, 1, 16, 165, 161, 215, 222, 111, 107, 143, 0, 0, 122, 77, 126, 1, 7, 71, 39, 245, 246, 123, 239, 189, 71, 94, 94, 158, 21, 35, 50, 174, 99, 199, 142, 252, 248, 227, 143, 21, 254, 246, 237, 234, 234, 202, 186, 117, 235, 74, 213, 251, 127, 253, 250, 245, 90, 139, 16, 181, 29, 52, 22, 191, 218, 37, 119, 26, 44, 73, 106, 66, 188, 250, 189, 167, 167, 39, 213, 170, 85, 43, 166, 182, 125, 169, 83, 167, 142, 230, 102, 233, 198, 88, 10, 97, 39, 36, 1, 16, 165, 161, 245, 236, 211, 211, 55, 160, 194, 195, 198, 204, 193, 55, 168, 30, 237, 135, 20, 13, 63, 59, 123, 246, 172, 77, 143, 8, 152, 56, 113, 34, 63, 255, 252, 51, 238, 238, 238, 229, 58, 190, 102, 205, 154, 236, 216, 177, 131, 206, 157, 75, 30, 142, 153, 151, 151, 199, 59, 239, 188, 163, 222, 86, 40, 20, 244, 157, 58, 187, 92, 215, 213, 149, 157, 94, 180, 12, 179, 45, 47, 249, 91, 30, 158, 158, 90, 75, 36, 155, 102, 189, 100, 33, 108, 132, 36, 0, 162, 36, 213, 1, 173, 217, 107, 130, 154, 90, 111, 249, 223, 146, 244, 157, 250, 146, 214, 212, 192, 179, 102, 205, 34, 41, 41, 201, 138, 17, 21, 111, 236, 216, 177, 236, 221, 187, 183, 204, 29, 231, 70, 141, 26, 197, 161, 67, 135, 74, 61, 162, 224, 219, 111, 191, 229, 216, 177, 99, 234, 237, 208, 190, 35, 8, 106, 98, 154, 191, 199, 236, 140, 52, 245, 123, 15, 15, 15, 147, 156, 211, 86, 232, 36, 52, 149, 43, 187, 17, 85, 158, 140, 107, 21, 37, 209, 155, 236, 223, 218, 243, 255, 23, 39, 176, 97, 51, 58, 13, 159, 196, 193, 95, 23, 3, 112, 227, 198, 13, 222, 124, 243, 77, 62, 255, 252, 115, 235, 6, 86, 140, 118, 237, 218, 241, 207, 63, 255, 176, 126, 253, 122, 22, 47, 94, 204, 182, 109, 219, 72, 79, 79, 215, 171, 231, 231, 231, 199, 224, 193, 131, 121, 242, 201, 39, 203, 148, 48, 220, 188, 121, 147, 215, 95, 127, 93, 189, 237, 224, 232, 196, 128, 25, 115, 76, 17, 186, 30, 115, 45, 116, 84, 86, 231, 206, 157, 99, 211, 166, 77, 156, 59, 119, 142, 184, 184, 56, 2, 3, 3, 105, 212, 168, 17, 195, 134, 13, 43, 83, 31, 5, 157, 207, 99, 27, 31, 78, 8, 19, 145, 4, 64, 148, 68, 239, 110, 31, 216, 208, 118, 19, 0, 128, 161, 51, 63, 224, 228, 142, 141, 164, 39, 221, 1, 96, 254, 252, 249, 132, 135, 135, 151, 170, 151, 188, 53, 141, 24, 49, 130, 17, 35, 70, 144, 158, 158, 206, 201, 147, 39, 185, 113, 227, 6, 183, 110, 221, 34, 48, 48, 144, 58, 117, 234, 208, 170, 85, 171, 114, 77, 34, 52, 115, 230, 76, 18, 19, 19, 213, 219, 61, 30, 122, 198, 100, 223, 254, 1, 156, 221, 139, 90, 198, 83, 83, 83, 139, 169, 105, 126, 199, 143, 31, 231, 165, 151, 94, 226, 143, 63, 254, 48, 184, 127, 230, 204, 153, 244, 239, 223, 159, 15, 63, 252, 144, 118, 237, 218, 149, 120, 190, 148, 148, 20, 173, 77, 211, 68, 41, 132, 109, 144, 4, 64, 148, 68, 191, 5, 160, 126, 197, 151, 142, 53, 39, 79, 223, 0, 6, 63, 251, 46, 191, 188, 61, 29, 128, 220, 220, 92, 198, 143, 31, 207, 145, 35, 71, 240, 246, 174, 248, 154, 247, 230, 230, 238, 238, 78, 88, 88, 88, 201, 21, 75, 97, 209, 162, 69, 90, 115, 34, 84, 171, 81, 135, 251, 166, 191, 81, 204, 17, 101, 231, 162, 145, 0, 232, 220, 48, 45, 106, 225, 194, 133, 60, 249, 228, 147, 100, 103, 103, 23, 91, 111, 219, 182, 109, 236, 222, 189, 155, 175, 190, 250, 170, 196, 57, 25, 116, 62, 143, 117, 179, 27, 33, 76, 76, 250, 0, 136, 146, 24, 104, 1, 176, 237, 4, 0, 160, 235, 232, 199, 104, 208, 166, 232, 249, 248, 217, 179, 103, 205, 50, 1, 143, 45, 139, 138, 138, 226, 169, 167, 158, 210, 42, 27, 249, 202, 23, 90, 55, 108, 83, 240, 242, 175, 161, 126, 159, 145, 145, 65, 124, 124, 124, 49, 181, 205, 99, 241, 226, 197, 60, 250, 232, 163, 37, 222, 252, 11, 101, 103, 103, 243, 196, 19, 79, 240, 227, 143, 63, 22, 91, 47, 54, 54, 86, 115, 243, 122, 249, 35, 20, 194, 246, 72, 2, 32, 74, 162, 117, 183, 119, 116, 118, 193, 183, 150, 237, 207, 243, 174, 80, 42, 121, 104, 222, 50, 220, 188, 125, 213, 101, 171, 86, 173, 226, 171, 175, 190, 178, 98, 84, 150, 147, 148, 148, 196, 216, 177, 99, 181, 250, 18, 116, 27, 251, 4, 161, 125, 71, 152, 252, 90, 1, 245, 181, 135, 199, 199, 196, 196, 152, 252, 26, 197, 137, 138, 138, 98, 218, 180, 105, 229, 58, 118, 198, 140, 25, 156, 56, 113, 194, 232, 254, 51, 103, 206, 104, 110, 90, 246, 131, 9, 97, 102, 146, 0, 136, 226, 40, 128, 6, 154, 5, 254, 117, 131, 81, 42, 29, 172, 19, 77, 25, 249, 213, 110, 192, 131, 239, 47, 209, 26, 107, 255, 220, 115, 207, 177, 102, 205, 26, 43, 70, 101, 126, 217, 217, 217, 140, 30, 61, 154, 83, 167, 78, 169, 203, 130, 154, 132, 50, 252, 165, 79, 204, 114, 61, 191, 58, 13, 181, 230, 95, 208, 188, 174, 37, 188, 242, 202, 43, 100, 101, 101, 149, 235, 216, 236, 236, 108, 102, 207, 54, 60, 28, 50, 61, 61, 157, 75, 151, 46, 105, 22, 73, 2, 32, 42, 21, 73, 0, 68, 113, 2, 1, 173, 233, 234, 252, 106, 213, 179, 82, 40, 229, 211, 162, 231, 16, 122, 132, 63, 171, 222, 206, 207, 207, 39, 60, 60, 156, 157, 59, 119, 90, 49, 42, 243, 201, 203, 203, 99, 194, 132, 9, 90, 75, 34, 187, 122, 122, 51, 249, 243, 53, 56, 185, 152, 102, 230, 65, 93, 14, 142, 78, 212, 212, 232, 84, 248, 247, 223, 127, 155, 229, 58, 134, 92, 191, 126, 157, 205, 155, 55, 87, 232, 28, 191, 253, 246, 155, 238, 141, 30, 128, 125, 251, 246, 145, 155, 155, 171, 89, 116, 76, 175, 146, 16, 118, 76, 18, 0, 81, 28, 189, 187, 189, 111, 144, 237, 55, 255, 235, 26, 246, 226, 60, 66, 251, 14, 87, 111, 103, 101, 101, 49, 98, 196, 8, 118, 237, 218, 101, 189, 160, 204, 32, 39, 39, 135, 201, 147, 39, 179, 110, 221, 58, 117, 153, 163, 179, 11, 143, 124, 177, 78, 175, 153, 222, 212, 130, 59, 220, 171, 126, 191, 103, 207, 30, 179, 94, 75, 211, 214, 173, 91, 43, 60, 244, 176, 160, 160, 128, 173, 91, 183, 234, 149, 235, 252, 126, 220, 1, 162, 244, 42, 9, 97, 199, 36, 1, 16, 197, 49, 144, 0, 216, 87, 11, 0, 168, 230, 188, 159, 244, 241, 207, 52, 237, 90, 180, 238, 125, 114, 114, 50, 3, 7, 14, 100, 245, 234, 213, 86, 140, 204, 116, 210, 210, 210, 24, 49, 98, 4, 203, 150, 45, 83, 151, 41, 148, 74, 30, 124, 127, 41, 77, 58, 247, 46, 249, 248, 196, 219, 164, 220, 190, 73, 70, 114, 2, 57, 89, 25, 101, 190, 126, 163, 142, 69, 75, 0, 95, 188, 120, 145, 147, 39, 79, 150, 249, 28, 229, 113, 246, 236, 89, 147, 156, 231, 220, 185, 115, 122, 101, 91, 182, 108, 209, 220, 220, 141, 204, 3, 32, 42, 25, 25, 6, 40, 138, 163, 247, 117, 191, 154, 29, 38, 0, 0, 14, 78, 206, 76, 254, 108, 13, 243, 39, 247, 225, 106, 180, 170, 37, 55, 43, 43, 139, 137, 19, 39, 18, 31, 31, 207, 140, 25, 51, 172, 28, 97, 249, 221, 184, 113, 131, 17, 35, 70, 112, 224, 192, 1, 117, 153, 66, 161, 96, 212, 43, 95, 210, 102, 192, 104, 131, 199, 156, 59, 188, 135, 83, 187, 54, 115, 38, 226, 47, 110, 93, 58, 75, 86, 186, 246, 8, 55, 239, 192, 90, 4, 214, 111, 66, 157, 150, 29, 104, 218, 165, 31, 141, 58, 222, 139, 179, 171, 241, 41, 139, 67, 186, 246, 195, 201, 197, 77, 157, 60, 172, 92, 185, 146, 247, 222, 123, 207, 4, 159, 174, 120, 154, 243, 27, 84, 196, 237, 219, 183, 181, 182, 163, 163, 163, 181, 102, 78, 4, 54, 154, 228, 66, 66, 216, 16, 105, 1, 16, 197, 209, 187, 219, 219, 91, 31, 0, 77, 174, 158, 62, 76, 95, 244, 23, 193, 26, 223, 86, 243, 242, 242, 120, 234, 169, 167, 24, 61, 122, 180, 77, 79, 25, 108, 204, 174, 93, 187, 104, 223, 190, 189, 214, 205, 95, 169, 116, 96, 204, 155, 255, 163, 251, 248, 39, 245, 234, 95, 58, 113, 144, 175, 38, 245, 100, 254, 228, 222, 236, 92, 252, 9, 87, 163, 143, 233, 221, 252, 1, 146, 227, 174, 113, 246, 208, 110, 118, 45, 254, 148, 239, 167, 13, 230, 181, 110, 213, 89, 250, 226, 4, 98, 34, 118, 80, 80, 80, 160, 87, 223, 197, 195, 139, 22, 189, 134, 168, 183, 87, 172, 88, 97, 145, 133, 152, 170, 87, 175, 110, 146, 243, 4, 4, 104, 175, 243, 243, 211, 79, 63, 105, 110, 102, 1, 235, 77, 114, 33, 33, 108, 136, 36, 0, 162, 56, 117, 116, 11, 170, 213, 172, 107, 141, 56, 76, 198, 205, 171, 26, 79, 124, 247, 155, 86, 159, 0, 128, 181, 107, 215, 18, 22, 22, 166, 251, 173, 207, 102, 229, 230, 230, 242, 214, 91, 111, 209, 175, 95, 63, 174, 95, 47, 26, 158, 238, 236, 234, 206, 148, 175, 215, 211, 101, 244, 84, 189, 99, 34, 183, 255, 202, 215, 15, 247, 226, 194, 209, 189, 101, 191, 94, 118, 22, 199, 126, 95, 205, 183, 83, 251, 243, 225, 253, 173, 136, 220, 182, 78, 47, 17, 104, 63, 184, 104, 57, 230, 139, 23, 47, 242, 235, 175, 191, 150, 249, 58, 101, 21, 18, 98, 154, 89, 41, 155, 53, 107, 166, 126, 159, 150, 150, 198, 119, 223, 125, 167, 185, 123, 43, 96, 154, 166, 6, 33, 108, 136, 36, 0, 162, 56, 53, 53, 55, 20, 10, 133, 77, 174, 2, 88, 86, 142, 46, 174, 60, 252, 217, 26, 238, 125, 240, 105, 173, 242, 51, 103, 206, 16, 22, 22, 198, 172, 89, 179, 172, 62, 165, 109, 113, 246, 239, 223, 79, 199, 142, 29, 153, 51, 103, 142, 214, 183, 236, 106, 53, 234, 240, 228, 130, 109, 180, 232, 49, 88, 239, 152, 107, 167, 143, 179, 244, 133, 241, 228, 102, 151, 111, 184, 156, 166, 184, 11, 209, 44, 126, 126, 12, 95, 62, 116, 15, 55, 206, 22, 245, 139, 107, 217, 107, 8, 126, 181, 27, 168, 183, 231, 205, 155, 87, 225, 107, 149, 100, 208, 160, 65, 229, 154, 30, 89, 147, 82, 169, 100, 240, 224, 162, 159, 217, 194, 133, 11, 117, 31, 9, 84, 141, 201, 35, 68, 149, 35, 9, 128, 40, 142, 214, 221, 222, 221, 199, 15, 7, 39, 103, 107, 197, 98, 82, 74, 165, 3, 35, 95, 254, 156, 201, 159, 255, 130, 155, 87, 209, 250, 245, 57, 57, 57, 124, 252, 241, 199, 52, 111, 222, 156, 213, 171, 87, 27, 108, 238, 182, 150, 155, 55, 111, 50, 117, 234, 84, 186, 119, 239, 206, 241, 227, 199, 181, 246, 53, 191, 119, 32, 47, 252, 114, 132, 250, 109, 186, 24, 60, 118, 235, 23, 175, 145, 159, 151, 107, 112, 95, 121, 197, 30, 143, 224, 179, 113, 157, 217, 187, 98, 62, 0, 74, 7, 71, 122, 78, 122, 78, 189, 255, 208, 161, 67, 172, 93, 187, 214, 164, 215, 212, 229, 239, 239, 207, 132, 9, 19, 74, 174, 88, 140, 177, 99, 199, 82, 163, 134, 106, 54, 195, 148, 148, 20, 222, 127, 255, 125, 205, 221, 135, 129, 202, 57, 102, 84, 84, 121, 146, 0, 136, 226, 212, 208, 220, 240, 244, 175, 97, 172, 158, 221, 106, 221, 111, 36, 47, 252, 114, 152, 122, 161, 218, 115, 239, 95, 185, 114, 133, 113, 227, 198, 209, 166, 77, 27, 86, 174, 92, 105, 145, 231, 217, 198, 92, 190, 124, 153, 103, 158, 121, 134, 134, 13, 27, 178, 96, 193, 2, 173, 164, 196, 193, 209, 137, 33, 207, 205, 101, 234, 55, 155, 241, 240, 53, 252, 60, 60, 53, 33, 142, 255, 246, 254, 110, 150, 216, 114, 178, 50, 88, 55, 247, 25, 214, 205, 125, 134, 130, 130, 2, 58, 143, 154, 162, 53, 53, 240, 172, 89, 179, 200, 204, 204, 52, 203, 181, 11, 189, 243, 206, 59, 229, 94, 227, 193, 203, 203, 139, 185, 115, 231, 170, 183, 223, 125, 247, 93, 173, 71, 42, 128, 249, 123, 50, 10, 97, 37, 246, 49, 165, 155, 176, 6, 15, 96, 142, 102, 65, 80, 147, 150, 116, 26, 254, 176, 117, 162, 49, 35, 55, 111, 95, 194, 70, 61, 130, 167, 111, 0, 23, 143, 237, 39, 55, 187, 232, 134, 21, 23, 23, 199, 218, 181, 107, 89, 185, 114, 37, 57, 57, 57, 52, 104, 208, 64, 119, 141, 120, 179, 200, 207, 207, 103, 247, 238, 221, 204, 153, 51, 135, 199, 31, 127, 156, 253, 251, 247, 235, 78, 74, 67, 227, 176, 94, 76, 157, 191, 145, 208, 190, 35, 180, 102, 59, 4, 72, 184, 22, 203, 225, 141, 63, 177, 245, 139, 87, 249, 117, 238, 179, 20, 152, 121, 153, 222, 75, 39, 14, 145, 28, 127, 141, 214, 253, 71, 225, 81, 205, 159, 147, 59, 85, 157, 230, 19, 19, 19, 201, 205, 205, 165, 95, 191, 126, 102, 187, 118, 181, 106, 213, 104, 221, 186, 53, 171, 86, 173, 42, 83, 139, 141, 82, 169, 100, 245, 234, 213, 116, 233, 162, 106, 53, 249, 247, 223, 127, 153, 58, 117, 170, 102, 178, 183, 19, 120, 217, 228, 1, 11, 97, 35, 20, 37, 87, 17, 85, 84, 35, 64, 107, 112, 116, 187, 65, 227, 8, 255, 104, 133, 145, 234, 149, 67, 202, 237, 155, 108, 250, 248, 255, 56, 178, 121, 185, 193, 155, 137, 131, 131, 3, 247, 221, 119, 31, 227, 199, 143, 167, 127, 255, 254, 4, 5, 153, 174, 79, 68, 110, 110, 46, 135, 15, 31, 102, 211, 166, 77, 44, 91, 182, 204, 224, 236, 116, 0, 222, 1, 65, 12, 123, 97, 30, 237, 135, 76, 208, 187, 241, 199, 70, 30, 96, 251, 247, 115, 57, 181, 103, 171, 217, 111, 250, 134, 140, 157, 243, 29, 157, 71, 77, 225, 243, 9, 93, 185, 28, 117, 24, 80, 221, 104, 183, 111, 223, 78, 239, 222, 37, 207, 71, 80, 17, 191, 253, 246, 27, 19, 38, 76, 40, 213, 104, 14, 111, 111, 111, 86, 172, 88, 193, 144, 33, 170, 145, 11, 105, 105, 105, 116, 236, 216, 145, 232, 232, 232, 194, 42, 185, 64, 123, 192, 248, 66, 1, 66, 216, 57, 73, 0, 132, 49, 93, 129, 125, 154, 5, 61, 30, 122, 134, 17, 179, 63, 179, 82, 56, 150, 117, 61, 230, 4, 219, 191, 127, 159, 227, 127, 252, 66, 126, 190, 241, 230, 255, 230, 205, 155, 211, 187, 119, 111, 58, 119, 238, 76, 179, 102, 205, 104, 218, 180, 41, 213, 170, 85, 51, 90, 191, 80, 94, 94, 30, 177, 177, 177, 156, 57, 115, 134, 147, 39, 79, 178, 123, 247, 110, 118, 239, 222, 93, 236, 114, 186, 213, 106, 212, 161, 247, 148, 23, 233, 50, 122, 170, 222, 180, 190, 25, 41, 137, 172, 255, 224, 121, 14, 111, 252, 201, 96, 226, 162, 84, 42, 200, 207, 55, 127, 127, 6, 39, 23, 55, 94, 222, 242, 31, 25, 41, 137, 124, 54, 190, 11, 185, 89, 170, 214, 148, 90, 181, 106, 113, 224, 192, 1, 234, 212, 209, 27, 88, 98, 82, 87, 174, 92, 225, 205, 55, 223, 100, 217, 178, 101, 6, 87, 6, 116, 118, 118, 230, 193, 7, 31, 228, 237, 183, 223, 86, 199, 82, 80, 80, 64, 120, 120, 56, 203, 151, 47, 215, 172, 250, 14, 96, 218, 117, 147, 133, 176, 49, 146, 0, 8, 99, 134, 161, 51, 249, 201, 224, 103, 222, 161, 223, 227, 175, 88, 41, 28, 235, 136, 143, 141, 97, 231, 194, 143, 248, 247, 247, 213, 100, 165, 149, 110, 173, 251, 192, 192, 64, 130, 130, 130, 240, 244, 244, 196, 211, 211, 19, 111, 111, 111, 114, 114, 114, 72, 77, 77, 37, 33, 33, 129, 212, 212, 84, 46, 94, 188, 88, 234, 5, 108, 130, 154, 132, 114, 239, 131, 79, 209, 105, 248, 36, 131, 157, 48, 175, 156, 58, 202, 226, 231, 199, 112, 231, 234, 69, 173, 242, 182, 45, 26, 208, 175, 91, 40, 125, 186, 133, 242, 243, 230, 127, 88, 186, 110, 119, 169, 174, 87, 81, 133, 137, 226, 158, 159, 190, 96, 253, 135, 51, 213, 229, 45, 91, 182, 228, 159, 127, 254, 193, 199, 199, 199, 236, 49, 36, 39, 39, 179, 109, 219, 54, 78, 159, 62, 205, 157, 59, 119, 240, 243, 243, 35, 36, 36, 132, 254, 253, 251, 235, 245, 23, 152, 51, 103, 14, 111, 189, 245, 150, 102, 209, 33, 160, 59, 144, 99, 246, 64, 133, 176, 34, 73, 0, 132, 49, 225, 192, 82, 205, 130, 7, 94, 251, 218, 224, 228, 50, 85, 65, 110, 86, 38, 81, 187, 54, 115, 104, 227, 82, 162, 247, 254, 97, 242, 30, 245, 186, 60, 170, 249, 211, 186, 255, 40, 58, 14, 123, 136, 134, 237, 239, 49, 90, 239, 220, 225, 61, 44, 120, 106, 56, 153, 169, 201, 234, 178, 206, 109, 26, 243, 220, 148, 161, 180, 106, 90, 52, 103, 67, 220, 237, 36, 6, 79, 121, 159, 172, 108, 243, 223, 211, 220, 188, 170, 241, 238, 190, 91, 0, 44, 121, 126, 12, 145, 219, 139, 230, 3, 232, 217, 179, 39, 155, 54, 109, 178, 72, 63, 138, 210, 248, 226, 139, 47, 120, 254, 249, 231, 53, 91, 77, 226, 129, 206, 192, 5, 235, 69, 37, 132, 101, 72, 2, 32, 140, 121, 26, 248, 82, 179, 224, 161, 15, 151, 209, 126, 72, 197, 134, 92, 85, 6, 105, 9, 183, 56, 123, 112, 23, 49, 7, 118, 16, 115, 112, 39, 241, 23, 207, 148, 124, 80, 9, 156, 93, 221, 105, 208, 174, 27, 77, 58, 247, 161, 73, 231, 222, 212, 105, 217, 161, 196, 101, 151, 175, 199, 156, 228, 171, 240, 30, 100, 166, 170, 158, 121, 187, 185, 58, 243, 242, 147, 35, 25, 121, 95, 152, 193, 250, 107, 127, 63, 192, 156, 47, 44, 179, 246, 193, 235, 219, 46, 224, 27, 84, 143, 156, 172, 12, 190, 121, 164, 47, 177, 145, 69, 51, 21, 118, 233, 210, 133, 173, 91, 183, 226, 235, 235, 107, 145, 88, 140, 121, 247, 221, 119, 121, 253, 245, 215, 53, 139, 50, 128, 62, 64, 132, 117, 34, 18, 194, 178, 100, 45, 0, 97, 140, 94, 59, 173, 171, 151, 249, 155, 110, 237, 129, 135, 111, 117, 218, 12, 24, 173, 158, 103, 63, 245, 78, 28, 55, 206, 157, 34, 254, 194, 25, 226, 98, 207, 16, 127, 241, 12, 89, 105, 41, 100, 165, 167, 146, 145, 156, 72, 86, 90, 10, 142, 46, 174, 184, 184, 123, 226, 236, 230, 129, 155, 183, 15, 94, 254, 53, 169, 209, 168, 25, 1, 245, 155, 18, 208, 176, 41, 53, 26, 54, 43, 211, 28, 11, 105, 9, 183, 248, 113, 250, 48, 245, 205, 223, 199, 203, 157, 111, 223, 121, 140, 208, 16, 227, 83, 53, 63, 48, 176, 51, 81, 49, 151, 89, 179, 117, 127, 197, 126, 0, 165, 144, 145, 146, 136, 111, 80, 61, 156, 92, 220, 120, 244, 235, 245, 204, 159, 220, 135, 155, 231, 255, 3, 32, 34, 34, 130, 206, 157, 59, 179, 118, 237, 90, 66, 67, 67, 205, 30, 139, 174, 244, 244, 116, 166, 77, 155, 166, 59, 221, 111, 14, 240, 32, 114, 243, 23, 85, 136, 36, 0, 194, 24, 189, 187, 189, 155, 36, 0, 6, 121, 250, 5, 210, 216, 47, 144, 198, 157, 122, 89, 236, 154, 27, 63, 158, 69, 194, 117, 213, 40, 1, 23, 103, 39, 230, 191, 245, 104, 177, 55, 255, 66, 175, 63, 245, 0, 142, 142, 14, 172, 220, 88, 246, 233, 128, 203, 66, 115, 69, 65, 79, 191, 64, 102, 44, 222, 193, 119, 143, 15, 82, 47, 196, 20, 19, 19, 67, 215, 174, 93, 249, 226, 139, 47, 152, 50, 101, 138, 222, 104, 6, 115, 137, 140, 140, 100, 210, 164, 73, 186, 19, 41, 101, 2, 163, 129, 45, 134, 143, 18, 162, 114, 146, 137, 128, 132, 49, 122, 93, 217, 93, 61, 37, 1, 176, 5, 215, 206, 68, 114, 120, 99, 209, 183, 215, 57, 207, 142, 161, 77, 243, 6, 165, 58, 86, 161, 80, 240, 202, 147, 35, 153, 61, 109, 4, 46, 206, 230, 203, 255, 23, 76, 31, 206, 245, 152, 162, 37, 129, 61, 253, 2, 153, 190, 112, 59, 141, 195, 122, 169, 203, 210, 210, 210, 152, 58, 117, 42, 125, 250, 244, 209, 28, 126, 103, 22, 105, 105, 105, 188, 242, 202, 43, 116, 236, 216, 81, 247, 230, 127, 11, 24, 132, 220, 252, 69, 21, 36, 9, 128, 48, 70, 111, 106, 53, 87, 207, 242, 205, 182, 38, 76, 235, 232, 150, 149, 234, 78, 107, 93, 219, 55, 101, 104, 159, 14, 101, 62, 199, 131, 195, 239, 101, 229, 23, 207, 105, 117, 20, 52, 165, 212, 132, 120, 254, 55, 245, 62, 18, 111, 94, 81, 151, 185, 121, 251, 50, 237, 135, 63, 232, 61, 249, 5, 173, 111, 252, 187, 118, 237, 34, 52, 52, 148, 201, 147, 39, 115, 238, 220, 57, 67, 167, 43, 183, 140, 140, 12, 62, 251, 236, 51, 130, 131, 131, 121, 255, 253, 247, 201, 201, 209, 234, 4, 121, 24, 232, 8, 236, 50, 233, 69, 133, 176, 19, 146, 0, 8, 99, 244, 186, 105, 75, 2, 96, 27, 226, 46, 20, 125, 91, 190, 191, 111, 199, 114, 159, 167, 73, 131, 32, 86, 124, 254, 44, 31, 191, 28, 78, 253, 218, 1, 37, 31, 80, 2, 7, 7, 37, 83, 199, 245, 195, 197, 217, 9, 80, 77, 170, 244, 211, 139, 19, 181, 38, 36, 82, 58, 56, 50, 236, 197, 121, 60, 250, 245, 6, 173, 149, 37, 115, 115, 115, 89, 178, 100, 9, 33, 33, 33, 12, 26, 52, 136, 53, 107, 214, 144, 145, 145, 161, 119, 141, 210, 40, 40, 40, 224, 240, 225, 195, 60, 245, 212, 83, 212, 174, 93, 155, 153, 51, 103, 114, 243, 230, 77, 205, 42, 57, 192, 135, 192, 189, 64, 108, 185, 46, 34, 68, 37, 32, 163, 0, 132, 49, 59, 0, 173, 169, 219, 230, 29, 77, 199, 209, 217, 197, 74, 225, 136, 66, 75, 95, 24, 207, 177, 63, 214, 0, 16, 62, 178, 7, 255, 247, 248, 240, 18, 142, 40, 89, 126, 65, 1, 251, 143, 158, 97, 237, 239, 17, 236, 140, 136, 34, 55, 183, 108, 107, 31, 180, 106, 90, 151, 89, 143, 221, 79, 251, 86, 141, 216, 117, 32, 138, 103, 222, 90, 164, 110, 165, 24, 253, 250, 124, 186, 141, 155, 166, 119, 76, 86, 90, 10, 91, 191, 124, 157, 125, 171, 254, 71, 94, 174, 254, 240, 68, 87, 87, 87, 186, 119, 239, 78, 239, 222, 189, 105, 217, 178, 37, 45, 91, 182, 164, 118, 237, 218, 184, 187, 187, 171, 235, 228, 230, 230, 18, 31, 31, 79, 116, 116, 52, 209, 209, 209, 236, 219, 183, 143, 237, 219, 183, 115, 227, 198, 13, 99, 161, 238, 3, 166, 3, 199, 141, 85, 16, 162, 170, 144, 4, 64, 24, 179, 15, 213, 108, 128, 128, 234, 217, 241, 199, 145, 57, 22, 235, 172, 37, 140, 59, 246, 251, 106, 150, 190, 88, 52, 28, 115, 112, 239, 246, 60, 61, 105, 32, 117, 106, 250, 155, 228, 252, 105, 25, 89, 28, 138, 60, 203, 129, 99, 49, 68, 157, 185, 66, 236, 181, 120, 238, 36, 106, 47, 143, 236, 226, 236, 72, 237, 26, 254, 116, 105, 215, 132, 62, 93, 91, 17, 214, 166, 177, 214, 239, 198, 188, 239, 54, 240, 211, 250, 61, 0, 120, 249, 215, 224, 181, 63, 207, 233, 205, 94, 88, 232, 206, 213, 139, 108, 251, 238, 61, 14, 111, 252, 201, 96, 34, 160, 203, 193, 193, 1, 111, 111, 111, 178, 178, 178, 72, 79, 79, 47, 237, 199, 74, 5, 198, 1, 91, 75, 123, 128, 16, 149, 157, 252, 111, 46, 140, 57, 10, 180, 43, 220, 112, 114, 113, 227, 195, 35, 169, 197, 84, 23, 150, 146, 151, 155, 195, 247, 79, 12, 34, 230, 64, 209, 42, 181, 74, 165, 146, 14, 173, 26, 210, 51, 172, 5, 189, 186, 180, 162, 126, 109, 195, 43, 3, 150, 87, 70, 102, 54, 105, 25, 89, 100, 102, 101, 227, 226, 236, 68, 128, 95, 241, 143, 131, 210, 50, 178, 24, 244, 200, 92, 18, 146, 84, 191, 51, 227, 222, 254, 129, 206, 163, 166, 20, 123, 76, 82, 220, 85, 14, 111, 248, 137, 131, 235, 23, 19, 31, 27, 99, 178, 216, 239, 122, 23, 120, 189, 196, 90, 66, 84, 33, 146, 0, 8, 99, 254, 3, 154, 21, 110, 184, 121, 251, 242, 222, 221, 217, 221, 132, 245, 101, 166, 38, 177, 242, 213, 71, 56, 241, 215, 6, 131, 251, 107, 6, 84, 163, 85, 211, 186, 180, 106, 90, 143, 86, 77, 235, 210, 188, 113, 29, 188, 61, 13, 127, 3, 55, 151, 111, 150, 253, 193, 183, 203, 255, 4, 160, 105, 215, 126, 76, 251, 225, 143, 82, 29, 87, 80, 80, 192, 245, 51, 145, 156, 222, 191, 157, 152, 136, 191, 184, 18, 117, 148, 212, 132, 248, 18, 143, 115, 116, 113, 165, 102, 112, 75, 178, 82, 147, 136, 191, 116, 86, 119, 119, 111, 164, 179, 159, 16, 90, 36, 1, 16, 198, 92, 0, 26, 20, 110, 120, 7, 4, 49, 103, 231, 21, 227, 181, 133, 85, 68, 237, 220, 196, 95, 63, 126, 200, 197, 227, 37, 79, 238, 83, 171, 134, 31, 33, 141, 106, 209, 172, 81, 45, 154, 5, 215, 166, 89, 163, 90, 212, 170, 225, 103, 182, 216, 206, 95, 190, 201, 240, 199, 231, 1, 170, 22, 164, 247, 15, 38, 162, 116, 40, 223, 208, 195, 180, 196, 219, 196, 95, 60, 77, 234, 157, 91, 100, 165, 167, 144, 157, 145, 142, 210, 209, 17, 87, 119, 47, 220, 125, 252, 240, 171, 211, 0, 191, 90, 13, 0, 120, 163, 103, 16, 105, 9, 90, 201, 106, 38, 224, 123, 247, 79, 33, 196, 93, 50, 17, 144, 48, 198, 85, 115, 195, 209, 217, 213, 88, 61, 97, 69, 45, 123, 15, 163, 101, 239, 97, 220, 190, 114, 158, 168, 93, 155, 137, 218, 181, 137, 243, 135, 255, 54, 248, 44, 253, 218, 205, 59, 92, 187, 121, 135, 157, 251, 139, 198, 231, 123, 123, 186, 209, 44, 184, 54, 33, 141, 106, 209, 181, 93, 83, 194, 218, 52, 49, 217, 252, 0, 13, 235, 4, 226, 229, 225, 70, 74, 90, 6, 57, 89, 25, 220, 186, 124, 142, 192, 6, 33, 229, 58, 151, 71, 53, 127, 60, 218, 118, 43, 177, 222, 149, 83, 71, 117, 111, 254, 0, 123, 209, 190, 249, 43, 129, 166, 168, 18, 92, 127, 32, 27, 136, 67, 181, 8, 80, 169, 59, 21, 8, 97, 239, 36, 1, 16, 198, 104, 205, 75, 235, 88, 134, 105, 106, 133, 229, 249, 215, 105, 68, 143, 135, 158, 161, 199, 67, 207, 144, 157, 145, 198, 149, 83, 71, 185, 28, 117, 132, 11, 71, 255, 33, 114, 251, 58, 163, 199, 37, 167, 102, 112, 240, 248, 89, 14, 30, 63, 203, 79, 191, 238, 193, 221, 205, 133, 123, 58, 132, 208, 183, 123, 107, 250, 117, 15, 197, 217, 169, 252, 255, 69, 40, 20, 10, 106, 84, 247, 33, 37, 77, 53, 156, 47, 35, 41, 161, 220, 231, 42, 173, 152, 3, 59, 12, 21, 255, 5, 212, 6, 70, 0, 247, 163, 234, 220, 106, 104, 53, 162, 44, 224, 27, 84, 125, 5, 210, 204, 20, 162, 16, 54, 67, 18, 0, 97, 140, 214, 227, 33, 133, 82, 166, 140, 176, 23, 206, 110, 30, 52, 234, 112, 47, 141, 58, 220, 75, 235, 254, 163, 212, 9, 128, 167, 167, 39, 31, 124, 240, 1, 199, 142, 29, 227, 223, 127, 255, 37, 42, 42, 138, 204, 76, 237, 86, 241, 244, 140, 44, 254, 220, 27, 201, 159, 123, 35, 241, 247, 245, 98, 204, 160, 174, 140, 29, 210, 181, 196, 78, 127, 198, 228, 231, 23, 148, 92, 201, 132, 98, 34, 254, 50, 84, 60, 16, 85, 39, 192, 226, 87, 87, 2, 23, 224, 121, 84, 45, 3, 15, 0, 150, 13, 94, 8, 11, 147, 4, 64, 24, 35, 253, 67, 42, 129, 236, 244, 162, 145, 27, 62, 62, 62, 204, 152, 49, 67, 189, 157, 155, 155, 75, 116, 116, 52, 199, 142, 29, 99, 223, 190, 125, 108, 220, 184, 145, 171, 87, 175, 170, 247, 223, 78, 72, 225, 127, 43, 254, 100, 193, 234, 191, 24, 212, 179, 29, 143, 140, 233, 77, 227, 250, 53, 75, 125, 237, 252, 130, 2, 110, 222, 74, 84, 111, 123, 250, 7, 86, 240, 211, 20, 47, 47, 39, 155, 243, 71, 255, 49, 180, 171, 39, 128, 82, 233, 64, 131, 118, 221, 104, 214, 253, 62, 234, 183, 233, 74, 96, 195, 166, 120, 250, 5, 146, 159, 151, 75, 124, 108, 12, 187, 151, 124, 198, 161, 13, 75, 1, 70, 2, 205, 129, 83, 102, 13, 88, 8, 43, 147, 255, 228, 133, 49, 137, 104, 44, 8, 84, 35, 184, 5, 47, 109, 56, 97, 197, 112, 68, 121, 92, 58, 113, 144, 207, 39, 168, 166, 115, 8, 9, 9, 41, 118, 206, 253, 194, 25, 244, 214, 172, 89, 195, 194, 133, 11, 185, 125, 251, 182, 214, 126, 133, 66, 65, 175, 46, 45, 121, 116, 76, 239, 82, 173, 61, 16, 125, 254, 42, 99, 102, 124, 10, 168, 214, 145, 120, 111, 223, 45, 179, 182, 36, 157, 59, 188, 135, 249, 147, 123, 235, 149, 251, 6, 213, 163, 219, 216, 39, 8, 27, 245, 8, 94, 254, 53, 140, 30, 159, 145, 156, 192, 171, 221, 212, 195, 39, 71, 1, 191, 154, 37, 80, 33, 108, 132, 180, 235, 10, 81, 137, 101, 103, 20, 245, 105, 243, 244, 244, 44, 182, 174, 66, 161, 160, 83, 167, 78, 204, 155, 55, 143, 203, 151, 47, 243, 195, 15, 63, 208, 186, 117, 107, 245, 254, 130, 130, 2, 118, 238, 63, 201, 67, 51, 191, 226, 193, 231, 191, 228, 247, 61, 199, 200, 203, 203, 55, 122, 190, 117, 191, 31, 80, 191, 15, 238, 120, 175, 217, 31, 35, 25, 122, 254, 223, 229, 129, 169, 188, 250, 199, 89, 250, 62, 54, 187, 216, 155, 63, 64, 244, 63, 127, 106, 110, 234, 141, 35, 20, 162, 178, 145, 4, 64, 24, 163, 221, 7, 64, 102, 0, 180, 75, 154, 243, 240, 59, 56, 148, 244, 8, 188, 136, 155, 155, 27, 83, 167, 78, 229, 248, 241, 227, 252, 254, 251, 239, 244, 234, 213, 75, 107, 127, 100, 116, 44, 179, 222, 255, 137, 1, 147, 223, 229, 163, 31, 54, 114, 40, 242, 28, 249, 26, 215, 218, 177, 255, 36, 171, 183, 22, 13, 77, 236, 242, 192, 212, 242, 127, 136, 82, 138, 137, 208, 79, 0, 122, 78, 122, 22, 165, 178, 228, 207, 157, 154, 16, 207, 198, 143, 102, 21, 110, 30, 5, 78, 22, 83, 93, 136, 74, 65, 250, 0, 8, 99, 164, 3, 84, 37, 160, 116, 44, 250, 39, 158, 155, 155, 91, 174, 115, 12, 24, 48, 0, 127, 127, 127, 58, 119, 238, 172, 117, 147, 7, 184, 121, 43, 137, 165, 235, 118, 179, 116, 221, 110, 60, 221, 93, 9, 10, 244, 37, 45, 61, 147, 107, 113, 69, 61, 254, 235, 183, 238, 76, 243, 158, 131, 203, 247, 1, 74, 41, 59, 35, 141, 75, 39, 14, 106, 149, 121, 7, 4, 17, 216, 168, 121, 137, 199, 102, 164, 36, 242, 227, 244, 251, 73, 138, 187, 10, 144, 7, 188, 128, 252, 254, 139, 42, 64, 90, 0, 132, 49, 242, 149, 191, 18, 112, 112, 114, 82, 191, 207, 203, 43, 219, 2, 63, 133, 210, 211, 211, 9, 15, 15, 87, 223, 252, 235, 180, 104, 79, 143, 240, 103, 245, 86, 135, 76, 77, 207, 36, 230, 226, 117, 173, 155, 191, 119, 96, 45, 194, 63, 90, 94, 170, 111, 225, 21, 17, 27, 121, 80, 111, 238, 131, 198, 97, 189, 74, 108, 185, 186, 125, 249, 28, 95, 79, 234, 165, 153, 60, 188, 129, 204, 24, 40, 170, 8, 105, 1, 16, 198, 104, 253, 111, 154, 95, 206, 111, 143, 194, 186, 28, 28, 139, 18, 128, 242, 182, 0, 60, 255, 252, 243, 234, 206, 131, 142, 206, 46, 76, 120, 111, 33, 65, 77, 66, 25, 244, 244, 219, 156, 254, 231, 79, 162, 118, 109, 226, 212, 158, 173, 90, 19, 240, 56, 186, 184, 210, 186, 223, 72, 70, 204, 254, 12, 79, 223, 138, 47, 53, 92, 146, 11, 255, 234, 247, 254, 111, 212, 225, 94, 163, 245, 11, 10, 10, 56, 188, 241, 39, 214, 127, 56, 147, 140, 100, 117, 194, 242, 46, 48, 215, 44, 1, 10, 97, 131, 36, 1, 16, 198, 100, 107, 110, 228, 230, 100, 27, 171, 39, 108, 152, 102, 2, 112, 245, 234, 85, 230, 207, 159, 207, 208, 161, 67, 169, 95, 191, 126, 169, 142, 95, 183, 110, 29, 223, 127, 255, 189, 122, 123, 200, 115, 115, 9, 106, 18, 10, 128, 139, 187, 39, 173, 251, 143, 162, 117, 255, 81, 228, 231, 231, 145, 112, 245, 34, 137, 55, 175, 226, 234, 225, 69, 245, 122, 141, 113, 241, 48, 52, 215, 142, 121, 92, 52, 144, 0, 52, 108, 215, 221, 96, 221, 43, 167, 142, 176, 233, 147, 151, 52, 23, 83, 202, 4, 102, 0, 11, 205, 21, 159, 16, 182, 72, 154, 121, 133, 49, 23, 1, 245, 93, 194, 39, 176, 54, 111, 238, 184, 100, 189, 104, 68, 185, 220, 56, 27, 197, 188, 17, 173, 245, 202, 91, 183, 110, 205, 208, 161, 67, 25, 54, 108, 24, 97, 97, 97, 40, 13, 244, 208, 143, 142, 142, 166, 115, 231, 206, 36, 39, 39, 3, 16, 210, 173, 63, 143, 127, 247, 155, 205, 117, 8, 45, 200, 207, 231, 213, 110, 213, 201, 76, 77, 82, 151, 233, 14, 59, 204, 207, 207, 35, 122, 239, 31, 236, 95, 243, 61, 167, 118, 109, 166, 160, 64, 253, 136, 255, 8, 240, 40, 112, 220, 210, 113, 11, 97, 109, 210, 2, 32, 140, 201, 210, 220, 200, 205, 201, 50, 86, 79, 216, 48, 55, 239, 106, 248, 213, 110, 200, 157, 171, 23, 180, 202, 35, 35, 35, 137, 140, 140, 100, 238, 220, 185, 4, 6, 6, 50, 116, 232, 80, 134, 14, 29, 202, 125, 247, 221, 135, 135, 135, 7, 201, 201, 201, 140, 28, 57, 82, 125, 243, 247, 170, 94, 147, 9, 239, 45, 178, 185, 155, 63, 192, 245, 152, 19, 90, 55, 127, 128, 134, 237, 84, 115, 31, 156, 59, 188, 135, 200, 109, 107, 57, 190, 109, 29, 201, 113, 215, 52, 171, 196, 163, 106, 242, 255, 6, 144, 231, 91, 162, 74, 178, 189, 127, 205, 194, 86, 156, 0, 90, 21, 110, 184, 122, 122, 51, 55, 194, 252, 115, 185, 11, 243, 184, 30, 115, 146, 83, 187, 183, 16, 181, 107, 51, 177, 145, 17, 90, 195, 3, 53, 185, 186, 186, 210, 171, 87, 47, 210, 211, 211, 217, 179, 103, 15, 0, 14, 78, 206, 76, 95, 184, 221, 104, 147, 186, 181, 253, 243, 243, 183, 172, 125, 247, 41, 173, 178, 38, 157, 251, 112, 251, 202, 121, 238, 92, 189, 168, 91, 253, 60, 240, 35, 240, 53, 144, 98, 145, 0, 133, 176, 81, 146, 0, 8, 99, 14, 3, 29, 10, 55, 28, 93, 92, 153, 119, 68, 214, 71, 169, 12, 82, 19, 226, 249, 111, 247, 86, 162, 118, 111, 230, 244, 190, 109, 100, 165, 21, 127, 31, 28, 59, 231, 59, 186, 140, 54, 255, 56, 254, 242, 90, 241, 202, 100, 14, 111, 252, 169, 184, 42, 183, 128, 245, 192, 10, 84, 61, 252, 101, 136, 159, 16, 200, 35, 0, 97, 156, 86, 175, 191, 60, 233, 4, 88, 105, 120, 250, 6, 208, 105, 196, 195, 116, 26, 241, 48, 185, 217, 89, 156, 59, 188, 135, 168, 157, 27, 137, 218, 189, 133, 132, 107, 177, 90, 117, 187, 141, 155, 102, 211, 55, 127, 128, 107, 167, 35, 13, 21, 103, 3, 27, 80, 117, 236, 219, 142, 52, 243, 11, 161, 71, 90, 0, 132, 49, 219, 129, 190, 154, 5, 31, 30, 73, 197, 201, 197, 205, 74, 225, 8, 75, 184, 125, 229, 60, 81, 187, 54, 115, 106, 247, 102, 242, 114, 114, 153, 246, 195, 239, 56, 216, 240, 82, 208, 121, 57, 217, 204, 14, 243, 209, 77, 80, 83, 81, 45, 230, 115, 197, 58, 81, 9, 97, 31, 164, 5, 64, 24, 163, 215, 46, 156, 153, 154, 44, 9, 64, 37, 231, 95, 167, 17, 61, 30, 122, 134, 30, 15, 61, 99, 237, 80, 74, 229, 230, 249, 255, 12, 181, 78, 109, 69, 110, 254, 66, 148, 72, 102, 2, 20, 198, 36, 235, 22, 100, 166, 234, 21, 9, 97, 85, 70, 154, 255, 101, 72, 159, 16, 165, 32, 9, 128, 48, 70, 239, 110, 95, 82, 103, 49, 33, 44, 237, 234, 105, 131, 247, 122, 73, 0, 132, 40, 5, 73, 0, 132, 49, 250, 45, 0, 105, 210, 2, 32, 108, 203, 53, 195, 9, 192, 49, 75, 199, 33, 132, 61, 146, 4, 64, 24, 163, 223, 2, 144, 42, 45, 0, 194, 182, 196, 95, 56, 163, 91, 116, 7, 184, 106, 133, 80, 132, 176, 59, 146, 0, 8, 99, 244, 59, 1, 74, 11, 128, 176, 33, 121, 185, 57, 36, 199, 95, 215, 45, 62, 107, 141, 88, 132, 176, 71, 146, 0, 8, 99, 146, 116, 11, 50, 146, 19, 173, 17, 135, 16, 6, 37, 221, 188, 66, 126, 190, 222, 18, 199, 178, 96, 133, 16, 165, 36, 9, 128, 48, 230, 182, 110, 65, 90, 226, 45, 67, 245, 132, 176, 138, 59, 58, 147, 22, 221, 101, 176, 80, 8, 161, 79, 18, 0, 97, 76, 188, 110, 65, 234, 29, 189, 34, 33, 172, 38, 241, 250, 101, 67, 197, 210, 2, 32, 68, 41, 73, 2, 32, 140, 49, 144, 0, 196, 89, 35, 14, 33, 12, 50, 210, 2, 32, 9, 128, 16, 165, 36, 9, 128, 48, 198, 64, 2, 32, 143, 0, 132, 237, 72, 188, 33, 45, 0, 66, 84, 132, 36, 0, 194, 152, 12, 84, 115, 170, 171, 165, 38, 200, 35, 0, 97, 59, 210, 18, 12, 38, 164, 122, 195, 2, 132, 16, 134, 73, 2, 32, 138, 163, 117, 199, 79, 147, 62, 0, 194, 134, 24, 25, 150, 42, 67, 85, 132, 40, 37, 73, 0, 68, 113, 180, 238, 248, 233, 73, 119, 12, 13, 187, 18, 194, 42, 12, 76, 76, 149, 131, 170, 229, 74, 8, 81, 10, 146, 0, 136, 226, 220, 212, 220, 200, 207, 207, 35, 245, 150, 170, 232, 114, 212, 97, 62, 25, 221, 129, 23, 66, 29, 249, 246, 209, 126, 242, 120, 64, 88, 92, 118, 102, 186, 110, 81, 154, 53, 226, 16, 194, 94, 73, 2, 32, 138, 163, 183, 164, 234, 165, 168, 195, 172, 255, 112, 38, 95, 62, 116, 47, 87, 163, 143, 81, 80, 80, 64, 204, 129, 157, 108, 249, 244, 101, 107, 196, 39, 170, 50, 133, 66, 183, 36, 223, 26, 97, 8, 97, 175, 28, 173, 29, 128, 176, 105, 122, 221, 172, 23, 61, 51, 138, 130, 130, 2, 189, 138, 113, 23, 245, 230, 100, 23, 194, 210, 252, 128, 195, 192, 151, 192, 114, 64, 158, 87, 9, 81, 12, 73, 0, 132, 49, 46, 64, 77, 221, 66, 67, 55, 127, 128, 150, 189, 135, 153, 59, 30, 81, 201, 36, 197, 93, 229, 214, 165, 115, 36, 222, 184, 66, 234, 237, 155, 164, 220, 137, 35, 39, 51, 131, 220, 236, 76, 50, 82, 146, 112, 73, 217, 131, 163, 0, 0, 32, 0, 73, 68, 65, 84, 114, 117, 195, 209, 201, 25, 103, 119, 79, 28, 157, 156, 241, 14, 8, 194, 191, 78, 35, 252, 106, 55, 192, 191, 110, 35, 148, 10, 131, 13, 152, 29, 128, 37, 192, 203, 192, 59, 192, 74, 192, 240, 47, 173, 16, 85, 156, 94, 27, 154, 168, 242, 90, 0, 83, 129, 73, 128, 127, 73, 149, 21, 74, 37, 253, 167, 189, 198, 128, 105, 175, 163, 80, 202, 19, 37, 81, 178, 171, 209, 199, 152, 63, 185, 15, 153, 169, 73, 184, 184, 123, 82, 173, 102, 93, 188, 170, 215, 192, 195, 183, 58, 46, 238, 158, 56, 56, 58, 225, 230, 237, 75, 110, 86, 38, 57, 89, 25, 228, 100, 101, 146, 147, 153, 78, 114, 252, 13, 110, 95, 57, 175, 94, 0, 200, 193, 209, 137, 188, 220, 156, 146, 46, 119, 24, 120, 17, 216, 109, 230, 143, 37, 132, 221, 145, 4, 64, 20, 10, 67, 245, 141, 233, 190, 210, 30, 208, 160, 77, 87, 238, 159, 53, 143, 6, 109, 187, 153, 47, 42, 81, 233, 228, 102, 101, 114, 225, 223, 125, 212, 104, 212, 12, 239, 192, 90, 229, 58, 254, 246, 149, 243, 156, 220, 177, 129, 131, 235, 151, 16, 31, 27, 83, 154, 195, 126, 5, 102, 3, 242, 172, 74, 136, 187, 36, 1, 16, 141, 129, 143, 128, 225, 148, 226, 247, 193, 217, 205, 131, 182, 3, 199, 208, 229, 129, 71, 229, 198, 47, 108, 194, 237, 43, 231, 217, 177, 224, 35, 14, 109, 88, 66, 110, 118, 86, 113, 85, 115, 128, 79, 129, 215, 239, 190, 23, 162, 74, 147, 4, 160, 234, 114, 4, 158, 7, 222, 2, 220, 138, 171, 232, 224, 228, 76, 211, 174, 125, 105, 123, 223, 24, 66, 251, 141, 196, 213, 211, 219, 34, 1, 10, 81, 22, 105, 137, 183, 217, 191, 250, 59, 118, 45, 249, 148, 244, 164, 132, 226, 170, 30, 6, 166, 0, 39, 44, 19, 153, 16, 182, 73, 18, 128, 170, 169, 30, 176, 6, 85, 179, 191, 81, 13, 218, 118, 163, 243, 3, 83, 104, 221, 119, 4, 110, 222, 190, 150, 137, 76, 136, 10, 202, 72, 73, 100, 251, 119, 115, 249, 123, 197, 215, 197, 181, 8, 100, 3, 111, 3, 239, 35, 195, 7, 69, 21, 37, 9, 64, 213, 211, 7, 248, 25, 8, 48, 180, 211, 193, 209, 137, 78, 35, 30, 166, 103, 248, 179, 212, 8, 110, 97, 217, 200, 132, 48, 161, 196, 27, 151, 217, 250, 197, 107, 28, 217, 188, 220, 232, 232, 21, 96, 39, 48, 1, 157, 73, 175, 132, 168, 10, 36, 1, 168, 90, 198, 0, 43, 48, 50, 252, 179, 221, 160, 113, 12, 122, 250, 109, 170, 215, 107, 108, 217, 168, 132, 48, 163, 243, 71, 254, 102, 213, 27, 143, 21, 215, 89, 240, 10, 170, 127, 27, 17, 150, 139, 74, 8, 235, 115, 176, 118, 0, 194, 98, 70, 163, 26, 19, 173, 119, 243, 247, 171, 221, 128, 240, 143, 87, 210, 247, 209, 255, 195, 221, 199, 207, 242, 145, 9, 97, 70, 190, 181, 234, 19, 54, 114, 50, 89, 169, 201, 92, 142, 58, 108, 168, 138, 55, 16, 14, 220, 0, 142, 90, 52, 56, 33, 172, 72, 18, 128, 170, 161, 29, 176, 5, 112, 214, 221, 209, 118, 192, 24, 30, 255, 238, 55, 106, 6, 55, 55, 201, 133, 82, 110, 223, 164, 32, 63, 31, 71, 103, 23, 147, 156, 79, 8, 83, 112, 116, 114, 166, 121, 143, 193, 4, 119, 234, 65, 76, 196, 14, 67, 43, 9, 58, 2, 247, 163, 106, 21, 221, 101, 233, 248, 132, 176, 6, 121, 4, 80, 249, 121, 163, 250, 86, 19, 172, 89, 168, 80, 40, 24, 242, 220, 92, 122, 79, 153, 133, 66, 127, 78, 245, 82, 41, 200, 207, 231, 210, 201, 67, 68, 237, 220, 68, 108, 100, 4, 215, 78, 71, 146, 150, 120, 27, 7, 71, 39, 6, 62, 245, 22, 125, 167, 190, 100, 130, 240, 133, 48, 173, 180, 132, 91, 44, 123, 233, 33, 78, 239, 219, 102, 172, 202, 98, 224, 49, 32, 215, 98, 65, 9, 97, 5, 146, 0, 84, 126, 175, 163, 234, 237, 172, 166, 80, 40, 24, 249, 202, 151, 220, 51, 97, 122, 185, 78, 120, 243, 252, 127, 68, 172, 93, 192, 209, 45, 43, 73, 185, 117, 195, 96, 29, 133, 66, 193, 203, 91, 162, 165, 63, 129, 176, 73, 249, 249, 121, 252, 254, 245, 28, 254, 250, 225, 125, 99, 29, 4, 55, 0, 227, 128, 98, 39, 22, 16, 194, 158, 73, 2, 80, 185, 121, 0, 23, 129, 234, 154, 133, 125, 31, 155, 205, 144, 103, 223, 43, 211, 137, 10, 242, 243, 57, 185, 115, 35, 187, 151, 126, 206, 249, 35, 127, 151, 234, 152, 233, 139, 254, 162, 113, 167, 94, 101, 186, 142, 16, 150, 20, 185, 109, 29, 203, 102, 135, 147, 155, 149, 105, 104, 247, 6, 84, 157, 3, 101, 210, 32, 81, 41, 73, 2, 80, 185, 245, 66, 53, 204, 73, 173, 118, 179, 182, 60, 191, 234, 0, 74, 135, 210, 173, 3, 149, 149, 150, 194, 225, 77, 203, 216, 243, 211, 23, 165, 157, 114, 21, 128, 234, 245, 26, 243, 127, 235, 35, 165, 47, 128, 176, 121, 177, 199, 35, 248, 113, 198, 253, 164, 37, 222, 54, 180, 123, 45, 48, 30, 121, 28, 32, 42, 33, 89, 13, 176, 114, 107, 169, 91, 208, 227, 161, 103, 74, 188, 249, 23, 20, 20, 112, 249, 228, 33, 14, 111, 252, 137, 195, 155, 150, 145, 153, 170, 215, 97, 202, 32, 79, 191, 64, 106, 133, 132, 82, 175, 85, 24, 247, 62, 248, 148, 220, 252, 133, 93, 168, 223, 166, 11, 211, 23, 237, 224, 135, 105, 67, 72, 188, 121, 69, 119, 247, 3, 192, 2, 224, 17, 100, 194, 32, 81, 201, 72, 2, 80, 185, 233, 205, 217, 27, 27, 121, 128, 102, 247, 14, 196, 203, 191, 134, 186, 172, 32, 63, 159, 219, 87, 206, 115, 233, 196, 33, 46, 30, 219, 199, 201, 157, 155, 72, 188, 113, 185, 84, 23, 168, 23, 26, 70, 235, 254, 163, 104, 213, 123, 24, 129, 13, 155, 153, 46, 114, 33, 44, 40, 168, 73, 43, 158, 94, 246, 55, 95, 63, 220, 139, 132, 107, 177, 186, 187, 39, 1, 87, 129, 87, 44, 31, 153, 16, 230, 35, 143, 0, 42, 183, 193, 168, 134, 255, 233, 241, 9, 172, 141, 163, 179, 11, 121, 57, 217, 36, 223, 186, 65, 126, 94, 217, 90, 56, 171, 213, 172, 195, 212, 111, 54, 81, 171, 105, 107, 83, 196, 41, 132, 77, 136, 143, 141, 97, 254, 228, 222, 234, 37, 135, 117, 60, 14, 252, 96, 225, 144, 132, 48, 27, 73, 0, 42, 55, 39, 32, 26, 104, 100, 234, 19, 123, 249, 215, 224, 173, 221, 215, 76, 125, 90, 33, 172, 46, 238, 226, 105, 230, 63, 220, 155, 148, 219, 122, 179, 3, 231, 0, 3, 208, 233, 87, 35, 132, 189, 82, 90, 59, 0, 97, 86, 57, 192, 195, 64, 122, 5, 207, 179, 11, 248, 79, 179, 32, 229, 246, 77, 82, 239, 196, 85, 240, 180, 66, 216, 158, 192, 6, 33, 60, 250, 245, 122, 156, 92, 244, 22, 201, 116, 2, 86, 3, 117, 45, 31, 149, 16, 166, 39, 9, 64, 229, 183, 23, 232, 9, 252, 83, 134, 99, 10, 80, 45, 153, 250, 10, 16, 2, 244, 198, 192, 183, 158, 235, 49, 81, 166, 136, 79, 8, 155, 83, 47, 52, 140, 9, 239, 45, 50, 52, 73, 86, 117, 138, 89, 79, 67, 8, 123, 34, 83, 1, 87, 13, 215, 128, 133, 192, 126, 160, 240, 107, 123, 14, 144, 1, 164, 2, 167, 81, 221, 240, 215, 2, 31, 3, 51, 129, 207, 80, 37, 15, 133, 99, 163, 234, 1, 67, 53, 79, 26, 212, 180, 21, 13, 219, 118, 51, 119, 236, 66, 88, 69, 205, 198, 45, 81, 40, 149, 156, 61, 184, 75, 119, 87, 61, 84, 255, 119, 238, 176, 120, 80, 194, 24, 23, 84, 157, 158, 221, 0, 47, 160, 38, 224, 11, 248, 221, 125, 121, 160, 250, 59, 203, 67, 70, 115, 168, 73, 22, 91, 181, 252, 121, 247, 85, 30, 71, 116, 11, 46, 159, 56, 84, 177, 104, 132, 176, 113, 253, 159, 120, 149, 107, 167, 143, 19, 185, 109, 157, 238, 174, 151, 81, 181, 138, 253, 101, 249, 168, 170, 156, 58, 168, 134, 52, 215, 190, 251, 170, 113, 183, 44, 80, 227, 207, 178, 140, 57, 190, 131, 106, 249, 231, 56, 224, 2, 112, 238, 238, 235, 212, 221, 87, 149, 153, 248, 73, 58, 1, 138, 210, 114, 6, 146, 209, 248, 135, 230, 87, 187, 33, 175, 253, 113, 214, 122, 17, 9, 97, 1, 25, 41, 137, 124, 50, 186, 35, 119, 174, 94, 208, 221, 117, 21, 213, 141, 41, 201, 242, 81, 85, 74, 254, 64, 107, 160, 5, 16, 138, 234, 103, 219, 10, 168, 102, 193, 24, 50, 129, 19, 168, 150, 134, 222, 117, 247, 117, 199, 130, 215, 255, 127, 246, 206, 59, 60, 170, 106, 235, 195, 239, 153, 158, 73, 239, 61, 164, 1, 33, 244, 80, 66, 23, 80, 20, 43, 22, 176, 247, 222, 175, 223, 181, 92, 187, 168, 92, 203, 85, 185, 118, 197, 114, 237, 13, 69, 5, 5, 84, 68, 122, 239, 132, 26, 90, 122, 79, 38, 189, 76, 57, 231, 251, 227, 144, 152, 6, 73, 38, 147, 204, 36, 57, 239, 243, 240, 8, 123, 246, 62, 103, 153, 204, 204, 94, 103, 237, 181, 126, 171, 91, 81, 28, 0, 133, 142, 176, 13, 24, 221, 120, 224, 185, 117, 121, 120, 248, 6, 58, 201, 28, 5, 133, 238, 33, 99, 223, 54, 222, 186, 110, 10, 54, 139, 185, 249, 75, 239, 0, 247, 58, 193, 164, 222, 64, 12, 48, 165, 209, 159, 211, 55, 14, 17, 64, 239, 173, 71, 239, 163, 199, 224, 239, 134, 222, 251, 228, 127, 125, 244, 232, 188, 245, 212, 167, 107, 104, 221, 255, 110, 122, 170, 113, 215, 32, 154, 69, 68, 139, 136, 165, 202, 140, 173, 206, 70, 93, 105, 29, 117, 165, 181, 212, 149, 214, 81, 157, 87, 77, 85, 94, 21, 230, 242, 83, 182, 124, 16, 249, 251, 120, 116, 17, 114, 164, 160, 215, 160, 56, 0, 10, 29, 225, 93, 224, 174, 198, 3, 183, 191, 191, 140, 132, 73, 231, 56, 201, 28, 5, 133, 238, 99, 229, 71, 47, 179, 244, 245, 22, 90, 64, 34, 114, 146, 237, 250, 238, 183, 168, 199, 17, 15, 156, 9, 76, 70, 254, 153, 69, 180, 54, 73, 80, 171, 240, 136, 240, 192, 51, 210, 19, 207, 126, 94, 120, 70, 122, 225, 25, 229, 133, 49, 200, 13, 65, 221, 53, 121, 235, 150, 42, 11, 229, 105, 101, 148, 29, 43, 163, 236, 88, 41, 37, 7, 138, 168, 41, 170, 105, 109, 234, 102, 96, 1, 114, 53, 72, 103, 171, 171, 156, 142, 226, 0, 40, 116, 132, 155, 145, 101, 81, 27, 56, 247, 190, 231, 152, 113, 199, 19, 78, 50, 71, 65, 161, 251, 16, 69, 27, 111, 92, 53, 129, 204, 253, 219, 155, 191, 116, 8, 24, 14, 180, 8, 15, 40, 48, 10, 184, 248, 228, 159, 33, 45, 94, 21, 192, 51, 194, 19, 191, 68, 127, 252, 18, 253, 241, 138, 246, 193, 35, 220, 189, 203, 54, 250, 142, 80, 149, 91, 69, 113, 74, 33, 185, 155, 115, 41, 218, 91, 136, 100, 107, 146, 59, 88, 10, 188, 141, 156, 44, 221, 99, 143, 8, 20, 7, 64, 161, 35, 12, 5, 246, 54, 30, 24, 60, 237, 66, 110, 121, 235, 103, 39, 153, 163, 160, 208, 189, 228, 30, 73, 97, 254, 156, 49, 216, 172, 45, 242, 196, 234, 43, 103, 250, 58, 106, 228, 112, 126, 253, 166, 31, 213, 248, 69, 65, 37, 224, 21, 237, 133, 95, 98, 0, 254, 131, 229, 77, 95, 231, 229, 250, 61, 67, 44, 149, 22, 114, 55, 231, 144, 241, 123, 26, 165, 71, 77, 141, 95, 170, 64, 254, 189, 191, 76, 15, 140, 8, 40, 14, 128, 66, 71, 80, 35, 39, 60, 185, 215, 15, 120, 6, 132, 240, 236, 234, 108, 231, 89, 164, 160, 208, 205, 44, 127, 235, 105, 86, 44, 104, 209, 78, 219, 4, 244, 231, 239, 178, 217, 190, 198, 16, 224, 22, 224, 90, 154, 181, 31, 215, 24, 52, 4, 38, 5, 17, 146, 28, 70, 208, 168, 96, 180, 238, 90, 167, 24, 232, 40, 202, 142, 151, 114, 236, 167, 163, 228, 110, 200, 70, 146, 164, 250, 225, 52, 224, 31, 192, 18, 167, 25, 102, 7, 138, 3, 160, 208, 81, 214, 34, 159, 225, 53, 240, 248, 178, 195, 4, 68, 157, 62, 127, 71, 65, 161, 183, 96, 173, 171, 229, 165, 139, 6, 83, 146, 157, 214, 252, 165, 183, 129, 251, 186, 223, 34, 167, 225, 5, 92, 129, 188, 241, 39, 55, 126, 65, 231, 165, 39, 100, 108, 8, 193, 201, 161, 4, 14, 15, 68, 165, 237, 125, 146, 51, 149, 217, 21, 28, 254, 250, 16, 185, 27, 155, 60, 0, 125, 138, 252, 30, 168, 116, 138, 81, 29, 164, 247, 253, 86, 20, 186, 154, 120, 154, 57, 0, 97, 9, 35, 136, 24, 52, 210, 73, 230, 40, 40, 116, 47, 42, 141, 6, 175, 192, 80, 246, 252, 177, 168, 249, 75, 73, 192, 55, 244, 224, 51, 225, 118, 50, 17, 120, 14, 89, 92, 236, 50, 78, 38, 243, 105, 220, 52, 68, 76, 139, 36, 241, 230, 161, 12, 185, 117, 24, 33, 201, 161, 120, 132, 121, 184, 196, 121, 126, 87, 160, 243, 210, 19, 54, 49, 28, 191, 65, 254, 148, 30, 49, 97, 174, 48, 3, 140, 64, 254, 153, 172, 2, 10, 157, 106, 96, 59, 80, 28, 0, 133, 142, 162, 7, 174, 105, 60, 224, 238, 227, 199, 144, 233, 179, 156, 100, 142, 130, 66, 247, 19, 18, 63, 152, 35, 91, 86, 97, 202, 205, 104, 60, 172, 66, 126, 42, 94, 236, 28, 171, 186, 20, 21, 178, 18, 232, 103, 192, 211, 200, 27, 157, 14, 192, 59, 206, 135, 129, 87, 38, 48, 252, 254, 36, 66, 199, 135, 99, 12, 50, 182, 38, 161, 220, 107, 49, 134, 184, 19, 53, 163, 31, 214, 106, 27, 165, 71, 76, 32, 31, 129, 92, 141, 172, 188, 218, 162, 183, 180, 43, 209, 119, 126, 75, 10, 142, 194, 11, 249, 9, 167, 193, 121, 244, 143, 140, 227, 137, 229, 169, 206, 179, 72, 65, 193, 9, 164, 239, 221, 194, 27, 87, 183, 144, 194, 182, 2, 131, 128, 222, 162, 144, 165, 69, 118, 248, 255, 5, 36, 212, 15, 234, 189, 245, 132, 79, 141, 36, 234, 172, 126, 120, 68, 120, 58, 205, 56, 87, 35, 103, 125, 54, 187, 223, 220, 137, 104, 177, 129, 44, 42, 116, 9, 240, 155, 115, 173, 58, 53, 138, 3, 160, 96, 15, 59, 144, 195, 157, 13, 60, 189, 50, 29, 159, 224, 86, 203, 122, 21, 20, 122, 45, 31, 221, 125, 33, 7, 214, 46, 107, 62, 252, 41, 112, 83, 247, 91, 227, 80, 220, 129, 91, 129, 7, 105, 212, 253, 208, 35, 194, 147, 248, 75, 251, 19, 54, 57, 2, 149, 166, 119, 134, 246, 59, 75, 201, 193, 98, 182, 189, 176, 25, 75, 165, 5, 160, 10, 56, 11, 89, 63, 192, 229, 80, 126, 131, 10, 246, 176, 182, 249, 192, 137, 157, 29, 105, 54, 168, 160, 208, 59, 152, 121, 239, 220, 214, 194, 221, 215, 210, 172, 252, 173, 7, 161, 69, 206, 102, 79, 3, 94, 231, 228, 230, 239, 211, 223, 151, 209, 255, 74, 230, 140, 55, 167, 19, 49, 45, 74, 217, 252, 79, 131, 223, 32, 127, 198, 61, 59, 9, 141, 81, 3, 178, 35, 245, 43, 16, 235, 92, 171, 90, 71, 201, 1, 80, 176, 7, 119, 228, 236, 223, 6, 60, 253, 131, 72, 156, 114, 158, 147, 204, 81, 80, 112, 14, 94, 129, 97, 100, 236, 219, 74, 81, 122, 147, 136, 191, 10, 185, 161, 204, 159, 206, 177, 202, 110, 102, 1, 63, 33, 135, 252, 141, 0, 1, 195, 3, 25, 118, 207, 72, 18, 174, 77, 196, 35, 194, 179, 79, 157, 237, 119, 6, 131, 175, 1, 223, 129, 126, 228, 172, 207, 66, 18, 37, 35, 48, 14, 57, 127, 194, 165, 58, 17, 42, 14, 128, 130, 61, 20, 1, 15, 209, 232, 8, 201, 92, 93, 197, 164, 171, 238, 118, 158, 69, 10, 10, 78, 194, 47, 52, 138, 173, 63, 125, 218, 124, 56, 17, 185, 44, 176, 39, 116, 150, 27, 9, 124, 141, 220, 225, 208, 31, 192, 127, 104, 0, 163, 30, 30, 67, 252, 165, 3, 48, 6, 187, 159, 118, 177, 66, 235, 24, 131, 140, 104, 61, 116, 20, 236, 200, 7, 185, 82, 66, 133, 139, 181, 144, 86, 226, 56, 10, 246, 80, 136, 220, 54, 179, 129, 252, 227, 7, 41, 205, 203, 116, 146, 57, 10, 10, 206, 35, 38, 105, 18, 145, 131, 71, 55, 31, 246, 5, 110, 236, 126, 107, 58, 68, 24, 114, 41, 223, 118, 96, 26, 128, 71, 184, 7, 99, 30, 27, 199, 248, 231, 38, 225, 19, 239, 235, 84, 227, 122, 3, 209, 231, 198, 16, 146, 28, 90, 255, 207, 135, 129, 1, 78, 52, 167, 5, 74, 4, 64, 193, 94, 226, 129, 241, 141, 7, 66, 226, 7, 43, 122, 0, 10, 125, 18, 189, 209, 131, 189, 127, 254, 216, 124, 56, 10, 120, 207, 9, 230, 180, 133, 0, 220, 142, 92, 174, 56, 14, 16, 180, 30, 58, 6, 94, 145, 192, 136, 251, 71, 225, 25, 165, 100, 245, 59, 18, 255, 33, 129, 100, 172, 72, 67, 180, 136, 106, 32, 26, 89, 43, 194, 37, 80, 34, 0, 10, 246, 242, 71, 243, 129, 67, 27, 126, 119, 134, 29, 10, 10, 78, 103, 248, 57, 179, 241, 14, 10, 111, 62, 60, 4, 24, 227, 4, 115, 78, 199, 32, 228, 36, 222, 5, 128, 151, 74, 163, 34, 246, 162, 120, 166, 191, 55, 131, 184, 75, 251, 163, 210, 42, 91, 130, 163, 209, 251, 232, 137, 191, 180, 225, 193, 255, 66, 160, 69, 237, 168, 179, 80, 126, 219, 10, 246, 178, 6, 104, 210, 47, 51, 117, 211, 74, 68, 209, 230, 36, 115, 20, 20, 156, 135, 90, 163, 101, 204, 197, 55, 180, 246, 146, 171, 148, 3, 10, 200, 217, 253, 59, 129, 73, 0, 126, 9, 254, 76, 249, 239, 52, 18, 111, 26, 130, 214, 163, 103, 235, 243, 187, 58, 177, 23, 197, 161, 247, 49, 212, 255, 179, 69, 79, 105, 103, 161, 56, 0, 10, 246, 82, 67, 179, 30, 232, 53, 229, 38, 50, 83, 182, 57, 201, 28, 5, 5, 231, 146, 124, 201, 141, 173, 101, 201, 95, 5, 184, 57, 193, 156, 198, 132, 1, 191, 35, 151, 245, 25, 52, 110, 26, 134, 220, 62, 140, 241, 47, 76, 82, 68, 124, 186, 9, 149, 86, 77, 236, 172, 134, 126, 41, 231, 33, 43, 41, 58, 29, 197, 1, 80, 232, 12, 45, 142, 1, 14, 111, 92, 225, 12, 59, 20, 20, 156, 142, 127, 100, 28, 177, 163, 166, 52, 31, 246, 1, 206, 117, 130, 57, 245, 204, 4, 246, 0, 51, 64, 174, 81, 159, 242, 250, 116, 162, 207, 141, 85, 74, 250, 186, 153, 126, 51, 163, 209, 122, 232, 64, 142, 198, 184, 68, 20, 64, 113, 0, 20, 58, 67, 139, 67, 255, 67, 27, 90, 248, 4, 10, 10, 125, 134, 177, 173, 31, 3, 92, 222, 221, 118, 32, 39, 120, 63, 15, 44, 5, 2, 4, 181, 138, 132, 107, 18, 25, 63, 111, 18, 198, 32, 163, 19, 204, 81, 208, 24, 52, 196, 92, 208, 160, 7, 116, 25, 141, 164, 149, 157, 133, 226, 0, 40, 116, 134, 125, 64, 78, 227, 129, 244, 189, 155, 169, 40, 202, 115, 146, 57, 10, 10, 206, 101, 200, 153, 179, 208, 232, 244, 205, 135, 47, 68, 22, 207, 234, 46, 2, 129, 101, 192, 147, 128, 202, 224, 239, 198, 132, 121, 147, 136, 159, 61, 0, 65, 165, 60, 245, 59, 147, 152, 243, 99, 209, 184, 105, 64, 222, 123, 255, 229, 100, 115, 148, 50, 64, 133, 78, 51, 24, 89, 72, 68, 70, 146, 8, 142, 29, 164, 148, 3, 42, 244, 73, 180, 122, 3, 233, 123, 183, 80, 152, 126, 164, 201, 48, 114, 255, 140, 131, 221, 96, 194, 88, 96, 53, 39, 207, 152, 131, 146, 130, 73, 158, 59, 1, 143, 112, 143, 110, 184, 117, 207, 192, 90, 109, 165, 50, 171, 2, 211, 225, 18, 242, 183, 231, 147, 191, 45, 151, 162, 189, 133, 20, 237, 45, 36, 119, 99, 14, 133, 187, 11, 40, 61, 106, 162, 34, 179, 130, 218, 226, 26, 204, 21, 102, 52, 110, 90, 212, 186, 206, 111, 151, 106, 157, 26, 75, 149, 5, 211, 161, 18, 144, 191, 59, 191, 0, 74, 59, 125, 97, 59, 209, 56, 235, 198, 10, 189, 134, 159, 129, 155, 27, 15, 164, 172, 252, 153, 228, 75, 111, 62, 197, 116, 5, 133, 222, 205, 136, 115, 230, 112, 96, 205, 210, 230, 195, 51, 129, 22, 66, 1, 14, 230, 90, 224, 67, 192, 32, 168, 4, 6, 92, 153, 32, 63, 245, 247, 225, 179, 126, 75, 165, 133, 146, 131, 197, 148, 30, 49, 81, 122, 196, 68, 217, 241, 82, 204, 229, 102, 187, 174, 229, 22, 104, 196, 43, 198, 11, 239, 24, 31, 2, 134, 6, 226, 155, 224, 135, 160, 238, 248, 207, 54, 246, 162, 56, 210, 150, 30, 199, 102, 182, 105, 145, 197, 129, 238, 177, 203, 32, 7, 208, 119, 223, 25, 10, 142, 66, 143, 172, 12, 216, 144, 78, 172, 209, 233, 121, 126, 93, 62, 122, 119, 37, 195, 88, 161, 239, 81, 85, 90, 204, 211, 83, 66, 144, 196, 38, 178, 239, 233, 200, 34, 48, 93, 129, 26, 120, 17, 121, 51, 65, 235, 161, 37, 233, 159, 99, 8, 28, 25, 212, 69, 183, 115, 97, 36, 48, 165, 150, 144, 191, 61, 143, 162, 61, 133, 148, 29, 43, 69, 18, 165, 83, 78, 55, 120, 122, 225, 225, 31, 0, 128, 209, 71, 86, 62, 20, 173, 86, 42, 138, 10, 169, 42, 41, 194, 106, 62, 181, 179, 160, 49, 106, 8, 24, 22, 72, 80, 82, 8, 161, 19, 194, 208, 186, 183, 191, 148, 114, 223, 135, 123, 73, 91, 118, 28, 228, 150, 193, 177, 64, 110, 187, 23, 59, 16, 197, 1, 80, 112, 4, 223, 3, 179, 27, 15, 220, 240, 218, 119, 12, 63, 103, 246, 41, 166, 43, 40, 244, 110, 254, 123, 69, 50, 153, 251, 183, 55, 31, 30, 0, 28, 105, 101, 122, 103, 240, 0, 190, 67, 46, 45, 195, 35, 220, 147, 49, 143, 39, 227, 30, 214, 135, 66, 254, 18, 148, 30, 49, 145, 179, 33, 155, 220, 77, 217, 212, 20, 214, 180, 152, 226, 225, 31, 64, 244, 168, 177, 244, 75, 26, 67, 112, 252, 64, 2, 162, 99, 9, 140, 137, 107, 216, 252, 79, 69, 117, 169, 9, 83, 118, 38, 217, 251, 246, 146, 181, 127, 47, 89, 41, 187, 73, 219, 185, 141, 154, 178, 166, 81, 123, 149, 86, 77, 200, 216, 16, 194, 167, 70, 18, 52, 50, 184, 205, 200, 64, 77, 97, 13, 127, 221, 181, 2, 201, 38, 2, 188, 202, 73, 231, 173, 187, 81, 28, 0, 5, 71, 112, 13, 240, 101, 227, 129, 164, 243, 175, 226, 218, 151, 191, 60, 197, 116, 5, 133, 222, 205, 178, 55, 158, 228, 207, 15, 95, 108, 62, 124, 47, 240, 142, 3, 111, 19, 130, 220, 106, 118, 20, 64, 208, 168, 96, 146, 254, 57, 26, 141, 177, 111, 136, 250, 148, 167, 149, 145, 189, 38, 147, 156, 13, 57, 212, 20, 86, 55, 121, 205, 205, 219, 135, 132, 41, 211, 73, 152, 122, 22, 3, 167, 76, 39, 184, 255, 64, 135, 221, 87, 180, 217, 72, 219, 190, 133, 253, 43, 127, 99, 255, 159, 191, 145, 177, 107, 59, 146, 244, 119, 148, 193, 24, 100, 36, 230, 194, 56, 34, 207, 234, 135, 198, 112, 234, 83, 246, 61, 111, 239, 36, 115, 101, 6, 64, 37, 114, 116, 168, 216, 97, 70, 182, 19, 197, 1, 80, 112, 4, 222, 64, 1, 160, 171, 31, 48, 120, 120, 243, 252, 186, 60, 212, 90, 221, 169, 87, 41, 40, 244, 82, 142, 110, 91, 205, 187, 55, 157, 217, 124, 120, 49, 112, 177, 131, 110, 49, 16, 88, 14, 196, 0, 196, 92, 16, 71, 226, 77, 67, 122, 125, 150, 191, 100, 19, 201, 221, 148, 75, 218, 178, 227, 148, 28, 108, 186, 95, 186, 121, 251, 48, 226, 252, 139, 73, 186, 120, 14, 9, 83, 207, 66, 163, 235, 158, 239, 158, 194, 19, 199, 216, 242, 237, 231, 108, 249, 246, 11, 138, 210, 79, 52, 140, 107, 61, 180, 196, 156, 23, 71, 236, 172, 120, 52, 198, 150, 142, 64, 85, 78, 37, 171, 239, 91, 89, 127, 68, 241, 28, 240, 76, 183, 24, 220, 136, 222, 253, 110, 113, 45, 130, 144, 117, 193, 163, 128, 80, 26, 109, 150, 64, 5, 144, 1, 156, 56, 249, 223, 140, 110, 183, 174, 243, 252, 14, 156, 221, 120, 224, 246, 5, 203, 73, 152, 120, 246, 41, 166, 43, 40, 244, 94, 108, 22, 51, 79, 78, 12, 164, 174, 186, 178, 241, 112, 57, 16, 64, 231, 91, 4, 79, 0, 150, 0, 254, 130, 32, 48, 232, 134, 193, 141, 85, 230, 122, 37, 117, 166, 90, 210, 255, 72, 35, 253, 247, 52, 234, 76, 181, 13, 227, 110, 94, 222, 12, 63, 111, 22, 73, 23, 207, 97, 208, 244, 179, 187, 109, 211, 111, 13, 73, 146, 56, 180, 106, 5, 43, 222, 122, 149, 131, 171, 254, 22, 68, 211, 121, 233, 232, 63, 123, 32, 253, 102, 198, 180, 232, 181, 176, 243, 181, 109, 228, 172, 207, 6, 48, 1, 253, 144, 247, 130, 110, 67, 113, 0, 186, 150, 112, 224, 14, 224, 82, 228, 254, 224, 237, 253, 121, 151, 0, 219, 128, 173, 39, 255, 172, 3, 202, 186, 194, 64, 7, 114, 23, 240, 110, 227, 129, 49, 179, 174, 231, 170, 127, 127, 226, 36, 115, 20, 20, 156, 203, 71, 119, 95, 200, 129, 181, 203, 154, 15, 79, 166, 153, 132, 118, 7, 185, 20, 249, 184, 205, 77, 165, 85, 49, 226, 254, 81, 132, 77, 106, 209, 132, 168, 215, 80, 153, 93, 201, 145, 239, 15, 147, 179, 62, 187, 254, 188, 28, 128, 240, 193, 195, 152, 122, 219, 61, 140, 185, 252, 26, 244, 198, 238, 148, 88, 104, 31, 89, 251, 246, 240, 251, 252, 23, 217, 241, 211, 194, 134, 227, 1, 99, 136, 59, 67, 110, 27, 70, 80, 82, 112, 195, 188, 242, 180, 50, 214, 254, 115, 21, 200, 83, 254, 5, 252, 167, 59, 237, 84, 28, 128, 174, 193, 11, 152, 135, 188, 41, 54, 137, 253, 8, 106, 1, 173, 167, 14, 173, 135, 22, 91, 157, 13, 73, 148, 16, 235, 108, 88, 42, 79, 251, 80, 96, 65, 254, 210, 88, 142, 172, 236, 117, 160, 139, 236, 238, 12, 193, 64, 22, 141, 254, 127, 245, 238, 158, 60, 183, 54, 23, 173, 222, 217, 82, 232, 10, 10, 221, 207, 186, 175, 222, 226, 167, 23, 31, 104, 62, 252, 36, 240, 111, 59, 47, 121, 31, 178, 158, 191, 74, 235, 174, 101, 244, 99, 201, 248, 15, 62, 125, 18, 91, 79, 165, 42, 167, 146, 212, 133, 135, 201, 89, 151, 213, 144, 197, 175, 214, 106, 25, 126, 254, 197, 76, 189, 253, 94, 250, 79, 104, 33, 185, 236, 146, 100, 236, 217, 201, 207, 115, 31, 109, 18, 17, 8, 29, 31, 198, 224, 91, 135, 97, 240, 147, 155, 3, 109, 123, 97, 51, 249, 219, 242, 0, 242, 145, 143, 116, 90, 102, 49, 118, 17, 138, 3, 224, 120, 18, 145, 55, 233, 104, 144, 55, 124, 191, 161, 254, 248, 12, 242, 195, 43, 206, 11, 189, 175, 161, 213, 159, 186, 104, 17, 169, 43, 169, 165, 182, 184, 150, 170, 204, 74, 42, 51, 42, 168, 204, 168, 196, 90, 213, 170, 99, 112, 28, 248, 26, 249, 73, 224, 112, 87, 253, 143, 216, 193, 111, 192, 57, 141, 7, 174, 127, 245, 27, 70, 204, 116, 134, 18, 170, 130, 130, 115, 201, 73, 221, 203, 171, 151, 182, 16, 196, 250, 9, 249, 41, 190, 35, 8, 192, 203, 156, 204, 20, 119, 11, 112, 99, 236, 83, 19, 240, 140, 234, 125, 101, 182, 85, 185, 242, 19, 127, 246, 154, 191, 55, 126, 131, 135, 39, 83, 111, 191, 151, 51, 110, 187, 7, 159, 208, 158, 25, 237, 216, 191, 98, 57, 223, 61, 114, 31, 133, 39, 142, 1, 160, 245, 208, 49, 244, 142, 97, 132, 77, 138, 192, 148, 90, 194, 134, 127, 173, 173, 159, 122, 31, 240, 118, 119, 217, 165, 56, 0, 142, 101, 16, 176, 1, 240, 69, 128, 160, 228, 16, 34, 207, 237, 135, 214, 211, 254, 172, 220, 154, 130, 26, 74, 15, 150, 80, 122, 208, 68, 249, 177, 50, 36, 91, 139, 154, 214, 173, 200, 142, 192, 183, 200, 245, 248, 206, 228, 122, 224, 179, 198, 3, 131, 167, 93, 200, 45, 111, 253, 236, 36, 115, 20, 20, 156, 135, 40, 218, 120, 124, 172, 15, 230, 218, 38, 25, 234, 29, 213, 3, 208, 2, 159, 2, 87, 3, 120, 69, 123, 49, 246, 201, 9, 24, 252, 13, 167, 93, 212, 211, 168, 51, 213, 114, 232, 171, 131, 100, 173, 206, 104, 248, 142, 211, 187, 123, 48, 237, 142, 251, 56, 235, 222, 7, 113, 247, 243, 119, 178, 133, 157, 199, 82, 91, 195, 111, 243, 95, 228, 143, 215, 95, 110, 208, 23, 8, 159, 28, 193, 208, 187, 71, 176, 253, 133, 45, 20, 165, 20, 130, 156, 255, 21, 79, 231, 243, 68, 218, 133, 226, 0, 56, 14, 55, 96, 23, 48, 80, 165, 85, 17, 127, 237, 64, 252, 134, 58, 246, 77, 107, 171, 179, 81, 150, 90, 74, 241, 238, 34, 76, 251, 138, 17, 45, 77, 132, 70, 44, 192, 34, 224, 45, 96, 163, 67, 111, 220, 126, 60, 129, 60, 160, 161, 219, 136, 90, 163, 101, 238, 234, 108, 220, 125, 122, 254, 7, 88, 65, 161, 163, 188, 113, 245, 4, 210, 247, 110, 105, 62, 28, 68, 251, 156, 117, 61, 240, 13, 112, 9, 64, 192, 208, 64, 70, 63, 58, 182, 87, 149, 249, 137, 86, 145, 244, 223, 78, 112, 248, 155, 131, 88, 171, 173, 0, 232, 141, 238, 76, 184, 254, 22, 102, 254, 223, 99, 120, 5, 135, 56, 217, 66, 199, 147, 125, 32, 133, 79, 239, 184, 158, 172, 148, 221, 0, 120, 132, 123, 16, 125, 110, 44, 251, 62, 218, 91, 63, 229, 102, 160, 91, 146, 167, 148, 94, 0, 142, 227, 121, 78, 150, 248, 196, 95, 59, 16, 255, 225, 142, 63, 155, 83, 105, 84, 184, 5, 27, 241, 31, 30, 64, 200, 228, 48, 220, 2, 221, 176, 213, 218, 168, 51, 213, 129, 252, 187, 28, 2, 220, 2, 92, 4, 152, 129, 67, 128, 213, 225, 134, 156, 26, 51, 48, 28, 89, 227, 26, 0, 73, 20, 9, 136, 140, 35, 114, 240, 168, 110, 52, 67, 65, 193, 53, 200, 62, 184, 139, 204, 125, 45, 4, 129, 254, 4, 142, 181, 177, 212, 3, 248, 133, 147, 2, 63, 161, 19, 194, 25, 253, 232, 88, 212, 250, 222, 163, 222, 158, 187, 41, 135, 237, 47, 110, 38, 103, 93, 54, 162, 69, 68, 173, 213, 50, 237, 142, 251, 185, 227, 139, 69, 140, 188, 232, 50, 244, 30, 189, 83, 204, 200, 43, 48, 152, 241, 87, 223, 72, 77, 121, 25, 105, 59, 182, 98, 174, 48, 83, 114, 72, 46, 105, 60, 121, 236, 49, 8, 120, 143, 250, 212, 192, 46, 68, 113, 0, 28, 131, 1, 185, 44, 71, 173, 243, 209, 19, 119, 69, 255, 46, 191, 161, 74, 163, 194, 61, 220, 131, 192, 49, 193, 4, 37, 7, 163, 210, 171, 169, 45, 168, 65, 52, 139, 32, 151, 25, 94, 140, 92, 129, 224, 139, 236, 8, 116, 87, 121, 137, 25, 184, 170, 241, 64, 93, 85, 57, 99, 47, 190, 177, 155, 110, 175, 160, 224, 58, 84, 20, 231, 179, 127, 213, 146, 230, 195, 7, 57, 125, 37, 128, 15, 114, 62, 205, 20, 128, 200, 51, 163, 24, 241, 143, 81, 168, 212, 189, 163, 121, 107, 121, 90, 25, 59, 95, 221, 206, 177, 159, 143, 96, 57, 153, 227, 52, 116, 230, 5, 220, 249, 213, 207, 140, 189, 252, 26, 116, 46, 152, 213, 239, 104, 212, 26, 13, 67, 102, 156, 71, 96, 116, 28, 41, 191, 255, 106, 181, 213, 89, 85, 141, 36, 139, 253, 145, 19, 189, 247, 119, 181, 29, 202, 17, 128, 99, 184, 5, 248, 8, 32, 120, 66, 40, 49, 179, 227, 156, 98, 132, 104, 21, 41, 217, 83, 68, 222, 186, 92, 42, 51, 154, 236, 247, 22, 224, 7, 228, 12, 226, 173, 93, 108, 134, 14, 89, 215, 218, 175, 126, 64, 16, 4, 30, 95, 118, 24, 255, 72, 231, 252, 92, 20, 20, 156, 69, 246, 161, 221, 188, 54, 187, 69, 244, 235, 7, 96, 206, 41, 150, 4, 33, 107, 106, 140, 0, 136, 62, 63, 150, 33, 183, 12, 235, 21, 223, 212, 54, 179, 141, 35, 223, 29, 230, 216, 226, 35, 13, 231, 252, 161, 9, 131, 153, 243, 194, 124, 6, 77, 239, 187, 122, 33, 89, 41, 187, 121, 243, 146, 179, 43, 42, 138, 10, 27, 103, 117, 238, 69, 126, 15, 116, 105, 20, 64, 137, 0, 56, 134, 207, 145, 63, 184, 68, 95, 18, 135, 222, 167, 69, 63, 240, 110, 65, 80, 9, 24, 195, 220, 9, 26, 23, 130, 79, 130, 47, 162, 85, 164, 166, 160, 6, 36, 212, 192, 80, 224, 54, 228, 44, 253, 10, 228, 234, 1, 241, 116, 215, 179, 19, 27, 114, 115, 139, 209, 141, 7, 117, 110, 70, 6, 140, 107, 161, 140, 166, 160, 208, 171, 113, 247, 241, 103, 213, 39, 175, 34, 218, 154, 156, 196, 233, 129, 55, 91, 153, 30, 14, 172, 66, 62, 202, 35, 126, 246, 0, 18, 111, 28, 210, 43, 54, 255, 162, 148, 66, 182, 205, 219, 36, 151, 187, 73, 224, 238, 235, 199, 101, 207, 191, 194, 117, 111, 125, 68, 80, 92, 215, 71, 76, 93, 25, 175, 224, 16, 38, 92, 119, 139, 126, 239, 111, 191, 20, 86, 22, 23, 213, 135, 63, 130, 145, 91, 72, 167, 118, 229, 189, 123, 193, 91, 203, 233, 196, 32, 159, 231, 9, 0, 195, 255, 149, 132, 91, 176, 241, 244, 43, 186, 17, 75, 185, 153, 188, 245, 185, 20, 108, 206, 107, 174, 53, 144, 129, 172, 75, 254, 33, 178, 10, 149, 35, 25, 11, 52, 201, 124, 242, 12, 8, 225, 233, 63, 211, 80, 107, 122, 79, 2, 147, 130, 66, 123, 120, 109, 206, 104, 178, 15, 238, 106, 60, 100, 69, 78, 148, 109, 252, 129, 140, 69, 206, 13, 136, 1, 24, 116, 221, 96, 226, 46, 237, 249, 27, 163, 165, 210, 204, 129, 79, 247, 145, 249, 87, 70, 195, 179, 236, 232, 75, 175, 96, 206, 75, 111, 224, 21, 20, 124, 250, 197, 125, 12, 115, 77, 53, 111, 94, 114, 78, 225, 177, 205, 235, 3, 79, 14, 109, 6, 198, 119, 229, 61, 123, 199, 161, 146, 115, 185, 147, 70, 142, 148, 181, 198, 230, 68, 83, 90, 162, 245, 210, 17, 121, 94, 63, 70, 62, 61, 134, 216, 203, 227, 49, 134, 52, 56, 39, 81, 200, 181, 197, 153, 200, 10, 126, 131, 28, 120, 219, 173, 200, 21, 17, 13, 84, 20, 229, 177, 127, 245, 175, 14, 188, 133, 130, 66, 207, 32, 32, 50, 182, 249, 144, 6, 89, 246, 181, 158, 68, 100, 181, 207, 24, 65, 16, 24, 114, 251, 240, 94, 177, 249, 231, 109, 206, 97, 245, 125, 43, 229, 134, 55, 18, 248, 134, 71, 114, 247, 183, 191, 112, 203, 255, 190, 85, 54, 255, 86, 208, 185, 25, 249, 231, 175, 171, 2, 135, 95, 112, 113, 253, 249, 237, 40, 228, 50, 208, 46, 67, 57, 2, 232, 28, 106, 228, 26, 252, 134, 172, 21, 175, 88, 111, 220, 195, 93, 47, 137, 69, 80, 9, 184, 71, 120, 16, 60, 33, 20, 207, 88, 47, 172, 213, 86, 106, 139, 106, 64, 62, 179, 31, 3, 220, 13, 76, 3, 234, 144, 91, 150, 118, 214, 147, 209, 112, 50, 131, 185, 158, 218, 202, 50, 70, 93, 112, 77, 39, 47, 171, 160, 208, 179, 200, 58, 176, 147, 19, 187, 54, 52, 31, 94, 134, 28, 57, 28, 14, 252, 5, 132, 8, 106, 129, 225, 247, 38, 17, 53, 163, 95, 243, 185, 61, 10, 107, 181, 149, 148, 247, 119, 115, 232, 203, 3, 216, 106, 109, 8, 42, 21, 83, 111, 187, 135, 59, 190, 252, 145, 240, 196, 161, 206, 54, 207, 165, 17, 84, 42, 146, 102, 205, 209, 187, 251, 249, 75, 35, 47, 186, 108, 243, 93, 223, 44, 89, 180, 244, 165, 103, 107, 219, 94, 105, 231, 253, 186, 234, 194, 125, 132, 51, 128, 213, 32, 43, 254, 73, 54, 201, 169, 73, 128, 29, 165, 182, 184, 150, 130, 77, 121, 20, 108, 202, 195, 90, 211, 228, 140, 178, 20, 88, 136, 124, 78, 105, 111, 38, 170, 55, 144, 77, 35, 231, 72, 80, 169, 120, 98, 249, 17, 252, 194, 163, 237, 188, 164, 130, 66, 207, 99, 243, 15, 31, 177, 112, 238, 29, 205, 135, 239, 65, 14, 241, 254, 1, 248, 171, 52, 42, 146, 30, 28, 77, 200, 184, 176, 110, 183, 207, 145, 148, 236, 47, 102, 247, 155, 59, 168, 46, 144, 197, 143, 130, 226, 7, 112, 195, 187, 159, 18, 59, 182, 75, 35, 217, 189, 153, 181, 254, 222, 204, 184, 92, 16, 204, 93, 113, 113, 229, 8, 160, 115, 204, 2, 208, 121, 234, 8, 73, 150, 63, 184, 101, 71, 74, 157, 106, 80, 71, 48, 248, 27, 136, 186, 32, 154, 145, 79, 141, 38, 250, 210, 56, 140, 97, 13, 123, 181, 15, 112, 59, 144, 2, 172, 4, 46, 167, 105, 247, 194, 246, 80, 134, 236, 68, 52, 32, 137, 34, 91, 127, 82, 154, 3, 41, 244, 45, 78, 81, 253, 50, 158, 198, 155, 255, 195, 99, 122, 244, 230, 47, 217, 68, 82, 191, 61, 196, 166, 167, 215, 55, 108, 254, 201, 87, 94, 199, 227, 107, 118, 40, 155, 127, 231, 152, 82, 84, 214, 180, 201, 154, 35, 81, 34, 0, 157, 99, 31, 48, 56, 98, 90, 20, 126, 131, 252, 216, 251, 238, 110, 16, 96, 212, 179, 201, 104, 61, 122, 102, 178, 91, 69, 90, 5, 5, 155, 114, 41, 222, 93, 212, 92, 105, 48, 31, 89, 146, 244, 99, 228, 35, 130, 246, 48, 14, 216, 212, 120, 192, 43, 40, 140, 167, 254, 56, 174, 36, 3, 42, 244, 25, 76, 185, 25, 60, 63, 35, 166, 249, 176, 5, 208, 170, 117, 106, 70, 63, 58, 150, 192, 145, 61, 247, 76, 188, 50, 187, 130, 157, 175, 110, 167, 60, 77, 110, 88, 234, 29, 18, 202, 117, 111, 255, 143, 193, 103, 205, 116, 178, 101, 189, 8, 129, 127, 220, 229, 45, 180, 86, 57, 210, 41, 148, 8, 128, 253, 184, 3, 9, 0, 126, 131, 252, 228, 15, 176, 0, 72, 80, 146, 82, 236, 92, 203, 58, 129, 103, 180, 39, 113, 87, 13, 32, 105, 238, 88, 162, 47, 137, 197, 237, 239, 164, 193, 96, 228, 118, 149, 169, 200, 253, 14, 238, 68, 22, 25, 58, 29, 155, 145, 163, 8, 13, 148, 23, 228, 176, 231, 247, 31, 28, 108, 181, 130, 130, 235, 226, 19, 28, 129, 70, 215, 162, 52, 88, 171, 214, 171, 25, 243, 248, 184, 30, 189, 249, 103, 173, 206, 100, 253, 67, 107, 26, 54, 255, 164, 89, 179, 121, 106, 99, 138, 178, 249, 59, 26, 137, 249, 239, 149, 75, 23, 56, 250, 178, 138, 3, 96, 63, 195, 57, 153, 68, 233, 29, 235, 131, 91, 128, 27, 190, 3, 100, 237, 155, 226, 221, 206, 238, 201, 211, 121, 52, 110, 26, 66, 38, 135, 49, 252, 145, 36, 6, 223, 55, 140, 128, 81, 65, 168, 52, 13, 111, 151, 9, 200, 82, 149, 185, 200, 97, 254, 11, 104, 214, 246, 184, 17, 11, 154, 15, 172, 249, 252, 245, 174, 48, 89, 65, 193, 37, 17, 84, 42, 188, 2, 67, 155, 140, 169, 13, 106, 198, 62, 57, 158, 128, 225, 129, 167, 88, 229, 218, 216, 234, 108, 236, 121, 123, 39, 187, 223, 216, 129, 181, 214, 138, 222, 232, 206, 245, 239, 126, 194, 109, 159, 125, 223, 43, 26, 247, 184, 32, 106, 68, 190, 94, 80, 34, 57, 52, 139, 82, 113, 0, 236, 103, 4, 200, 146, 188, 158, 81, 94, 0, 132, 77, 146, 91, 85, 150, 31, 45, 171, 215, 231, 239, 21, 120, 198, 120, 17, 127, 205, 0, 146, 158, 29, 75, 236, 156, 120, 60, 99, 188, 234, 15, 143, 244, 200, 138, 102, 191, 32, 39, 252, 253, 23, 104, 222, 255, 244, 19, 160, 73, 72, 36, 115, 255, 118, 142, 239, 88, 215, 245, 134, 43, 40, 184, 8, 141, 35, 0, 26, 55, 13, 201, 79, 77, 192, 127, 136, 227, 251, 133, 116, 7, 21, 25, 21, 172, 127, 120, 181, 92, 222, 7, 132, 13, 26, 194, 163, 171, 182, 49, 254, 234, 27, 157, 107, 88, 239, 199, 211, 166, 98, 241, 219, 229, 146, 195, 60, 44, 197, 1, 176, 159, 104, 0, 183, 64, 55, 84, 90, 249, 199, 24, 62, 37, 82, 254, 187, 4, 5, 155, 243, 156, 105, 91, 151, 160, 113, 211, 16, 52, 62, 132, 193, 247, 13, 99, 196, 99, 163, 137, 56, 59, 10, 189, 95, 67, 91, 210, 32, 224, 1, 96, 39, 178, 140, 229, 227, 200, 109, 45, 171, 57, 41, 147, 220, 24, 37, 10, 160, 208, 87, 72, 221, 244, 39, 69, 233, 71, 1, 208, 24, 181, 36, 63, 61, 1, 191, 196, 158, 249, 148, 156, 181, 42, 131, 245, 143, 172, 166, 34, 83, 46, 85, 159, 120, 253, 173, 60, 186, 106, 43, 33, 3, 29, 41, 35, 162, 112, 42, 4, 136, 81, 139, 124, 49, 87, 146, 28, 178, 119, 43, 14, 128, 253, 68, 0, 24, 252, 221, 26, 6, 116, 94, 58, 66, 79, 102, 242, 22, 110, 201, 111, 208, 187, 238, 141, 24, 2, 12, 68, 204, 140, 98, 228, 19, 163, 73, 188, 103, 40, 65, 201, 193, 168, 13, 13, 178, 18, 67, 129, 127, 35, 39, 11, 238, 68, 150, 28, 110, 34, 67, 184, 111, 213, 18, 138, 51, 219, 106, 136, 166, 160, 208, 179, 57, 180, 254, 119, 62, 186, 119, 22, 162, 104, 67, 235, 174, 101, 220, 220, 9, 248, 38, 248, 181, 189, 208, 197, 16, 173, 34, 41, 11, 246, 176, 251, 205, 157, 216, 234, 108, 24, 60, 60, 185, 233, 131, 47, 185, 246, 205, 15, 209, 26, 220, 218, 190, 128, 130, 35, 57, 55, 168, 140, 135, 29, 113, 33, 197, 1, 176, 159, 22, 14, 0, 64, 191, 115, 228, 108, 95, 115, 185, 153, 162, 157, 5, 221, 111, 85, 119, 35, 128, 87, 156, 55, 177, 87, 244, 103, 212, 179, 201, 196, 95, 59, 16, 223, 68, 63, 4, 117, 67, 129, 201, 72, 224, 49, 154, 41, 90, 73, 162, 200, 218, 47, 223, 234, 110, 107, 21, 20, 186, 141, 3, 107, 150, 242, 191, 251, 47, 193, 90, 87, 139, 214, 67, 199, 184, 103, 39, 226, 211, 191, 173, 188, 89, 215, 163, 182, 184, 150, 77, 79, 172, 39, 253, 183, 19, 192, 201, 144, 255, 234, 237, 140, 189, 92, 17, 245, 114, 22, 2, 204, 123, 183, 68, 154, 220, 217, 235, 40, 14, 128, 253, 132, 129, 92, 75, 223, 24, 191, 193, 254, 248, 12, 144, 63, 228, 57, 127, 101, 119, 67, 71, 103, 215, 65, 165, 85, 17, 144, 20, 200, 192, 91, 19, 25, 245, 92, 50, 113, 87, 245, 199, 103, 144, 111, 99, 103, 160, 9, 91, 127, 250, 132, 154, 138, 158, 163, 155, 160, 160, 208, 94, 82, 86, 46, 230, 147, 7, 102, 99, 53, 215, 161, 243, 210, 49, 254, 185, 137, 120, 199, 249, 56, 219, 172, 14, 83, 188, 175, 136, 117, 15, 173, 194, 148, 90, 2, 192, 232, 203, 174, 228, 145, 149, 155, 9, 142, 31, 224, 100, 203, 250, 60, 26, 65, 197, 194, 119, 42, 165, 144, 206, 92, 68, 113, 0, 236, 199, 11, 64, 235, 222, 178, 158, 61, 254, 18, 89, 199, 187, 38, 191, 154, 146, 125, 61, 183, 36, 176, 51, 104, 220, 52, 4, 142, 9, 38, 225, 182, 193, 178, 51, 112, 245, 0, 52, 198, 166, 133, 2, 117, 213, 149, 108, 248, 166, 203, 52, 46, 20, 20, 156, 194, 158, 63, 22, 241, 249, 131, 87, 96, 179, 152, 209, 123, 235, 25, 255, 220, 36, 188, 98, 188, 157, 109, 86, 135, 57, 190, 228, 40, 155, 231, 110, 160, 174, 180, 14, 149, 70, 195, 236, 23, 230, 115, 203, 199, 223, 160, 55, 186, 158, 212, 121, 31, 37, 68, 101, 229, 171, 133, 146, 100, 183, 164, 191, 226, 0, 216, 143, 17, 64, 173, 111, 249, 179, 15, 78, 14, 197, 35, 66, 110, 237, 156, 181, 60, 189, 79, 69, 1, 90, 67, 227, 166, 33, 112, 116, 16, 177, 87, 180, 108, 112, 178, 250, 179, 255, 82, 87, 85, 209, 202, 42, 5, 133, 158, 199, 238, 223, 22, 242, 197, 195, 87, 99, 179, 90, 208, 251, 24, 24, 247, 220, 36, 60, 251, 121, 57, 219, 172, 14, 33, 90, 108, 236, 126, 115, 7, 7, 62, 217, 135, 100, 147, 240, 10, 10, 230, 129, 37, 43, 57, 243, 238, 255, 115, 182, 105, 10, 45, 153, 94, 82, 198, 35, 246, 46, 86, 28, 0, 251, 16, 0, 55, 0, 181, 190, 101, 249, 187, 32, 8, 12, 188, 42, 1, 128, 234, 188, 106, 138, 118, 246, 124, 93, 0, 71, 224, 55, 196, 191, 113, 55, 66, 0, 170, 203, 74, 88, 175, 68, 1, 20, 122, 1, 59, 126, 253, 154, 47, 31, 185, 22, 209, 102, 69, 239, 107, 96, 252, 188, 137, 120, 70, 121, 58, 219, 172, 14, 81, 103, 170, 101, 211, 147, 27, 200, 90, 149, 9, 64, 204, 152, 113, 60, 182, 102, 7, 253, 39, 76, 113, 178, 101, 10, 167, 66, 130, 185, 246, 234, 3, 40, 14, 128, 125, 24, 56, 249, 179, 107, 45, 2, 0, 16, 58, 62, 28, 239, 88, 57, 236, 151, 245, 91, 58, 162, 85, 108, 117, 94, 159, 66, 128, 176, 51, 35, 91, 12, 175, 254, 244, 53, 37, 10, 160, 208, 163, 217, 182, 248, 115, 190, 121, 252, 70, 68, 209, 134, 193, 223, 141, 9, 243, 38, 225, 17, 222, 179, 54, 255, 210, 163, 165, 172, 123, 120, 117, 195, 121, 255, 184, 43, 175, 231, 159, 191, 174, 198, 39, 52, 220, 201, 150, 41, 180, 129, 78, 82, 241, 249, 2, 73, 234, 176, 190, 186, 226, 0, 216, 71, 195, 99, 253, 233, 77, 202, 0, 0, 32, 0, 73, 68, 65, 84, 236, 169, 28, 0, 4, 24, 120, 77, 34, 32, 103, 209, 230, 173, 205, 233, 22, 195, 92, 29, 255, 145, 1, 184, 5, 53, 173, 156, 168, 42, 45, 102, 227, 194, 22, 130, 129, 10, 10, 61, 130, 205, 139, 62, 230, 219, 167, 110, 65, 20, 109, 184, 5, 202, 155, 191, 123, 152, 135, 179, 205, 234, 16, 217, 107, 179, 216, 244, 196, 58, 106, 139, 107, 81, 169, 213, 92, 54, 239, 85, 110, 120, 255, 51, 52, 250, 22, 18, 198, 10, 46, 136, 4, 35, 68, 59, 142, 2, 20, 7, 192, 62, 132, 191, 255, 114, 234, 126, 74, 65, 73, 193, 4, 37, 201, 58, 223, 217, 43, 50, 177, 148, 119, 73, 71, 199, 30, 133, 160, 18, 8, 59, 171, 181, 40, 192, 124, 204, 181, 213, 78, 176, 72, 65, 193, 126, 54, 46, 92, 192, 247, 115, 239, 64, 18, 69, 140, 65, 70, 198, 63, 63, 25, 99, 72, 15, 74, 146, 147, 32, 245, 155, 67, 236, 122, 125, 59, 54, 179, 13, 55, 111, 31, 238, 254, 246, 23, 206, 186, 247, 65, 103, 91, 166, 208, 113, 158, 238, 232, 81, 192, 169, 244, 219, 21, 78, 143, 244, 247, 95, 78, 159, 225, 151, 120, 243, 80, 10, 247, 20, 98, 171, 179, 145, 241, 107, 26, 113, 87, 43, 229, 51, 1, 73, 129, 100, 175, 200, 164, 182, 176, 166, 97, 172, 162, 56, 159, 77, 11, 63, 224, 140, 235, 31, 112, 162, 101, 10, 10, 237, 103, 205, 231, 175, 179, 228, 149, 135, 144, 36, 9, 99, 176, 59, 227, 159, 159, 136, 91, 160, 177, 237, 133, 46, 130, 104, 17, 217, 243, 246, 78, 178, 215, 102, 1, 16, 20, 63, 128, 187, 191, 89, 66, 112, 255, 129, 78, 182, 172, 231, 34, 73, 18, 37, 199, 143, 144, 183, 111, 55, 166, 244, 227, 84, 230, 229, 96, 169, 173, 65, 180, 137, 168, 212, 42, 116, 70, 15, 180, 110, 70, 60, 66, 66, 241, 8, 10, 33, 112, 64, 34, 254, 113, 3, 80, 105, 28, 178, 21, 235, 68, 21, 159, 45, 144, 164, 228, 59, 4, 193, 210, 246, 116, 197, 1, 176, 151, 191, 15, 244, 219, 200, 240, 247, 8, 247, 32, 230, 252, 88, 142, 47, 57, 74, 225, 142, 2, 2, 198, 4, 225, 221, 191, 231, 213, 3, 59, 18, 65, 37, 16, 54, 61, 130, 227, 223, 53, 237, 42, 252, 215, 255, 94, 97, 220, 236, 91, 209, 27, 123, 86, 248, 84, 161, 239, 241, 231, 7, 47, 176, 236, 205, 167, 0, 112, 15, 243, 96, 252, 115, 19, 91, 136, 130, 185, 50, 230, 114, 51, 219, 95, 220, 66, 201, 33, 185, 76, 185, 255, 196, 51, 184, 227, 203, 31, 113, 247, 237, 121, 42, 133, 174, 128, 165, 186, 146, 148, 69, 223, 112, 112, 233, 143, 148, 231, 100, 117, 104, 173, 70, 175, 39, 116, 88, 18, 49, 83, 206, 34, 118, 202, 89, 184, 117, 238, 119, 48, 82, 42, 227, 31, 192, 171, 237, 153, 172, 28, 1, 216, 135, 212, 202, 223, 78, 201, 192, 171, 6, 97, 12, 54, 130, 4, 199, 191, 59, 138, 104, 182, 117, 161, 105, 61, 131, 192, 209, 65, 141, 251, 8, 0, 80, 81, 148, 199, 234, 79, 231, 59, 201, 34, 5, 133, 246, 241, 251, 187, 207, 54, 108, 254, 30, 17, 158, 140, 127, 126, 82, 143, 218, 252, 171, 114, 171, 216, 248, 248, 218, 134, 205, 127, 244, 165, 87, 112, 223, 162, 223, 250, 204, 230, 47, 137, 142, 77, 200, 62, 177, 238, 47, 190, 186, 234, 124, 182, 124, 248, 102, 135, 55, 127, 0, 107, 93, 29, 153, 219, 54, 177, 246, 181, 231, 249, 98, 206, 12, 254, 122, 241, 41, 74, 51, 211, 237, 182, 71, 130, 167, 23, 84, 73, 161, 109, 207, 228, 52, 7, 216, 10, 167, 195, 27, 40, 5, 24, 245, 200, 88, 66, 199, 135, 181, 185, 160, 96, 71, 30, 91, 231, 109, 6, 32, 108, 122, 4, 81, 23, 68, 119, 165, 125, 61, 130, 130, 45, 249, 45, 162, 0, 122, 163, 7, 143, 47, 59, 140, 103, 64, 167, 4, 174, 20, 20, 28, 142, 36, 73, 44, 254, 207, 131, 172, 253, 226, 13, 64, 110, 3, 158, 252, 204, 4, 116, 94, 58, 39, 91, 214, 126, 138, 83, 138, 216, 254, 159, 45, 88, 42, 45, 8, 130, 192, 133, 143, 63, 199, 204, 135, 158, 64, 16, 122, 223, 86, 96, 173, 171, 197, 148, 126, 2, 83, 218, 113, 76, 105, 71, 49, 165, 31, 167, 228, 196, 81, 44, 213, 85, 92, 251, 253, 10, 212, 218, 14, 39, 205, 183, 96, 207, 119, 159, 177, 233, 221, 215, 144, 36, 199, 138, 189, 168, 117, 58, 70, 94, 115, 11, 163, 174, 187, 221, 222, 227, 129, 15, 239, 242, 17, 110, 111, 107, 82, 239, 251, 173, 119, 15, 158, 64, 57, 192, 168, 135, 199, 16, 58, 161, 125, 101, 50, 59, 95, 219, 78, 206, 250, 44, 4, 149, 192, 224, 251, 134, 225, 209, 175, 103, 149, 9, 57, 26, 73, 148, 216, 251, 234, 46, 106, 242, 154, 38, 255, 77, 184, 226, 78, 102, 63, 245, 142, 147, 172, 82, 80, 104, 137, 36, 138, 252, 48, 239, 30, 54, 45, 252, 0, 0, 223, 129, 126, 140, 125, 106, 124, 171, 74, 160, 174, 74, 230, 95, 25, 164, 188, 183, 27, 209, 42, 162, 53, 24, 184, 254, 157, 79, 24, 125, 217, 149, 206, 54, 171, 211, 212, 150, 153, 48, 101, 156, 160, 52, 253, 132, 188, 225, 167, 31, 163, 52, 253, 4, 21, 249, 185, 167, 124, 218, 159, 249, 239, 215, 137, 153, 124, 102, 167, 238, 123, 108, 213, 239, 172, 152, 251, 176, 195, 55, 255, 198, 132, 141, 24, 205, 57, 207, 205, 199, 224, 211, 225, 30, 18, 54, 81, 98, 244, 61, 190, 194, 238, 211, 77, 82, 28, 0, 251, 208, 3, 181, 0, 35, 255, 111, 20, 225, 83, 90, 102, 181, 183, 70, 93, 89, 29, 107, 255, 241, 23, 117, 101, 117, 24, 2, 12, 12, 125, 112, 228, 169, 203, 8, 251, 8, 166, 253, 37, 28, 254, 248, 64, 147, 49, 149, 74, 205, 195, 63, 237, 38, 56, 46, 209, 73, 86, 41, 40, 252, 141, 40, 218, 248, 246, 201, 91, 216, 190, 228, 11, 0, 252, 135, 4, 48, 230, 137, 113, 104, 12, 61, 36, 133, 74, 130, 67, 95, 31, 224, 232, 15, 169, 0, 120, 4, 4, 114, 215, 87, 63, 19, 155, 60, 193, 201, 134, 117, 142, 194, 212, 131, 252, 250, 224, 237, 212, 150, 117, 188, 159, 72, 204, 164, 105, 204, 124, 225, 77, 187, 239, 93, 91, 106, 226, 155, 107, 47, 164, 182, 188, 204, 238, 107, 180, 23, 239, 240, 72, 46, 124, 253, 99, 60, 131, 219, 21, 213, 111, 204, 234, 187, 124, 132, 105, 167, 155, 208, 67, 222, 193, 46, 71, 29, 114, 34, 160, 202, 86, 215, 254, 243, 36, 189, 183, 158, 17, 255, 72, 98, 203, 243, 155, 168, 45, 170, 37, 125, 241, 9, 98, 47, 143, 239, 50, 35, 123, 2, 190, 131, 253, 240, 138, 243, 166, 252, 216, 223, 31, 36, 81, 180, 177, 236, 205, 167, 184, 233, 141, 69, 78, 180, 76, 193, 85, 17, 109, 86, 42, 139, 11, 40, 43, 204, 161, 188, 48, 143, 242, 194, 92, 170, 74, 139, 168, 173, 40, 163, 166, 162, 148, 234, 114, 19, 181, 21, 101, 72, 146, 72, 109, 101, 5, 162, 77, 206, 185, 177, 90, 234, 176, 212, 252, 29, 109, 170, 46, 55, 1, 96, 244, 146, 159, 174, 84, 106, 53, 122, 119, 89, 182, 87, 163, 213, 161, 51, 26, 145, 36, 137, 130, 227, 135, 40, 43, 144, 117, 60, 2, 71, 6, 51, 250, 209, 177, 168, 117, 61, 195, 113, 111, 158, 233, 31, 24, 27, 207, 61, 11, 151, 246, 138, 102, 62, 1, 113, 3, 176, 247, 25, 54, 125, 243, 58, 106, 75, 77, 246, 60, 89, 3, 176, 247, 135, 47, 187, 101, 243, 7, 40, 203, 206, 100, 201, 253, 55, 49, 251, 163, 133, 232, 61, 59, 36, 43, 61, 245, 125, 147, 116, 233, 157, 190, 194, 143, 167, 154, 160, 56, 0, 246, 83, 3, 184, 219, 204, 214, 14, 45, 10, 28, 25, 76, 212, 140, 104, 50, 254, 72, 163, 96, 115, 30, 62, 131, 124, 241, 27, 234, 223, 53, 22, 246, 16, 162, 46, 140, 102, 223, 27, 123, 154, 36, 84, 166, 172, 252, 153, 19, 187, 54, 16, 51, 114, 162, 243, 12, 83, 112, 42, 53, 229, 38, 178, 14, 236, 34, 39, 117, 47, 69, 25, 71, 79, 254, 57, 134, 41, 55, 3, 209, 214, 177, 207, 93, 91, 247, 105, 15, 33, 201, 161, 36, 61, 56, 6, 149, 182, 103, 228, 78, 155, 203, 235, 216, 246, 194, 22, 76, 135, 101, 101, 191, 129, 83, 166, 115, 251, 231, 63, 96, 180, 115, 211, 115, 53, 4, 181, 154, 152, 201, 211, 57, 248, 107, 199, 31, 20, 68, 171, 149, 212, 63, 151, 50, 108, 246, 181, 29, 94, 43, 73, 18, 135, 150, 253, 220, 225, 117, 157, 33, 114, 236, 196, 142, 110, 254, 0, 72, 2, 255, 158, 43, 73, 63, 207, 21, 132, 86, 159, 84, 21, 7, 192, 126, 170, 1, 119, 177, 3, 17, 128, 122, 18, 111, 26, 66, 113, 74, 17, 85, 185, 149, 28, 255, 246, 8, 238, 97, 238, 232, 155, 181, 21, 238, 75, 120, 68, 121, 226, 63, 60, 128, 226, 221, 69, 77, 198, 151, 188, 242, 48, 247, 127, 185, 30, 65, 213, 51, 190, 112, 21, 236, 71, 146, 36, 114, 143, 164, 112, 116, 203, 42, 78, 236, 218, 64, 214, 129, 93, 20, 103, 29, 119, 182, 89, 13, 132, 77, 138, 96, 228, 3, 163, 78, 217, 218, 218, 213, 168, 204, 174, 100, 235, 188, 77, 84, 231, 85, 1, 48, 254, 154, 155, 184, 230, 245, 5, 14, 73, 124, 115, 37, 226, 166, 206, 176, 203, 1, 0, 56, 180, 108, 177, 93, 14, 64, 105, 250, 49, 170, 138, 10, 236, 186, 167, 61, 244, 27, 63, 153, 201, 255, 247, 132, 189, 203, 19, 66, 202, 184, 20, 248, 161, 181, 23, 21, 7, 192, 126, 106, 0, 108, 118, 148, 244, 105, 12, 26, 146, 30, 26, 195, 134, 71, 215, 98, 173, 177, 146, 250, 217, 33, 6, 223, 63, 12, 149, 166, 239, 110, 116, 145, 231, 71, 83, 146, 82, 140, 100, 251, 59, 12, 144, 190, 119, 11, 91, 126, 250, 132, 113, 151, 221, 226, 68, 203, 20, 186, 138, 162, 140, 163, 28, 217, 178, 138, 35, 155, 87, 114, 116, 235, 106, 42, 77, 174, 217, 52, 43, 242, 204, 126, 12, 187, 123, 4, 130, 170, 103, 108, 254, 37, 7, 138, 217, 246, 226, 22, 44, 149, 102, 4, 65, 224, 162, 39, 231, 49, 243, 193, 199, 157, 109, 86, 151, 16, 158, 148, 140, 193, 219, 199, 174, 60, 128, 226, 163, 135, 40, 58, 122, 152, 128, 248, 142, 9, 31, 149, 156, 56, 214, 225, 123, 217, 139, 119, 120, 36, 103, 62, 249, 82, 103, 31, 130, 158, 144, 36, 105, 145, 32, 8, 45, 178, 21, 21, 7, 192, 126, 236, 118, 0, 0, 188, 99, 189, 25, 114, 235, 80, 246, 190, 183, 155, 170, 172, 74, 210, 127, 62, 78, 204, 236, 190, 155, 15, 96, 240, 55, 16, 60, 49, 180, 69, 207, 132, 165, 175, 63, 206, 208, 51, 47, 198, 221, 167, 111, 31, 147, 244, 6, 68, 209, 198, 241, 29, 235, 216, 247, 215, 18, 246, 253, 181, 132, 146, 236, 19, 221, 118, 111, 85, 163, 51, 123, 65, 173, 66, 208, 52, 221, 204, 69, 139, 8, 162, 4, 146, 220, 14, 183, 158, 184, 139, 227, 25, 116, 253, 144, 30, 147, 46, 157, 189, 54, 139, 61, 111, 239, 68, 180, 136, 104, 244, 122, 174, 127, 231, 19, 198, 204, 190, 202, 217, 102, 117, 25, 42, 141, 134, 232, 137, 211, 56, 180, 236, 39, 187, 214, 167, 254, 182, 152, 128, 123, 59, 38, 161, 95, 93, 82, 220, 225, 251, 8, 130, 64, 236, 212, 25, 36, 204, 156, 133, 119, 84, 12, 130, 32, 96, 169, 173, 161, 52, 253, 56, 233, 155, 214, 113, 124, 205, 138, 38, 249, 41, 0, 106, 157, 158, 115, 158, 159, 111, 87, 232, 191, 49, 18, 140, 88, 80, 206, 57, 192, 111, 205, 95, 83, 28, 0, 251, 233, 148, 3, 0, 16, 117, 118, 52, 197, 7, 138, 201, 94, 147, 73, 254, 198, 60, 220, 35, 60, 8, 26, 215, 119, 235, 223, 35, 206, 137, 162, 120, 87, 33, 150, 138, 191, 85, 44, 171, 76, 69, 44, 125, 253, 113, 46, 159, 171, 52, 11, 234, 137, 152, 107, 171, 57, 188, 97, 5, 251, 86, 45, 230, 192, 154, 165, 84, 153, 138, 218, 94, 116, 26, 84, 90, 53, 106, 79, 29, 26, 15, 29, 106, 163, 22, 149, 155, 6, 181, 155, 22, 181, 155, 22, 65, 167, 70, 165, 83, 163, 210, 169, 16, 180, 234, 38, 155, 126, 123, 176, 152, 106, 40, 250, 227, 56, 88, 108, 32, 192, 160, 235, 135, 16, 119, 113, 207, 113, 202, 143, 254, 144, 202, 161, 175, 15, 128, 4, 238, 190, 126, 220, 249, 245, 207, 196, 143, 159, 236, 108, 179, 186, 156, 184, 105, 103, 219, 239, 0, 172, 88, 202, 184, 59, 255, 217, 161, 90, 123, 209, 218, 46, 149, 221, 38, 76, 184, 239, 145, 86, 143, 27, 252, 99, 251, 19, 55, 237, 28, 198, 221, 113, 63, 155, 223, 127, 131, 195, 191, 47, 105, 120, 109, 202, 131, 79, 225, 31, 159, 208, 225, 123, 181, 134, 36, 241, 24, 138, 3, 224, 80, 74, 1, 74, 83, 219, 151, 64, 116, 42, 134, 221, 57, 130, 138, 244, 50, 202, 211, 202, 57, 177, 232, 24, 134, 64, 55, 188, 226, 188, 29, 98, 96, 79, 67, 227, 166, 33, 234, 194, 24, 142, 125, 157, 218, 100, 124, 203, 143, 255, 35, 249, 146, 155, 232, 55, 124, 156, 147, 44, 83, 232, 8, 54, 139, 153, 131, 235, 126, 99, 231, 210, 175, 217, 191, 250, 87, 44, 117, 53, 109, 47, 106, 134, 160, 22, 208, 250, 186, 161, 245, 119, 67, 231, 111, 68, 235, 99, 64, 227, 165, 71, 213, 69, 165, 119, 117, 249, 85, 20, 175, 60, 142, 104, 182, 33, 168, 5, 134, 223, 51, 146, 136, 105, 81, 93, 114, 47, 71, 35, 217, 68, 82, 222, 223, 67, 198, 159, 178, 122, 92, 96, 76, 156, 156, 233, 223, 71, 52, 253, 35, 70, 141, 195, 224, 229, 109, 87, 86, 126, 141, 169, 132, 140, 45, 235, 136, 158, 120, 218, 106, 185, 38, 232, 60, 58, 166, 223, 226, 25, 28, 202, 208, 203, 174, 57, 237, 28, 163, 127, 16, 211, 159, 248, 55, 33, 67, 71, 178, 118, 254, 243, 140, 184, 234, 70, 18, 206, 157, 213, 161, 251, 180, 193, 148, 119, 202, 165, 9, 247, 120, 9, 27, 27, 15, 42, 14, 128, 253, 124, 14, 76, 45, 61, 98, 162, 96, 71, 62, 65, 163, 130, 237, 186, 136, 218, 160, 102, 204, 227, 227, 88, 255, 240, 26, 234, 202, 234, 72, 253, 244, 32, 67, 30, 24, 129, 161, 143, 38, 5, 6, 142, 10, 162, 112, 75, 126, 147, 178, 64, 89, 132, 229, 94, 254, 239, 187, 45, 168, 84, 61, 163, 252, 170, 175, 33, 137, 34, 199, 119, 174, 99, 231, 210, 111, 216, 243, 199, 34, 170, 203, 74, 58, 180, 94, 80, 11, 232, 2, 221, 209, 135, 122, 96, 8, 245, 68, 27, 96, 236, 182, 51, 247, 218, 204, 114, 138, 87, 167, 33, 217, 68, 84, 90, 21, 73, 255, 28, 77, 200, 184, 182, 213, 61, 93, 1, 107, 173, 149, 157, 175, 110, 163, 96, 71, 62, 0, 49, 163, 147, 185, 235, 155, 37, 120, 6, 6, 57, 217, 178, 238, 67, 165, 209, 208, 111, 226, 52, 14, 47, 183, 47, 51, 63, 245, 143, 165, 29, 114, 0, 124, 34, 163, 59, 116, 125, 207, 208, 240, 118, 43, 45, 38, 94, 52, 155, 128, 1, 131, 8, 28, 232, 120, 13, 20, 149, 196, 109, 64, 19, 7, 160, 135, 156, 108, 185, 36, 26, 224, 32, 16, 239, 19, 239, 203, 164, 255, 156, 209, 169, 159, 102, 201, 129, 98, 54, 63, 179, 1, 209, 42, 98, 12, 49, 50, 248, 254, 97, 168, 123, 138, 208, 136, 131, 169, 206, 173, 38, 101, 254, 174, 38, 9, 129, 0, 151, 60, 246, 6, 147, 175, 185, 215, 73, 86, 41, 180, 70, 193, 137, 67, 108, 253, 233, 83, 118, 46, 251, 150, 210, 188, 204, 14, 173, 213, 250, 185, 97, 136, 240, 194, 16, 226, 129, 46, 216, 29, 65, 221, 253, 73, 176, 85, 71, 74, 48, 109, 202, 4, 81, 66, 235, 161, 101, 204, 19, 227, 240, 75, 232, 25, 249, 38, 53, 69, 53, 108, 157, 183, 137, 138, 244, 114, 0, 146, 102, 205, 230, 198, 5, 159, 163, 53, 244, 156, 190, 4, 142, 34, 125, 211, 58, 150, 253, 235, 110, 187, 214, 106, 141, 30, 220, 244, 203, 218, 118, 87, 72, 88, 106, 107, 248, 120, 230, 184, 118, 247, 20, 48, 250, 7, 114, 253, 143, 43, 157, 46, 183, 44, 65, 149, 222, 66, 232, 45, 129, 66, 69, 253, 88, 223, 77, 59, 239, 60, 86, 224, 121, 128, 210, 163, 38, 242, 183, 231, 117, 234, 98, 126, 137, 254, 12, 189, 115, 56, 0, 213, 121, 213, 28, 254, 223, 65, 68, 171, 99, 155, 86, 244, 20, 140, 161, 70, 66, 166, 180, 124, 2, 91, 254, 214, 211, 148, 23, 228, 180, 178, 66, 161, 59, 49, 215, 86, 179, 125, 201, 23, 188, 117, 253, 25, 188, 116, 225, 96, 254, 250, 223, 43, 237, 219, 252, 5, 1, 125, 136, 7, 62, 99, 195, 9, 153, 157, 72, 240, 69, 3, 241, 78, 10, 69, 31, 230, 233, 148, 205, 191, 34, 37, 31, 211, 134, 12, 16, 37, 12, 126, 6, 38, 252, 123, 114, 143, 217, 252, 75, 143, 152, 88, 255, 200, 154, 134, 205, 127, 198, 125, 15, 113, 235, 39, 223, 245, 201, 205, 31, 32, 98, 244, 184, 14, 135, 230, 235, 177, 84, 87, 146, 187, 119, 103, 187, 231, 107, 13, 110, 248, 197, 180, 63, 55, 164, 186, 184, 144, 204, 173, 27, 236, 49, 205, 161, 8, 224, 94, 167, 97, 118, 227, 49, 37, 158, 218, 57, 246, 1, 87, 2, 254, 85, 217, 149, 244, 59, 59, 186, 83, 81, 0, 239, 88, 31, 68, 171, 72, 201, 193, 98, 234, 74, 234, 168, 45, 168, 193, 127, 120, 64, 159, 140, 211, 120, 70, 123, 81, 180, 163, 16, 91, 237, 223, 73, 150, 86, 115, 29, 37, 217, 105, 140, 152, 121, 185, 19, 45, 235, 187, 100, 31, 220, 197, 138, 5, 255, 230, 155, 199, 111, 98, 215, 242, 239, 40, 205, 205, 104, 115, 141, 160, 22, 112, 139, 244, 198, 107, 104, 48, 190, 19, 35, 241, 72, 8, 64, 23, 232, 222, 225, 4, 61, 135, 34, 73, 148, 110, 203, 161, 98, 175, 28, 54, 119, 15, 243, 96, 252, 243, 147, 240, 8, 235, 25, 109, 168, 179, 215, 101, 177, 253, 197, 173, 88, 171, 44, 168, 212, 106, 174, 124, 245, 29, 102, 62, 248, 184, 211, 159, 48, 157, 137, 74, 173, 166, 248, 200, 33, 74, 78, 28, 181, 107, 189, 193, 211, 139, 168, 228, 73, 237, 158, 95, 145, 151, 67, 94, 202, 174, 118, 207, 47, 60, 188, 159, 1, 51, 46, 64, 163, 215, 219, 99, 158, 195, 16, 84, 248, 45, 125, 233, 217, 79, 234, 255, 173, 68, 0, 58, 135, 141, 147, 81, 128, 178, 227, 165, 228, 108, 232, 120, 43, 200, 230, 36, 92, 147, 72, 228, 116, 57, 249, 168, 120, 79, 17, 105, 63, 187, 142, 24, 74, 119, 162, 214, 171, 233, 55, 43, 166, 197, 248, 222, 63, 127, 98, 239, 159, 246, 101, 252, 42, 116, 28, 115, 109, 53, 155, 23, 125, 204, 252, 203, 199, 242, 218, 156, 209, 108, 248, 246, 61, 106, 43, 219, 72, 182, 18, 64, 31, 226, 129, 239, 132, 72, 66, 175, 24, 130, 255, 244, 24, 140, 241, 126, 168, 244, 206, 63, 210, 146, 44, 34, 69, 127, 157, 160, 242, 128, 172, 57, 224, 29, 239, 195, 196, 23, 39, 99, 12, 50, 58, 217, 178, 118, 32, 193, 225, 175, 14, 178, 235, 191, 219, 17, 45, 54, 220, 188, 125, 184, 103, 225, 82, 166, 220, 114, 151, 179, 45, 115, 9, 34, 199, 218, 175, 26, 154, 190, 105, 77, 135, 230, 15, 56, 231, 194, 14, 205, 47, 205, 72, 99, 241, 253, 55, 81, 124, 252, 72, 219, 147, 187, 18, 137, 73, 239, 151, 73, 253, 235, 255, 169, 68, 0, 58, 207, 126, 224, 18, 32, 216, 116, 184, 132, 168, 25, 209, 157, 147, 10, 21, 32, 120, 116, 48, 101, 71, 75, 169, 202, 173, 162, 50, 67, 62, 174, 241, 138, 239, 123, 149, 1, 198, 16, 35, 85, 89, 149, 212, 22, 54, 205, 34, 63, 182, 125, 45, 227, 46, 189, 25, 173, 190, 111, 134, 59, 187, 131, 226, 172, 227, 252, 245, 241, 127, 248, 250, 241, 27, 217, 189, 252, 59, 202, 11, 115, 219, 92, 163, 245, 49, 224, 145, 16, 128, 239, 196, 40, 60, 7, 7, 162, 243, 55, 58, 37, 180, 127, 42, 108, 85, 22, 10, 255, 56, 134, 57, 95, 86, 199, 11, 26, 21, 204, 216, 39, 198, 163, 245, 112, 253, 118, 190, 182, 90, 27, 59, 231, 111, 35, 227, 143, 52, 0, 130, 226, 250, 243, 192, 226, 149, 196, 140, 81, 42, 99, 234, 113, 243, 245, 99, 239, 194, 47, 236, 90, 91, 87, 81, 78, 252, 180, 153, 184, 181, 83, 38, 217, 205, 215, 143, 244, 13, 171, 169, 46, 110, 127, 89, 107, 141, 169, 152, 3, 75, 190, 167, 224, 96, 10, 6, 111, 31, 188, 66, 194, 157, 161, 114, 42, 72, 144, 189, 244, 165, 103, 55, 128, 226, 0, 56, 2, 9, 249, 40, 224, 70, 107, 141, 85, 16, 173, 54, 2, 71, 218, 87, 17, 80, 143, 160, 18, 8, 73, 14, 163, 40, 165, 144, 218, 226, 90, 202, 143, 149, 33, 168, 5, 188, 98, 251, 158, 19, 224, 21, 239, 77, 193, 230, 124, 164, 70, 249, 16, 230, 234, 74, 106, 202, 77, 12, 158, 122, 129, 19, 45, 235, 125, 72, 146, 196, 145, 205, 127, 241, 235, 235, 143, 177, 232, 249, 123, 57, 190, 115, 61, 150, 218, 234, 211, 174, 81, 233, 212, 24, 227, 252, 240, 155, 24, 133, 215, 200, 80, 244, 33, 30, 168, 92, 176, 195, 165, 185, 168, 154, 162, 223, 143, 97, 173, 168, 3, 32, 230, 252, 56, 70, 220, 159, 212, 35, 154, 250, 212, 20, 214, 176, 249, 217, 13, 20, 239, 147, 55, 155, 132, 51, 206, 228, 190, 31, 127, 199, 47, 162, 125, 93, 72, 251, 10, 58, 163, 59, 199, 215, 172, 160, 198, 212, 177, 10, 148, 122, 60, 67, 194, 8, 25, 58, 178, 221, 243, 85, 26, 13, 105, 235, 87, 117, 236, 38, 146, 68, 89, 86, 6, 169, 127, 252, 202, 222, 69, 95, 81, 116, 248, 0, 230, 170, 74, 244, 158, 222, 157, 22, 252, 105, 47, 42, 80, 255, 250, 210, 179, 159, 65, 159, 60, 93, 238, 50, 62, 5, 110, 16, 212, 2, 147, 95, 153, 138, 87, 76, 231, 55, 107, 107, 181, 133, 77, 207, 108, 160, 236, 168, 44, 115, 25, 121, 126, 52, 225, 103, 70, 116, 250, 186, 61, 141, 188, 245, 185, 164, 253, 216, 84, 126, 83, 16, 4, 238, 252, 104, 5, 253, 147, 219, 95, 190, 163, 208, 58, 117, 85, 21, 108, 249, 233, 19, 54, 124, 243, 46, 133, 233, 237, 11, 81, 234, 67, 60, 112, 239, 239, 143, 91, 180, 183, 75, 61, 229, 183, 70, 77, 122, 41, 37, 235, 50, 144, 172, 34, 130, 74, 96, 240, 205, 67, 137, 62, 63, 214, 217, 102, 181, 11, 83, 106, 9, 219, 95, 220, 74, 93, 105, 45, 0, 147, 110, 188, 157, 43, 95, 121, 187, 215, 105, 250, 59, 138, 77, 239, 190, 198, 238, 111, 63, 181, 107, 109, 216, 240, 81, 204, 122, 171, 253, 107, 37, 81, 228, 187, 27, 47, 197, 148, 230, 24, 105, 96, 239, 200, 126, 196, 76, 156, 198, 128, 115, 46, 196, 63, 174, 75, 187, 53, 154, 69, 51, 254, 247, 4, 9, 149, 138, 3, 224, 56, 130, 128, 67, 128, 175, 95, 130, 63, 19, 94, 152, 236, 16, 247, 202, 92, 110, 102, 211, 211, 235, 27, 178, 125, 251, 93, 28, 75, 104, 43, 25, 242, 189, 26, 9, 246, 191, 181, 151, 138, 180, 242, 38, 195, 126, 225, 209, 60, 252, 227, 110, 244, 238, 246, 101, 255, 246, 117, 74, 243, 179, 88, 247, 229, 91, 108, 250, 254, 195, 182, 207, 245, 1, 149, 65, 131, 123, 188, 31, 238, 3, 252, 209, 120, 57, 55, 153, 169, 189, 84, 236, 205, 167, 108, 167, 124, 124, 161, 49, 106, 25, 245, 208, 24, 2, 71, 246, 140, 26, 249, 172, 85, 153, 236, 125, 111, 23, 162, 69, 68, 165, 209, 112, 249, 75, 111, 112, 198, 173, 246, 149, 186, 245, 21, 178, 182, 111, 230, 151, 127, 222, 102, 215, 90, 65, 173, 230, 198, 197, 107, 48, 120, 181, 255, 225, 45, 109, 195, 42, 150, 63, 118, 191, 93, 247, 59, 29, 254, 241, 9, 12, 58, 255, 18, 6, 93, 112, 41, 26, 125, 23, 104, 194, 168, 184, 240, 46, 47, 225, 215, 30, 237, 0, 44, 40, 151, 2, 172, 54, 34, 212, 18, 145, 162, 154, 40, 21, 248, 75, 224, 135, 136, 31, 224, 11, 248, 1, 62, 8, 24, 0, 61, 96, 4, 180, 192, 233, 210, 125, 203, 144, 147, 251, 74, 37, 176, 10, 80, 129, 68, 45, 96, 2, 76, 168, 40, 1, 76, 2, 148, 136, 96, 18, 36, 114, 4, 53, 121, 26, 145, 188, 91, 189, 133, 43, 128, 119, 1, 134, 221, 61, 130, 168, 25, 209, 14, 249, 255, 172, 43, 171, 99, 243, 83, 235, 169, 200, 148, 243, 1, 250, 98, 36, 160, 38, 191, 154, 148, 215, 118, 183, 40, 141, 28, 63, 231, 54, 230, 60, 243, 190, 147, 172, 234, 153, 100, 31, 220, 197, 234, 207, 230, 179, 251, 183, 239, 177, 181, 67, 214, 84, 23, 104, 196, 35, 33, 16, 183, 104, 159, 30, 211, 13, 79, 18, 37, 76, 27, 51, 169, 62, 42, 135, 131, 141, 193, 70, 198, 60, 62, 30, 207, 40, 215, 119, 22, 69, 139, 200, 254, 143, 83, 72, 255, 93, 238, 149, 224, 238, 235, 199, 173, 159, 46, 36, 225, 140, 51, 157, 108, 153, 235, 99, 173, 171, 227, 147, 11, 38, 97, 173, 171, 181, 107, 253, 153, 79, 190, 196, 128, 179, 207, 239, 208, 154, 85, 47, 62, 201, 161, 229, 139, 237, 186, 95, 91, 24, 253, 3, 24, 121, 205, 45, 36, 94, 56, 199, 161, 21, 4, 130, 196, 59, 119, 250, 10, 247, 186, 252, 167, 249, 191, 38, 201, 71, 175, 34, 65, 16, 73, 20, 4, 18, 36, 137, 65, 146, 192, 0, 1, 34, 1, 151, 202, 2, 147, 68, 177, 238, 197, 169, 163, 133, 204, 189, 187, 116, 26, 131, 134, 201, 243, 167, 226, 30, 234, 152, 210, 162, 58, 83, 45, 155, 159, 217, 208, 224, 4, 132, 157, 25, 65, 212, 249, 209, 14, 185, 118, 79, 33, 123, 101, 22, 153, 75, 211, 154, 140, 9, 130, 192, 173, 239, 254, 202, 160, 201, 51, 157, 99, 84, 15, 161, 254, 124, 127, 237, 151, 111, 112, 112, 237, 50, 36, 169, 69, 99, 176, 38, 8, 106, 1, 67, 164, 55, 158, 137, 129, 232, 130, 220, 187, 201, 74, 199, 32, 214, 89, 41, 254, 43, 141, 186, 252, 74, 0, 124, 19, 252, 24, 253, 104, 50, 122, 111, 215, 143, 90, 84, 23, 84, 179, 227, 149, 173, 13, 199, 126, 161, 9, 137, 220, 249, 213, 207, 4, 197, 245, 111, 99, 165, 66, 61, 203, 254, 117, 55, 233, 155, 214, 217, 181, 182, 255, 89, 231, 113, 214, 211, 47, 119, 104, 141, 205, 92, 199, 143, 119, 93, 71, 209, 145, 131, 118, 221, 179, 61, 120, 71, 244, 99, 234, 191, 158, 37, 108, 248, 40, 71, 93, 50, 229, 46, 31, 97, 152, 75, 57, 0, 11, 36, 73, 107, 43, 99, 152, 74, 98, 146, 4, 163, 128, 81, 8, 12, 162, 7, 229, 42, 228, 165, 30, 226, 197, 51, 70, 97, 174, 169, 198, 39, 222, 151, 137, 47, 77, 118, 216, 25, 169, 185, 220, 204, 150, 231, 54, 82, 118, 76, 254, 114, 8, 158, 24, 74, 204, 165, 113, 61, 232, 167, 211, 57, 36, 81, 98, 255, 155, 123, 27, 42, 35, 234, 241, 9, 137, 228, 225, 159, 118, 227, 230, 233, 227, 36, 203, 92, 23, 209, 102, 101, 215, 242, 239, 88, 249, 209, 203, 228, 29, 221, 223, 230, 124, 181, 187, 14, 143, 4, 127, 220, 251, 251, 119, 153, 238, 126, 87, 98, 41, 171, 165, 120, 229, 9, 172, 229, 114, 178, 95, 248, 228, 8, 134, 223, 55, 18, 149, 214, 245, 147, 253, 10, 118, 228, 179, 235, 245, 29, 88, 42, 205, 0, 140, 153, 115, 53, 215, 188, 190, 0, 189, 123, 207, 208, 39, 112, 21, 82, 126, 248, 138, 245, 111, 190, 100, 215, 90, 175, 176, 8, 174, 249, 118, 121, 135, 215, 85, 21, 21, 240, 235, 131, 119, 216, 173, 67, 208, 30, 4, 65, 96, 216, 156, 235, 24, 127, 247, 131, 142, 168, 30, 176, 248, 123, 227, 225, 212, 173, 99, 174, 36, 105, 130, 76, 140, 23, 4, 102, 72, 2, 103, 11, 242, 166, 223, 243, 190, 117, 154, 177, 238, 147, 5, 124, 253, 127, 119, 2, 16, 63, 123, 0, 9, 215, 56, 78, 215, 217, 90, 109, 97, 235, 188, 205, 148, 28, 148, 91, 82, 6, 36, 5, 18, 123, 101, 127, 84, 26, 215, 78, 196, 114, 20, 53, 5, 53, 164, 188, 38, 159, 139, 54, 102, 240, 212, 11, 184, 249, 173, 159, 251, 180, 24, 74, 99, 108, 22, 51, 187, 150, 127, 199, 138, 15, 94, 160, 48, 45, 181, 205, 249, 90, 127, 55, 60, 19, 3, 49, 198, 248, 66, 15, 233, 123, 223, 156, 154, 180, 82, 74, 54, 100, 32, 89, 68, 16, 96, 192, 21, 9, 12, 184, 60, 193, 229, 29, 100, 73, 148, 72, 253, 246, 16, 71, 126, 56, 12, 18, 104, 244, 122, 102, 207, 123, 141, 51, 110, 187, 199, 217, 166, 245, 72, 74, 51, 210, 248, 230, 218, 142, 213, 233, 215, 35, 8, 2, 55, 47, 223, 132, 206, 216, 241, 168, 151, 185, 178, 130, 229, 143, 223, 79, 206, 238, 237, 118, 221, 187, 189, 196, 77, 61, 155, 51, 159, 124, 17, 181, 174, 115, 229, 171, 42, 145, 238, 143, 0, 188, 91, 42, 249, 170, 224, 82, 73, 224, 66, 36, 166, 3, 174, 127, 40, 103, 7, 239, 95, 125, 49, 123, 150, 45, 70, 80, 9, 140, 127, 110, 18, 126, 131, 29, 39, 49, 106, 171, 179, 177, 253, 229, 45, 20, 238, 42, 0, 192, 43, 206, 155, 1, 55, 13, 66, 99, 236, 241, 190, 211, 105, 177, 154, 37, 42, 242, 235, 200, 95, 159, 67, 233, 182, 150, 146, 192, 151, 62, 254, 38, 147, 174, 238, 219, 95, 154, 230, 154, 42, 54, 47, 250, 152, 213, 159, 188, 70, 105, 126, 27, 194, 84, 2, 24, 34, 188, 240, 28, 20, 136, 62, 172, 7, 127, 12, 69, 137, 178, 157, 185, 84, 236, 147, 63, 15, 106, 189, 154, 225, 247, 142, 36, 108, 146, 235, 231, 201, 212, 149, 214, 177, 235, 191, 219, 41, 218, 43, 11, 19, 249, 69, 246, 227, 182, 79, 23, 18, 61, 106, 172, 147, 45, 235, 217, 124, 49, 123, 6, 149, 5, 246, 201, 179, 95, 242, 238, 151, 132, 12, 25, 110, 215, 90, 201, 102, 99, 207, 194, 207, 217, 246, 191, 119, 237, 206, 67, 104, 15, 177, 83, 206, 226, 236, 231, 231, 119, 234, 129, 71, 18, 184, 166, 91, 28, 128, 249, 153, 146, 155, 209, 157, 179, 80, 113, 157, 4, 179, 0, 215, 87, 222, 232, 36, 149, 197, 69, 204, 155, 56, 140, 178, 188, 92, 220, 2, 221, 152, 50, 127, 154, 67, 5, 71, 68, 171, 200, 158, 183, 119, 145, 189, 70, 214, 96, 119, 11, 114, 35, 225, 182, 193, 232, 123, 97, 23, 65, 73, 132, 202, 66, 51, 229, 133, 102, 36, 17, 144, 36, 10, 127, 59, 74, 221, 73, 65, 151, 122, 52, 58, 61, 247, 127, 185, 142, 136, 68, 135, 157, 147, 245, 24, 106, 43, 203, 88, 255, 245, 59, 172, 253, 226, 77, 42, 77, 133, 167, 157, 43, 104, 84, 184, 199, 251, 225, 145, 24, 216, 99, 178, 249, 79, 133, 88, 107, 165, 120, 77, 26, 117, 185, 242, 121, 191, 123, 168, 59, 163, 30, 73, 198, 43, 186, 123, 106, 170, 59, 67, 225, 158, 2, 118, 191, 190, 131, 186, 82, 249, 184, 98, 240, 140, 115, 185, 105, 193, 23, 184, 251, 245, 140, 126, 4, 174, 204, 242, 71, 239, 35, 109, 227, 106, 187, 214, 158, 241, 208, 51, 36, 94, 52, 187, 237, 137, 167, 161, 44, 59, 147, 29, 159, 45, 224, 200, 159, 75, 17, 173, 214, 78, 93, 235, 84, 140, 187, 227, 1, 70, 94, 115, 75, 103, 46, 241, 82, 151, 30, 140, 189, 111, 146, 134, 159, 255, 216, 220, 121, 90, 61, 95, 32, 112, 61, 48, 152, 62, 34, 62, 164, 51, 26, 9, 79, 28, 202, 214, 133, 95, 97, 169, 178, 80, 118, 188, 148, 240, 201, 145, 14, 107, 113, 42, 168, 4, 66, 147, 195, 144, 36, 137, 146, 253, 197, 88, 171, 172, 20, 239, 42, 196, 43, 214, 27, 157, 79, 207, 254, 82, 175, 71, 146, 160, 218, 100, 161, 56, 173, 134, 154, 10, 155, 44, 185, 4, 32, 8, 24, 194, 188, 168, 62, 110, 106, 34, 16, 36, 218, 108, 164, 110, 254, 139, 177, 23, 95, 143, 70, 215, 212, 17, 170, 50, 21, 113, 104, 195, 239, 88, 205, 117, 120, 5, 132, 116, 227, 255, 69, 215, 82, 105, 42, 228, 207, 15, 94, 228, 139, 71, 174, 229, 224, 186, 229, 152, 79, 35, 220, 163, 210, 107, 240, 28, 22, 132, 255, 25, 253, 112, 139, 246, 113, 9, 105, 222, 206, 96, 46, 172, 162, 240, 143, 99, 88, 74, 228, 39, 173, 224, 209, 33, 140, 125, 122, 130, 203, 203, 250, 74, 54, 145, 195, 95, 29, 100, 223, 251, 123, 176, 214, 90, 81, 105, 52, 92, 248, 216, 179, 92, 61, 255, 61, 187, 66, 207, 10, 45, 49, 165, 29, 35, 119, 207, 14, 187, 214, 122, 133, 134, 17, 53, 110, 114, 167, 238, 111, 240, 242, 38, 102, 242, 116, 6, 206, 156, 133, 90, 171, 165, 50, 63, 23, 115, 85, 101, 167, 174, 217, 156, 236, 93, 91, 137, 155, 122, 54, 110, 62, 126, 246, 94, 34, 207, 225, 155, 241, 92, 73, 82, 93, 249, 143, 185, 23, 156, 255, 216, 220, 119, 16, 120, 25, 249, 92, 191, 79, 170, 86, 4, 198, 196, 33, 138, 54, 142, 108, 88, 75, 117, 126, 53, 150, 42, 11, 65, 73, 157, 83, 9, 108, 130, 0, 1, 67, 3, 113, 11, 50, 82, 176, 35, 31, 91, 157, 141, 162, 29, 133, 232, 188, 117, 184, 135, 247, 236, 196, 161, 186, 42, 27, 197, 105, 181, 84, 149, 88, 229, 167, 254, 102, 168, 116, 106, 180, 190, 6, 170, 79, 152, 154, 140, 215, 148, 155, 200, 72, 217, 74, 210, 121, 87, 162, 82, 203, 111, 239, 252, 99, 7, 120, 109, 206, 104, 182, 45, 254, 140, 205, 223, 127, 136, 70, 171, 35, 118, 84, 251, 27, 127, 184, 34, 21, 197, 249, 172, 88, 240, 111, 190, 124, 248, 26, 142, 108, 93, 133, 213, 92, 119, 202, 185, 42, 131, 6, 207, 33, 65, 248, 157, 209, 15, 67, 184, 23, 66, 47, 200, 23, 169, 216, 87, 128, 105, 93, 6, 162, 217, 134, 32, 8, 12, 184, 42, 129, 97, 119, 142, 112, 121, 101, 191, 234, 130, 106, 182, 205, 219, 76, 206, 122, 249, 120, 198, 55, 60, 146, 187, 191, 251, 133, 228, 43, 174, 85, 242, 87, 28, 72, 77, 169, 137, 227, 171, 87, 216, 181, 86, 163, 55, 144, 112, 238, 197, 14, 177, 67, 239, 225, 73, 196, 232, 241, 12, 155, 125, 45, 225, 73, 201, 232, 61, 188, 48, 87, 85, 80, 91, 86, 218, 249, 139, 75, 18, 85, 69, 5, 196, 159, 121, 174, 189, 87, 200, 115, 216, 59, 110, 174, 36, 169, 130, 202, 184, 76, 128, 127, 3, 74, 205, 202, 73, 36, 81, 100, 193, 117, 151, 177, 103, 233, 207, 128, 99, 245, 1, 26, 83, 180, 183, 144, 29, 175, 108, 197, 82, 41, 215, 117, 7, 79, 8, 37, 250, 146, 216, 30, 83, 183, 93, 143, 36, 66, 121, 65, 29, 149, 5, 22, 78, 95, 168, 38, 83, 182, 35, 135, 138, 148, 130, 22, 227, 163, 46, 184, 154, 171, 95, 252, 28, 65, 16, 88, 248, 204, 237, 108, 94, 244, 113, 195, 107, 106, 141, 150, 135, 127, 222, 67, 80, 244, 64, 7, 90, 222, 61, 148, 100, 167, 177, 230, 243, 215, 217, 244, 195, 135, 109, 158, 49, 170, 61, 116, 120, 38, 6, 226, 62, 208, 223, 229, 213, 250, 218, 139, 88, 107, 165, 100, 125, 6, 181, 89, 178, 40, 148, 206, 75, 199, 200, 7, 70, 117, 90, 126, 187, 59, 200, 221, 152, 195, 222, 119, 119, 97, 169, 146, 63, 163, 195, 207, 155, 197, 117, 239, 252, 15, 119, 95, 187, 159, 224, 20, 78, 129, 41, 237, 24, 223, 94, 111, 223, 38, 174, 247, 244, 226, 230, 165, 93, 219, 190, 183, 34, 47, 135, 204, 109, 27, 201, 218, 182, 137, 204, 109, 27, 59, 21, 29, 184, 226, 179, 159, 58, 212, 158, 184, 30, 1, 118, 59, 196, 93, 126, 183, 84, 58, 203, 163, 142, 239, 5, 184, 31, 80, 14, 176, 26, 33, 8, 2, 67, 103, 94, 64, 202, 111, 191, 82, 81, 88, 64, 225, 174, 2, 252, 135, 4, 224, 22, 232, 216, 48, 165, 49, 216, 157, 208, 241, 225, 20, 239, 43, 194, 92, 86, 71, 85, 102, 37, 229, 71, 202, 240, 25, 228, 139, 218, 5, 181, 217, 91, 163, 174, 194, 70, 81, 90, 13, 181, 229, 182, 182, 39, 159, 196, 16, 226, 65, 109, 110, 37, 182, 170, 166, 130, 54, 185, 169, 41, 8, 2, 196, 143, 153, 202, 190, 85, 75, 200, 62, 248, 119, 235, 78, 73, 20, 41, 201, 78, 35, 233, 252, 171, 28, 102, 123, 87, 147, 119, 116, 63, 139, 255, 243, 16, 223, 63, 123, 39, 233, 123, 54, 35, 218, 78, 125, 174, 168, 241, 54, 224, 51, 38, 12, 223, 9, 145, 232, 131, 220, 29, 118, 236, 228, 108, 234, 242, 171, 40, 90, 113, 12, 75, 177, 220, 28, 202, 111, 176, 63, 227, 230, 78, 196, 59, 214, 181, 203, 63, 173, 181, 86, 246, 125, 176, 135, 67, 95, 28, 64, 180, 136, 114, 150, 255, 11, 255, 101, 246, 11, 243, 209, 185, 185, 246, 113, 69, 79, 197, 224, 233, 205, 238, 111, 63, 65, 180, 181, 255, 187, 164, 30, 155, 185, 142, 65, 23, 92, 134, 174, 11, 203, 47, 245, 30, 158, 4, 14, 76, 36, 110, 218, 57, 12, 191, 252, 122, 130, 18, 135, 162, 82, 107, 40, 207, 206, 236, 112, 206, 128, 187, 127, 32, 97, 35, 70, 219, 99, 70, 93, 167, 190, 25, 222, 175, 144, 18, 69, 145, 5, 130, 68, 207, 142, 167, 118, 3, 69, 233, 39, 120, 121, 250, 88, 42, 139, 139, 208, 123, 235, 153, 244, 202, 25, 14, 119, 2, 64, 238, 26, 182, 231, 157, 157, 228, 172, 207, 6, 64, 235, 165, 163, 255, 181, 3, 93, 186, 155, 160, 36, 66, 105, 78, 29, 213, 38, 11, 109, 232, 211, 180, 138, 173, 198, 66, 225, 210, 35, 88, 79, 214, 79, 215, 35, 8, 2, 151, 63, 247, 33, 161, 253, 135, 240, 198, 85, 227, 91, 136, 223, 220, 253, 201, 74, 226, 199, 76, 237, 132, 229, 93, 79, 73, 246, 9, 150, 191, 245, 12, 59, 151, 125, 131, 36, 182, 114, 22, 210, 8, 173, 159, 27, 94, 195, 130, 113, 235, 231, 227, 242, 165, 111, 29, 66, 148, 40, 79, 201, 167, 124, 79, 62, 136, 18, 130, 32, 16, 127, 217, 0, 6, 92, 153, 224, 242, 17, 46, 211, 225, 18, 118, 191, 177, 131, 170, 92, 57, 97, 53, 184, 255, 64, 110, 253, 223, 183, 68, 12, 29, 225, 100, 203, 122, 63, 223, 223, 114, 185, 221, 226, 60, 231, 255, 231, 61, 162, 198, 117, 255, 182, 86, 91, 106, 98, 231, 87, 31, 179, 247, 135, 47, 145, 218, 233, 188, 4, 38, 12, 97, 246, 7, 223, 216, 115, 187, 114, 187, 30, 13, 223, 148, 36, 253, 69, 255, 156, 251, 36, 34, 95, 10, 208, 51, 186, 106, 56, 25, 163, 143, 47, 209, 163, 198, 178, 237, 251, 175, 177, 84, 155, 41, 216, 145, 79, 216, 164, 112, 52, 14, 22, 91, 81, 105, 84, 132, 78, 8, 71, 227, 166, 161, 40, 165, 16, 91, 173, 156, 23, 32, 217, 36, 188, 226, 188, 93, 238, 156, 209, 102, 22, 41, 58, 81, 67, 109, 69, 199, 61, 245, 122, 84, 90, 53, 134, 8, 47, 170, 79, 148, 54, 73, 10, 4, 56, 184, 102, 41, 81, 195, 146, 209, 27, 61, 201, 77, 77, 105, 242, 90, 85, 105, 241, 41, 163, 0, 18, 206, 221, 67, 171, 74, 139, 249, 227, 189, 231, 249, 234, 177, 27, 228, 232, 197, 105, 60, 35, 93, 144, 59, 190, 227, 34, 241, 25, 27, 142, 214, 199, 208, 171, 54, 127, 107, 121, 29, 69, 43, 79, 80, 125, 204, 4, 18, 232, 125, 244, 140, 126, 52, 153, 168, 25, 253, 92, 58, 178, 33, 217, 68, 82, 191, 59, 196, 158, 183, 119, 97, 46, 151, 29, 211, 73, 55, 220, 198, 29, 95, 44, 194, 47, 34, 202, 201, 214, 245, 13, 242, 247, 239, 161, 248, 232, 97, 187, 214, 250, 199, 14, 32, 116, 88, 251, 59, 3, 58, 10, 141, 193, 141, 200, 177, 19, 136, 24, 61, 158, 19, 107, 255, 196, 118, 154, 220, 158, 122, 106, 76, 197, 140, 186, 254, 118, 187, 190, 219, 59, 236, 0, 188, 87, 38, 141, 85, 215, 176, 28, 152, 67, 47, 16, 237, 233, 78, 252, 163, 162, 241, 14, 13, 35, 229, 183, 95, 48, 87, 152, 41, 218, 83, 72, 216, 164, 136, 46, 73, 92, 242, 77, 240, 195, 127, 80, 0, 69, 123, 10, 176, 86, 91, 169, 56, 94, 78, 121, 106, 41, 222, 253, 125, 208, 184, 185, 198, 175, 173, 174, 194, 74, 97, 90, 45, 86, 179, 29, 143, 253, 205, 80, 233, 53, 232, 2, 221, 169, 57, 81, 218, 100, 179, 148, 36, 137, 3, 107, 126, 101, 248, 217, 115, 72, 219, 189, 177, 73, 72, 176, 40, 243, 24, 35, 207, 189, 2, 119, 159, 128, 22, 215, 115, 212, 214, 114, 120, 227, 10, 142, 110, 93, 141, 209, 219, 23, 55, 175, 182, 123, 141, 215, 85, 85, 176, 242, 163, 151, 249, 252, 161, 171, 56, 186, 101, 213, 105, 67, 253, 134, 48, 79, 124, 39, 70, 225, 61, 50, 20, 77, 15, 144, 185, 237, 40, 85, 135, 139, 41, 94, 117, 2, 219, 201, 200, 78, 224, 200, 96, 146, 159, 25, 143, 87, 180, 235, 70, 179, 0, 42, 179, 43, 229, 68, 191, 117, 89, 32, 129, 103, 96, 16, 183, 124, 252, 53, 51, 238, 123, 8, 77, 39, 197, 91, 20, 218, 79, 121, 78, 38, 153, 219, 54, 218, 181, 214, 61, 48, 152, 152, 201, 211, 29, 108, 81, 251, 241, 8, 10, 193, 63, 110, 32, 71, 86, 44, 109, 123, 178, 36, 49, 120, 214, 28, 180, 29, 175, 32, 233, 88, 4, 224, 253, 82, 233, 78, 96, 33, 2, 189, 167, 142, 170, 155, 137, 26, 158, 132, 206, 104, 228, 224, 170, 21, 212, 149, 214, 97, 58, 80, 66, 216, 228, 136, 46, 81, 242, 51, 6, 27, 137, 152, 22, 69, 101, 118, 5, 85, 57, 149, 152, 75, 205, 20, 110, 45, 192, 224, 167, 199, 24, 234, 220, 114, 163, 138, 66, 51, 166, 172, 218, 86, 51, 252, 237, 69, 227, 161, 67, 235, 103, 160, 38, 173, 89, 103, 59, 73, 226, 232, 214, 85, 4, 70, 15, 160, 202, 84, 212, 100, 92, 163, 211, 51, 112, 194, 217, 118, 223, 243, 116, 145, 130, 197, 175, 60, 196, 162, 121, 247, 178, 127, 245, 47, 108, 254, 225, 35, 18, 38, 157, 131, 87, 96, 235, 157, 28, 109, 86, 11, 91, 22, 253, 143, 79, 30, 152, 205, 129, 53, 75, 177, 89, 204, 173, 206, 3, 48, 68, 122, 225, 55, 185, 31, 158, 195, 130, 209, 56, 80, 91, 194, 85, 176, 213, 88, 40, 89, 147, 78, 229, 129, 66, 16, 37, 212, 122, 53, 131, 111, 25, 198, 144, 91, 134, 186, 140, 243, 218, 42, 18, 164, 253, 118, 130, 29, 255, 217, 74, 77, 161, 92, 142, 57, 108, 230, 133, 220, 183, 232, 55, 162, 134, 39, 57, 217, 184, 190, 135, 165, 166, 134, 212, 63, 126, 177, 107, 173, 103, 112, 24, 253, 207, 178, 59, 187, 222, 33, 120, 71, 68, 113, 98, 205, 159, 212, 152, 74, 218, 156, 27, 127, 230, 121, 184, 7, 116, 184, 203, 101, 81, 187, 28, 128, 79, 36, 201, 48, 243, 255, 230, 190, 7, 60, 77, 31, 169, 227, 239, 74, 226, 146, 39, 34, 90, 173, 28, 221, 184, 142, 154, 162, 26, 202, 142, 149, 17, 54, 49, 188, 75, 66, 154, 106, 189, 154, 240, 73, 17, 232, 60, 117, 20, 165, 20, 33, 154, 109, 148, 236, 45, 166, 38, 175, 26, 175, 120, 239, 238, 47, 155, 146, 160, 44, 183, 142, 242, 2, 11, 93, 17, 171, 214, 122, 27, 80, 27, 181, 13, 89, 226, 141, 105, 178, 249, 159, 164, 44, 63, 155, 41, 215, 253, 195, 238, 163, 145, 83, 173, 74, 223, 179, 153, 133, 115, 239, 104, 136, 70, 216, 172, 22, 172, 22, 51, 67, 167, 207, 106, 50, 79, 18, 69, 118, 46, 253, 154, 79, 254, 49, 155, 29, 191, 124, 137, 185, 250, 212, 217, 192, 134, 72, 47, 252, 167, 197, 224, 153, 24, 136, 218, 189, 119, 86, 214, 86, 31, 51, 81, 252, 215, 9, 44, 37, 114, 162, 159, 79, 127, 95, 198, 61, 51, 65, 46, 159, 117, 221, 136, 63, 213, 249, 213, 108, 255, 207, 22, 210, 151, 159, 64, 178, 73, 232, 141, 238, 92, 249, 202, 219, 92, 54, 239, 85, 244, 238, 74, 109, 191, 51, 208, 26, 220, 216, 243, 221, 103, 118, 173, 53, 6, 4, 145, 112, 238, 172, 182, 39, 118, 49, 133, 169, 7, 41, 74, 109, 59, 143, 33, 110, 234, 217, 120, 135, 71, 118, 244, 242, 185, 109, 186, 252, 213, 31, 16, 0, 0, 32, 0, 73, 68, 65, 84, 211, 31, 86, 74, 193, 181, 101, 44, 3, 20, 23, 214, 129, 92, 244, 228, 60, 170, 203, 74, 89, 243, 225, 59, 20, 238, 202, 103, 215, 127, 119, 144, 244, 224, 232, 174, 57, 215, 20, 32, 250, 252, 88, 252, 6, 251, 179, 107, 254, 118, 42, 50, 43, 40, 222, 83, 68, 217, 209, 50, 162, 47, 137, 37, 32, 41, 208, 241, 247, 60, 5, 165, 185, 117, 84, 22, 181, 221, 130, 182, 51, 184, 15, 240, 71, 80, 11, 148, 172, 207, 60, 237, 217, 57, 128, 41, 55, 131, 244, 189, 155, 137, 30, 62, 222, 97, 247, 151, 68, 145, 31, 95, 248, 71, 139, 164, 61, 173, 190, 105, 243, 202, 3, 107, 150, 178, 236, 141, 39, 201, 73, 221, 123, 218, 235, 233, 130, 220, 241, 30, 21, 134, 62, 184, 247, 110, 36, 182, 42, 51, 166, 141, 89, 212, 102, 203, 142, 155, 160, 22, 232, 63, 103, 32, 253, 103, 15, 116, 237, 68, 191, 147, 79, 253, 135, 62, 223, 143, 181, 86, 62, 174, 137, 77, 158, 192, 13, 239, 126, 170, 116, 240, 115, 50, 70, 255, 0, 116, 238, 30, 118, 149, 216, 89, 107, 170, 218, 158, 212, 13, 232, 140, 237, 75, 20, 87, 107, 237, 122, 32, 168, 58, 173, 3, 240, 94, 145, 20, 110, 181, 178, 18, 232, 121, 5, 211, 61, 128, 43, 94, 126, 147, 154, 242, 50, 182, 126, 247, 37, 185, 27, 179, 217, 41, 73, 140, 252, 231, 232, 46, 107, 236, 227, 21, 237, 205, 228, 215, 166, 146, 250, 237, 97, 142, 45, 62, 130, 181, 202, 194, 209, 47, 15, 83, 188, 171, 144, 216, 57, 241, 104, 189, 186, 54, 156, 92, 145, 111, 238, 242, 205, 191, 30, 99, 156, 31, 130, 78, 77, 201, 154, 244, 22, 137, 129, 205, 57, 188, 225, 15, 135, 58, 0, 155, 23, 125, 76, 230, 254, 166, 13, 65, 212, 26, 45, 147, 175, 190, 7, 73, 146, 56, 184, 110, 57, 127, 126, 240, 34, 105, 187, 79, 127, 62, 169, 245, 49, 224, 53, 42, 20, 183, 72, 215, 62, 243, 238, 20, 18, 84, 30, 42, 162, 108, 103, 142, 220, 196, 7, 240, 138, 241, 102, 248, 189, 35, 93, 190, 188, 175, 58, 191, 154, 61, 239, 236, 164, 56, 69, 142, 44, 105, 13, 110, 92, 244, 212, 60, 166, 223, 249, 143, 6, 17, 42, 5, 231, 98, 240, 246, 177, 203, 1, 48, 87, 187, 134, 3, 80, 93, 82, 220, 174, 121, 122, 15, 59, 74, 22, 5, 178, 78, 249, 46, 93, 80, 34, 69, 73, 106, 86, 1, 3, 58, 126, 101, 133, 246, 32, 8, 2, 195, 207, 189, 136, 220, 131, 251, 201, 75, 61, 72, 101, 86, 5, 165, 71, 75, 9, 29, 23, 214, 101, 78, 128, 160, 86, 17, 48, 60, 144, 160, 164, 16, 76, 169, 37, 152, 203, 234, 168, 45, 172, 161, 112, 75, 62, 42, 189, 26, 247, 8, 143, 46, 169, 20, 168, 41, 183, 82, 154, 211, 118, 70, 171, 35, 209, 122, 27, 112, 139, 240, 162, 54, 171, 188, 97, 115, 105, 13, 73, 20, 25, 123, 201, 141, 14, 185, 103, 117, 89, 9, 159, 252, 227, 50, 44, 205, 36, 121, 39, 95, 115, 47, 130, 160, 226, 171, 71, 175, 103, 237, 23, 175, 83, 154, 151, 121, 202, 107, 168, 221, 117, 248, 36, 135, 227, 59, 62, 2, 173, 119, 239, 235, 237, 80, 143, 165, 180, 150, 226, 85, 39, 168, 74, 45, 6, 81, 66, 165, 85, 51, 240, 170, 65, 140, 184, 63, 9, 131, 191, 91, 219, 23, 112, 18, 146, 36, 145, 190, 92, 62, 235, 175, 202, 145, 55, 151, 216, 228, 9, 220, 183, 232, 55, 134, 205, 188, 192, 17, 173, 90, 21, 28, 196, 145, 63, 126, 165, 170, 232, 244, 189, 49, 90, 67, 165, 209, 49, 226, 202, 27, 186, 192, 162, 246, 35, 73, 18, 27, 223, 254, 79, 187, 28, 152, 228, 219, 31, 64, 163, 239, 112, 34, 240, 207, 173, 58, 0, 111, 155, 164, 126, 130, 138, 117, 64, 76, 71, 175, 168, 208, 49, 4, 149, 138, 145, 179, 102, 99, 202, 206, 36, 43, 101, 55, 213, 121, 85, 148, 28, 40, 38, 116, 124, 88, 151, 246, 48, 55, 248, 25, 232, 55, 163, 31, 26, 189, 134, 146, 3, 197, 216, 204, 54, 74, 15, 154, 48, 237, 43, 193, 24, 234, 142, 222, 215, 113, 89, 229, 54, 171, 72, 209, 241, 26, 135, 38, 252, 181, 23, 181, 155, 22, 247, 56, 63, 44, 166, 90, 172, 21, 173, 59, 32, 229, 69, 121, 156, 121, 203, 35, 14, 121, 106, 91, 252, 202, 67, 28, 223, 190, 182, 201, 152, 206, 205, 157, 194, 140, 99, 108, 95, 242, 5, 149, 37, 45, 85, 11, 235, 81, 25, 52, 120, 141, 12, 197, 111, 114, 20, 186, 0, 35, 184, 88, 201, 166, 163, 144, 172, 34, 21, 123, 243, 49, 173, 77, 111, 200, 240, 247, 77, 240, 35, 249, 169, 241, 132, 140, 11, 115, 233, 242, 190, 138, 140, 114, 182, 191, 180, 149, 140, 63, 210, 16, 173, 34, 90, 131, 129, 89, 79, 206, 227, 218, 55, 63, 196, 51, 160, 251, 142, 210, 20, 218, 199, 241, 213, 127, 80, 158, 211, 70, 87, 204, 86, 16, 4, 72, 186, 246, 86, 187, 239, 155, 181, 109, 35, 251, 126, 250, 150, 200, 177, 19, 237, 190, 70, 198, 230, 245, 236, 95, 252, 93, 155, 243, 140, 126, 254, 140, 186, 225, 142, 142, 223, 64, 226, 227, 22, 71, 0, 159, 231, 73, 238, 85, 2, 139, 1, 165, 88, 181, 155, 80, 169, 213, 92, 247, 214, 71, 104, 13, 6, 214, 126, 252, 30, 37, 7, 138, 217, 244, 204, 6, 146, 159, 154, 128, 174, 11, 195, 242, 130, 90, 69, 220, 165, 253, 9, 76, 10, 102, 223, 135, 123, 40, 57, 80, 76, 117, 78, 21, 251, 223, 222, 75, 192, 168, 32, 250, 93, 24, 141, 214, 179, 243, 247, 47, 203, 53, 35, 218, 95, 230, 223, 105, 84, 6, 13, 1, 51, 98, 169, 220, 95, 72, 217, 174, 220, 22, 71, 2, 54, 139, 153, 194, 244, 35, 132, 196, 15, 238, 212, 125, 178, 15, 238, 98, 243, 247, 31, 182, 24, 55, 215, 84, 97, 62, 205, 153, 162, 160, 85, 225, 153, 24, 132, 199, 144, 192, 46, 117, 250, 92, 129, 234, 99, 37, 148, 109, 207, 197, 86, 35, 31, 5, 105, 140, 90, 18, 174, 29, 68, 191, 153, 49, 46, 167, 81, 209, 24, 155, 217, 198, 145, 239, 15, 115, 236, 167, 163, 72, 54, 249, 253, 19, 63, 110, 18, 215, 190, 253, 49, 193, 241, 74, 144, 212, 85, 209, 123, 218, 119, 124, 102, 169, 173, 65, 146, 36, 187, 222, 147, 71, 254, 92, 198, 170, 23, 159, 68, 180, 90, 137, 26, 55, 153, 136, 209, 227, 236, 186, 255, 166, 247, 94, 107, 215, 220, 160, 65, 195, 58, 124, 125, 0, 73, 197, 161, 38, 223, 54, 146, 36, 9, 27, 109, 124, 35, 192, 52, 187, 174, 168, 96, 55, 130, 32, 48, 228, 236, 243, 176, 212, 213, 114, 108, 243, 6, 234, 74, 106, 41, 216, 153, 71, 200, 216, 48, 52, 198, 174, 45, 125, 210, 251, 232, 137, 156, 222, 15, 247, 48, 119, 74, 83, 77, 88, 107, 172, 84, 231, 84, 81, 176, 57, 15, 65, 16, 228, 99, 1, 59, 19, 177, 44, 181, 34, 101, 217, 221, 27, 250, 63, 21, 186, 32, 119, 140, 177, 190, 212, 102, 149, 35, 214, 53, 245, 72, 226, 199, 78, 35, 36, 46, 209, 238, 107, 75, 146, 196, 23, 15, 95, 77, 73, 118, 90, 187, 215, 8, 26, 21, 30, 131, 2, 8, 152, 26, 131, 33, 210, 171, 215, 232, 245, 183, 134, 165, 164, 134, 226, 85, 105, 84, 30, 44, 146, 29, 48, 1, 34, 167, 71, 49, 230, 177, 113, 4, 12, 13, 116, 233, 205, 191, 104, 111, 33, 219, 230, 109, 34, 127, 107, 30, 72, 18, 110, 222, 62, 92, 254, 210, 235, 92, 241, 234, 219, 120, 248, 183, 212, 144, 80, 112, 29, 178, 119, 110, 165, 240, 208, 254, 142, 47, 148, 36, 134, 95, 121, 3, 106, 109, 199, 30, 128, 82, 22, 125, 205, 154, 87, 159, 109, 208, 27, 201, 220, 186, 129, 184, 105, 231, 160, 247, 240, 108, 247, 53, 108, 22, 11, 127, 62, 247, 47, 114, 247, 182, 175, 155, 225, 144, 75, 175, 34, 56, 113, 104, 135, 236, 4, 108, 106, 145, 135, 155, 56, 0, 161, 15, 204, 125, 30, 1, 59, 98, 9, 10, 142, 64, 16, 4, 6, 77, 61, 11, 181, 70, 203, 225, 181, 127, 97, 46, 51, 147, 179, 33, 11, 191, 196, 0, 12, 126, 93, 124, 22, 44, 128, 87, 63, 111, 162, 206, 142, 70, 178, 137, 148, 30, 41, 69, 180, 136, 148, 165, 150, 82, 184, 45, 31, 181, 65, 141, 49, 204, 189, 195, 95, 212, 21, 249, 102, 204, 53, 78, 136, 253, 159, 2, 149, 78, 141, 218, 77, 75, 77, 90, 211, 110, 92, 161, 253, 135, 16, 55, 122, 138, 221, 215, 221, 190, 228, 11, 214, 126, 249, 102, 251, 108, 208, 107, 240, 28, 122, 178, 45, 111, 148, 79, 175, 232, 206, 119, 42, 108, 85, 102, 74, 183, 102, 83, 186, 57, 171, 161, 95, 131, 119, 156, 15, 163, 31, 25, 75, 244, 121, 177, 14, 87, 194, 116, 36, 117, 101, 117, 236, 251, 96, 15, 7, 62, 221, 135, 165, 66, 182, 61, 233, 226, 57, 220, 243, 221, 175, 12, 152, 60, 213, 165, 157, 22, 5, 153, 252, 253, 123, 236, 110, 11, 60, 236, 178, 171, 58, 36, 174, 179, 245, 227, 183, 217, 178, 224, 245, 38, 149, 71, 214, 218, 26, 210, 214, 175, 162, 223, 248, 41, 24, 188, 218, 142, 70, 84, 230, 231, 240, 251, 147, 255, 71, 230, 214, 246, 53, 35, 18, 4, 129, 41, 15, 62, 213, 33, 7, 227, 36, 219, 239, 244, 21, 222, 110, 248, 244, 189, 103, 146, 166, 74, 2, 79, 116, 244, 42, 10, 142, 231, 220, 135, 158, 192, 232, 227, 203, 194, 127, 221, 79, 109, 113, 45, 155, 158, 92, 199, 136, 251, 71, 17, 58, 161, 117, 17, 25, 71, 162, 113, 211, 48, 232, 134, 33, 68, 158, 217, 143, 3, 159, 238, 167, 96, 71, 30, 230, 50, 51, 199, 23, 30, 37, 119, 117, 54, 145, 231, 246, 195, 111, 88, 64, 187, 107, 178, 107, 202, 59, 214, 216, 162, 59, 208, 248, 180, 116, 166, 242, 143, 29, 176, 251, 122, 181, 149, 101, 252, 250, 223, 199, 218, 156, 167, 11, 48, 98, 140, 243, 195, 189, 191, 95, 175, 222, 244, 1, 68, 179, 141, 138, 189, 249, 242, 19, 255, 201, 144, 185, 206, 83, 199, 192, 107, 18, 137, 58, 187, 159, 75, 111, 158, 146, 40, 145, 254, 219, 9, 14, 127, 125, 176, 161, 115, 159, 95, 68, 20, 87, 190, 250, 14, 67, 103, 94, 224, 100, 235, 20, 58, 130, 190, 29, 155, 238, 169, 48, 87, 85, 99, 108, 71, 107, 59, 73, 20, 89, 59, 255, 121, 14, 44, 249, 161, 213, 215, 203, 115, 179, 249, 254, 214, 43, 24, 119, 251, 253, 36, 156, 127, 105, 171, 201, 122, 101, 89, 233, 236, 95, 252, 61, 7, 150, 124, 143, 165, 166, 186, 149, 171, 180, 78, 248, 232, 241, 120, 6, 135, 182, 123, 126, 61, 130, 196, 31, 112, 82, 202, 119, 126, 166, 228, 134, 192, 7, 184, 180, 212, 70, 223, 226, 140, 91, 239, 38, 160, 95, 12, 31, 221, 124, 37, 181, 21, 229, 236, 120, 117, 43, 3, 175, 26, 68, 255, 217, 3, 187, 229, 183, 228, 17, 225, 201, 216, 39, 199, 81, 114, 160, 152, 131, 95, 236, 199, 116, 168, 132, 154, 130, 26, 82, 63, 59, 132, 71, 148, 7, 225, 51, 162, 240, 77, 244, 59, 173, 45, 162, 77, 194, 102, 233, 188, 204, 175, 163, 209, 120, 233, 65, 37, 128, 248, 183, 109, 5, 105, 169, 118, 95, 239, 183, 119, 158, 165, 162, 40, 175, 245, 123, 121, 232, 48, 198, 249, 98, 140, 245, 69, 211, 139, 51, 250, 235, 145, 108, 34, 149, 7, 139, 168, 216, 155, 143, 104, 150, 195, 160, 42, 173, 154, 152, 243, 99, 137, 191, 108, 0, 90, 15, 215, 22, 48, 50, 29, 42, 33, 229, 131, 61, 148, 159, 144, 213, 36, 85, 26, 13, 211, 110, 191, 143, 11, 159, 120, 14, 125, 23, 118, 135, 83, 232, 26, 236, 205, 1, 0, 249, 28, 190, 45, 108, 102, 51, 127, 62, 255, 40, 199, 215, 172, 56, 253, 181, 170, 43, 89, 247, 250, 11, 108, 253, 248, 109, 66, 135, 141, 194, 51, 36, 4, 149, 70, 71, 69, 94, 14, 197, 199, 82, 41, 203, 74, 183, 203, 198, 97, 179, 175, 181, 107, 157, 216, 216, 1, 48, 122, 114, 143, 4, 138, 106, 133, 139, 49, 120, 198, 185, 60, 178, 98, 35, 239, 94, 121, 17, 69, 105, 199, 57, 252, 181, 92, 42, 56, 252, 222, 145, 221, 150, 44, 230, 151, 232, 207, 196, 23, 167, 144, 183, 53, 151, 195, 95, 29, 160, 34, 163, 130, 202, 140, 74, 14, 127, 124, 0, 99, 168, 59, 97, 103, 70, 224, 63, 34, 160, 213, 204, 109, 209, 230, 122, 155, 63, 128, 160, 18, 80, 27, 181, 13, 25, 232, 0, 229, 5, 57, 118, 93, 43, 247, 200, 62, 214, 127, 253, 78, 139, 113, 99, 140, 15, 158, 195, 67, 228, 230, 60, 125, 0, 201, 38, 82, 149, 90, 76, 69, 74, 1, 182, 106, 249, 169, 89, 16, 4, 34, 166, 69, 50, 224, 170, 65, 184, 5, 184, 110, 89, 31, 200, 225, 254, 67, 159, 239, 39, 115, 85, 134, 172, 239, 12, 244, 159, 48, 229, 255, 219, 187, 243, 248, 184, 202, 122, 143, 227, 159, 231, 76, 50, 147, 125, 223, 155, 166, 109, 186, 210, 210, 214, 110, 64, 75, 129, 182, 32, 5, 5, 81, 46, 112, 5, 197, 123, 21, 161, 236, 234, 85, 47, 224, 21, 197, 237, 162, 160, 160, 2, 197, 94, 228, 162, 34, 42, 40, 114, 89, 101, 95, 74, 129, 150, 165, 165, 44, 221, 215, 52, 109, 246, 125, 155, 100, 102, 206, 115, 255, 72, 154, 54, 54, 109, 150, 38, 153, 73, 250, 125, 191, 94, 121, 205, 244, 204, 89, 126, 133, 102, 206, 115, 158, 229, 247, 227, 95, 127, 126, 55, 163, 250, 62, 190, 42, 17, 194, 151, 152, 212, 239, 99, 123, 234, 164, 10, 52, 55, 242, 143, 239, 124, 157, 189, 107, 215, 244, 250, 156, 173, 13, 245, 236, 122, 227, 149, 126, 199, 116, 176, 156, 233, 179, 24, 51, 255, 148, 254, 28, 90, 239, 73, 101, 53, 64, 212, 35, 214, 122, 171, 234, 248, 218, 128, 68, 36, 3, 46, 119, 202, 52, 110, 120, 105, 13, 43, 190, 120, 62, 219, 222, 122, 157, 189, 43, 139, 105, 42, 109, 98, 206, 183, 79, 24, 210, 47, 213, 156, 19, 114, 201, 158, 155, 67, 241, 171, 123, 216, 246, 183, 205, 52, 149, 52, 209, 92, 210, 196, 182, 63, 110, 166, 248, 31, 187, 201, 93, 146, 79, 230, 188, 172, 46, 249, 11, 156, 8, 206, 224, 230, 137, 141, 234, 210, 0, 104, 172, 169, 192, 13, 5, 113, 60, 125, 27, 147, 126, 236, 214, 175, 29, 82, 176, 199, 155, 30, 71, 218, 169, 99, 143, 137, 254, 52, 27, 112, 105, 220, 92, 73, 227, 199, 229, 132, 90, 14, 252, 119, 200, 154, 147, 205, 113, 151, 78, 35, 113, 76, 255, 191, 128, 135, 130, 27, 116, 217, 245, 204, 14, 182, 62, 178, 185, 179, 187, 63, 57, 39, 151, 243, 127, 112, 27, 243, 46, 250, 66, 68, 15, 85, 72, 207, 122, 51, 238, 126, 56, 78, 212, 225, 123, 171, 90, 106, 170, 121, 250, 219, 87, 82, 209, 139, 52, 189, 131, 193, 120, 60, 44, 188, 254, 198, 126, 30, 204, 223, 151, 25, 19, 0, 136, 170, 174, 97, 41, 14, 249, 3, 25, 156, 12, 172, 132, 244, 12, 190, 246, 248, 139, 252, 233, 235, 203, 120, 235, 79, 191, 163, 118, 75, 13, 175, 255, 199, 43, 204, 188, 110, 54, 217, 243, 134, 174, 46, 147, 113, 12, 163, 151, 20, 144, 191, 104, 52, 37, 111, 238, 99, 219, 163, 91, 168, 223, 85, 135, 191, 202, 207, 206, 191, 110, 163, 248, 185, 34, 114, 22, 228, 146, 53, 63, 155, 232, 68, 47, 142, 199, 16, 229, 51, 4, 91, 35, 175, 39, 192, 19, 211, 245, 151, 219, 186, 46, 141, 213, 21, 36, 101, 246, 126, 60, 109, 237, 211, 127, 102, 219, 219, 175, 118, 221, 104, 32, 101, 126, 254, 136, 191, 249, 187, 129, 16, 141, 27, 43, 105, 220, 80, 129, 235, 63, 112, 227, 207, 152, 153, 201, 164, 139, 166, 144, 54, 181, 23, 131, 167, 97, 86, 186, 166, 132, 141, 191, 255, 152, 166, 146, 246, 68, 43, 78, 84, 20, 139, 46, 191, 150, 115, 191, 243, 3, 98, 142, 226, 201, 81, 34, 71, 116, 108, 255, 31, 146, 98, 146, 186, 255, 55, 80, 95, 178, 151, 167, 190, 185, 172, 223, 221, 246, 3, 97, 246, 23, 191, 74, 230, 228, 254, 173, 90, 114, 13, 157, 107, 149, 163, 172, 67, 120, 75, 30, 73, 175, 68, 121, 189, 124, 105, 249, 3, 228, 79, 159, 201, 99, 223, 191, 129, 182, 134, 54, 222, 185, 117, 53, 133, 231, 78, 96, 202, 165, 83, 7, 45, 115, 96, 119, 140, 99, 200, 91, 56, 138, 188, 147, 71, 81, 190, 182, 140, 109, 143, 110, 161, 122, 99, 21, 129, 250, 54, 246, 60, 187, 155, 226, 23, 138, 72, 155, 145, 65, 206, 194, 92, 226, 82, 98, 168, 47, 59, 124, 101, 187, 112, 113, 186, 89, 90, 89, 95, 81, 210, 235, 6, 64, 115, 93, 53, 143, 223, 246, 205, 67, 182, 199, 79, 76, 111, 79, 226, 51, 66, 5, 27, 90, 105, 220, 88, 73, 243, 214, 106, 220, 64, 199, 82, 74, 3, 89, 179, 115, 152, 120, 209, 36, 82, 39, 165, 133, 55, 192, 94, 168, 219, 81, 199, 134, 7, 62, 164, 234, 163, 3, 197, 161, 166, 44, 58, 131, 11, 111, 189, 147, 188, 227, 142, 15, 99, 100, 50, 208, 130, 173, 253, 251, 238, 241, 198, 197, 19, 151, 222, 125, 117, 189, 218, 162, 157, 97, 189, 249, 231, 207, 61, 137, 121, 95, 190, 186, 191, 135, 127, 116, 77, 146, 233, 204, 65, 30, 101, 96, 126, 228, 61, 159, 201, 225, 44, 185, 234, 235, 20, 158, 176, 128, 251, 191, 242, 121, 42, 119, 239, 100, 199, 19, 219, 168, 222, 88, 197, 236, 111, 206, 35, 46, 123, 136, 111, 60, 166, 189, 171, 55, 107, 78, 54, 213, 27, 170, 216, 249, 212, 118, 74, 223, 46, 193, 134, 44, 85, 235, 42, 168, 90, 87, 65, 92, 110, 60, 49, 133, 105, 196, 142, 77, 197, 68, 71, 206, 204, 247, 127, 238, 1, 0, 104, 168, 44, 235, 245, 241, 143, 223, 246, 77, 26, 170, 186, 238, 239, 120, 61, 36, 207, 238, 251, 140, 220, 225, 160, 181, 164, 145, 198, 13, 21, 180, 20, 215, 117, 142, 145, 99, 32, 231, 196, 92, 38, 94, 56, 57, 226, 243, 246, 3, 248, 171, 252, 108, 254, 243, 70, 138, 95, 46, 194, 118, 44, 213, 202, 153, 52, 133, 243, 127, 120, 187, 102, 247, 143, 80, 253, 45, 234, 147, 58, 110, 252, 97, 63, 43, 56, 113, 33, 163, 231, 205, 103, 207, 59, 111, 245, 55, 172, 126, 203, 152, 120, 28, 75, 127, 120, 71, 191, 211, 77, 155, 246, 201, 254, 157, 162, 44, 140, 26, 144, 200, 100, 200, 140, 157, 115, 2, 223, 89, 185, 150, 7, 175, 251, 42, 235, 158, 120, 148, 218, 173, 53, 188, 254, 205, 87, 152, 113, 205, 44, 114, 231, 15, 254, 82, 193, 238, 164, 77, 77, 39, 109, 106, 58, 254, 170, 22, 118, 63, 183, 139, 162, 23, 118, 211, 90, 235, 167, 185, 99, 174, 128, 89, 179, 151, 216, 49, 201, 196, 21, 166, 17, 147, 151, 16, 246, 52, 183, 158, 238, 122, 0, 170, 186, 159, 201, 255, 207, 182, 189, 253, 42, 239, 62, 241, 224, 33, 219, 147, 231, 230, 225, 68, 240, 186, 246, 190, 114, 3, 33, 90, 118, 214, 210, 184, 177, 130, 64, 141, 191, 115, 187, 199, 231, 97, 212, 105, 163, 25, 247, 233, 241, 36, 22, 244, 121, 253, 241, 144, 11, 52, 5, 216, 254, 247, 173, 236, 124, 122, 59, 161, 142, 4, 80, 241, 105, 233, 124, 250, 134, 239, 115, 234, 87, 174, 236, 111, 37, 53, 25, 6, 42, 182, 108, 234, 215, 113, 163, 102, 159, 120, 196, 207, 79, 190, 238, 63, 249, 235, 87, 63, 79, 168, 109, 232, 146, 156, 101, 78, 57, 158, 115, 110, 91, 142, 183, 239, 107, 254, 1, 48, 80, 219, 226, 210, 229, 139, 43, 10, 136, 252, 166, 187, 28, 34, 54, 57, 133, 43, 254, 240, 55, 94, 187, 239, 30, 30, 189, 249, 91, 4, 154, 252, 188, 119, 219, 219, 228, 47, 46, 96, 218, 87, 166, 135, 109, 185, 85, 76, 122, 44, 147, 47, 57, 142, 137, 23, 77, 166, 116, 245, 62, 118, 61, 179, 147, 234, 141, 85, 216, 160, 75, 243, 246, 26, 154, 183, 215, 224, 137, 141, 38, 182, 48, 133, 248, 241, 105, 68, 167, 133, 103, 118, 184, 19, 219, 77, 15, 64, 69, 73, 143, 199, 249, 27, 235, 248, 203, 205, 95, 237, 124, 130, 220, 207, 151, 147, 64, 252, 164, 200, 31, 247, 238, 145, 133, 214, 210, 70, 154, 182, 85, 211, 178, 187, 182, 75, 218, 228, 216, 140, 88, 198, 156, 61, 142, 130, 79, 142, 197, 59, 0, 41, 162, 7, 91, 168, 45, 196, 174, 103, 118, 176, 237, 209, 173, 4, 58, 38, 124, 70, 121, 189, 156, 122, 217, 213, 124, 250, 134, 239, 17, 151, 146, 26, 230, 8, 101, 176, 237, 123, 255, 157, 126, 29, 55, 122, 222, 130, 35, 126, 158, 58, 118, 2, 167, 125, 235, 102, 94, 254, 239, 239, 246, 235, 252, 125, 53, 126, 209, 153, 44, 190, 233, 71, 68, 199, 246, 191, 151, 215, 133, 219, 191, 145, 106, 186, 100, 64, 139, 2, 234, 129, 17, 240, 205, 117, 108, 58, 237, 242, 107, 40, 60, 113, 1, 191, 253, 202, 231, 41, 223, 182, 133, 226, 87, 138, 168, 120, 191, 156, 233, 87, 206, 36, 231, 132, 240, 117, 71, 59, 81, 14, 121, 11, 243, 201, 91, 152, 79, 67, 81, 3, 123, 95, 219, 195, 222, 149, 123, 104, 169, 108, 33, 212, 18, 160, 241, 227, 10, 26, 63, 174, 32, 58, 37, 134, 216, 49, 41, 196, 20, 36, 225, 77, 31, 186, 33, 12, 79, 55, 13, 128, 250, 195, 172, 229, 63, 216, 195, 223, 187, 156, 234, 189, 59, 187, 108, 51, 30, 135, 212, 5, 163, 7, 44, 182, 112, 8, 54, 180, 209, 188, 173, 154, 230, 237, 213, 4, 27, 187, 142, 155, 166, 77, 73, 103, 236, 57, 133, 228, 158, 148, 215, 239, 148, 208, 67, 201, 186, 150, 226, 87, 138, 216, 242, 151, 77, 180, 84, 182, 175, 229, 54, 142, 195, 188, 11, 46, 230, 220, 255, 250, 17, 25, 99, 84, 227, 44, 146, 5, 90, 154, 121, 243, 238, 219, 41, 92, 116, 38, 163, 231, 245, 191, 76, 119, 125, 201, 94, 246, 188, 115, 228, 146, 219, 221, 137, 75, 207, 36, 119, 198, 236, 30, 247, 155, 124, 214, 121, 84, 109, 223, 202, 250, 135, 127, 223, 159, 240, 122, 197, 27, 159, 192, 41, 95, 255, 14, 147, 150, 158, 123, 180, 167, 42, 77, 240, 243, 171, 127, 222, 24, 101, 12, 165, 214, 170, 1, 48, 156, 141, 158, 49, 139, 239, 174, 122, 159, 167, 126, 250, 3, 94, 188, 235, 231, 180, 214, 248, 121, 247, 214, 53, 228, 46, 24, 197, 244, 101, 51, 240, 38, 13, 92, 101, 191, 254, 72, 44, 72, 100, 202, 165, 83, 153, 252, 197, 227, 168, 217, 84, 77, 201, 27, 251, 216, 187, 114, 15, 109, 13, 109, 4, 106, 253, 4, 106, 75, 169, 95, 95, 138, 19, 19, 69, 204, 168, 36, 226, 198, 166, 224, 203, 75, 28, 212, 155, 141, 39, 238, 208, 6, 64, 213, 158, 237, 71, 60, 102, 213, 159, 151, 179, 254, 249, 71, 15, 217, 158, 52, 51, 187, 61, 185, 208, 48, 19, 108, 108, 195, 95, 84, 71, 243, 174, 90, 218, 202, 187, 142, 149, 250, 82, 99, 200, 91, 48, 138, 209, 103, 20, 144, 52, 182, 255, 75, 169, 134, 146, 181, 150, 210, 183, 246, 181, 231, 203, 216, 123, 160, 132, 234, 148, 69, 103, 112, 254, 15, 111, 99, 244, 140, 89, 97, 140, 78, 122, 163, 124, 227, 135, 188, 248, 163, 155, 168, 43, 222, 205, 182, 87, 158, 227, 252, 229, 127, 36, 117, 108, 97, 191, 206, 181, 250, 55, 119, 98, 221, 190, 167, 33, 159, 120, 250, 217, 189, 30, 99, 95, 112, 205, 183, 240, 37, 38, 241, 206, 253, 119, 31, 210, 43, 120, 52, 140, 49, 140, 93, 184, 132, 147, 175, 191, 161, 95, 153, 254, 14, 61, 33, 63, 249, 82, 142, 57, 100, 66, 132, 185, 183, 214, 222, 7, 244, 191, 238, 161, 68, 148, 109, 171, 87, 241, 224, 181, 151, 81, 190, 173, 61, 179, 157, 47, 217, 199, 241, 87, 204, 32, 119, 65, 100, 77, 245, 8, 181, 133, 40, 123, 187, 148, 146, 55, 247, 82, 190, 174, 140, 144, 191, 107, 113, 30, 39, 218, 67, 76, 126, 34, 190, 188, 68, 124, 57, 137, 68, 13, 66, 151, 243, 222, 135, 62, 192, 6, 14, 124, 65, 36, 103, 141, 226, 251, 47, 23, 117, 187, 239, 238, 245, 171, 185, 231, 203, 75, 8, 254, 211, 152, 159, 55, 43, 158, 172, 179, 38, 180, 103, 22, 28, 6, 2, 53, 126, 90, 118, 215, 210, 178, 187, 142, 64, 77, 215, 76, 103, 78, 180, 67, 246, 188, 92, 70, 47, 41, 32, 115, 86, 86, 68, 151, 229, 61, 152, 117, 45, 251, 86, 21, 179, 245, 145, 205, 93, 110, 252, 227, 230, 158, 200, 103, 111, 249, 41, 147, 22, 46, 10, 95, 112, 210, 43, 214, 117, 89, 247, 208, 253, 188, 243, 192, 114, 220, 224, 129, 101, 165, 73, 121, 249, 156, 119, 215, 239, 72, 200, 204, 238, 211, 249, 222, 255, 203, 239, 120, 107, 121, 239, 170, 233, 29, 204, 56, 14, 23, 255, 241, 73, 146, 243, 251, 86, 12, 119, 199, 202, 23, 89, 245, 203, 91, 105, 170, 60, 124, 185, 239, 222, 94, 127, 252, 226, 165, 204, 190, 244, 114, 210, 11, 7, 44, 55, 223, 214, 244, 100, 142, 191, 200, 152, 67, 150, 68, 152, 229, 117, 246, 18, 99, 121, 104, 160, 174, 36, 225, 23, 240, 183, 240, 228, 79, 190, 199, 75, 203, 239, 236, 172, 74, 149, 115, 98, 46, 83, 191, 60, 125, 232, 87, 10, 244, 130, 27, 8, 81, 249, 65, 5, 165, 107, 74, 41, 123, 167, 148, 214, 90, 255, 33, 251, 120, 18, 188, 196, 228, 38, 224, 203, 73, 192, 151, 155, 216, 237, 19, 124, 95, 149, 63, 179, 245, 144, 39, 223, 27, 159, 252, 152, 172, 113, 83, 186, 108, 43, 219, 190, 129, 187, 255, 125, 49, 77, 53, 149, 93, 182, 59, 49, 81, 100, 159, 59, 25, 79, 124, 228, 78, 34, 115, 91, 67, 180, 150, 54, 224, 223, 215, 72, 235, 190, 6, 130, 13, 93, 27, 48, 78, 148, 67, 250, 244, 12, 114, 231, 231, 145, 59, 63, 143, 232, 132, 200, 31, 219, 223, 239, 112, 55, 254, 130, 153, 179, 249, 244, 141, 223, 103, 250, 89, 231, 42, 145, 207, 48, 208, 80, 86, 194, 203, 63, 190, 137, 125, 135, 41, 218, 147, 148, 151, 207, 217, 183, 222, 69, 218, 184, 9, 61, 158, 171, 181, 161, 158, 53, 255, 243, 43, 62, 126, 252, 145, 126, 197, 50, 126, 241, 82, 206, 252, 193, 207, 251, 117, 108, 91, 115, 19, 107, 31, 188, 143, 77, 207, 60, 70, 75, 77, 117, 175, 143, 51, 142, 67, 246, 212, 233, 140, 59, 229, 116, 10, 23, 157, 73, 82, 238, 128, 62, 172, 185, 214, 101, 209, 213, 105, 230, 245, 110, 175, 189, 98, 159, 141, 179, 113, 236, 181, 154, 12, 56, 226, 236, 124, 103, 53, 15, 94, 119, 25, 37, 155, 218, 11, 221, 56, 209, 30, 198, 127, 118, 2, 19, 206, 159, 132, 39, 38, 50, 235, 206, 91, 107, 169, 221, 90, 67, 217, 154, 82, 202, 215, 150, 81, 191, 251, 160, 101, 103, 7, 137, 74, 246, 225, 203, 138, 39, 58, 61, 14, 111, 70, 28, 209, 105, 177, 125, 126, 98, 173, 95, 87, 66, 253, 250, 174, 75, 249, 78, 249, 194, 117, 124, 238, 166, 95, 118, 254, 121, 215, 250, 183, 184, 255, 218, 207, 30, 114, 243, 7, 200, 56, 163, 144, 152, 252, 200, 74, 24, 99, 67, 46, 109, 229, 77, 237, 55, 252, 146, 6, 218, 170, 90, 186, 84, 39, 3, 240, 120, 61, 100, 206, 202, 34, 231, 164, 60, 178, 231, 229, 16, 29, 193, 13, 152, 238, 184, 1, 151, 226, 87, 247, 176, 253, 177, 173, 157, 73, 124, 160, 125, 40, 236, 156, 155, 110, 209, 141, 127, 24, 217, 254, 234, 243, 188, 118, 251, 15, 104, 109, 168, 63, 226, 126, 81, 190, 24, 102, 93, 242, 21, 166, 95, 240, 133, 67, 210, 251, 90, 215, 165, 98, 203, 70, 182, 189, 248, 12, 155, 159, 123, 2, 127, 93, 237, 97, 206, 114, 100, 78, 84, 20, 255, 250, 187, 199, 72, 41, 24, 219, 175, 227, 59, 227, 9, 133, 40, 122, 251, 77, 138, 214, 188, 78, 237, 238, 29, 84, 239, 218, 73, 115, 85, 121, 103, 220, 222, 248, 4, 82, 199, 77, 32, 189, 112, 2, 233, 227, 38, 50, 106, 206, 9, 135, 205, 57, 112, 180, 172, 229, 174, 171, 83, 205, 245, 135, 251, 220, 0, 220, 91, 103, 127, 137, 85, 58, 224, 145, 40, 216, 218, 202, 115, 119, 254, 148, 231, 126, 249, 179, 206, 226, 22, 49, 233, 177, 28, 247, 111, 211, 24, 181, 48, 242, 51, 214, 181, 213, 183, 81, 245, 81, 37, 149, 31, 86, 80, 245, 97, 37, 141, 123, 27, 186, 221, 207, 56, 134, 232, 180, 216, 246, 198, 64, 122, 28, 222, 140, 88, 162, 146, 99, 142, 216, 40, 8, 84, 183, 80, 246, 196, 230, 46, 219, 28, 199, 195, 25, 87, 220, 68, 225, 156, 83, 217, 176, 242, 105, 222, 248, 243, 114, 66, 193, 192, 33, 199, 38, 205, 200, 38, 41, 220, 107, 254, 173, 37, 80, 215, 74, 91, 69, 51, 129, 170, 246, 159, 182, 170, 22, 172, 123, 104, 139, 41, 54, 51, 142, 204, 153, 153, 100, 206, 202, 38, 107, 118, 118, 196, 54, 0, 143, 36, 232, 15, 82, 244, 220, 46, 118, 60, 177, 13, 127, 245, 129, 94, 162, 252, 233, 159, 224, 156, 155, 110, 97, 198, 217, 159, 209, 141, 127, 152, 8, 182, 250, 121, 227, 215, 63, 99, 195, 147, 221, 87, 208, 59, 28, 227, 241, 144, 53, 101, 26, 241, 25, 217, 24, 99, 104, 170, 42, 167, 122, 251, 86, 218, 154, 251, 183, 222, 255, 96, 51, 47, 250, 18, 11, 174, 253, 246, 81, 159, 39, 130, 236, 136, 247, 51, 163, 187, 177, 255, 253, 218, 27, 0, 13, 54, 139, 16, 219, 128, 200, 95, 212, 43, 253, 82, 93, 92, 196, 223, 111, 254, 54, 239, 61, 118, 160, 107, 44, 237, 184, 116, 166, 125, 117, 250, 176, 72, 226, 178, 159, 191, 218, 79, 213, 135, 21, 84, 125, 92, 73, 237, 214, 26, 26, 246, 52, 96, 15, 87, 116, 200, 49, 68, 37, 120, 137, 74, 246, 17, 157, 28, 67, 84, 146, 143, 168, 228, 24, 162, 147, 125, 157, 235, 245, 43, 254, 177, 149, 214, 178, 190, 125, 121, 196, 141, 79, 35, 237, 148, 190, 141, 17, 30, 45, 215, 31, 192, 109, 106, 35, 212, 216, 74, 168, 190, 149, 214, 138, 22, 90, 74, 155, 58, 43, 238, 253, 179, 232, 4, 47, 25, 211, 51, 200, 152, 153, 73, 198, 140, 44, 226, 115, 123, 95, 215, 60, 210, 180, 213, 183, 178, 243, 233, 29, 236, 122, 102, 103, 231, 114, 62, 128, 194, 19, 23, 176, 244, 27, 55, 50, 125, 233, 57, 186, 241, 15, 35, 85, 59, 182, 242, 194, 45, 223, 166, 102, 215, 145, 39, 221, 14, 165, 212, 49, 227, 184, 224, 183, 143, 16, 229, 27, 49, 197, 187, 66, 214, 114, 250, 213, 169, 230, 181, 35, 237, 212, 249, 91, 115, 111, 173, 189, 17, 184, 117, 208, 195, 146, 176, 218, 250, 198, 107, 60, 114, 195, 215, 40, 254, 104, 61, 208, 81, 177, 109, 73, 1, 147, 254, 117, 50, 177, 153, 145, 55, 63, 160, 39, 161, 214, 16, 117, 59, 107, 169, 219, 86, 75, 237, 182, 246, 215, 166, 125, 141, 61, 206, 200, 117, 124, 30, 60, 241, 94, 28, 175, 167, 189, 1, 208, 203, 25, 188, 113, 133, 169, 164, 46, 44, 24, 208, 9, 114, 142, 99, 32, 24, 196, 109, 13, 64, 107, 136, 144, 63, 64, 168, 57, 72, 160, 206, 79, 107, 165, 159, 214, 170, 150, 67, 38, 73, 118, 97, 32, 33, 47, 145, 148, 73, 41, 164, 76, 72, 37, 117, 114, 26, 73, 133, 201, 195, 254, 166, 216, 80, 84, 207, 206, 167, 182, 83, 252, 106, 113, 103, 218, 97, 99, 12, 83, 79, 95, 202, 210, 111, 220, 200, 196, 147, 79, 11, 115, 132, 210, 87, 27, 158, 248, 27, 171, 126, 253, 211, 33, 77, 160, 211, 19, 143, 215, 199, 249, 247, 62, 72, 198, 196, 227, 194, 29, 202, 128, 49, 240, 237, 43, 83, 76, 143, 147, 25, 58, 191, 33, 110, 177, 54, 42, 187, 142, 149, 64, 255, 23, 94, 202, 176, 224, 134, 66, 172, 250, 253, 125, 60, 249, 147, 155, 105, 172, 106, 31, 219, 118, 162, 28, 70, 159, 49, 134, 137, 23, 76, 34, 38, 61, 178, 75, 183, 246, 36, 216, 28, 164, 126, 87, 29, 141, 197, 13, 52, 238, 109, 236, 124, 109, 169, 104, 238, 182, 123, 188, 183, 162, 18, 188, 237, 115, 13, 188, 30, 156, 104, 15, 142, 215, 131, 137, 114, 192, 180, 39, 54, 220, 127, 191, 53, 198, 96, 60, 6, 199, 235, 96, 131, 110, 123, 50, 157, 80, 251, 171, 219, 22, 194, 134, 92, 220, 128, 75, 200, 31, 34, 216, 208, 70, 160, 49, 208, 235, 184, 140, 199, 16, 151, 21, 79, 226, 152, 68, 82, 198, 167, 146, 50, 41, 149, 148, 9, 41, 68, 13, 192, 164, 200, 136, 96, 161, 124, 93, 25, 59, 159, 216, 78, 197, 7, 229, 157, 243, 63, 28, 143, 135, 217, 231, 93, 192, 210, 111, 220, 72, 254, 244, 79, 132, 55, 70, 233, 151, 55, 239, 190, 157, 245, 143, 252, 33, 220, 97, 116, 97, 28, 135, 79, 222, 114, 59, 227, 23, 157, 25, 238, 80, 6, 210, 95, 174, 76, 230, 18, 99, 76, 143, 95, 42, 93, 30, 17, 150, 215, 218, 66, 3, 107, 129, 225, 177, 240, 87, 142, 74, 115, 109, 13, 255, 184, 253, 199, 188, 118, 255, 189, 157, 243, 3, 156, 104, 15, 99, 206, 28, 203, 132, 11, 38, 226, 27, 97, 181, 236, 221, 128, 75, 227, 222, 70, 154, 246, 53, 208, 180, 175, 9, 127, 117, 11, 45, 149, 45, 248, 171, 253, 248, 171, 253, 237, 171, 15, 34, 160, 48, 134, 49, 6, 111, 138, 143, 184, 204, 56, 18, 70, 39, 144, 144, 151, 72, 124, 94, 2, 9, 163, 19, 136, 207, 137, 199, 120, 34, 167, 166, 194, 64, 9, 54, 7, 41, 126, 173, 136, 93, 79, 239, 236, 50, 207, 35, 38, 49, 137, 147, 47, 189, 140, 69, 203, 174, 83, 2, 159, 97, 110, 205, 255, 252, 138, 181, 127, 252, 109, 184, 195, 232, 100, 28, 135, 83, 255, 227, 187, 76, 253, 204, 133, 225, 14, 101, 32, 125, 16, 239, 103, 193, 145, 198, 253, 15, 118, 72, 31, 225, 111, 106, 237, 39, 45, 60, 13, 140, 144, 71, 10, 233, 73, 93, 233, 62, 158, 189, 227, 86, 86, 253, 254, 62, 130, 173, 237, 93, 115, 30, 159, 135, 49, 103, 141, 99, 194, 249, 19, 195, 158, 72, 104, 168, 216, 144, 139, 191, 166, 21, 127, 101, 11, 254, 26, 63, 129, 142, 167, 243, 64, 115, 128, 96, 115, 128, 64, 83, 176, 253, 181, 227, 207, 193, 150, 246, 110, 233, 96, 115, 224, 8, 67, 14, 134, 232, 248, 104, 60, 62, 15, 142, 215, 233, 124, 239, 241, 122, 136, 142, 247, 226, 75, 241, 225, 75, 245, 17, 147, 22, 123, 224, 53, 217, 55, 44, 50, 238, 13, 132, 250, 93, 117, 236, 126, 118, 39, 123, 95, 43, 38, 120, 112, 89, 225, 177, 133, 44, 94, 118, 29, 11, 190, 248, 21, 149, 230, 29, 33, 172, 235, 178, 250, 55, 119, 178, 254, 225, 223, 15, 104, 210, 156, 254, 112, 162, 162, 88, 244, 159, 183, 48, 249, 172, 243, 194, 26, 199, 0, 43, 181, 112, 242, 213, 41, 102, 71, 111, 15, 232, 246, 91, 102, 121, 141, 253, 178, 49, 220, 127, 184, 207, 101, 100, 170, 217, 187, 135, 103, 110, 255, 49, 111, 61, 244, 0, 161, 64, 251, 204, 247, 168, 152, 40, 242, 79, 47, 160, 240, 156, 241, 196, 229, 12, 223, 137, 100, 18, 57, 220, 64, 136, 146, 55, 247, 177, 235, 217, 157, 212, 108, 58, 176, 94, 218, 24, 195, 196, 147, 79, 99, 241, 149, 215, 51, 227, 236, 207, 224, 120, 134, 223, 74, 5, 233, 217, 142, 215, 94, 224, 229, 91, 191, 71, 160, 185, 177, 231, 157, 7, 65, 108, 106, 26, 103, 254, 240, 14, 242, 102, 206, 9, 203, 245, 7, 73, 181, 227, 178, 104, 89, 154, 249, 176, 47, 7, 29, 246, 6, 127, 111, 173, 189, 28, 248, 13, 48, 242, 250, 27, 229, 136, 170, 139, 139, 248, 199, 207, 127, 194, 155, 127, 252, 223, 206, 172, 92, 198, 24, 178, 230, 100, 51, 238, 156, 241, 100, 204, 204, 12, 115, 132, 50, 28, 53, 20, 53, 80, 252, 90, 17, 123, 94, 220, 77, 91, 253, 129, 217, 252, 49, 137, 73, 204, 253, 151, 207, 179, 232, 138, 107, 25, 53, 117, 122, 24, 35, 148, 161, 82, 187, 103, 55, 207, 125, 247, 235, 84, 239, 220, 54, 164, 215, 29, 183, 112, 49, 167, 126, 243, 123, 196, 165, 103, 12, 233, 117, 7, 147, 133, 38, 227, 176, 244, 170, 36, 243, 70, 95, 143, 61, 226, 19, 254, 242, 90, 123, 133, 129, 123, 81, 35, 224, 152, 84, 190, 125, 43, 47, 221, 115, 7, 171, 255, 252, 7, 218, 90, 154, 59, 183, 39, 79, 72, 161, 240, 220, 9, 228, 157, 156, 55, 34, 199, 163, 101, 224, 180, 214, 182, 178, 119, 229, 30, 246, 188, 180, 155, 134, 162, 174, 57, 28, 70, 207, 152, 197, 169, 151, 93, 197, 188, 11, 47, 193, 23, 167, 222, 165, 99, 77, 160, 165, 153, 87, 126, 122, 51, 219, 95, 121, 126, 208, 175, 149, 60, 106, 52, 39, 93, 249, 13, 10, 79, 251, 228, 160, 95, 107, 136, 181, 88, 203, 217, 61, 45, 247, 59, 156, 30, 187, 248, 127, 83, 99, 207, 183, 134, 7, 129, 225, 183, 70, 76, 6, 68, 83, 117, 21, 175, 63, 176, 130, 87, 239, 187, 155, 186, 210, 3, 37, 115, 99, 210, 99, 25, 251, 169, 113, 140, 94, 92, 128, 47, 117, 100, 77, 24, 148, 254, 11, 250, 131, 148, 189, 93, 194, 222, 149, 197, 84, 188, 95, 129, 13, 29, 168, 183, 16, 151, 146, 202, 188, 11, 46, 102, 254, 37, 255, 206, 152, 217, 243, 194, 24, 165, 68, 2, 107, 45, 27, 159, 122, 148, 183, 127, 123, 87, 159, 210, 231, 246, 86, 74, 193, 88, 102, 94, 244, 111, 76, 62, 251, 60, 60, 209, 35, 110, 90, 91, 131, 129, 207, 93, 153, 98, 94, 234, 239, 9, 122, 53, 198, 127, 79, 157, 157, 231, 184, 60, 142, 33, 204, 169, 207, 36, 156, 130, 109, 109, 124, 240, 204, 227, 188, 120, 247, 47, 216, 249, 238, 154, 206, 237, 198, 24, 210, 103, 100, 144, 127, 90, 1, 185, 39, 231, 225, 241, 106, 236, 246, 88, 99, 93, 75, 213, 135, 149, 20, 191, 90, 68, 233, 234, 146, 46, 19, 250, 140, 227, 48, 249, 212, 37, 156, 248, 249, 75, 153, 125, 222, 5, 120, 143, 162, 166, 185, 140, 76, 109, 77, 141, 172, 253, 227, 111, 217, 248, 228, 223, 240, 215, 215, 29, 213, 185, 188, 113, 241, 140, 61, 101, 9, 147, 151, 126, 134, 81, 179, 79, 232, 117, 101, 191, 97, 166, 220, 194, 167, 174, 78, 49, 221, 23, 80, 232, 165, 94, 79, 242, 235, 200, 22, 248, 16, 112, 198, 209, 92, 80, 70, 134, 45, 175, 191, 194, 203, 247, 254, 138, 15, 159, 127, 186, 75, 245, 46, 111, 146, 151, 81, 167, 142, 102, 244, 233, 195, 167, 140, 172, 244, 143, 27, 8, 81, 177, 174, 130, 146, 213, 251, 40, 123, 167, 180, 75, 150, 62, 128, 81, 211, 102, 48, 239, 130, 139, 57, 225, 162, 47, 144, 58, 106, 116, 152, 162, 148, 225, 36, 20, 8, 80, 180, 122, 37, 91, 95, 120, 134, 178, 13, 31, 208, 88, 94, 218, 227, 49, 209, 177, 113, 100, 78, 158, 74, 246, 180, 153, 228, 207, 157, 79, 238, 140, 217, 35, 241, 105, 255, 96, 187, 60, 14, 75, 175, 72, 50, 91, 142, 246, 68, 125, 154, 229, 223, 145, 44, 232, 71, 192, 13, 125, 61, 86, 70, 166, 186, 210, 18, 214, 60, 252, 32, 111, 61, 244, 0, 165, 91, 54, 117, 249, 44, 185, 48, 153, 209, 75, 198, 144, 51, 63, 143, 152, 52, 13, 17, 140, 4, 193, 230, 32, 229, 107, 75, 41, 121, 171, 132, 138, 181, 101, 93, 158, 244, 161, 125, 249, 222, 188, 11, 46, 102, 222, 5, 23, 147, 59, 101, 90, 152, 162, 148, 145, 194, 95, 87, 67, 197, 150, 77, 180, 53, 54, 208, 214, 212, 128, 117, 45, 214, 13, 17, 159, 153, 69, 124, 70, 54, 241, 25, 89, 35, 106, 66, 95, 79, 12, 188, 111, 162, 249, 212, 178, 120, 83, 210, 243, 222, 189, 58, 95, 223, 173, 168, 182, 11, 93, 135, 223, 1, 227, 7, 34, 8, 25, 25, 118, 188, 253, 22, 111, 61, 244, 0, 239, 254, 253, 97, 252, 7, 85, 247, 50, 198, 144, 50, 49, 149, 156, 249, 185, 228, 158, 148, 167, 229, 132, 195, 76, 211, 190, 70, 202, 222, 45, 163, 124, 109, 41, 213, 31, 87, 225, 6, 221, 46, 159, 167, 23, 140, 101, 214, 185, 231, 51, 251, 179, 23, 50, 118, 238, 137, 195, 62, 5, 177, 72, 36, 178, 240, 112, 130, 159, 203, 122, 155, 228, 167, 55, 250, 253, 155, 250, 235, 42, 155, 228, 245, 112, 187, 133, 43, 6, 42, 24, 25, 25, 218, 90, 154, 89, 247, 248, 163, 188, 245, 167, 223, 177, 101, 213, 171, 88, 183, 235, 13, 35, 105, 108, 18, 57, 39, 230, 145, 115, 82, 174, 134, 9, 34, 80, 200, 31, 162, 122, 99, 37, 229, 239, 149, 83, 190, 182, 148, 166, 146, 67, 191, 111, 114, 38, 31, 199, 172, 115, 207, 231, 19, 231, 158, 79, 193, 204, 217, 97, 136, 82, 228, 152, 17, 178, 240, 95, 87, 37, 115, 91, 111, 210, 251, 246, 197, 81, 55, 213, 239, 173, 177, 139, 48, 220, 5, 28, 63, 0, 241, 200, 8, 83, 95, 86, 202, 250, 103, 30, 103, 221, 147, 127, 103, 203, 235, 175, 116, 38, 24, 218, 47, 46, 39, 158, 236, 185, 57, 100, 76, 207, 36, 253, 248, 244, 145, 147, 211, 126, 24, 177, 33, 151, 154, 205, 53, 84, 126, 88, 65, 229, 7, 21, 212, 108, 174, 233, 50, 115, 31, 218, 51, 167, 77, 56, 105, 33, 199, 159, 249, 41, 166, 159, 117, 46, 57, 147, 166, 132, 41, 90, 145, 99, 74, 149, 129, 139, 175, 76, 49, 47, 12, 198, 201, 7, 164, 175, 110, 133, 181, 209, 110, 45, 215, 99, 248, 30, 160, 188, 157, 210, 173, 150, 186, 90, 62, 124, 246, 41, 222, 127, 234, 49, 62, 126, 241, 217, 46, 185, 5, 0, 140, 99, 72, 30, 159, 66, 198, 140, 76, 50, 166, 103, 146, 122, 92, 154, 86, 20, 12, 130, 160, 63, 72, 237, 230, 26, 170, 55, 85, 81, 189, 177, 138, 154, 205, 213, 221, 86, 27, 76, 202, 202, 102, 218, 25, 103, 115, 252, 153, 159, 226, 184, 37, 103, 18, 155, 164, 222, 26, 145, 33, 180, 210, 113, 185, 116, 89, 154, 41, 26, 172, 11, 12, 232, 96, 221, 138, 38, 155, 235, 182, 241, 51, 12, 95, 28, 232, 115, 203, 200, 210, 214, 210, 204, 134, 151, 158, 99, 253, 211, 143, 179, 241, 149, 23, 168, 43, 221, 119, 200, 62, 78, 180, 67, 234, 228, 180, 206, 198, 64, 114, 97, 10, 209, 241, 234, 33, 232, 11, 107, 45, 77, 123, 27, 59, 75, 37, 87, 111, 170, 162, 126, 87, 29, 54, 116, 104, 79, 98, 76, 66, 34, 19, 23, 158, 198, 148, 211, 78, 103, 242, 169, 75, 200, 155, 58, 93, 227, 249, 34, 67, 47, 136, 229, 39, 233, 41, 252, 232, 34, 99, 142, 80, 7, 252, 232, 13, 202, 111, 247, 61, 117, 118, 158, 99, 185, 27, 56, 97, 48, 206, 47, 35, 79, 233, 150, 77, 108, 94, 249, 50, 155, 87, 190, 204, 150, 85, 175, 210, 84, 93, 117, 232, 78, 6, 226, 115, 18, 72, 30, 159, 66, 202, 132, 20, 146, 199, 167, 144, 92, 152, 66, 84, 92, 212, 208, 7, 28, 129, 108, 200, 165, 97, 79, 35, 13, 69, 117, 212, 237, 168, 163, 110, 91, 45, 117, 59, 106, 9, 182, 4, 187, 221, 63, 54, 41, 153, 241, 39, 157, 204, 248, 19, 79, 102, 210, 41, 139, 25, 59, 123, 30, 78, 148, 254, 91, 138, 132, 209, 70, 7, 190, 184, 44, 197, 172, 29, 138, 139, 13, 90, 243, 254, 17, 107, 61, 85, 245, 92, 130, 229, 102, 96, 226, 96, 93, 71, 70, 30, 235, 186, 20, 127, 180, 190, 179, 65, 176, 245, 141, 149, 180, 54, 117, 95, 56, 196, 24, 67, 124, 94, 60, 201, 227, 83, 73, 26, 155, 68, 124, 110, 2, 241, 121, 241, 196, 231, 198, 227, 68, 143, 204, 225, 3, 55, 224, 210, 84, 210, 72, 211, 190, 38, 26, 247, 53, 208, 176, 187, 158, 250, 221, 245, 52, 22, 55, 30, 50, 118, 127, 176, 140, 177, 133, 20, 158, 48, 159, 194, 19, 22, 48, 97, 193, 41, 228, 77, 153, 54, 82, 147, 164, 136, 12, 55, 33, 12, 119, 183, 212, 115, 211, 127, 140, 54, 45, 67, 117, 209, 65, 239, 223, 187, 197, 90, 39, 171, 142, 127, 1, 126, 108, 96, 210, 96, 95, 79, 70, 30, 55, 20, 162, 108, 235, 217, 22, 246, 107, 0, 0, 10, 81, 73, 68, 65, 84, 38, 118, 175, 123, 151, 162, 247, 223, 99, 247, 186, 119, 217, 243, 193, 251, 4, 252, 71, 248, 61, 49, 16, 155, 17, 219, 222, 32, 232, 108, 20, 36, 16, 159, 27, 79, 76, 90, 76, 68, 79, 54, 180, 174, 197, 95, 237, 167, 165, 162, 185, 227, 167, 133, 150, 138, 102, 154, 75, 155, 105, 220, 215, 136, 191, 178, 165, 199, 114, 170, 105, 249, 5, 20, 124, 98, 14, 99, 102, 205, 101, 204, 172, 185, 20, 204, 154, 75, 124, 106, 218, 16, 253, 13, 68, 164, 15, 214, 59, 134, 101, 203, 146, 205, 154, 158, 119, 29, 88, 67, 54, 192, 183, 191, 33, 96, 224, 191, 129, 9, 67, 117, 93, 25, 153, 220, 96, 144, 146, 205, 27, 186, 52, 10, 74, 183, 108, 234, 146, 127, 224, 72, 156, 104, 7, 95, 178, 15, 95, 74, 12, 222, 100, 31, 190, 100, 47, 190, 148, 24, 124, 41, 62, 188, 201, 94, 188, 73, 62, 60, 62, 15, 30, 175, 7, 39, 218, 131, 199, 231, 180, 191, 122, 61, 56, 94, 231, 136, 147, 19, 3, 77, 1, 232, 184, 63, 7, 154, 3, 96, 45, 216, 246, 237, 129, 198, 64, 251, 107, 83, 27, 129, 166, 0, 109, 13, 109, 180, 214, 180, 210, 86, 215, 138, 191, 198, 79, 107, 109, 43, 109, 245, 173, 221, 142, 209, 119, 39, 54, 41, 153, 220, 41, 83, 25, 53, 109, 6, 163, 166, 78, 39, 111, 218, 116, 70, 77, 157, 78, 92, 74, 106, 175, 142, 23, 145, 176, 105, 182, 240, 195, 242, 100, 126, 113, 139, 49, 221, 143, 211, 13, 178, 33, 159, 225, 179, 194, 218, 104, 183, 142, 127, 7, 254, 11, 24, 51, 212, 215, 151, 145, 173, 190, 188, 140, 242, 237, 91, 40, 223, 190, 149, 138, 29, 219, 40, 223, 190, 181, 243, 125, 107, 243, 128, 229, 207, 0, 232, 108, 24, 0, 4, 26, 3, 61, 236, 221, 127, 241, 169, 105, 164, 230, 23, 144, 57, 182, 144, 204, 241, 19, 201, 26, 63, 145, 236, 9, 147, 200, 26, 63, 137, 164, 172, 236, 65, 187, 174, 136, 12, 154, 199, 220, 16, 95, 187, 38, 221, 236, 9, 103, 16, 97, 155, 226, 251, 136, 181, 158, 234, 26, 62, 101, 13, 55, 98, 88, 16, 174, 56, 228, 216, 81, 91, 178, 151, 202, 93, 59, 168, 47, 47, 163, 190, 172, 148, 198, 170, 138, 206, 247, 13, 85, 21, 52, 116, 188, 31, 232, 134, 66, 119, 140, 49, 196, 38, 37, 147, 144, 153, 69, 98, 70, 38, 137, 25, 89, 36, 231, 230, 145, 152, 158, 73, 82, 118, 14, 169, 121, 249, 164, 143, 25, 75, 218, 232, 49, 248, 226, 19, 6, 61, 30, 17, 25, 18, 107, 93, 203, 183, 174, 73, 53, 175, 132, 59, 16, 136, 144, 165, 122, 203, 107, 237, 28, 99, 249, 26, 134, 139, 1, 77, 67, 150, 176, 106, 107, 105, 166, 161, 162, 28, 55, 24, 164, 165, 190, 14, 107, 45, 45, 117, 181, 157, 175, 0, 205, 181, 53, 157, 175, 190, 132, 68, 60, 29, 179, 231, 125, 241, 9, 157, 133, 72, 188, 177, 113, 68, 249, 124, 0, 68, 199, 196, 226, 139, 143, 39, 54, 57, 133, 184, 148, 84, 173, 169, 23, 57, 182, 236, 181, 240, 195, 140, 100, 238, 31, 236, 165, 125, 125, 17, 17, 13, 128, 253, 238, 173, 177, 99, 49, 92, 11, 124, 21, 208, 55, 164, 136, 136, 12, 103, 13, 192, 207, 156, 102, 238, 92, 150, 103, 154, 123, 220, 123, 136, 69, 84, 3, 96, 191, 251, 43, 108, 98, 155, 151, 47, 99, 185, 14, 77, 24, 20, 17, 145, 225, 165, 1, 184, 43, 228, 112, 199, 181, 73, 166, 155, 164, 38, 145, 33, 34, 27, 0, 7, 91, 94, 107, 231, 56, 112, 133, 11, 95, 48, 160, 50, 114, 34, 34, 18, 145, 44, 52, 25, 203, 253, 81, 209, 252, 247, 229, 9, 166, 44, 220, 241, 244, 36, 226, 27, 0, 251, 45, 175, 181, 169, 198, 240, 5, 99, 185, 204, 194, 39, 194, 29, 143, 136, 136, 72, 135, 122, 96, 185, 227, 240, 139, 101, 73, 166, 50, 220, 193, 244, 214, 176, 105, 0, 28, 236, 238, 42, 59, 45, 202, 195, 165, 182, 125, 174, 64, 122, 184, 227, 17, 17, 145, 99, 82, 41, 150, 21, 214, 240, 171, 171, 83, 76, 77, 184, 131, 233, 171, 97, 217, 0, 216, 239, 142, 61, 54, 54, 46, 129, 11, 173, 225, 223, 128, 211, 128, 145, 153, 251, 85, 68, 68, 34, 201, 7, 198, 242, 139, 180, 20, 254, 114, 145, 49, 109, 225, 14, 166, 191, 134, 117, 3, 224, 96, 43, 154, 108, 174, 109, 227, 34, 107, 248, 60, 112, 34, 35, 232, 239, 38, 34, 34, 97, 231, 90, 120, 222, 194, 157, 87, 39, 243, 130, 49, 166, 119, 233, 58, 35, 216, 136, 188, 73, 174, 168, 182, 5, 214, 240, 57, 11, 23, 98, 56, 57, 220, 241, 136, 136, 200, 240, 100, 160, 214, 90, 254, 208, 209, 205, 191, 35, 220, 241, 12, 164, 17, 217, 0, 56, 216, 111, 26, 236, 84, 27, 228, 34, 107, 184, 88, 197, 136, 68, 68, 164, 151, 222, 179, 240, 63, 254, 6, 30, 28, 202, 10, 125, 67, 105, 196, 55, 0, 14, 118, 111, 157, 61, 193, 88, 206, 183, 112, 30, 48, 37, 220, 241, 136, 136, 72, 68, 41, 195, 240, 39, 66, 252, 239, 85, 105, 230, 163, 112, 7, 51, 216, 142, 169, 6, 192, 193, 150, 215, 218, 66, 199, 114, 110, 199, 48, 193, 124, 64, 133, 209, 69, 68, 142, 61, 109, 88, 158, 183, 134, 63, 120, 146, 249, 191, 101, 198, 12, 94, 101, 175, 8, 115, 204, 54, 0, 14, 118, 87, 133, 205, 139, 142, 230, 51, 22, 62, 11, 44, 6, 188, 225, 142, 73, 68, 68, 6, 213, 219, 24, 30, 12, 25, 254, 28, 201, 217, 250, 6, 147, 26, 0, 255, 228, 215, 85, 54, 201, 235, 225, 108, 23, 62, 231, 192, 82, 11, 41, 225, 142, 73, 68, 68, 142, 158, 129, 247, 129, 71, 92, 120, 120, 164, 77, 232, 235, 15, 53, 0, 142, 224, 17, 107, 61, 149, 117, 124, 2, 56, 195, 88, 206, 213, 80, 129, 136, 200, 48, 99, 217, 0, 252, 213, 241, 240, 151, 101, 73, 102, 83, 184, 195, 137, 36, 106, 0, 244, 193, 125, 141, 54, 59, 20, 96, 169, 53, 156, 5, 156, 137, 178, 16, 138, 136, 68, 26, 11, 188, 7, 252, 95, 40, 196, 255, 93, 155, 110, 62, 14, 119, 64, 145, 74, 13, 128, 126, 186, 197, 90, 39, 171, 142, 89, 192, 25, 6, 206, 0, 22, 1, 81, 225, 141, 74, 68, 228, 152, 20, 194, 178, 218, 192, 95, 61, 46, 143, 94, 158, 110, 138, 195, 29, 208, 112, 160, 6, 192, 0, 89, 81, 111, 51, 92, 151, 37, 22, 22, 155, 246, 137, 132, 147, 195, 29, 147, 136, 200, 8, 86, 110, 44, 207, 99, 120, 202, 184, 60, 187, 44, 205, 212, 133, 59, 160, 225, 70, 13, 128, 65, 114, 95, 163, 205, 14, 4, 57, 213, 129, 51, 92, 248, 164, 129, 113, 225, 142, 73, 68, 100, 24, 11, 1, 239, 27, 120, 209, 184, 60, 85, 146, 202, 155, 183, 24, 227, 134, 59, 168, 225, 76, 13, 128, 33, 178, 188, 214, 22, 58, 176, 216, 194, 18, 44, 139, 49, 228, 134, 59, 38, 17, 145, 72, 102, 97, 167, 129, 151, 12, 60, 215, 98, 121, 241, 27, 169, 166, 54, 220, 49, 141, 36, 106, 0, 132, 201, 61, 13, 246, 56, 143, 203, 98, 107, 89, 8, 156, 2, 228, 135, 59, 38, 17, 145, 176, 178, 148, 0, 47, 239, 255, 185, 42, 213, 236, 10, 111, 64, 35, 155, 26, 0, 17, 98, 69, 181, 45, 112, 61, 156, 130, 101, 1, 237, 13, 130, 105, 104, 201, 161, 136, 140, 108, 101, 192, 42, 107, 120, 197, 58, 188, 124, 77, 162, 217, 24, 238, 128, 142, 37, 106, 0, 68, 168, 123, 202, 109, 130, 241, 114, 146, 177, 44, 236, 168, 104, 120, 50, 16, 27, 238, 184, 68, 68, 142, 194, 14, 44, 111, 88, 195, 42, 55, 196, 27, 215, 164, 177, 97, 36, 148, 213, 29, 174, 212, 0, 24, 38, 86, 88, 27, 109, 27, 152, 107, 67, 44, 192, 112, 18, 48, 23, 24, 27, 230, 176, 68, 68, 14, 167, 5, 203, 90, 235, 176, 198, 184, 172, 34, 138, 55, 174, 74, 52, 229, 225, 14, 74, 14, 80, 3, 96, 24, 91, 81, 109, 147, 129, 233, 214, 48, 199, 194, 28, 96, 14, 134, 169, 225, 142, 75, 68, 142, 65, 150, 18, 99, 88, 133, 229, 13, 99, 121, 175, 53, 149, 119, 174, 55, 166, 53, 220, 97, 201, 225, 169, 1, 48, 194, 220, 215, 104, 179, 3, 33, 230, 26, 152, 107, 97, 142, 113, 153, 171, 21, 7, 34, 50, 128, 44, 176, 221, 192, 58, 11, 107, 13, 188, 103, 92, 222, 214, 58, 252, 225, 71, 13, 128, 99, 192, 189, 149, 118, 148, 245, 48, 199, 180, 15, 27, 204, 178, 134, 105, 166, 125, 248, 64, 255, 255, 69, 228, 72, 130, 192, 70, 107, 88, 135, 203, 58, 7, 214, 182, 185, 188, 127, 125, 186, 169, 15, 119, 96, 114, 244, 116, 3, 56, 70, 61, 98, 173, 183, 188, 154, 137, 142, 135, 169, 198, 50, 141, 3, 195, 7, 227, 208, 191, 11, 145, 99, 81, 77, 71, 225, 156, 247, 172, 225, 99, 143, 203, 6, 252, 172, 93, 150, 103, 154, 195, 29, 152, 12, 14, 125, 209, 75, 23, 43, 170, 109, 178, 141, 98, 170, 117, 57, 30, 195, 52, 108, 199, 143, 134, 17, 68, 70, 138, 98, 96, 139, 133, 45, 142, 97, 67, 200, 229, 35, 235, 225, 131, 107, 147, 76, 85, 184, 3, 147, 161, 165, 6, 128, 244, 202, 111, 235, 108, 90, 208, 229, 120, 107, 56, 206, 194, 4, 3, 227, 129, 9, 180, 191, 198, 133, 57, 60, 17, 233, 170, 1, 216, 2, 108, 193, 176, 25, 216, 236, 88, 182, 4, 219, 216, 114, 77, 150, 105, 12, 115, 108, 18, 33, 212, 0, 144, 163, 182, 188, 214, 166, 2, 133, 251, 127, 28, 40, 180, 150, 105, 198, 48, 205, 66, 74, 152, 195, 19, 25, 185, 44, 37, 24, 62, 54, 237, 235, 235, 55, 184, 134, 143, 129, 29, 229, 201, 236, 82, 158, 124, 233, 137, 26, 0, 50, 168, 86, 212, 219, 140, 144, 61, 208, 99, 96, 92, 38, 88, 24, 135, 97, 12, 144, 131, 74, 40, 139, 28, 142, 197, 82, 138, 97, 151, 133, 34, 3, 187, 173, 161, 200, 24, 118, 187, 134, 237, 153, 9, 108, 191, 200, 152, 182, 112, 7, 41, 195, 151, 26, 0, 18, 86, 251, 123, 15, 28, 151, 60, 215, 33, 183, 163, 247, 32, 175, 99, 206, 65, 33, 48, 6, 240, 132, 55, 74, 145, 65, 17, 192, 82, 137, 97, 159, 129, 29, 192, 14, 23, 118, 56, 46, 37, 174, 195, 190, 4, 63, 155, 190, 148, 99, 154, 194, 29, 164, 140, 92, 106, 0, 72, 68, 123, 196, 90, 111, 93, 29, 163, 2, 46, 249, 198, 67, 129, 181, 228, 27, 75, 190, 117, 24, 99, 44, 185, 180, 247, 34, 100, 2, 190, 48, 135, 42, 210, 201, 64, 173, 133, 18, 160, 28, 216, 107, 12, 229, 214, 82, 98, 13, 197, 198, 176, 155, 54, 118, 149, 165, 83, 162, 110, 122, 9, 39, 53, 0, 100, 68, 184, 99, 143, 141, 141, 142, 33, 213, 19, 77, 174, 227, 146, 135, 33, 213, 53, 228, 26, 75, 30, 144, 138, 33, 151, 3, 239, 179, 81, 161, 37, 233, 187, 86, 44, 213, 24, 246, 117, 84, 173, 171, 49, 29, 239, 93, 195, 62, 143, 75, 73, 192, 82, 227, 194, 30, 173, 147, 151, 225, 64, 13, 0, 57, 230, 60, 96, 109, 76, 115, 53, 153, 68, 145, 99, 32, 203, 184, 100, 2, 169, 251, 127, 76, 251, 196, 197, 84, 219, 254, 231, 20, 218, 27, 13, 169, 64, 76, 24, 195, 150, 129, 227, 2, 213, 64, 13, 80, 99, 161, 218, 116, 188, 55, 182, 125, 123, 151, 109, 30, 202, 9, 82, 162, 76, 119, 50, 210, 168, 1, 32, 210, 75, 15, 88, 27, 211, 214, 76, 170, 117, 73, 13, 89, 82, 140, 237, 218, 96, 0, 82, 177, 36, 26, 72, 182, 150, 88, 12, 49, 29, 219, 125, 180, 47, 149, 76, 182, 16, 99, 32, 62, 140, 127, 141, 225, 204, 2, 181, 64, 51, 208, 2, 212, 97, 104, 54, 182, 253, 189, 53, 52, 89, 139, 223, 129, 58, 219, 222, 5, 95, 13, 212, 56, 29, 55, 115, 7, 106, 92, 151, 106, 221, 200, 69, 218, 169, 1, 32, 18, 6, 247, 87, 216, 196, 166, 24, 98, 162, 66, 36, 226, 146, 96, 61, 248, 140, 37, 25, 75, 28, 6, 31, 144, 234, 24, 124, 174, 237, 200, 177, 96, 136, 118, 44, 9, 0, 22, 162, 12, 36, 2, 184, 224, 113, 32, 169, 99, 187, 3, 36, 119, 92, 194, 208, 117, 9, 102, 10, 3, 251, 251, 30, 0, 26, 59, 174, 27, 52, 237, 235, 206, 1, 66, 6, 234, 59, 182, 119, 190, 119, 193, 117, 160, 174, 99, 187, 181, 237, 55, 114, 204, 65, 239, 129, 26, 227, 208, 98, 92, 90, 92, 168, 117, 60, 180, 120, 66, 180, 4, 160, 54, 152, 76, 179, 10, 203, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 136, 200, 49, 232, 255, 1, 187, 79, 208, 14, 222, 178, 69, 133, 0, 0, 0, 0, 73, 69, 78, 68, 174, 66, 96, 130},
		"vue.webapp/public/favicon.ico":   {0, 0, 1, 0, 1, 0, 16, 16, 0, 0, 1, 0, 32, 0, 86, 4, 0, 0, 22, 0, 0, 0, 40, 0, 0, 0, 16, 0, 0, 0, 32, 0, 0, 0, 1, 0, 32, 0, 0, 0, 0, 0, 0, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 19, 18, 11, 17, 65, 61, 38, 60, 111, 103, 65, 102, 147, 134, 86, 133, 167, 153, 98, 153, 176, 161, 103, 164, 173, 160, 102, 165, 162, 153, 96, 157, 144, 135, 85, 139, 115, 109, 68, 110, 75, 71, 44, 69, 30, 28, 17, 23, 0, 0, 0, 0, 0, 0, 0, 0, 59, 56, 35, 55, 171, 161, 101, 164, 237, 222, 140, 230, 255, 248, 157, 255, 242, 237, 146, 255, 193, 214, 124, 255, 168, 198, 111, 255, 187, 210, 121, 255, 230, 232, 141, 255, 255, 250, 158, 255, 255, 255, 161, 255, 255, 253, 159, 255, 255, 244, 154, 255, 216, 204, 130, 236, 171, 161, 101, 162, 37, 35, 22, 30, 225, 212, 133, 217, 255, 255, 168, 255, 255, 248, 157, 255, 159, 189, 106, 255, 66, 140, 59, 255, 47, 100, 36, 255, 95, 112, 59, 255, 38, 94, 29, 255, 48, 118, 42, 255, 111, 170, 84, 255, 210, 216, 130, 255, 255, 244, 155, 255, 253, 238, 150, 255, 171, 163, 110, 255, 181, 172, 116, 255, 164, 156, 102, 210, 37, 35, 22, 34, 100, 90, 58, 86, 104, 130, 71, 185, 58, 146, 58, 255, 48, 135, 49, 255, 146, 145, 85, 255, 255, 210, 143, 255, 205, 175, 114, 255, 116, 120, 64, 255, 70, 108, 47, 255, 81, 158, 71, 255, 170, 201, 112, 255, 255, 247, 157, 255, 205, 192, 126, 255, 149, 142, 99, 255, 179, 170, 111, 228, 0, 0, 0, 0, 0, 0, 0, 0, 14, 44, 16, 114, 42, 108, 41, 206, 55, 80, 39, 172, 245, 204, 137, 254, 255, 216, 145, 255, 255, 217, 146, 255, 255, 216, 146, 255, 255, 209, 142, 255, 182, 168, 106, 243, 53, 137, 56, 231, 144, 196, 103, 255, 232, 227, 140, 255, 230, 222, 139, 255, 82, 78, 49, 96, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 5, 107, 86, 59, 149, 255, 222, 149, 255, 253, 212, 142, 255, 253, 212, 142, 255, 254, 213, 142, 255, 255, 225, 151, 255, 129, 104, 71, 164, 0, 5, 0, 27, 38, 97, 40, 162, 74, 141, 64, 217, 60, 100, 48, 168, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 18, 159, 133, 89, 221, 255, 220, 148, 255, 255, 213, 142, 255, 255, 214, 143, 255, 255, 221, 147, 255, 165, 137, 90, 206, 0, 0, 0, 19, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 166, 139, 93, 199, 255, 221, 146, 255, 245, 206, 140, 255, 194, 172, 130, 255, 147, 139, 118, 255, 148, 130, 100, 221, 0, 0, 0, 19, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 151, 126, 82, 187, 246, 214, 159, 255, 218, 216, 213, 255, 138, 134, 124, 255, 174, 161, 138, 255, 211, 213, 217, 255, 39, 39, 39, 82, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 83, 68, 42, 122, 235, 211, 168, 255, 246, 250, 255, 255, 203, 200, 194, 255, 251, 215, 155, 255, 166, 155, 136, 223, 8, 9, 11, 20, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 6, 136, 114, 78, 200, 193, 175, 145, 255, 194, 171, 133, 235, 152, 126, 82, 190, 88, 72, 44, 159, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 42, 35, 23, 73, 40, 31, 17, 81, 1, 0, 0, 16, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
}
