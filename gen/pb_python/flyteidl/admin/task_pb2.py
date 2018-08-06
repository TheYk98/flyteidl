# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: flyteidl/admin/task.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from flyteidl.core import tasks_pb2 as flyteidl_dot_core_dot_tasks__pb2
from flyteidl.admin import common_pb2 as flyteidl_dot_admin_dot_common__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='flyteidl/admin/task.proto',
  package='flyteidl.admin',
  syntax='proto3',
  serialized_pb=_b('\n\x19\x66lyteidl/admin/task.proto\x12\x0e\x66lyteidl.admin\x1a\x19\x66lyteidl/core/tasks.proto\x1a\x1b\x66lyteidl/admin/common.proto\"t\n\x11TaskCreateRequest\x12&\n\x02id\x18\x01 \x01(\x0b\x32\x1a.flyteidl.admin.Identifier\x12\x0f\n\x07version\x18\x02 \x01(\t\x12&\n\x04spec\x18\x03 \x01(\x0b\x32\x18.flyteidl.admin.TaskSpec\"p\n\x0fTaskListRequest\x12\x0f\n\x07project\x18\x01 \x01(\t\x12\x0e\n\x06\x64omain\x18\x02 \x01(\t\x12\x0c\n\x04name\x18\x03 \x01(\t\x12\r\n\x05limit\x18\x04 \x01(\r\x12\x0e\n\x06offset\x18\x05 \x01(\r\x12\x0f\n\x07\x66ilters\x18\x06 \x01(\t\"!\n\x12TaskCreateResponse\x12\x0b\n\x03urn\x18\x01 \x01(\t\"t\n\x04Task\x12&\n\x02id\x18\x01 \x01(\x0b\x32\x1a.flyteidl.admin.Identifier\x12\x0f\n\x07version\x18\x02 \x01(\t\x12\x0b\n\x03urn\x18\x03 \x01(\t\x12&\n\x04spec\x18\x06 \x01(\x0b\x32\x18.flyteidl.admin.TaskSpec\"?\n\x08TaskList\x12#\n\x05tasks\x18\x01 \x03(\x0b\x32\x14.flyteidl.admin.Task\x12\x0e\n\x06offset\x18\x02 \x01(\r\"9\n\x08TaskSpec\x12-\n\x08template\x18\x01 \x01(\x0b\x32\x1b.flyteidl.core.TaskTemplateB3Z1github.com/lyft/flyteidl/gen/pb-go/flyteidl/adminb\x06proto3')
  ,
  dependencies=[flyteidl_dot_core_dot_tasks__pb2.DESCRIPTOR,flyteidl_dot_admin_dot_common__pb2.DESCRIPTOR,])




_TASKCREATEREQUEST = _descriptor.Descriptor(
  name='TaskCreateRequest',
  full_name='flyteidl.admin.TaskCreateRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='flyteidl.admin.TaskCreateRequest.id', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='version', full_name='flyteidl.admin.TaskCreateRequest.version', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='spec', full_name='flyteidl.admin.TaskCreateRequest.spec', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=101,
  serialized_end=217,
)


_TASKLISTREQUEST = _descriptor.Descriptor(
  name='TaskListRequest',
  full_name='flyteidl.admin.TaskListRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='project', full_name='flyteidl.admin.TaskListRequest.project', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='domain', full_name='flyteidl.admin.TaskListRequest.domain', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='name', full_name='flyteidl.admin.TaskListRequest.name', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='limit', full_name='flyteidl.admin.TaskListRequest.limit', index=3,
      number=4, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='offset', full_name='flyteidl.admin.TaskListRequest.offset', index=4,
      number=5, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='filters', full_name='flyteidl.admin.TaskListRequest.filters', index=5,
      number=6, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=219,
  serialized_end=331,
)


_TASKCREATERESPONSE = _descriptor.Descriptor(
  name='TaskCreateResponse',
  full_name='flyteidl.admin.TaskCreateResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='urn', full_name='flyteidl.admin.TaskCreateResponse.urn', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=333,
  serialized_end=366,
)


