# Generated by Django 4.0 on 2021-12-16 09:40

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('AssicurazioneMedica', '0004_amministratore_session_valid_until_and_more'),
    ]

    operations = [
        migrations.AlterField(
            model_name='amministratore',
            name='email',
            field=models.CharField(max_length=40, unique=True),
        ),
        migrations.AlterField(
            model_name='cliente',
            name='email',
            field=models.CharField(max_length=40, unique=True),
        ),
    ]