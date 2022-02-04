FROM golang:1.17.6-bullseye

RUN groupadd -g 1000 -o goxygenuser

RUN useradd -m -u 1000 -g 1000 -o -s /bin/bash goxygenuser

USER goxygenuser

WORKDIR /goxygen

RUN mkdir generated 

COPY . .

# Without this line, the docker image will not be able to run and produce error: 
# `Failed to setup a Git repository: exit status 128`
# Because there is a line of code that stage files to the git repository and commit on
# codegen > codegen.go > function initGitRepo 
RUN git config --global user.email "arlhba@gmail.com" && git config --global user.name "Luqmanul Hakim"

RUN go build -o goxygen

ENV GOXYGEN_DOCKER=true

ENTRYPOINT [ "/goxygen/goxygen" ]