_TASK = _descriptor.Descriptor(
  name='Task',
  full_name='flyteidl.admin.Task',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='flyteidl.admin.Task.id', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='version', full_name='flyteidl.admin.Task.version', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='urn', full_name='flyteidl.admin.Task.urn', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='spec', full_name='flyteidl.admin.Task.spec', index=3,
      number=6, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=368,
  serialized_end=484,
)


_TASKLIST = _descriptor.Descriptor(
  name='TaskList',
  full_name='flyteidl.admin.TaskList',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='tasks', full_name='flyteidl.admin.TaskList.tasks', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='offset', full_name='flyteidl.admin.TaskList.offset', index=1,
      number=2, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=486,
  serialized_end=549,
)


_TASKSPEC = _descriptor.Descriptor(
  name='TaskSpec',
  full_name='flyteidl.admin.TaskSpec',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='template', full_name='flyteidl.admin.TaskSpec.template', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=551,
  serialized_end=608,
)

_TASKCREATEREQUEST.fields_by_name['id'].message_type = flyteidl_dot_admin_dot_common__pb2._IDENTIFIER
_TASKCREATEREQUEST.fields_by_name['spec'].message_type = _TASKSPEC
_TASK.fields_by_name['id'].message_type = flyteidl_dot_admin_dot_common__pb2._IDENTIFIER
_TASK.fields_by_name['spec'].message_type = _TASKSPEC
_TASKLIST.fields_by_name['tasks'].message_type = _TASK
_TASKSPEC.fields_by_name['template'].message_type = flyteidl_dot_core_dot_tasks__pb2._TASKTEMPLATE
DESCRIPTOR.message_types_by_name['TaskCreateRequest'] = _TASKCREATEREQUEST
DESCRIPTOR.message_types_by_name['TaskListRequest'] = _TASKLISTREQUEST
DESCRIPTOR.message_types_by_name['TaskCreateResponse'] = _TASKCREATERESPONSE
DESCRIPTOR.message_types_by_name['Task'] = _TASK
DESCRIPTOR.message_types_by_name['TaskList'] = _TASKLIST
DESCRIPTOR.message_types_by_name['TaskSpec'] = _TASKSPEC
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

TaskCreateRequest = _reflection.GeneratedProtocolMessageType('TaskCreateRequest', (_message.Message,), dict(
  DESCRIPTOR = _TASKCREATEREQUEST,
  __module__ = 'flyteidl.admin.task_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.admin.TaskCreateRequest)
  ))
_sym_db.RegisterMessage(TaskCreateRequest)

TaskListRequest = _reflection.GeneratedProtocolMessageType('TaskListRequest', (_message.Message,), dict(
  DESCRIPTOR = _TASKLISTREQUEST,
  __module__ = 'flyteidl.admin.task_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.admin.TaskListRequest)
  ))
_sym_db.RegisterMessage(TaskListRequest)

TaskCreateResponse = _reflection.GeneratedProtocolMessageType('TaskCreateResponse', (_message.Message,), dict(
  DESCRIPTOR = _TASKCREATERESPONSE,
  __module__ = 'flyteidl.admin.task_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.admin.TaskCreateResponse)
  ))
_sym_db.RegisterMessage(TaskCreateResponse)

Task = _reflection.GeneratedProtocolMessageType('Task', (_message.Message,), dict(
  DESCRIPTOR = _TASK,
  __module__ = 'flyteidl.admin.task_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.admin.Task)
  ))
_sym_db.RegisterMessage(Task)

TaskList = _reflection.GeneratedProtocolMessageType('TaskList', (_message.Message,), dict(
  DESCRIPTOR = _TASKLIST,
  __module__ = 'flyteidl.admin.task_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.admin.TaskList)
  ))
_sym_db.RegisterMessage(TaskList)

TaskSpec = _reflection.GeneratedProtocolMessageType('TaskSpec', (_message.Message,), dict(
  DESCRIPTOR = _TASKSPEC,
  __module__ = 'flyteidl.admin.task_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.admin.TaskSpec)
  ))
_sym_db.RegisterMessage(TaskSpec)


DESCRIPTOR.has_options = True
DESCRIPTOR._options = _descriptor._ParseOptions(descriptor_pb2.FileOptions(), _b('Z1github.com/lyft/flyteidl/gen/pb-go/flyteidl/admin'))
# @@protoc_insertion_point(module_scope)
