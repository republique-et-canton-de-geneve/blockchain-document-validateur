<?php

/**
 * Implements hook_panels_layouts().
 */
// Plugin definition
$plugin = array(
  'title' => t('One Column Multiple'),
  'category' => t('Bootstrap'),
  'icon' => 'bootstrap_onecol_multiple.png',
  'theme' => 'bootstrap_onecol_multiple',
  'admin css' => 'bootstrap_onecol_multiple.css',
  'regions' => array(
    'navigation' => t('Navigation'),
    'top' => t('Top'),
    'middle' => t('Middle'),
    'bottom' => t('Bottom'),
  ),
);
