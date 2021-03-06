.. _servers:

servers
=======

Commands for Rackspace Cloud Servers, dedicated and virtual.

Synopsis
--------

::

   rack servers <resource> <action> [command flags]

Commands
--------

``instance``
~~~~~~~~~~~~

  Server Instance operations

``list``
^^^^^^^^
Usage::

    rack servers instance list [optional flags]

``create``
^^^^^^^^^^
Usage::

    rack servers instance create --name <serverName> [optional flags]
    (echo serverName1 && echo serverName2) | rack servers instance create --stdin name [optional flags]

``get``
^^^^^^^
Usage::

    rack servers instance get --id <serverID> [optional flags]
    rack servers instance get --name <serverName> [optional flags]
    (echo serverID1 && echo serverID2) | rack servers instance get --stdin id [optional flags]

``update``
^^^^^^^^^^
Usage::

    rack servers instance update --id <serverID> [optional flags]
    rack servers instance update --name <serverName> [optional flags]

``delete``
^^^^^^^^^^
Usage::

    rack servers instance delete --id <serverID> [optional flags]
    rack servers instance delete --name <serverName> [optional flags]
    (echo serverID1 && echo serverID2) | rack servers instance delete --stdin id [optional flags]

``reboot``
^^^^^^^^^^
Usage::

    rack servers instance reboot --id <serverID> [--soft | --hard] [optional flags]
    rack servers instance reboot --name <serverName> [--soft | --hard] [optional flags]
    (echo serverID1 && echo serverID2) | rack servers instance reboot --stdin id [--soft | --hard] [optional flags]

``rebuild``
^^^^^^^^^^^
Usage::

    rack servers instance rebuild --id <serverID> --image-id <imageID> --admin-pass <adminPass> [optional flags]
    rack servers instance rebuild --name <serverName> --image-id <imageID> --admin-pass <adminPass> [optional flags]

``resize``
^^^^^^^^^^
Usage::

    rack servers instance resize --id <serverID> --flavor-id <flavorID> [optional flags]
    rack servers instance resize --name <serverName> --flavor-id <flavorID> [optional flags]
    (echo serverID1 && echo serverID2) | rack servers instance resize --stdin id --flavor-id <flavorID> [optional flags]

``set-metadata``
^^^^^^^^^^^^^^^^
Usage::

    rack servers instance set-metadata --id <serverID> --metadata <key1=val1,key2=val2,...> [optional flags]
    rack servers instance set-metadata --name <serverName> --metadata <key1=val1,key2=val2,...> [optional flags]

``get-metadata``
^^^^^^^^^^^^^^^^
Usage::

    rack servers instance get-metadata --id <serverID> [optional flags]
    rack servers instance get-metadata --name <serverName> [optional flags]

``update-metadata``
^^^^^^^^^^^^^^^^
Usage::

    rack servers instance update-metadata --id <serverID> --metadata <key1=val1,key2=val2,...> [optional flags]
    rack servers instance update-metadata --name <serverName> --metadata <key1=val1,key2=val2,...> [optional flags]

``delete-metadata``
^^^^^^^^^^^^^^^^
Usage::

    rack servers instance delete-metadata --id <serverID> --metadata <key1=val1,key2=val2,...> [optional flags]
    rack servers instance delete-metadata --name <serverName> --metadata <key1=val1,key2=val2,...> [optional flags]


``image``
~~~~~~~~~

  Server Image operations

``list``
^^^^^^^^
Usage::

    rack servers image list [optional flags]

``get``
^^^^^^^^
Usage::

    rack servers image get --id <imageID> [optional flags]
    rack servers image get --name <imageName>] [optional flags]
    (echo imageID1 && echo imageID2) | rack servers image get --stdin id [optional flags]

Note: To guarantee usage of the same image every time, use the `id` flag. Images often
are updated with security patches, and the updated images will have a different ID but
the same name.


``flavor``
~~~~~~~~~~

  Server Flavor operations

``list``
^^^^^^^^
Usage::

    rack servers flavor list [optional flags]

``get``
^^^^^^^
Usage::

    rack servers flavor get --id <flavorID> [optional flags]
    rack servers flavor get --name <flavorName>] [optional flags]
    (echo flavorID1 && echo flavorID2) | rack servers flavor get --stdin id [optional flags]

``keypair``
~~~~~~~~~~~

  Server Keypair operations

``list``
^^^^^^^^
Usage::

    rack servers keypair list [flags]

``generate``
^^^^^^^^^^
Usage::

    rack servers keypair generate --name <keypairName> [optional flags]
    (echo keypairName1 && echo keypairName2) | rack servers keypair generate --stdin name [optional flags]

``upload``
^^^^^^^^^^
Usage::

    rack servers keypair upload --name <keypairName> --public-key <publicKeyData> [optional flags]
    rack servers keypair upload --name <keypairName> --file <publicKeyfile> [optional flags]

``get``
^^^^^^^
Usage::

    rack [globals] servers keypair get --name <keypairName> [optional flags]
    (echo keypairName1 && echo keypairName2) | rack servers keypair get --stdin name [optional flags]

``delete``
^^^^^^^^^^
Usage::

    rack servers keypair delete --name <keypairName> [optional flags]
    (echo keypairName1 && echo keypairName2) | rack servers keypair delete --stdin name [optional flags]


``volume-attachment``
~~~~~~~~~~~~~~~~~~~~~

  Volume Attachment operations

``list``
^^^^^^^^
Usage::

    rack servers volume-attachment list --server-id <serverID> [optional flags]
    rack servers volume-attachment list --server-name <serverName> [optional flags]
    rack servers volume-attachment list --stdin server-id [optional flags]

``create``
^^^^^^^^^^
Usage::

    rack servers volume-attachment create --server-id <serverID> --volume-id <volumeID> [optional flags]
    rack servers volume-attachment create --server-name <serverName> --volume-id <volumeID> [optional flags]
    rack servers volume-attachment create --server-id <serverID> --volume-name <volumeName> [optional flags]
    rack servers volume-attachment create --server-name <serverName> --volume-name <volumeName> [optional flags]
    (echo volumeID1 && echo volumeID2) | rack servers volume-attachment create --server-id <serverID> --stdin volume-id [optional flags]
    (echo volumeID1 && echo volumeID2) | rack servers volume-attachment create --server-name <serverName> --stdin volume-id [optional flags]

``get``
^^^^^^^
Usage::

    rack servers volume-attachment get --server-id <serverID> --id <attachmentID> [optional flags]
    rack servers volume-attachment get --server-name <serverName> --id <attachmentID> [optional flags]

``delete``
^^^^^^^^^^
Usage::

    rack servers volume-attachment delete --server-id <serverID> --id <attachmentID> [optional flags]
    rack servers volume-attachment delete --server-name <serverName> --id <attachmentID> [optional flags]
