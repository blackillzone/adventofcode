#!/bin/bash
# ===================================================
#
# Script to init a day from template folder, for Advent of Code
#
# ===================================================

# Output colors
readonly NORMAL="\\033[0;39m"
readonly RED="\\033[1;31m"
readonly YELLOW="\\033[1;33m"
readonly BLUE="\\033[1;34m"

# logging functions
log() {
  echo -e "$BLUE$(date +"%Y-%m-%d %H:%M:%S,%3N") [INFO] $1 $NORMAL"
}

warn() {
  echo -e "$YELLOW$(date +"%Y-%m-%d %H:%M:%S,%3N") [WARN] $1 $NORMAL"
}

error() {
  echo -e "$RED$(date +"%Y-%m-%d %H:%M:%S,%3N") [ERROR] $1$NORMAL"
}

# Scripts param
readonly BASE_DIR=$(readlink -m "$(dirname "$0")")
readonly SCRIPT_NAME=$(basename $0)
readonly BASE_OPTS=$*

# ===================================================
# Parameters checking

readonly USAGE="Usage: ${SCRIPT_NAME} \n\
  -action=<String corresponding to the action [mandatory: init/run/publish]>\n\
      init: Init the folder, following the template\n\
      run: Run the code, and output the results\n\
      publish: Publish the result to AOC website\n\
  -part=<Integer corresponding to the part needed to be published> [mandatory for publish action: current part available are 1 and 2]\n\
  -language=<String corresponding to the language used for the template> [mandatory: current languages available are go and python]\n\
  -day=<Integer corresponding to the day of the month> [optional if launched in December: default set to current day]\n\
  -year=<Integer corresponding to the yearh> [optional if launched in December: default set to current year]"

while [ $# -gt 0 ]
do
  case $1
  in

    -action\=*)
      readonly ACTION=$(echo $1 | sed -r 's/^[^=]+=//')
      shift 1
    ;;

    -year\=*)
      readonly YEAR=$(echo $1 | sed -r 's/^[^=]+=//')
      shift 1
    ;;

    -day\=*)
      readonly DAY=$(echo $1 | sed -r 's/^[^=]+=//')
      shift 1
    ;;

    -language\=*)
      readonly LANGUAGE=$(echo $1 | sed -r 's/^[^=]+=//')
      shift 1
    ;;

    -part\=*)
      readonly PART=$(echo $1 | sed -r 's/^[^=]+=//')
      shift 1
    ;;

    *)
      error "This parameter is unknown: $1"
      echo -e ${USAGE} && exit 1
	    shift 1
    ;;
  esac
done

[ -z "${YEAR}" ] && [ -z "${DAY}" ] && [ $(date -u +"%m") != 12 ] && error "We're currently not in December, so '-year' and '-day' parameters are mandatory." && echo -e ${USAGE} && exit 1
[ -z "${YEAR}" ] && log "'-year' parameter not set. Set to current year: $(date -u +"%Y")" && YEAR=$(date -u +"%Y")
[ -z "${DAY}" ] && log "'-day' parameter not set. Set to current day: $(date -u +"%d")" && DAY=$(date -u +"%d")
[ -z "${ACTION}" ] && error "'-action' parameter not set, and is mandatory." && echo -e ${USAGE} && exit 1
[ ${ACTION} != "init" ] && [ ${ACTION} != "run" ] && [ ${ACTION} != "publish" ] && error "${ACTION} is not supported. Try init, run or publish." && echo -e ${USAGE} && exit 1
[ ${ACTION} == "publish" ] && [ -z "${PART}" ] && error "'-part' parameter not set, and is mandatory, when publishing." && echo -e ${USAGE} && exit 1
[ ${ACTION} == "publish" ] && [ ${PART} != "1" ] && [ ${PART} != "2" ] && error "${PART} is not 1 or 2." && echo -e ${USAGE} && exit 1
[ -z "${LANGUAGE}" ] && error "'-language' parameter not set, and is mandatory." && echo -e ${USAGE} && exit 1
[ ${LANGUAGE} != "go" ] && [ ${LANGUAGE} != "python" ] && error "${LANGUAGE} is not supported. Try python or go" && echo -e ${USAGE} && exit 1

function init_aoc {
  if [ ! -d "$BASE_DIR/adventofcode_${YEAR}" ]; then
    log "Creating the yearly folder: adventofcode_${YEAR}"
    mkdir -p $BASE_DIR/adventofcode_${YEAR}
  fi

  if [ ! -d "$BASE_DIR/adventofcode_${YEAR}/day${DAY}" ]; then
    log "Creating the daily folder: day${DAY}"
    mkdir -p $BASE_DIR/adventofcode_${YEAR}/day${DAY}
  fi

  if [ ! -f "$BASE_DIR/adventofcode_${YEAR}/day${DAY}/readme.md" ]; then
    log "Retrieving puzzle description..."
    aoc r -d "${DAY}" -y "${YEAR}" >> "$BASE_DIR/adventofcode_${YEAR}/day${DAY}/readme.md"
  fi

  if [ ! -f "$BASE_DIR/adventofcode_${YEAR}/day${DAY}/input" ]; then
    log "Retrieving input..."
    aoc d -d "${DAY}" -y "${YEAR}" -f "$BASE_DIR/adventofcode_${YEAR}/day${DAY}/input"
  fi

  if [ ${LANGUAGE} == "go" ]; then
    if [ ! -f "$BASE_DIR/adventofcode_${YEAR}/day${DAY}/main.go" ]; then
      log "Copy go file... You should be ready to start and code !"
      cp "${BASE_DIR}/adventofcode_xxxx/day_xx/go/main.go" "$BASE_DIR/adventofcode_${YEAR}/day${DAY}/"
    fi
  fi

  if [ ${LANGUAGE} == "python" ]; then
    log "To do python"
  fi
}

function run_aoc_go {
  cd "$BASE_DIR/adventofcode_${YEAR}/day${DAY}/"
  go run "main.go"
}

function publish_aoc_go {
  cd "$BASE_DIR/adventofcode_${YEAR}/day${DAY}/"
  if [ ${PART} == "1" ]; then
    RESULT=$(go run "main.go" | sed -n '2p')
    aoc s 1 ${RESULT} -y ${YEAR} -d ${DAY}
  elif [ ${PART} == "2" ]; then
    RESULT=$(go run "main.go" | sed -n '4p')
    aoc s 2 ${RESULT} -y ${YEAR} -d ${DAY}
  fi
}

if [ ${ACTION} == "init" ]; then
  init_aoc
elif [ ${ACTION} == "run" ]; then
  if [ ${LANGUAGE} == "go" ]; then
    run_aoc_go
  elif [ ${LANGUAGE} == "python" ]; then
    log "To do python"
    # run_aoc_python
  fi
elif [ ${ACTION} == "publish" ]; then
  if [ ${LANGUAGE} == "go" ]; then
    publish_aoc_go
  elif [ ${LANGUAGE} == "python" ]; then
    log "To do python"
    # run_aoc_python
  fi
fi
