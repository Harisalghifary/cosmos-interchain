# Project description

You should read this whole page before attacking the problem, including the part about Docker and your working copy.

The exercise is divided into 4 parts that you should take in order. The 4 parts of the exercise are disclosed bit by bit, and here are the 1st and 2nd parts. The 4 parts are for the same project, this project.

This is the beginning of a larger project dedicated to the future payments of [toll roads](https://en.wikipedia.org/wiki/Toll_road). The project, built with Ignite 0.22.1 and CosmJS 0.28.13, is far from complete.

However it already contains:

1. A single `SystemInfo`. It is meant to keep track of system-wide information, in this case, the ID of the next road operator created. It was created with:

    ```sh
    $ ignite scaffold single SystemInfo nextOperatorId:uint --no-message
    ```

2. A map of `RoadOperator`s. It was created with:

    ```sh
    $ ignite scaffold map RoadOperator name token active:bool
    ```

    The default behavior of the scaffolding command is to have the index of the operator come from `MsgCreateRoadOperator`. However, `index` was removed from `MsgCreateRoadOperator`. That's because when a user creates a new road operator, the user does not choose the ID. Instead, it is chosen by the system on creation. After rebuild, various compilation errors were "fixed" in a lazy way.

3. More actions to be disclosed in subsequent parts.

## Do now

Before doing any work, and before you forget, do these steps:

1. Clone the repository.
2. Build the Docker image with:

    ```sh
    $ docker build -f Dockerfile-exam . -t exam_i
    ```

3. Read this document in full.

This is not to save you time. Instead this is to save you mistakes and surprises down the road.

## Your work

The following steps are in order of increasing difficulty:

1. Adjustments on system info.
2. Adjustments on road operators.
3. More actions to be disclosed in subsequent parts.

The tests have been divided into different packages to avoid compilation errors while your project is incomplete.

### On system info

Adjust `x/tollroad/types/genesis.go` so that the `x/tollroad/genesis_test.go` tests pass. To confirm, run:

```sh
$ go test github.com/b9lab/toll-road/x/tollroad
```

### On road operators

When a road operator is created, its ID has to be taken from `SystemInfo`. For this part, you are going to work only in `x/tollroad/keeper/msg_server_road_operator.go` and in it, only adjust the `CreateRoadOperator` function body. **Not the function signature, not another function, not another file.**

This is what you have to implement:

1. Make sure that the new road operator has its ID taken from `SystemInfo`.
2. Have this ID returned by the message server function.
3. Make sure the next id in `SystemInfo` is incremented.
4. Emit an event with the expected type and attributes.

To confirm, run:

```sh
$ go test github.com/b9lab/toll-road/x/tollroad/roadoperatorstudent
```

Look into the `x/tollroad/roadoperatorstudent/msg_server_road_operator_test.go` file to see what is expected, in particular the details of the expected event.

### On subsequent parts

They will be disclosed at some point.

## Preparing your working copy

This is your personal repository, it is named like `IDA-P5-final-exam/student-projects/YOUR_NAME-exam-code`, is forked from a public _upstream_ repository named like `IDA-P5-final-exam/exam-code`, and this is where you need to upload your work.

To prepare your working copy:

1. Clone your repository:

    ```sh 
    $ git clone https://YOUR_NAME@git.academy.b9lab.com/ida-p5-final-exam/student-projects/YOUR_NAME-exam-code.git
    ```

    Where you have replaced `YOUR_NAME` with its actual value in both places. If you are unsure of your username, you can find it [here](https://git.academy.b9lab.com/-/profile/account). For instance:

    ```sh
    $ git clone https://alice@git.academy.b9lab.com/ida-p5-final-exam/student-projects/alice-exam-code.git 
    ```

2. Work on the exercise. When you are happy with the result(s). Commit (the messages matter for you but not for the grading) and push.

    ```sh
    $ git add THE_FILES_YOU_ARE_COMMITTING
    $ git commit -m "Add my submission."
    $ git push
    ```

3. You can keep `push`ing as many further commits as you want before the deadline.

### Subsequent parts

When new parts are disclosed, you will see a [merge request](../../merge_requests) from the upstream repository towards your personal repository. When you are ready to work on the next parts, you can simply merge the MR. There should not be any conflict if you did not modify files mentioned in `fileconfig.yml`.

Note that even after accepting a merge request, you can still work on the previous parts of the exercise. It is only a disclosure mechanism, it does not freeze previous scores in any way.

### Deadline

After the submission deadline, we will run the tests on your latest version of the `main` branch, not on an intermediate commit you made.

If you wish to use the SSH protocol instead of supplying your username and password over HTTPS to perform Git operations like clone/pull/push, you can learn how to [handle SSH keys](https://docs.gitlab.com/ee/user/ssh.html#generate-an-ssh-key-pair) and add your SSH keys [here](https://git.academy.b9lab.com/-/profile/keys).

## Looking at test files

Make sure you **read the section about Docker** before you jump into running anything.

For your convenience, here are the `go test` commands that the grading scripts will run:

```sh
$ go test github.com/b9lab/toll-road/x/tollroad
$ go test github.com/b9lab/toll-road/x/tollroad/roadoperatorstudent
$ go test github.com/b9lab/toll-road/x/tollroad/uservaultstudent
```

Each of them is actually launched from a script, respectively:

```sh
$ ./x/tollroad/testing.sh
$ ./x/tollroad/roadoperatorstudent/testing.sh
$ ./x/tollroad/uservaultstudent/testing.sh
```

Inside each of them (after disclosure) you can see how much weight is given to each test.

The NPM tests are launched from `./testing-cosmjs.sh` and each of the 4 tests has the same weight.

In turn, these 4 `testing` scripts are launched from this script:

```sh
$ ./score.sh
```

When `score.sh` has run well, irrespective of the completeness of the exercise, it outputs a score such as:

```txt
FS_SCORE:4%
```

This is your score as we will record it.

Inside `score.sh` and the 4 `testing`, you can see how much weight is given to each individual test and to each part of the exercise. Namely:

* About `SystemInfo`: [1](score.sh#L12)
* About `RoadOperator`: [2](score.sh#L22)
* About `UserVault`: [4](score.sh#L32)
* About ComsJS: [2](score.sh#L48)

You can run these scripts (inside Docker or not, see below) as many times as you want, and see what score that would give you. We do not track how many times you run them and do not keep a score on our end when _you_ run them.

Passing all tests gives you 100%, passing none gives you 0%, and from the start you may get 0% or 4%. Passing a single one gives you a score that depends on the weights applied.

How do you test the outcome? That's the object of testing it in Docker. The paragraphs also discuss read-only files.

## Running tests in Docker

This is the right way to run the tests, as it will let you test the outcome just the way we will test it after you have submitted your work to us. Install [Docker](https://docs.docker.com/engine/install/).

This project was actually built within a container described by the Docker file `Dockerfile-exam`. Do the following to run your tests in Docker too.

* Create the image:
  
    ```sh
    $ docker build -f Dockerfile-exam . -t exam_i
    ```

    It will take time because it downloads the [Go](Dockerfile-exam#L48) and [NPM](Dockerfile-exam#L53) dependencies onto the image so that the containers do not have to do it every time you run the tests.

    In fact you should create the image **as soon as you cloned it** and **after every merge request**. Indeed, it makes [a copy](Dockerfile-exam#L47) of the current directory into `/original` so as to have a snapshot of the files that have to remain read-only (more about that below).

* Run the scoring in a throwaway container just like you did earlier:

    ```sh
    $ docker run --rm -it -v $(pwd):/exam -w /exam exam_i ./score.sh
    ```

* Run the scoring like we will do when you have submitted your work:

    ```sh
    $ docker run --rm -it -v $(pwd):/exam -w /exam exam_i /original/score-ci.sh
    ```

Notes on `score-ci.sh`:

* It is launched from the `/original` folder which is within the container, and then runs `score.sh` on `/exam` which is your working folder.
* It starts by [overwriting](score-ci.sh#L5) all the files that should be [read-only](fileconfig.yml#L7), by taking the original ones in `/original`. So be sure to make a commit so as not to lose work. If it appears that you have edited a file that is supposedly read-only, then you have to make further commits to undo your changes.
* This is the script that we run on our Gitlab runner and it will overwrite any read-only file that you updated.
* The score you get here is the closest to the score we will get when running it on our side.

Now, if you want something more permanent to debug your code:

* Build a reusable container:

    ```sh
    $ docker create --name exam -i -v $(pwd):/exam -w /exam -p 1317:1317 -p 4500:4500 -p 26657:26657 exam_i
    $ docker start exam
    ```

* Test a single part on the reusable container. Connect to it:

    ```sh
    $ docker exec -it exam bash
    ```

    Then in this connected shell:

    ```sh
    $ go test github.com/b9lab/toll-road/x/tollroad
    ```

    Or, in the same shell:

    ```sh
    $ ./score.sh
    ```

    Or even:

    ```sh
    $ /original/score-ci.sh
    ```

It is possible that you could do without Docker if you have all the right tools and versions installed already. But better use Docker and `/original/score-ci.sh` at the end at least to avoid any surprises.

## Official grading

To grade your project, we run the same `score-ci.sh` file via a Gitlab pipeline. In fact, you can see yourself the status of the pipeline's jobs [here](../../jobs) or from the "CI/CD -> Jobs" menu in the navigation on the left.

1. If a job named `test` has a _Passed_ status, it means that the job itself ran as expected. To see your score, click on it and look at the bottom of the log where you can find the `FS_SCORE` information.
2. If a job named `test` has a _Failed_ status, it means that the job failed, usually because of a timeout, and there is no score to collect. You can send it back to the job queue by clicking the relaunch button. If it fails twice in a row, you should alert us.

Have a look inside the [`fileconfig.yml`](./fileconfig.yml) file. Under the `readonly_paths` label is the list of files and folders that our script will overwrite in your repository before running the tests. If you modify any of these files, you may have inconsistent results between what you see _on your machine_ and _on the pipeline's job_.

After the submission deadline, we will make sure all `test` jobs _Passed_ and will collect all your scores.

## Submission checklist

Before you consider your work done, check everything off this list:

1. There are no pending [merge requests](../../merge_requests).
2. You committed your work in the `main` branch.
3. You pushed your `main` branch to the `main` branch on our Academy Gitlab server.
4. On the [jobs](../../jobs) page, your latest job has completed and at the end of it, you see the same `FS_SCORE` that you got when running locally on your computer. If the [jobs](../../jobs) link does not work:

    * It is not because the page does not exist.
    * You need to open it from the repository itself.

## Cleanup

* If you have worked a bit on the project, are not happy and want to start over, you can use Git's:

    ```sh
    $ git stash -u && git stash drop
    ```

* If you want to remove the Docker elements you created:

    ```sh
    # If you created and started the reusable container
    $ docker stop exam
    # If you created the reusable container
    $ docker rm exam
    # If you created the image
    $ docker rmi exam_i
    ```

Please do not publish your solution to the outside world so that we can reuse the exercise or parts of it.

Good luck.
