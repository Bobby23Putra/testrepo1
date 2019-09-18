import pymodbus
import serial, struct, sys
from pymodbus.pdu import ModbusRequest
from pymodbus.client.sync import ModbusSerialClient as ModbusClient #initialize a serial RTU client instance
from pymodbus.transaction import ModbusRtuFramer
from pymodbus.constants import Endian
from pymodbus.payload import BinaryPayloadDecoder
import json
register = [3392, 2000, 4000] 
x = {
  "Cell1": register[0],
  "Cell2": register[1],
  "Cell3": register[2]
}

client = ModbusClient(method='rtu', port='/dev/ttyS2', timeout=1, stopbits = 1, bytesize = 8, parity = 'N', baudrate = 19200)
if client.connect() :
    print("berhasil")
else :
    print("gagal")
y = json.dumps(x)

print(y)