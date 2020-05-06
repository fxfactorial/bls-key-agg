FROM algebr/ostn-node

WORKDIR /src/
COPY main.go go.* /src/

ARG HMY_PATH=/src

ENV MCL_DIR=${HMY_PATH}/mcl
ENV BLS_DIR=${HMY_PATH}/bls
ENV CGO_CFLAGS="-I${BLS_DIR}/include -I${MCL_DIR}/include"
ENV CGO_LDFLAGS="-L${BLS_DIR}/lib"
ENV LD_LIBRARY_PATH=${BLS_DIR}/lib:${MCL_DIR}/lib

RUN apt install bash

RUN git clone https://github.com/harmony-one/mcl.git mcl

RUN cd mcl bls && make -j8

RUN git clone https://github.com/harmony-one/bls.git bls

RUN cd bls && make -j8 BLS_SWAP_G=1

RUN go build -o /bin/demo

CMD /bin/demo
