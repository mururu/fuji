#########
ChangeLog
#########

0.2.0
=====

**Paied support and customize has been started**

:release: 2015-04-22

- MQTT 3.1.1
- MQTT over TLS
- Sending Device Status using `gopsutil <https://github.com/shirou/gopsutil>`_

  - CPU and memory information
  - Interval

- Will Message

  - ``will_message = ""`` in config.
  - topic_prefix
  - will message will be sent to fixed topic. ex: ``<Gateway-Name>/will``
- Retain

  - ``retain = true`` in config.
- Binary Message

  - To send binary, specify like ``\\x00\\x12`` 
- Work on `Armadillo-IoT <http://armadillo.atmark-techno.com/armadillo-iot>`_

  - MQTT over TLS via 3G
- Work on Intel Edison
- Infinity connection retry
- Subscribe
  - ``subscribe = true`` in config file.
  - Subscribe fixed topic (ex: ``<Gateway-Name>/<Device-Name>``). When a message comes, write the payload to the device.
- Binary for `Armadillo-IoT <http://armadillo.atmark-techno.com/armadillo-iot>`_

  - ARM5
- Packaged file for Raspberry Pi Model B+

  - Raspbian, ARM6

- Packaged file for Intel Edison

  - Yocto Linux

- Message sending Retry

  - If sending message failed, retry sometimes.
- Redundant server and switching each other

  - After retry to a primary broker failed, switch to secondary broker.
  - ``[device "sango/1"]`` and ``[device "sango/2"]``
- Multiple Broker

  - Connect to each broker independently.

- Japanese Document

0.1.0
=====

:release: 2015-02-13

- First internal release

- Serial port device

  - Acquire real sensro data from EnOcean Device

- Dummy Device
- MQTT 3.1
- MQTT Broker

  - single broker only

- Topic_prefix
- dpkg install

  - Raspbian

- Work on `Raspberry Pi <http://www.raspberrypi.org/>`_
- Package file for `Raspberry Pi`
- Japanese Document
