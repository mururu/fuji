###########################
MQTT Gateway: Fuji
###########################

:version: 0.2.0

.. image:: https://circleci.com/gh/shiguredo/fuji/tree/develop.svg?style=svg&circle-token=203d959fffaf8dcdc0c68642dde5329e55a47792
    :target: https://circleci.com/gh/shiguredo/fuji/tree/develop

What is MQTT Gateway
=====================

**This definition is Shiguredo original**

A MQTT gateway is a sensor-MQTT gateway which receives data from sensors and sends that data to a MQTT broker.

Architecture::

    <sensor> -> (BLE)     ->             +-------+
    <sensor> -> (EnOcean) -> (Serial) -> |Gateway| -> (MQTT) -> <MQTT Broker>
    <sensor> -> (USB)     ->             +-------+

fuji is a MQTT gateway which is written by Golang.

Supported Hardware
====================

- `Raspberry Pi <http://www.raspberrypi.org/>`_ series
- `Armadillo-IoT <http://armadillo.atmark-techno.com/armadillo-iot>`_
- `Intel Edison <http://www.intel.com/content/www/us/en/do-it-yourself/edison.html?_ga=1.251267654.1109522025.1429502791>`_
- Linux i686/i386

Comming Soon

- Mac OS X
- FreeBSD
- Windows (7 or later)

Downloads
=========

:URL: https://github.com/shiguredo/fuji/releases/tag/0.2.0

- `fuji-gw_0.2.0_arm5.tar.gz <https://github.com/shiguredo/fuji/releases/download/0.2.0/fuji-gw_0.2.0_arm5.tar.gz>`_
- `fuji-gw_0.2.0_arm6.tar.gz <https://github.com/shiguredo/fuji/releases/download/0.2.0/fuji-gw_0.2.0_arm6.tar.gz>`_
- `fuji-gw_0.2.0_arm7.tar.gz <https://github.com/shiguredo/fuji/releases/download/0.2.0/fuji-gw_0.2.0_arm7.tar.gz>`_
- `fuji-gw_0.2.0_edison_386.ipk <https://github.com/shiguredo/fuji/releases/download/0.2.0/fuji-gw_0.2.0_edison_386.ipk>`_
- `fuji-gw_0.2.0_linux_386.tar.gz <https://github.com/shiguredo/fuji/releases/download/0.2.0/fuji-gw_0.2.0_linux_386.tar.gz>`_
- `fuji-gw_0.2.0_linux_amd64.tar.gz <https://github.com/shiguredo/fuji/releases/download/0.2.0/fuji-gw_0.2.0_linux_amd64.tar.gz>`_
- `fuji-gw_0.2.0_raspi2_arm7.deb <https://github.com/shiguredo/fuji/releases/download/0.2.0/fuji-gw_0.2.0_raspi2_arm7.deb>`_
- `fuji-gw_0.2.0_raspi_arm6.deb <https://github.com/shiguredo/fuji/releases/download/0.2.0/fuji-gw_0.2.0_raspi_arm6.deb>`_

ChangeLog
=========

see `CHANGELOG.rst <https://github.com/shiguredo/fuji/blob/develop/CHANGELOG.rst>`_

Build
=====

see `BUILD.rst <https://github.com/shiguredo/fuji/blob/develop/BUILD.rst>`_

Install
=======

see `INSTALL.rst <https://github.com/shiguredo/fuji/blob/develop/INSTALL.rst>`_

How to Contribute
=================

see `CONTRIBUTING.rst <https://github.com/shiguredo/fuji/blob/develop/CONTRIBUTING.rst>`_

Development Logs
========================

**Sorry for written in Japanese.**

**開発に関する詳細については開発ログをご覧ください**

`時雨堂 MQTT ゲートウェイ Fuji 開発ログ <https://gist.github.com/voluntas/23132cd3848af5b3ee1e>`_


License
========

::

  Copyright 2015 Shiguredo Inc. <fuji@shiguredo.jp>

  Licensed under the Apache License, Version 2.0 (the "License");
  you may not use this file except in compliance with the License.
  You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an "AS IS" BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  See the License for the specific language governing permissions and
  limitations under the License.
