FROM golang:1.23 as backend-base
WORKDIR /emacs-web-doctor-backend
RUN apt update && apt install -y emacs


FROM backend-base as backend-build
COPY ./backend .
RUN CORS_ACTIVE=no go test --shuffle on ./... \
    # As doctor.el executes the (progn ...) simply provide an empty string.
    && echo "" | emacs --batch --quick --load ert --load doctor-test.el --load doctor.el --funcall ert-run-tests-batch-and-exit \
    && go build .


FROM node:22 AS frontend-base
WORKDIR /emacs-web-doctor-frontend

FROM frontend-base as frontend-build
COPY ./frontend .
RUN npm ci && PUBLIC_API_BASE_URL="" npm run build


FROM alpine:3
WORKDIR /emacs-web-doctor

RUN apk update \
    && apk add --no-cache gcompat \
    && apk add --no-cache bash \
    && apk add --no-cache emacs-nox

COPY --from=backend-build /emacs-web-doctor-backend .
COPY --from=frontend-build /emacs-web-doctor-frontend/build ./public
CMD ./backend
